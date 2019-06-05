/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package coredump

import (
	"context"
	"fmt"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apiserver/pkg/registry/rest"
	coreclient "k8s.io/client-go/kubernetes/typed/core/v1"
)

var (
	InternalCoredumpEndpoint = builders.NewInternalResource(
		"coredumpendpoints",
		"CoredumpEndpoint",
		func() runtime.Object { return &CoredumpEndpoint{} },
		func() runtime.Object { return &CoredumpEndpointList{} },
	)
	InternalCoredumpEndpointStatus = builders.NewInternalResourceStatus(
		"coredumpendpoints",
		"CoredumpEndpointStatus",
		func() runtime.Object { return &CoredumpEndpoint{} },
		func() runtime.Object { return &CoredumpEndpointList{} },
	)
	InternalCoredumpEndpointDumpREST = builders.NewInternalSubresource(
		"coredumpendpoints", "CoredumpEndpointDump", "dump",
		func() runtime.Object { return &CoredumpEndpointDump{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("coredump.fujitsu.com").WithKinds(
		InternalCoredumpEndpoint,
		InternalCoredumpEndpointStatus,
		InternalCoredumpEndpointDumpREST,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CoredumpEndpoint struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   CoredumpEndpointSpec
	Status CoredumpEndpointStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CoredumpEndpointDump struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

type CoredumpEndpointStatus struct {
}

type CoredumpEndpointSpec struct {
	PodUID types.UID
}

//
// CoredumpEndpoint Functions and Structs
//
// +k8s:deepcopy-gen=false
type CoredumpEndpointStrategy struct {
	builders.DefaultStorageStrategy
	PodClient coreclient.PodsGetter
}

// +k8s:deepcopy-gen=false
type CoredumpEndpointStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CoredumpEndpointList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []CoredumpEndpoint
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CoredumpEndpointDumpList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []CoredumpEndpointDump
}

func (CoredumpEndpoint) NewStatus() interface{} {
	return CoredumpEndpointStatus{}
}

func (pc *CoredumpEndpoint) GetStatus() interface{} {
	return pc.Status
}

func (pc *CoredumpEndpoint) SetStatus(s interface{}) {
	pc.Status = s.(CoredumpEndpointStatus)
}

func (pc *CoredumpEndpoint) GetSpec() interface{} {
	return pc.Spec
}

func (pc *CoredumpEndpoint) SetSpec(s interface{}) {
	pc.Spec = s.(CoredumpEndpointSpec)
}

func (pc *CoredumpEndpoint) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *CoredumpEndpoint) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc CoredumpEndpoint) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store CoredumpEndpoint.
// +k8s:deepcopy-gen=false
type CoredumpEndpointRegistry interface {
	ListCoredumpEndpoints(ctx context.Context, options *internalversion.ListOptions) (*CoredumpEndpointList, error)
	GetCoredumpEndpoint(ctx context.Context, id string, options *metav1.GetOptions) (*CoredumpEndpoint, error)
	CreateCoredumpEndpoint(ctx context.Context, id *CoredumpEndpoint) (*CoredumpEndpoint, error)
	UpdateCoredumpEndpoint(ctx context.Context, id *CoredumpEndpoint) (*CoredumpEndpoint, error)
	DeleteCoredumpEndpoint(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewCoredumpEndpointRegistry(sp builders.StandardStorageProvider) CoredumpEndpointRegistry {
	return &storageCoredumpEndpoint{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageCoredumpEndpoint struct {
	builders.StandardStorageProvider
}

func (s *storageCoredumpEndpoint) ListCoredumpEndpoints(ctx context.Context, options *internalversion.ListOptions) (*CoredumpEndpointList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*CoredumpEndpointList), err
}

func (s *storageCoredumpEndpoint) GetCoredumpEndpoint(ctx context.Context, id string, options *metav1.GetOptions) (*CoredumpEndpoint, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*CoredumpEndpoint), nil
}

func (s *storageCoredumpEndpoint) CreateCoredumpEndpoint(ctx context.Context, object *CoredumpEndpoint) (*CoredumpEndpoint, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*CoredumpEndpoint), nil
}

func (s *storageCoredumpEndpoint) UpdateCoredumpEndpoint(ctx context.Context, object *CoredumpEndpoint) (*CoredumpEndpoint, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*CoredumpEndpoint), nil
}

func (s *storageCoredumpEndpoint) DeleteCoredumpEndpoint(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}
