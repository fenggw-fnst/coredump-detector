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

package v1alpha1

import (
	"log"

	"github.com/WanLinghao/coredump-detector/pkg/apis/coredump"
	"github.com/WanLinghao/coredump-detector/pkg/k8sclient"
	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	coredumpapi "github.com/WanLinghao/api/coredump"
)

var (
	coredumpCoredumpEndpointStorage = builders.NewApiResource( // Resource status endpoint
		coredump.InternalCoredumpEndpoint,
		CoredumpEndpointSchemeFns{},
		func() runtime.Object { return &coredumpapi.CoredumpEndpoint{} },     // Register versioned resource
		func() runtime.Object { return &coredumpapi.CoredumpEndpointList{} }, // Register versioned resource list
		&coredump.CoredumpEndpointStrategy{builders.StorageStrategySingleton, k8sclient.GetClient().CoreV1()},
	)
	ApiVersion = builders.NewApiVersion("coredump.fujitsu.com", "v1alpha1").WithResources(
		coredumpCoredumpEndpointStorage,
		builders.NewApiResource( // Resource status endpoint
			coredump.InternalCoredumpEndpointStatus,
			CoredumpEndpointSchemeFns{},
			func() runtime.Object { return &coredumpapi.CoredumpEndpoint{} },     // Register versioned resource
			func() runtime.Object { return &coredumpapi.CoredumpEndpointList{} }, // Register versioned resource list
			&coredump.CoredumpEndpointStatusStrategy{builders.StatusStorageStrategySingleton},
		), builders.NewApiResourceWithStorage(
			coredump.InternalCoredumpEndpointDumpREST,
			builders.SchemeFnsSingleton,
			func() runtime.Object { return &coredumpapi.CoredumpEndpointDump{} }, // Register versioned resource
			nil,
			func(generic.RESTOptionsGetter) rest.Storage {
				return &coredump.CoredumpEndpointDumpREST{coredump.NewCoredumpEndpointRegistry(coredumpCoredumpEndpointStorage)}
			},
		),
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

//
// CoredumpEndpoint Functions and Structs
//
// +k8s:deepcopy-gen=false
type CoredumpEndpointSchemeFns struct {
	builders.DefaultSchemeFns
}

// DefaultingFunction sets default CoredumpEndpoint field values
func (CoredumpEndpointSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*coredumpapi.CoredumpEndpoint)
	// set default field values here
	log.Printf("Defaulting fields for CoredumpEndpoint %s\n", obj.Name)
}
