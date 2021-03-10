// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscaling) DeepCopyInto(out *Autoscaling) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.AverageCPUUtilization != nil {
		in, out := &in.AverageCPUUtilization, &out.AverageCPUUtilization
		*out = new(int32)
		**out = **in
	}
	if in.ScaleDownStabilizationWindowSeconds != nil {
		in, out := &in.ScaleDownStabilizationWindowSeconds, &out.ScaleDownStabilizationWindowSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscaling.
func (in *Autoscaling) DeepCopy() *Autoscaling {
	if in == nil {
		return nil
	}
	out := new(Autoscaling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIImageDefinition) DeepCopyInto(out *OCIImageDefinition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIImageDefinition.
func (in *OCIImageDefinition) DeepCopy() *OCIImageDefinition {
	if in == nil {
		return nil
	}
	out := new(OCIImageDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayCluster) DeepCopyInto(out *RayCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayCluster.
func (in *RayCluster) DeepCopy() *RayCluster {
	if in == nil {
		return nil
	}
	out := new(RayCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterHead) DeepCopyInto(out *RayClusterHead) {
	*out = *in
	in.RayClusterNode.DeepCopyInto(&out.RayClusterNode)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterHead.
func (in *RayClusterHead) DeepCopy() *RayClusterHead {
	if in == nil {
		return nil
	}
	out := new(RayClusterHead)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterList) DeepCopyInto(out *RayClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RayCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterList.
func (in *RayClusterList) DeepCopy() *RayClusterList {
	if in == nil {
		return nil
	}
	out := new(RayClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RayClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterNetworkPolicy) DeepCopyInto(out *RayClusterNetworkPolicy) {
	*out = *in
	if in.ClientServerLabels != nil {
		in, out := &in.ClientServerLabels, &out.ClientServerLabels
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.DashboardLabels != nil {
		in, out := &in.DashboardLabels, &out.DashboardLabels
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterNetworkPolicy.
func (in *RayClusterNetworkPolicy) DeepCopy() *RayClusterNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(RayClusterNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterNode) DeepCopyInto(out *RayClusterNode) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterNode.
func (in *RayClusterNode) DeepCopy() *RayClusterNode {
	if in == nil {
		return nil
	}
	out := new(RayClusterNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterSpec) DeepCopyInto(out *RayClusterSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(OCIImageDefinition)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(Autoscaling)
		(*in).DeepCopyInto(*out)
	}
	if in.RedisShardPorts != nil {
		in, out := &in.RedisShardPorts, &out.RedisShardPorts
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.ObjectStoreMemoryBytes != nil {
		in, out := &in.ObjectStoreMemoryBytes, &out.ObjectStoreMemoryBytes
		*out = new(int64)
		**out = **in
	}
	if in.EnableDashboard != nil {
		in, out := &in.EnableDashboard, &out.EnableDashboard
		*out = new(bool)
		**out = **in
	}
	if in.EnableNetworkPolicy != nil {
		in, out := &in.EnableNetworkPolicy, &out.EnableNetworkPolicy
		*out = new(bool)
		**out = **in
	}
	if in.NetworkPolicyClientLabels != nil {
		in, out := &in.NetworkPolicyClientLabels, &out.NetworkPolicyClientLabels
		*out = make([]map[string]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(map[string]string, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	if in.PodSecurityContext != nil {
		in, out := &in.PodSecurityContext, &out.PodSecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Head.DeepCopyInto(&out.Head)
	in.Worker.DeepCopyInto(&out.Worker)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterSpec.
func (in *RayClusterSpec) DeepCopy() *RayClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RayClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterStatus) DeepCopyInto(out *RayClusterStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterStatus.
func (in *RayClusterStatus) DeepCopy() *RayClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RayClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RayClusterWorker) DeepCopyInto(out *RayClusterWorker) {
	*out = *in
	in.RayClusterNode.DeepCopyInto(&out.RayClusterNode)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RayClusterWorker.
func (in *RayClusterWorker) DeepCopy() *RayClusterWorker {
	if in == nil {
		return nil
	}
	out := new(RayClusterWorker)
	in.DeepCopyInto(out)
	return out
}