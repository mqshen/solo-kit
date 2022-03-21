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
	"encoding/json"
	"errors"
	"fmt"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/golang/protobuf/ptypes"
	"time"

	"github.com/golang/protobuf/proto"
)

type TypedResources map[string]Resources

// compile-time assertion on type
var _ Snapshot = &GenericSnapshot{}

type GenericSnapshot struct {
	typedResources TypedResources
}

func (s *GenericSnapshot) Deserialize(bytes []byte) {
	//newTyped := &TypedResources{}

	newTypedResourcesMap := map[string]interface{}{}
	err := json.Unmarshal(bytes, &newTypedResourcesMap)
	if err != nil {
		panic(err)
	}

	newTypedResourcesMap2 := map[string]Resources{}
	_ = json.Unmarshal(bytes, &newTypedResourcesMap2)
	// ignore error, all other fields worked
	//if err != nil {
	//	panic(err)
	//}

	// go through and set resource for each!
	for typeurl, resources := range newTypedResourcesMap2 {
		for b, _ := range resources.Items {
			switch typeurl {
			case "type.googleapis.com/envoy.config.cluster.v3.Cluster":
				er := NewEnvoyResource(makeCluster("todo"))
				resources.Items[b] = er
			}
		}
	}


	//newResources := newTypedResourcesMap["type.googleapis.com/envoy.config.cluster.v3.Cluster"]
	//err := json.Unmarshal(bytes, &newResources)
	//if err != nil {
	//	panic(err)
	//}

	c := makeCluster("test")
	resourcesCopy := Resources{
		Version: "",
		Items:   make(map[string]Resource, 1),
	}

	//into := c.rtype.EmptyProto()
	//err = proto.Unmarshal(r.Value, into)
	//if err != nil {
	//	break
	//}
	//resource := c.rtype.ProtoToResource(into)
	//resources.Items[resource.Self().Name] = resource

    er := NewEnvoyResource(c)
	resourcesCopy.Items["test"] = er//proto.Clone(c).(Resource)
	fmt.Printf("envoy resource name %v\n", er.Name())
	fmt.Printf("envoy refs %v\n", er.References())
	s.typedResources = TypedResources{}
	s.typedResources["test"] = resourcesCopy//proto.Clone(v.ResourceProto()).(Resource)
	//panic("implement me")
	fmt.Printf("%v\n", s.typedResources)
	panic("exit")
}

const (
	ListenerName = "listener_0"
	UpstreamHost = "127.0.0.1"
)

func makeCluster(clusterName string) *cluster.Cluster {
	return &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       ptypes.DurationProto(5 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
	}
}

func makeEndpoint(clusterName string) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  UpstreamHost,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: uint32(8080),
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func (s *GenericSnapshot) Serialize() []byte {

	// let's start with a single resource

	// convert to json then write out
	b, err := json.Marshal(s.typedResources)
	if err != nil {
		panic(err)
	}
	fmt.Printf("output json: %v\n", string(b))
	return b

	typedResourcesCopy := make(TypedResources)
	for typeName, resources := range s.typedResources {
		resourcesCopy := Resources{
			Version: resources.Version,
			Items:   make(map[string]Resource, len(resources.Items)),
		}
		for k, v := range resources.Items {

			v.ResourceProto().String()

			resourcesCopy.Items[k] = proto.Clone(v.ResourceProto()).(Resource)
		}
		typedResourcesCopy[typeName] = resourcesCopy
	}
	return nil
	//return &GenericSnapshot{typedResources: typedResourcesCopy}
}

// Combine snapshots with distinct types to one.
func (s *GenericSnapshot) Combine(a *GenericSnapshot) (*GenericSnapshot, error) {
	if s.typedResources == nil {
		return a, nil
	} else if a.typedResources == nil {
		return s, nil
	}
	combined := TypedResources{}
	for k, v := range s.typedResources {
		combined[k] = v
	}
	for k, v := range a.typedResources {
		if _, ok := combined[k]; ok {
			return nil, errors.New("overlapping types found")
		}
		combined[k] = v
	}
	return NewGenericSnapshot(combined), nil
}

// Combine snapshots with distinct types to one.
func (s *GenericSnapshot) Merge(newSnap *GenericSnapshot) (*GenericSnapshot, error) {
	if s.typedResources == nil {
		return newSnap, nil
	}
	combined := TypedResources{}
	for k, v := range s.typedResources {
		combined[k] = v
	}
	for k, v := range newSnap.typedResources {
		combined[k] = v
	}
	return NewGenericSnapshot(combined), nil
}

// NewSnapshot creates a snapshot from response types and a version.
func NewGenericSnapshot(resources TypedResources) *GenericSnapshot {
	return &GenericSnapshot{
		typedResources: resources,
	}
}
func NewEasyGenericSnapshot(version string, resources ...[]Resource) *GenericSnapshot {
	t := TypedResources{}

	for _, resources := range resources {
		for _, resource := range resources {
			r := t[resource.Self().Type]
			if r.Items == nil {
				r.Items = make(map[string]Resource)
				r.Version = version
			}
			r.Items[resource.Self().Name] = resource
			t[resource.Self().Type] = r
		}
	}

	return &GenericSnapshot{
		typedResources: t,
	}
}

func (s *GenericSnapshot) Consistent() error {
	if s == nil {
		return nil
	}

	var required []XdsResourceReference

	for _, resources := range s.typedResources {
		for _, resource := range resources.Items {
			required = append(required, resource.References()...)
		}
	}

	for _, ref := range required {
		if resources, ok := s.typedResources[ref.Type]; ok {
			if _, ok := resources.Items[ref.Name]; ok {
				return fmt.Errorf("required resource not in snapshot: %s %s", ref.Type, ref.Name)
			}
		} else {
			return fmt.Errorf("required resource not in snapshot: %s %s", ref.Type, ref.Name)
		}
	}

	return nil
}

// GetResources selects snapshot resources by type.
func (s *GenericSnapshot) GetResources(typ string) Resources {
	if s == nil {
		return Resources{}
	}

	return s.typedResources[typ]
}

func (s *GenericSnapshot) Clone() Snapshot {
	typedResourcesCopy := make(TypedResources)
	for typeName, resources := range s.typedResources {
		resourcesCopy := Resources{
			Version: resources.Version,
			Items:   make(map[string]Resource, len(resources.Items)),
		}
		for k, v := range resources.Items {
			resourcesCopy.Items[k] = proto.Clone(v.ResourceProto()).(Resource)
		}
		typedResourcesCopy[typeName] = resourcesCopy
	}
	return &GenericSnapshot{typedResources: typedResourcesCopy}
}
