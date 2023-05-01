//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgram) DeepCopyInto(out *BpfProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgram.
func (in *BpfProgram) DeepCopy() *BpfProgram {
	if in == nil {
		return nil
	}
	out := new(BpfProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BpfProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramCommon) DeepCopyInto(out *BpfProgramCommon) {
	*out = *in
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	in.ByteCode.DeepCopyInto(&out.ByteCode)
	if in.GlobalData != nil {
		in, out := &in.GlobalData, &out.GlobalData
		*out = make(map[string][]byte, len(*in))
		for key, val := range *in {
			var outVal []byte
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]byte, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramCommon.
func (in *BpfProgramCommon) DeepCopy() *BpfProgramCommon {
	if in == nil {
		return nil
	}
	out := new(BpfProgramCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramList) DeepCopyInto(out *BpfProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BpfProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramList.
func (in *BpfProgramList) DeepCopy() *BpfProgramList {
	if in == nil {
		return nil
	}
	out := new(BpfProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BpfProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramSpec) DeepCopyInto(out *BpfProgramSpec) {
	*out = *in
	if in.Programs != nil {
		in, out := &in.Programs, &out.Programs
		*out = make(map[string]map[string]string, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramSpec.
func (in *BpfProgramSpec) DeepCopy() *BpfProgramSpec {
	if in == nil {
		return nil
	}
	out := new(BpfProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BpfProgramStatus) DeepCopyInto(out *BpfProgramStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BpfProgramStatus.
func (in *BpfProgramStatus) DeepCopy() *BpfProgramStatus {
	if in == nil {
		return nil
	}
	out := new(BpfProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BytecodeImage) DeepCopyInto(out *BytecodeImage) {
	*out = *in
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(ImagePullSecretSelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BytecodeImage.
func (in *BytecodeImage) DeepCopy() *BytecodeImage {
	if in == nil {
		return nil
	}
	out := new(BytecodeImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BytecodeSelector) DeepCopyInto(out *BytecodeSelector) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(BytecodeImage)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BytecodeSelector.
func (in *BytecodeSelector) DeepCopy() *BytecodeSelector {
	if in == nil {
		return nil
	}
	out := new(BytecodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePullSecretSelector) DeepCopyInto(out *ImagePullSecretSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePullSecretSelector.
func (in *ImagePullSecretSelector) DeepCopy() *ImagePullSecretSelector {
	if in == nil {
		return nil
	}
	out := new(ImagePullSecretSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSelector) DeepCopyInto(out *InterfaceSelector) {
	*out = *in
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.PrimaryNodeInterface != nil {
		in, out := &in.PrimaryNodeInterface, &out.PrimaryNodeInterface
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSelector.
func (in *InterfaceSelector) DeepCopy() *InterfaceSelector {
	if in == nil {
		return nil
	}
	out := new(InterfaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgram) DeepCopyInto(out *TcProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgram.
func (in *TcProgram) DeepCopy() *TcProgram {
	if in == nil {
		return nil
	}
	out := new(TcProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TcProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramList) DeepCopyInto(out *TcProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TcProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramList.
func (in *TcProgramList) DeepCopy() *TcProgramList {
	if in == nil {
		return nil
	}
	out := new(TcProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TcProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramSpec) DeepCopyInto(out *TcProgramSpec) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	in.InterfaceSelector.DeepCopyInto(&out.InterfaceSelector)
	if in.ProceedOn != nil {
		in, out := &in.ProceedOn, &out.ProceedOn
		*out = make([]TcProceedOnValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramSpec.
func (in *TcProgramSpec) DeepCopy() *TcProgramSpec {
	if in == nil {
		return nil
	}
	out := new(TcProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TcProgramStatus) DeepCopyInto(out *TcProgramStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TcProgramStatus.
func (in *TcProgramStatus) DeepCopy() *TcProgramStatus {
	if in == nil {
		return nil
	}
	out := new(TcProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgram) DeepCopyInto(out *TracepointProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgram.
func (in *TracepointProgram) DeepCopy() *TracepointProgram {
	if in == nil {
		return nil
	}
	out := new(TracepointProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracepointProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramList) DeepCopyInto(out *TracepointProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TracepointProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramList.
func (in *TracepointProgramList) DeepCopy() *TracepointProgramList {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracepointProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramSpec) DeepCopyInto(out *TracepointProgramSpec) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramSpec.
func (in *TracepointProgramSpec) DeepCopy() *TracepointProgramSpec {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracepointProgramStatus) DeepCopyInto(out *TracepointProgramStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracepointProgramStatus.
func (in *TracepointProgramStatus) DeepCopy() *TracepointProgramStatus {
	if in == nil {
		return nil
	}
	out := new(TracepointProgramStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgram) DeepCopyInto(out *XdpProgram) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgram.
func (in *XdpProgram) DeepCopy() *XdpProgram {
	if in == nil {
		return nil
	}
	out := new(XdpProgram)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XdpProgram) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramList) DeepCopyInto(out *XdpProgramList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]XdpProgram, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramList.
func (in *XdpProgramList) DeepCopy() *XdpProgramList {
	if in == nil {
		return nil
	}
	out := new(XdpProgramList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *XdpProgramList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramSpec) DeepCopyInto(out *XdpProgramSpec) {
	*out = *in
	in.BpfProgramCommon.DeepCopyInto(&out.BpfProgramCommon)
	in.InterfaceSelector.DeepCopyInto(&out.InterfaceSelector)
	if in.ProceedOn != nil {
		in, out := &in.ProceedOn, &out.ProceedOn
		*out = make([]XdpProceedOnValue, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramSpec.
func (in *XdpProgramSpec) DeepCopy() *XdpProgramSpec {
	if in == nil {
		return nil
	}
	out := new(XdpProgramSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XdpProgramStatus) DeepCopyInto(out *XdpProgramStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XdpProgramStatus.
func (in *XdpProgramStatus) DeepCopy() *XdpProgramStatus {
	if in == nil {
		return nil
	}
	out := new(XdpProgramStatus)
	in.DeepCopyInto(out)
	return out
}
