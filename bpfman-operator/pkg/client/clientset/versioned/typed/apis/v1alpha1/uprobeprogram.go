/*
Copyright 2023 The bpfman Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/bpfman/bpfman/bpfman-operator/apis/v1alpha1"
	scheme "github.com/bpfman/bpfman/bpfman-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UprobeProgramsGetter has a method to return a UprobeProgramInterface.
// A group's client should implement this interface.
type UprobeProgramsGetter interface {
	UprobePrograms() UprobeProgramInterface
}

// UprobeProgramInterface has methods to work with UprobeProgram resources.
type UprobeProgramInterface interface {
	Create(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.CreateOptions) (*v1alpha1.UprobeProgram, error)
	Update(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.UpdateOptions) (*v1alpha1.UprobeProgram, error)
	UpdateStatus(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.UpdateOptions) (*v1alpha1.UprobeProgram, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.UprobeProgram, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.UprobeProgramList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UprobeProgram, err error)
	UprobeProgramExpansion
}

// uprobePrograms implements UprobeProgramInterface
type uprobePrograms struct {
	client rest.Interface
}

// newUprobePrograms returns a UprobePrograms
func newUprobePrograms(c *BpfmanV1alpha1Client) *uprobePrograms {
	return &uprobePrograms{
		client: c.RESTClient(),
	}
}

// Get takes name of the uprobeProgram, and returns the corresponding uprobeProgram object, and an error if there is any.
func (c *uprobePrograms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UprobeProgram, err error) {
	result = &v1alpha1.UprobeProgram{}
	err = c.client.Get().
		Resource("uprobeprograms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UprobePrograms that match those selectors.
func (c *uprobePrograms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UprobeProgramList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UprobeProgramList{}
	err = c.client.Get().
		Resource("uprobeprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested uprobePrograms.
func (c *uprobePrograms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("uprobeprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a uprobeProgram and creates it.  Returns the server's representation of the uprobeProgram, and an error, if there is any.
func (c *uprobePrograms) Create(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.CreateOptions) (result *v1alpha1.UprobeProgram, err error) {
	result = &v1alpha1.UprobeProgram{}
	err = c.client.Post().
		Resource("uprobeprograms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uprobeProgram).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a uprobeProgram and updates it. Returns the server's representation of the uprobeProgram, and an error, if there is any.
func (c *uprobePrograms) Update(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.UpdateOptions) (result *v1alpha1.UprobeProgram, err error) {
	result = &v1alpha1.UprobeProgram{}
	err = c.client.Put().
		Resource("uprobeprograms").
		Name(uprobeProgram.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uprobeProgram).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *uprobePrograms) UpdateStatus(ctx context.Context, uprobeProgram *v1alpha1.UprobeProgram, opts v1.UpdateOptions) (result *v1alpha1.UprobeProgram, err error) {
	result = &v1alpha1.UprobeProgram{}
	err = c.client.Put().
		Resource("uprobeprograms").
		Name(uprobeProgram.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(uprobeProgram).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the uprobeProgram and deletes it. Returns an error if one occurs.
func (c *uprobePrograms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("uprobeprograms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *uprobePrograms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("uprobeprograms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched uprobeProgram.
func (c *uprobePrograms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UprobeProgram, err error) {
	result = &v1alpha1.UprobeProgram{}
	err = c.client.Patch(pt).
		Resource("uprobeprograms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
