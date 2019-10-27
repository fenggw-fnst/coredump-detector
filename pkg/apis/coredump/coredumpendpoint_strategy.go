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

package coredump

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// Validate checks that an instance of CoredumpEndpoint is well formed
func (c *CoredumpEndpointStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	ce := obj.(*CoredumpEndpoint)
	log.Printf("Validating fields for CoredumpEndpoint %s/%s\n", ce.Namespace, ce.Name)
	errors := field.ErrorList{}

	if len(ce.Spec.PodUID) == 0 {
		pod, err := c.PodClient.Pods(ce.Namespace).Get(ce.Name, metav1.GetOptions{})
		if err != nil {
			fieldError := field.InternalError(field.NewPath("spec").Child("podUID"), fmt.Errorf("get pod failed: %v", err))
			errors = append(errors, fieldError)
			return errors
		}
		ce.Spec.PodUID = pod.UID
	}

	// perform validation here and add to errors using field.Invalid
	return errors
}
