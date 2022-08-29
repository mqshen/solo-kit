package consul

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/yaml"
)

type ResourceClient struct {
	consul       *api.Client
	root         string
	resourceType resources.VersionedResource
}

var jsonpbMarshaler = &jsonpb.Marshaler{OrigName: false}

func NewResourceClient(client *api.Client, rootKey string, resourceType resources.VersionedResource) *ResourceClient {
	return &ResourceClient{
		consul:       client,
		root:         rootKey,
		resourceType: resourceType,
	}
}

var _ clients.ResourceClient = &ResourceClient{}

func (rc *ResourceClient) Kind() string {
	return resources.Kind(rc.resourceType)
}

func (rc *ResourceClient) NewResource() resources.Resource {
	return resources.Clone(rc.resourceType)
}

func (rc *ResourceClient) Register() error {
	return nil
}

func (rc *ResourceClient) Read(namespace, name string, opts clients.ReadOpts) (resources.Resource, error) {
	if err := resources.ValidateName(name); err != nil {
		return nil, errors.Wrapf(err, "validation error")
	}
	opts = opts.WithDefaults()
	key := rc.resourceKey(namespace, name)

	kvPair, _, err := rc.consul.KV().Get(key, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "performing consul KV get")
	}
	if kvPair == nil {
		return nil, errors.NewNotExistErr(namespace, name)
	}
	resource := rc.NewResource()
	if err := protoutils.UnmarshalYAML(kvPair.Value, resource); err != nil {
		return nil, errors.Wrapf(err, "reading KV into %v", rc.Kind())
	}
	resources.UpdateMetadata(resource, func(meta *core.Metadata) {
		meta.ResourceVersion = fmt.Sprintf("%v", kvPair.ModifyIndex)
	})
	return resource, nil
}

func printStatus(key string, resource resources.Resource, modifyIndex uint64, flag string) {
	if key == "gloo/gloo.solo.io/v1/Proxy/ifp/ifp3-gateway-proxy" {
		if res, ok := resource.(resources.InputResource); ok {
			out := &bytes.Buffer{}
			if err := jsonpbMarshaler.Marshal(out, res.GetStatus()); err != nil {
				fmt.Printf("failed to marshal proto to bytes")
			}

			data, err := yaml.JSONToYAML(out.Bytes())
			if err != nil {
				fmt.Printf("failed to marshal proto to bytes")
			}
			fmt.Printf("%s key: %s modifyIndex: %d index: %s status: %s", flag, key, modifyIndex, res.GetMetadata().ResourceVersion, string(data))

		}
	}
}

func (rc *ResourceClient) Write(resource resources.Resource, opts clients.WriteOpts) (resources.Resource, error) {
	opts = opts.WithDefaults()
	if err := resources.Validate(resource); err != nil {
		return nil, errors.Wrapf(err, "validation error")
	}
	meta := resource.GetMetadata()
	if meta.Namespace == "" {
		return nil, errors.Errorf("namespace cannot be empty for consul-backed resources")
	}
	key := rc.resourceKey(meta.Namespace, meta.Name)

	original, err := rc.Read(meta.Namespace, meta.Name, clients.ReadOpts{
		Ctx: opts.Ctx,
	})
	if original != nil && err == nil {
		if !opts.OverwriteExisting {
			return nil, errors.NewExistErr(meta)
		}
		if meta.ResourceVersion != original.GetMetadata().ResourceVersion {
			return nil, errors.NewResourceVersionErr(meta.Namespace, meta.Name, meta.ResourceVersion, original.GetMetadata().ResourceVersion)
		}
	}

	// mutate and return clone
	clone := resources.Clone(resource)
	clone.SetMetadata(meta)

	data, err := protoutils.MarshalYAML(clone)
	if err != nil {
		panic(errors.Wrapf(err, "internal err: failed to marshal resource"))
	}
	var modifyIndex uint64
	if meta.GetResourceVersion() != "" {
		i, err := strconv.Atoi(meta.GetResourceVersion())
		if err != nil {
			return nil, errors.Wrapf(err, "invalid resource version: %v (must be int)", meta.GetResourceVersion())
		}
		modifyIndex = uint64(i)
	}

	// printStatus(key, original, modifyIndex, "original")
	kvPair := &api.KVPair{
		Key:         key,
		Value:       data,
		ModifyIndex: modifyIndex,
	}
	// if key == "gloo/gloo.solo.io/v1/Proxy/ifp/ifp3-gateway-proxy" {
	// 	fmt.Printf("write key %s with modifyIndex: %d\n", key, modifyIndex)
	// }
	if success, _, err := rc.consul.KV().CAS(kvPair, nil); err != nil {
		return nil, errors.Wrapf(err, "writing to KV")
	} else if !success {
		currentResource, err := rc.Read(meta.Namespace, meta.Name, clients.ReadOpts{Ctx: opts.Ctx})
		if err != nil {
			return nil, errors.Wrapf(err, "reading KV into %v", rc.Kind())
		}

		if res, ok := resource.(resources.InputResource); ok {
			if currentRes, ok := currentResource.(resources.InputResource); ok {
				if res.GetStatus() == nil || currentRes.GetStatus() == nil ||
					res.GetStatus().ReportedBy != currentRes.GetStatus().ReportedBy {
					if original != nil && original.GetMetadata() != nil {
						return nil, errors.NewResourceVersionErr(meta.Namespace, meta.Name, meta.ResourceVersion, original.GetMetadata().ResourceVersion)
					}
				}
			}
		}
		// printStatus(key, currentResource, modifyIndex, "conflict")

		if currentResource != nil && currentResource.GetMetadata() != nil {
			return nil, errors.Errorf("writing to KV failed, lastModifyIndex: %d  currentIndex: %s unknown error", modifyIndex, currentResource.GetMetadata().ResourceVersion)
		}
		return nil, errors.Errorf("writing to KV failed, lastModifyIndex: %d unknown error", modifyIndex)
	}
	// return a read object to update the modify index

	result, err := rc.Read(meta.Namespace, meta.Name, clients.ReadOpts{Ctx: opts.Ctx})
	// printStatus(key, result, modifyIndex, "")
	return result, err
}

