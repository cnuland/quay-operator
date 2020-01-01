// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureRegistryBackendSource) DeepCopyInto(out *AzureRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureRegistryBackendSource.
func (in *AzureRegistryBackendSource) DeepCopy() *AzureRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(AzureRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clair) DeepCopyInto(out *Clair) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clair.
func (in *Clair) DeepCopy() *Clair {
	if in == nil {
		return nil
	}
	out := new(Clair)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudfrontS3RegistryBackendSource) DeepCopyInto(out *CloudfrontS3RegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudfrontS3RegistryBackendSource.
func (in *CloudfrontS3RegistryBackendSource) DeepCopy() *CloudfrontS3RegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(CloudfrontS3RegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleCloudRegistryBackendSource) DeepCopyInto(out *GoogleCloudRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleCloudRegistryBackendSource.
func (in *GoogleCloudRegistryBackendSource) DeepCopy() *GoogleCloudRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(GoogleCloudRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalRegistryBackendSource) DeepCopyInto(out *LocalRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalRegistryBackendSource.
func (in *LocalRegistryBackendSource) DeepCopy() *LocalRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(LocalRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Quay) DeepCopyInto(out *Quay) {
	*out = *in
	if in.ConfigEnvVars != nil {
		in, out := &in.ConfigEnvVars, &out.ConfigEnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ConfigResources.DeepCopyInto(&out.ConfigResources)
	if in.RepoMirrorEnvVars != nil {
		in, out := &in.RepoMirrorEnvVars, &out.RepoMirrorEnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RepoMirrorResources.DeepCopyInto(&out.RepoMirrorResources)
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(Database)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	if in.ConfigNodePort != nil {
		in, out := &in.ConfigNodePort, &out.ConfigNodePort
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryBackends != nil {
		in, out := &in.RegistryBackends, &out.RegistryBackends
		*out = make([]RegistryBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RegistryStorage != nil {
		in, out := &in.RegistryStorage, &out.RegistryStorage
		*out = new(RegistryStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ConfigFiles != nil {
		in, out := &in.ConfigFiles, &out.ConfigFiles
		*out = make([]QuayConfigFiles, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Quay.
func (in *Quay) DeepCopy() *Quay {
	if in == nil {
		return nil
	}
	out := new(Quay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayConfigFile) DeepCopyInto(out *QuayConfigFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayConfigFile.
func (in *QuayConfigFile) DeepCopy() *QuayConfigFile {
	if in == nil {
		return nil
	}
	out := new(QuayConfigFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayConfigFiles) DeepCopyInto(out *QuayConfigFiles) {
	*out = *in
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]QuayConfigFile, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayConfigFiles.
func (in *QuayConfigFiles) DeepCopy() *QuayConfigFiles {
	if in == nil {
		return nil
	}
	out := new(QuayConfigFiles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystem) DeepCopyInto(out *QuayEcosystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystem.
func (in *QuayEcosystem) DeepCopy() *QuayEcosystem {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuayEcosystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemCondition) DeepCopyInto(out *QuayEcosystemCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemCondition.
func (in *QuayEcosystemCondition) DeepCopy() *QuayEcosystemCondition {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemList) DeepCopyInto(out *QuayEcosystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuayEcosystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemList.
func (in *QuayEcosystemList) DeepCopy() *QuayEcosystemList {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuayEcosystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemSpec) DeepCopyInto(out *QuayEcosystemSpec) {
	*out = *in
	if in.Quay != nil {
		in, out := &in.Quay, &out.Quay
		*out = new(Quay)
		(*in).DeepCopyInto(*out)
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = new(Redis)
		(*in).DeepCopyInto(*out)
	}
	if in.Clair != nil {
		in, out := &in.Clair, &out.Clair
		*out = new(Clair)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemSpec.
func (in *QuayEcosystemSpec) DeepCopy() *QuayEcosystemSpec {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuayEcosystemStatus) DeepCopyInto(out *QuayEcosystemStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]QuayEcosystemCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuayEcosystemStatus.
func (in *QuayEcosystemStatus) DeepCopy() *QuayEcosystemStatus {
	if in == nil {
		return nil
	}
	out := new(QuayEcosystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RADOSRegistryBackendSource) DeepCopyInto(out *RADOSRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RADOSRegistryBackendSource.
func (in *RADOSRegistryBackendSource) DeepCopy() *RADOSRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(RADOSRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RHOCSRegistryBackendSource) DeepCopyInto(out *RHOCSRegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RHOCSRegistryBackendSource.
func (in *RHOCSRegistryBackendSource) DeepCopy() *RHOCSRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(RHOCSRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryBackend) DeepCopyInto(out *RegistryBackend) {
	*out = *in
	in.RegistryBackendSource.DeepCopyInto(&out.RegistryBackendSource)
	if in.ReplicateByDefault != nil {
		in, out := &in.ReplicateByDefault, &out.ReplicateByDefault
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryBackend.
func (in *RegistryBackend) DeepCopy() *RegistryBackend {
	if in == nil {
		return nil
	}
	out := new(RegistryBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryBackendSource) DeepCopyInto(out *RegistryBackendSource) {
	*out = *in
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(LocalRegistryBackendSource)
		**out = **in
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3RegistryBackendSource)
		**out = **in
	}
	if in.GoogleCloud != nil {
		in, out := &in.GoogleCloud, &out.GoogleCloud
		*out = new(GoogleCloudRegistryBackendSource)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureRegistryBackendSource)
		**out = **in
	}
	if in.RADOS != nil {
		in, out := &in.RADOS, &out.RADOS
		*out = new(RADOSRegistryBackendSource)
		**out = **in
	}
	if in.RHOCS != nil {
		in, out := &in.RHOCS, &out.RHOCS
		*out = new(RHOCSRegistryBackendSource)
		**out = **in
	}
	if in.Swift != nil {
		in, out := &in.Swift, &out.Swift
		*out = new(SwiftRegistryBackendSource)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudfrontS3 != nil {
		in, out := &in.CloudfrontS3, &out.CloudfrontS3
		*out = new(CloudfrontS3RegistryBackendSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryBackendSource.
func (in *RegistryBackendSource) DeepCopy() *RegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(RegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStorage) DeepCopyInto(out *RegistryStorage) {
	*out = *in
	if in.PersistentVolumeAccessModes != nil {
		in, out := &in.PersistentVolumeAccessModes, &out.PersistentVolumeAccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStorage.
func (in *RegistryStorage) DeepCopy() *RegistryStorage {
	if in == nil {
		return nil
	}
	out := new(RegistryStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3RegistryBackendSource) DeepCopyInto(out *S3RegistryBackendSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3RegistryBackendSource.
func (in *S3RegistryBackendSource) DeepCopy() *S3RegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(S3RegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftRegistryBackendSource) DeepCopyInto(out *SwiftRegistryBackendSource) {
	*out = *in
	if in.OSOptions != nil {
		in, out := &in.OSOptions, &out.OSOptions
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftRegistryBackendSource.
func (in *SwiftRegistryBackendSource) DeepCopy() *SwiftRegistryBackendSource {
	if in == nil {
		return nil
	}
	out := new(SwiftRegistryBackendSource)
	in.DeepCopyInto(out)
	return out
}
