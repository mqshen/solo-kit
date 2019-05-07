// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"fmt"

	"go.opencensus.io/trace"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/eventloop"
	"github.com/solo-io/solo-kit/pkg/errors"
)

// a Syncer which implements this interface
// can make smarter decisions over whether
// it should be restarted (including having its context cancelled)
// based on a diff of the previous and current snapshot
type KubeconfigsSyncDecider interface {
	KubeconfigsSyncer
	ShouldSync(old, new *KubeconfigsSnapshot) bool
}

type kubeconfigsSimpleEventLoop struct {
	emitter KubeconfigsSimpleEmitter
	syncers []KubeconfigsSyncer
}

func NewKubeconfigsSimpleEventLoop(emitter KubeconfigsSimpleEmitter, syncers ...KubeconfigsSyncer) eventloop.SimpleEventLoop {
	return &kubeconfigsSimpleEventLoop{
		emitter: emitter,
		syncers: syncers,
	}
}

func (el *kubeconfigsSimpleEventLoop) Run(ctx context.Context) (<-chan error, error) {
	ctx = contextutils.WithLogger(ctx, "v1.event_loop")
	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("event loop started")

	errs := make(chan error)

	watch, emitterErrs, err := el.emitter.Snapshots(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "starting snapshot watch")
	}

	go errutils.AggregateErrs(ctx, errs, emitterErrs, "v1.emitter errors")
	go func() {
		// create a new context for each syncer for each loop, cancel each before each loop
		syncerCancels := make(map[KubeconfigsSyncer]context.CancelFunc)

		// use closure to allow cancel function to be updated as context changes
		defer func() {
			for _, cancel := range syncerCancels {
				cancel()
			}
		}()

		// cache the previous snapshot for comparison
		var previousSnapshot *KubeconfigsSnapshot

		for {
			select {
			case snapshot, ok := <-watch:
				if !ok {
					return
				}

				// cancel any open watches from previous loop
				for _, syncer := range el.syncers {
					// allow the syncer to decide if we should sync it + cancel its previous context
					if syncDecider, isDecider := syncer.(KubeconfigsSyncDecider); isDecider {
						if shouldSync := syncDecider.ShouldSync(previousSnapshot, snapshot); !shouldSync {
							continue // skip syncing this syncer
						}
					}

					// if this syncer had a previous context, cancel it
					cancel, ok := syncerCancels[syncer]
					if ok {
						cancel()
					}

					ctx, span := trace.StartSpan(ctx, fmt.Sprintf("kubeconfigs.multicluster.solo.io.SimpleEventLoopSync-%T", syncer))
					ctx, canc := context.WithCancel(ctx)
					err := syncer.Sync(ctx, snapshot)
					span.End()

					if err != nil {
						select {
						case errs <- err:
						default:
							logger.Errorf("write error channel is full! could not propagate err: %v", err)
						}
					}

					syncerCancels[syncer] = canc
					previousSnapshot = snapshot
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return errs, nil
}
