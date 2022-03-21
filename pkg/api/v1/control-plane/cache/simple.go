// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/jsonpb"
	gogoproto "github.com/golang/protobuf/proto"
	goproto "github.com/golang/protobuf/proto"
	proto2 "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	pany "github.com/golang/protobuf/ptypes/any"
	"github.com/solo-io/solo-kit/pkg/api/v1/control-plane/log"
	"google.golang.org/protobuf/types/known/anypb"
	"sync"
	"sync/atomic"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

var (
	MResponses = stats.Int64("xds/responses", "The responses for envoy", "1")

	KeyType = tag.MustNewKey("type")

	ResponsesView = &view.View{
		Name:        "xds/responses",
		Measure:     MResponses,
		Description: "The times gloo responded to xds request",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			KeyType,
		},
	}

	VersionUpToDateError = errors.New("skip fetch: version up to date")
)

func init() {
	view.Register(ResponsesView)
}

// SnapshotCache is a snapshot-based cache that maintains a single versioned
// snapshot of responses per node. SnapshotCache consistently replies with the
// latest snapshot. For the protocol to work correctly in ADS mode, EDS/RDS
// requests are responded only when all resources in the snapshot xDS response
// are named as part of the request. It is expected that the CDS response names
// all EDS clusters, and the LDS response names all RDS routes in a snapshot,
// to ensure that Envoy makes the request for all EDS clusters or RDS routes
// eventually.
//
// SnapshotCache can operate as a REST or regular xDS backend. The snapshot
// can be partial, e.g. only include RDS or EDS resources.
type SnapshotCache interface {
	Cache

	// SetSnapshot sets a response snapshot for a node. For ADS, the snapshots
	// should have distinct versions and be internally consistent (e.g. all
	// referenced resources must be included in the snapshot).
	//
	// This method will cause the server to respond to all open watches, for which
	// the version differs from the snapshot version.
	SetSnapshot(node string, snapshot Snapshot) error

	// GetSnapshots gets the snapshot for a node.
	GetSnapshot(node string) (Snapshot, error)

	// ClearSnapshot removes all status and snapshot information associated with a node.
	ClearSnapshot(node string)

	// deprecated: used to persist snapshot to disk
	Serialize() []byte

	// deprecated: used to initialize snapshot from disk
	Deserialize(bytes []byte)
}

// compile-time assertion
var _ SnapshotCache = &snapshotCache{}
type snapshotCache struct {
	log log.Logger

	// ads flag to hold responses until all resources are named
	ads bool

	// snapshots are cached resources indexed by node IDs
	snapshots map[string]Snapshot

	// status information for all nodes indexed by node IDs
	status map[string]*statusInfo

	// hash is the hashing function for Envoy nodes
	hash NodeHash

	// watchCount is an atomic counter incremented for each watch
	watchCount int64

	mu sync.RWMutex

	// deprecated: remove me
	filePath string
}

func (cache *snapshotCache) Deserialize(bytes []byte) {
	gs := &GenericSnapshot{}
	gs.Deserialize(bytes)
	cache.snapshots["test"] = gs
	//panic("implement me")
}

