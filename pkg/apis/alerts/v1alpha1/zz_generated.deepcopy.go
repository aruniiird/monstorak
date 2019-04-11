// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertDetails) DeepCopyInto(out *AlertDetails) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertDetails.
func (in *AlertDetails) DeepCopy() *AlertDetails {
	if in == nil {
		return nil
	}
	out := new(AlertDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardDetails) DeepCopyInto(out *DashboardDetails) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardDetails.
func (in *DashboardDetails) DeepCopy() *DashboardDetails {
	if in == nil {
		return nil
	}
	out := new(DashboardDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAlerts) DeepCopyInto(out *StorageAlerts) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAlerts.
func (in *StorageAlerts) DeepCopy() *StorageAlerts {
	if in == nil {
		return nil
	}
	out := new(StorageAlerts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageAlerts) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAlertsList) DeepCopyInto(out *StorageAlertsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageAlerts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAlertsList.
func (in *StorageAlertsList) DeepCopy() *StorageAlertsList {
	if in == nil {
		return nil
	}
	out := new(StorageAlertsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageAlertsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAlertsSpec) DeepCopyInto(out *StorageAlertsSpec) {
	*out = *in
	in.StorageAlert.DeepCopyInto(&out.StorageAlert)
	in.StorageDashboard.DeepCopyInto(&out.StorageDashboard)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAlertsSpec.
func (in *StorageAlertsSpec) DeepCopy() *StorageAlertsSpec {
	if in == nil {
		return nil
	}
	out := new(StorageAlertsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageAlertsStatus) DeepCopyInto(out *StorageAlertsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageAlertsStatus.
func (in *StorageAlertsStatus) DeepCopy() *StorageAlertsStatus {
	if in == nil {
		return nil
	}
	out := new(StorageAlertsStatus)
	in.DeepCopyInto(out)
	return out
}