func (rc *ResourceClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	if namespace == "" {
		return errors.Errorf("namespace cannot be empty for consul-backed resources")
	}
	key := rc.resourceKey(namespace, name)
	if !opts.IgnoreNotExist {
		if _, err := rc.Read(namespace, name, clients.ReadOpts{Ctx: opts.Ctx}); err != nil {
			return err
		}
	}
	_, err := rc.consul.KV().Delete(key, nil)
	if err != nil {
		return errors.Wrapf(err, "deleting resource %v", name)
	}
	return nil
}

func (rc *ResourceClient) List(namespace string, opts clients.ListOpts) (resources.ResourceList, error) {
	opts = opts.WithDefaults()

	resourceDir := rc.resourceDir(namespace)
	kvPairs, _, err := rc.consul.KV().List(resourceDir, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "reading namespace root")
	}

	var resourceList resources.ResourceList
	for _, kvPair := range kvPairs {
		resource := rc.NewResource()
		if err := protoutils.UnmarshalYAML(kvPair.Value, resource); err != nil {
			return nil, errors.Wrapf(err, "reading KV into %v", rc.Kind())
		}
		resources.UpdateMetadata(resource, func(meta *core.Metadata) {
			meta.ResourceVersion = fmt.Sprintf("%v", kvPair.ModifyIndex)
		})
		if labels.SelectorFromSet(opts.Selector).Matches(labels.Set(resource.GetMetadata().Labels)) {
			resourceList = append(resourceList, resource)
		}
	}

	sort.SliceStable(resourceList, func(i, j int) bool {
		return resourceList[i].GetMetadata().Name < resourceList[j].GetMetadata().Name
	})

	return resourceList, nil
}

func (rc *ResourceClient) Watch(namespace string, opts clients.WatchOpts) (<-chan resources.ResourceList, <-chan error, error) {
	opts = opts.WithDefaults()
	var lastIndex uint64
	resourceDir := rc.resourceDir(namespace)
	resourcesChan := make(chan resources.ResourceList)
	errs := make(chan error)
	go func() {
		// watch should open up with an initial read
		list, err := rc.List(namespace, clients.ListOpts{
			Ctx:      opts.Ctx,
			Selector: opts.Selector,
		})
		if err != nil {
			errs <- err
			return
		}
		resourcesChan <- list
	}()
	updatedResourceList := func() (resources.ResourceList, error) {
		kvPairs, meta, err := rc.consul.KV().List(resourceDir,
			&api.QueryOptions{
				RequireConsistent: true,
				WaitIndex:         lastIndex,
				WaitTime:          opts.RefreshRate,
			})
		if err != nil {
			return nil, errors.Wrapf(err, "getting kv-pairs list")
		}
		// no change since last poll
		if lastIndex == meta.LastIndex {
			return nil, nil
		}
		var resourceList resources.ResourceList
		for _, kvPair := range kvPairs {
			resource := rc.NewResource()
			if err := protoutils.UnmarshalYAML(kvPair.Value, resource); err != nil {
				return nil, errors.Wrapf(err, "reading KV into %v", rc.Kind())
			}
			resources.UpdateMetadata(resource, func(meta *core.Metadata) {
				meta.ResourceVersion = fmt.Sprintf("%v", kvPair.ModifyIndex)
			})
			if labels.SelectorFromSet(opts.Selector).Matches(labels.Set(resource.GetMetadata().Labels)) {
				resourceList = append(resourceList, resource)
			}
		}

		sort.SliceStable(resourceList, func(i, j int) bool {
			return resourceList[i].GetMetadata().Name < resourceList[j].GetMetadata().Name
		})

		// update index
		lastIndex = meta.LastIndex
		return resourceList, nil
	}

	go func() {
		for {
			select {
			default:
				list, err := updatedResourceList()
				if err != nil {
					errs <- err
				}
				if list != nil {
					resourcesChan <- list
				}
			case <-opts.Ctx.Done():
				close(resourcesChan)
				close(errs)
				return
			}
		}
	}()

	return resourcesChan, errs, nil
}

// works with "" (NamespaceAll)
func (rc *ResourceClient) resourceDir(namespace string) string {
	return strings.Join([]string{
		rc.root,
		rc.resourceType.GroupVersionKind().Group,
		rc.resourceType.GroupVersionKind().Version,
		rc.resourceType.GroupVersionKind().Kind,
		namespace,
	}, "/")
}

func (rc *ResourceClient) resourceKey(namespace, name string) string {
	return strings.Join([]string{
		rc.resourceDir(namespace),
		name}, "/")
}