func (cache *snapshotCache) Serialize() []byte {
	// let's start with a single resource
	snap := cache.snapshots["test"]
	clusters := snap.GetResources("type.googleapis.com/envoy.config.cluster.v3.Cluster")//resource.ClusterTypeV3)

	item := clusters.Items["test"]
	//m, err := proto.Marshal(proto2.MessageV2(item.ResourceProto()))
	//if err != nil {
	//	panic(err)
	//}
	//json123, err := MarshalBytes(item)
	//if err != nil {
	//	panic(err)
	//}

	//any := MustMessageToAny(item.ResourceProto())
	//if err != nil {
	//	panic(err)
	//}

	any, err := anypb.New(proto2.MessageV2(item.ResourceProto()))
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = jsonpbMarshaler.Marshal(buf, any)
	if err != nil {
		panic(err)
	}
	//bts2 := buf.Bytes()

	//self := item.Self()
	topLvlBytes, err := json.Marshal(&cache.snapshots)
	if err != nil {
		panic(err)
	}

	//y, _ := yaml.Marshal(cache.snapshots)
	//fmt.Printf("yaml:\n%v\n", string(y))
	//fmt.Printf("item self %v\n", item.Self())
	//fmt.Printf("output higher level json: %v\n", string(topLvlBytes))

	ifce := map[string]interface{}{}
	yaml.Unmarshal(topLvlBytes, &ifce)

	//c := ifce["Clusters"]
	//cluster := &cluster.Cluster{}
	//yaml.Unmarshal(c, &cluster)
	//NewEnvoyResource(cluster)


	serialized := map[string][]byte{}
	for k, v := range cache.snapshots {
		serialized[k] = v.Serialize()
	}

	byts, err := json.Marshal(&serialized)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("output any json: %v\n", string(bts2))
	//fmt.Printf("output any json: %v\n", any.String())
	//fmt.Printf("output json: %v\n", string(json123))
	//fmt.Printf("output string: %v\n", item.ResourceProto().String())
	//fmt.Printf("output marshal: %v\n", m)
	return byts
	//b, err := json.Marshal(item)
	//if err != nil {
	//	panic(err)
	//}
	//return b
}

func MessageToAny(msg proto2.Message) (*pany.Any, error) {

	name, err := protoToMessageName(msg)
	if err != nil {
		return nil, err
	}
	buf, err := protoToMessageBytes(msg)
	if err != nil {
		return nil, err
	}
	return &pany.Any{
		TypeUrl: name,
		Value:   buf,
	}, nil
}

func MustAnyToMessage(a *pany.Any) proto2.Message {
	var x ptypes.DynamicAny
	err := ptypes.UnmarshalAny(a, &x)
	if err != nil {
		panic(err)
	}
	return x.Message
}

func MustMessageToAny(msg proto2.Message) *pany.Any {
	anymsg, err := MessageToAny(msg)
	if err != nil {
		panic(err)
	}
	return anymsg
}

func protoToMessageName(msg proto2.Message) (string, error) {
	typeUrlPrefix := "type.googleapis.com/"

	if s := gogoproto.MessageName(msg); s != "" {
		return typeUrlPrefix + s, nil
	} else if s := goproto.MessageName(msg); s != "" {
		return typeUrlPrefix + s, nil
	}
	return "", fmt.Errorf("can't determine message name")
}

func protoToMessageBytes(msg proto2.Message) ([]byte, error) {
	if b, err := protoToMessageBytesGolang(msg); err == nil {
		return b, nil
	}
	return protoToMessageBytesGogo(msg)
}

func protoToMessageBytesGogo(msg proto2.Message) ([]byte, error) {
	b := gogoproto.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(msg)
	return b.Bytes(), err
}

func protoToMessageBytesGolang(msg proto2.Message) ([]byte, error) {
	b := proto2.NewBuffer(nil)
	b.SetDeterministic(true)
	err := b.Marshal(msg)
	return b.Bytes(), err
}

var jsonpbMarshaler = &jsonpb.Marshaler{OrigName: false}
var jsonpbMarshalerEmitZeroValues = &jsonpb.Marshaler{OrigName: false, EmitDefaults: true}
var jsonpbMarshalerEnumsAsInts = &jsonpb.Marshaler{OrigName: false, EnumsAsInts: true}

func MarshalBytes(res Resource) ([]byte, error) {
	if pb, ok := res.(proto2.Message); ok {
		buf := &bytes.Buffer{}
		err := jsonpbMarshaler.Marshal(buf, pb)
		return buf.Bytes(), err
	}
	return json.Marshal(res)
}

// NewSnapshotCache initializes a simple cache.
//
// ADS flag forces a delay in responding to streaming requests until all
// resources are explicitly named in the request. This avoids the problem of a
// partial request over a single stream for a subset of resources which would
// require generating a fresh version for acknowledgement. ADS flag requires
// snapshot consistency. For non-ADS case (and fetch), mutliple partial
// requests are sent across multiple streams and re-using the snapshot version
// is OK.
//
// Logger is optional.
func NewSnapshotCache(ads bool, hash NodeHash, logger log.Logger) SnapshotCache {
	return &snapshotCache{
		log:       logger,
		ads:       ads,
		snapshots: make(map[string]Snapshot),
		status:    make(map[string]*statusInfo),
		hash:      hash,
	}
}

