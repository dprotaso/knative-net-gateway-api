/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// ParentReferenceApplyConfiguration represents a declarative configuration of the ParentReference type for use
// with apply.
type ParentReferenceApplyConfiguration struct {
	Group       *v1.Group       `json:"group,omitempty"`
	Kind        *v1.Kind        `json:"kind,omitempty"`
	Namespace   *v1.Namespace   `json:"namespace,omitempty"`
	Name        *v1.ObjectName  `json:"name,omitempty"`
	SectionName *v1.SectionName `json:"sectionName,omitempty"`
	Port        *v1.PortNumber  `json:"port,omitempty"`
}

// ParentReferenceApplyConfiguration constructs a declarative configuration of the ParentReference type for use with
// apply.
func ParentReference() *ParentReferenceApplyConfiguration {
	return &ParentReferenceApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithGroup(value v1.Group) *ParentReferenceApplyConfiguration {
	b.Group = &value
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithKind(value v1.Kind) *ParentReferenceApplyConfiguration {
	b.Kind = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithNamespace(value v1.Namespace) *ParentReferenceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithName(value v1.ObjectName) *ParentReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithSectionName sets the SectionName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SectionName field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithSectionName(value v1.SectionName) *ParentReferenceApplyConfiguration {
	b.SectionName = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *ParentReferenceApplyConfiguration) WithPort(value v1.PortNumber) *ParentReferenceApplyConfiguration {
	b.Port = &value
	return b
}