// deprecated: prefer NewSnapshotCache once https://github.com/solo-io/gloo/issues/6114 is resolved
func NewSnapshotCacheFromBackup(ads bool, hash NodeHash, logger log.Logger, filePath string) SnapshotCache {
	sc := &snapshotCache{
		log:       logger,
		ads:       ads,
		snapshots: make(map[string]Snapshot),
		status:    make(map[string]*statusInfo),
		hash:      hash,
		filePath: filePath,
	}

	// TODO(kdorosh) initialize snapshot cache
	return sc
}

// SetSnapshotCache updates a snapshot for a node.
func (cache *snapshotCache) SetSnapshot(node string, snapshot Snapshot) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if len(cache.filePath) > 0 {
		// TODO(kdorosh)
	}

	// update the existing entry
	cache.snapshots[node] = snapshot

	// trigger existing watches for which version changed
	if info, ok := cache.status[node]; ok {
		info.mu.Lock()
		for id, watch := range info.watches {
			version := snapshot.GetResources(watch.Request.TypeUrl).Version
			if version != watch.Request.VersionInfo {
				if cache.log != nil {
					cache.log.Infof("respond open watch %d%v with new version %q", id, watch.Request.ResourceNames, version)
				}

				resources := snapshot.GetResources(watch.Request.TypeUrl).Items
				// Before sending a response, need to be able to differentiate between a resource in the snapshot that does not exist vs. should be empty
				// nil resource - not set in the snapshot and should not be updated
				// empty resource - intended behavior is for the resources to be cleared
				if resources != nil {
					// snapshot has been initialized and exists
					cache.respond(watch.Request, watch.Response, resources, version)

					// discard the watch
					delete(info.watches, id)
				}
			}
		}
		info.mu.Unlock()
	}

	return nil
}

// Returns a copy of the snapshot for the given node, or an error if not found.
func (cache *snapshotCache) GetSnapshot(node string) (Snapshot, error) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	snap, ok := cache.snapshots[node]
	if !ok {
		return nil, fmt.Errorf("no snapshot found for node %s", node)
	}

	return snap.Clone(), nil
}

// ClearSnapshot clears snapshot and info for a node.
func (cache *snapshotCache) ClearSnapshot(node string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.snapshots, node)
	delete(cache.status, node)
}

// nameSet creates a map from a string slice to value true.
func nameSet(names []string) map[string]bool {
	set := make(map[string]bool)
	for _, name := range names {
		set[name] = true
	}
	return set
}

// Superset checks that all resources are listed in the names set.
func Superset(names map[string]bool, resources map[string]Resource) error {
	for resourceName := range resources {
		if _, exists := names[resourceName]; !exists {
			return fmt.Errorf("%q not listed", resourceName)
		}
	}
	return nil
}

// CreateWatch returns a watch for an xDS request.
func (cache *snapshotCache) CreateWatch(request Request) (chan Response, func()) {
	nodeID := cache.hash.ID(request.Node)

	cache.mu.Lock()
	defer cache.mu.Unlock()

	info, ok := cache.status[nodeID]
	if !ok {
		info = NewStatusInfo(request.Node)
		cache.status[nodeID] = info
	}

	// update last watch request time
	info.mu.Lock()
	info.lastWatchRequestTime = time.Now()
	info.mu.Unlock()

	// allocate capacity 1 to allow one-time non-blocking use
	value := make(chan Response, 1)

	snapshot, exists := cache.snapshots[nodeID]
	if snapshot == nil {
		snapshot = NilSnapshot{}
	}
	version := snapshot.GetResources(request.TypeUrl).Version

	// if the requested version is up-to-date or missing a response, leave an open watch
	if !exists || request.VersionInfo == version {
		watchID := cache.nextWatchID()
		if cache.log != nil {
			cache.log.Infof("open watch %d for %s%v from nodeID %q, version %q", watchID,
				request.TypeUrl, request.ResourceNames, nodeID, request.VersionInfo)
		}
		info.mu.Lock()
		info.watches[watchID] = ResponseWatch{Request: request, Response: value}
		info.mu.Unlock()
		return value, cache.cancelWatch(nodeID, watchID)
	}

	// otherwise, the watch may be responded immediately
	cache.respond(request, value, snapshot.GetResources(request.TypeUrl).Items, version)

	return value, func() {
		close(value)
	}
}

func (cache *snapshotCache) nextWatchID() int64 {
	return atomic.AddInt64(&cache.watchCount, 1)
}

// cancellation function for cleaning stale watches
func (cache *snapshotCache) cancelWatch(nodeID string, watchID int64) func() {
	return func() {
		// uses the cache mutex
		cache.mu.Lock()
		defer cache.mu.Unlock()
		if info, ok := cache.status[nodeID]; ok {
			info.mu.Lock()
			if watch, ok := info.watches[watchID]; ok {
				close(watch.Response)
			}
			delete(info.watches, watchID)
			info.mu.Unlock()
		}
	}
}

// Respond to a watch with the snapshot value. The value channel should have capacity not to block.
// TODO(kuat) do not respond always, see issue https://github.com/envoyproxy/go-control-plane/issues/46
func (cache *snapshotCache) respond(request Request, value chan Response, resources map[string]Resource, version string) {
	// for ADS, the request names must match the snapshot names
	// if they do not, then the watch is never responded, and it is expected that envoy makes another request
	if len(request.ResourceNames) != 0 && cache.ads {
		if err := Superset(nameSet(request.ResourceNames), resources); err != nil {
			if cache.log != nil {
				cache.log.Debugf("ADS mode: not responding to request: %v", err)
			}
			return
		}
	}
	if cache.log != nil {
		cache.log.Debugf("respond %s; resources: %v. version %q with version %q",
			request.TypeUrl, request.ResourceNames, request.VersionInfo, version)
	}
	stats.RecordWithTags(context.TODO(), []tag.Mutator{
		tag.Insert(KeyType, request.GetTypeUrl()),
	}, MResponses.M(1))
	value <- createResponse(request, resources, version)
}

func createResponse(request Request, resources map[string]Resource, version string) Response {
	filtered := make([]Resource, 0, len(resources))

	// Reply only with the requested resources. Envoy may ask each resource
	// individually in a separate stream. It is ok to reply with the same version
	// on separate streams since requests do not share their response versions.
	if len(request.ResourceNames) != 0 {
		set := nameSet(request.ResourceNames)
		for name, resource := range resources {
			if set[name] {
				filtered = append(filtered, resource)
			}
		}
	} else {
		for _, resource := range resources {
			filtered = append(filtered, resource)
		}
	}

	return Response{
		Request:   request,
		Version:   version,
		Resources: filtered,
	}
}

// Fetch implements the cache fetch function.
// Fetch is called on multiple streams, so responding to individual names with the same version works.
func (cache *snapshotCache) Fetch(ctx context.Context, request Request) (*Response, error) {
	nodeID := cache.hash.ID(request.Node)

	cache.mu.RLock()
	defer cache.mu.RUnlock()

	if snapshot, exists := cache.snapshots[nodeID]; exists {
		// Respond only if the request version is distinct from the current snapshot state.
		// It might be beneficial to hold the request since Envoy will re-attempt the refresh.
		version := snapshot.GetResources(request.TypeUrl).Version
		if request.VersionInfo == version {
			return nil, VersionUpToDateError
		}

		resources := snapshot.GetResources(request.TypeUrl).Items
		out := createResponse(request, resources, version)
		return &out, nil
	}

	return nil, fmt.Errorf("missing snapshot for %q", nodeID)
}

// GetStatusInfo retrieves the status info for the node.
func (cache *snapshotCache) GetStatusInfo(node string) StatusInfo {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	info, exists := cache.status[node]
	if !exists {
		return nil
	}

	return info
}

// GetStatusKeys retrieves all node IDs in the status map.
func (cache *snapshotCache) GetStatusKeys() []string {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	var out []string
	for id := range cache.status {
		out = append(out, id)
	}

	return out
}
