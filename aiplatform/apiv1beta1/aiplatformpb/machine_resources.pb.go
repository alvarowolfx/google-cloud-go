// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: google/cloud/aiplatform/v1beta1/machine_resources.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specification of a single machine.
type MachineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The type of the machine.
	//
	// See the [list of machine types supported for
	// prediction](https://cloud.google.com/vertex-ai/docs/predictions/configure-compute#machine-types)
	//
	// See the [list of machine types supported for custom
	// training](https://cloud.google.com/vertex-ai/docs/training/configure-compute#machine-types).
	//
	// For [DeployedModel][google.cloud.aiplatform.v1beta1.DeployedModel] this
	// field is optional, and the default value is `n1-standard-2`. For
	// [BatchPredictionJob][google.cloud.aiplatform.v1beta1.BatchPredictionJob] or
	// as part of [WorkerPoolSpec][google.cloud.aiplatform.v1beta1.WorkerPoolSpec]
	// this field is required.
	MachineType string `protobuf:"bytes,1,opt,name=machine_type,json=machineType,proto3" json:"machine_type,omitempty"`
	// Immutable. The type of accelerator(s) that may be attached to the machine
	// as per
	// [accelerator_count][google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count].
	AcceleratorType AcceleratorType `protobuf:"varint,2,opt,name=accelerator_type,json=acceleratorType,proto3,enum=google.cloud.aiplatform.v1beta1.AcceleratorType" json:"accelerator_type,omitempty"`
	// The number of accelerators to attach to the machine.
	AcceleratorCount int32 `protobuf:"varint,3,opt,name=accelerator_count,json=acceleratorCount,proto3" json:"accelerator_count,omitempty"`
}

func (x *MachineSpec) Reset() {
	*x = MachineSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineSpec) ProtoMessage() {}

func (x *MachineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineSpec.ProtoReflect.Descriptor instead.
func (*MachineSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{0}
}

func (x *MachineSpec) GetMachineType() string {
	if x != nil {
		return x.MachineType
	}
	return ""
}

func (x *MachineSpec) GetAcceleratorType() AcceleratorType {
	if x != nil {
		return x.AcceleratorType
	}
	return AcceleratorType_ACCELERATOR_TYPE_UNSPECIFIED
}

func (x *MachineSpec) GetAcceleratorCount() int32 {
	if x != nil {
		return x.AcceleratorCount
	}
	return 0
}

// A description of resources that are dedicated to a DeployedModel, and
// that need a higher degree of manual configuration.
type DedicatedResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The specification of a single machine used by the
	// prediction.
	MachineSpec *MachineSpec `protobuf:"bytes,1,opt,name=machine_spec,json=machineSpec,proto3" json:"machine_spec,omitempty"`
	// Required. Immutable. The minimum number of machine replicas this
	// DeployedModel will be always deployed on. This value must be greater than
	// or equal to 1.
	//
	// If traffic against the DeployedModel increases, it may dynamically be
	// deployed onto more replicas, and as traffic decreases, some of these extra
	// replicas may be freed.
	MinReplicaCount int32 `protobuf:"varint,2,opt,name=min_replica_count,json=minReplicaCount,proto3" json:"min_replica_count,omitempty"`
	// Immutable. The maximum number of replicas this DeployedModel may be
	// deployed on when the traffic against it increases. If the requested value
	// is too large, the deployment will error, but if deployment succeeds then
	// the ability to scale the model to that many replicas is guaranteed (barring
	// service outages). If traffic against the DeployedModel increases beyond
	// what its replicas at maximum may handle, a portion of the traffic will be
	// dropped. If this value is not provided, will use
	// [min_replica_count][google.cloud.aiplatform.v1beta1.DedicatedResources.min_replica_count]
	// as the default value.
	//
	// The value of this field impacts the charge against Vertex CPU and GPU
	// quotas. Specifically, you will be charged for (max_replica_count *
	// number of cores in the selected machine type) and (max_replica_count *
	// number of GPUs per replica in the selected machine type).
	MaxReplicaCount int32 `protobuf:"varint,3,opt,name=max_replica_count,json=maxReplicaCount,proto3" json:"max_replica_count,omitempty"`
	// Immutable. The metric specifications that overrides a resource
	// utilization metric (CPU utilization, accelerator's duty cycle, and so on)
	// target value (default to 60 if not set). At most one entry is allowed per
	// metric.
	//
	// If
	// [machine_spec.accelerator_count][google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count]
	// is above 0, the autoscaling will be based on both CPU utilization and
	// accelerator's duty cycle metrics and scale up when either metrics exceeds
	// its target value while scale down if both metrics are under their target
	// value. The default target value is 60 for both metrics.
	//
	// If
	// [machine_spec.accelerator_count][google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_count]
	// is 0, the autoscaling will be based on CPU utilization metric only with
	// default target value 60 if not explicitly set.
	//
	// For example, in the case of Online Prediction, if you want to override
	// target CPU utilization to 80, you should set
	// [autoscaling_metric_specs.metric_name][google.cloud.aiplatform.v1beta1.AutoscalingMetricSpec.metric_name]
	// to `aiplatform.googleapis.com/prediction/online/cpu/utilization` and
	// [autoscaling_metric_specs.target][google.cloud.aiplatform.v1beta1.AutoscalingMetricSpec.target]
	// to `80`.
	AutoscalingMetricSpecs []*AutoscalingMetricSpec `protobuf:"bytes,4,rep,name=autoscaling_metric_specs,json=autoscalingMetricSpecs,proto3" json:"autoscaling_metric_specs,omitempty"`
}

func (x *DedicatedResources) Reset() {
	*x = DedicatedResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DedicatedResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DedicatedResources) ProtoMessage() {}

func (x *DedicatedResources) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DedicatedResources.ProtoReflect.Descriptor instead.
func (*DedicatedResources) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{1}
}

func (x *DedicatedResources) GetMachineSpec() *MachineSpec {
	if x != nil {
		return x.MachineSpec
	}
	return nil
}

func (x *DedicatedResources) GetMinReplicaCount() int32 {
	if x != nil {
		return x.MinReplicaCount
	}
	return 0
}

func (x *DedicatedResources) GetMaxReplicaCount() int32 {
	if x != nil {
		return x.MaxReplicaCount
	}
	return 0
}

func (x *DedicatedResources) GetAutoscalingMetricSpecs() []*AutoscalingMetricSpec {
	if x != nil {
		return x.AutoscalingMetricSpecs
	}
	return nil
}

// A description of resources that to large degree are decided by Vertex AI,
// and require only a modest additional configuration.
// Each Model supporting these resources documents its specific guidelines.
type AutomaticResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The minimum number of replicas this DeployedModel will be always
	// deployed on. If traffic against it increases, it may dynamically be
	// deployed onto more replicas up to
	// [max_replica_count][google.cloud.aiplatform.v1beta1.AutomaticResources.max_replica_count],
	// and as traffic decreases, some of these extra replicas may be freed. If the
	// requested value is too large, the deployment will error.
	MinReplicaCount int32 `protobuf:"varint,1,opt,name=min_replica_count,json=minReplicaCount,proto3" json:"min_replica_count,omitempty"`
	// Immutable. The maximum number of replicas this DeployedModel may be
	// deployed on when the traffic against it increases. If the requested value
	// is too large, the deployment will error, but if deployment succeeds then
	// the ability to scale the model to that many replicas is guaranteed (barring
	// service outages). If traffic against the DeployedModel increases beyond
	// what its replicas at maximum may handle, a portion of the traffic will be
	// dropped. If this value is not provided, a no upper bound for scaling under
	// heavy traffic will be assume, though Vertex AI may be unable to scale
	// beyond certain replica number.
	MaxReplicaCount int32 `protobuf:"varint,2,opt,name=max_replica_count,json=maxReplicaCount,proto3" json:"max_replica_count,omitempty"`
}

func (x *AutomaticResources) Reset() {
	*x = AutomaticResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutomaticResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutomaticResources) ProtoMessage() {}

func (x *AutomaticResources) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutomaticResources.ProtoReflect.Descriptor instead.
func (*AutomaticResources) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{2}
}

func (x *AutomaticResources) GetMinReplicaCount() int32 {
	if x != nil {
		return x.MinReplicaCount
	}
	return 0
}

func (x *AutomaticResources) GetMaxReplicaCount() int32 {
	if x != nil {
		return x.MaxReplicaCount
	}
	return 0
}

// A description of resources that are used for performing batch operations, are
// dedicated to a Model, and need manual configuration.
type BatchDedicatedResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The specification of a single machine.
	MachineSpec *MachineSpec `protobuf:"bytes,1,opt,name=machine_spec,json=machineSpec,proto3" json:"machine_spec,omitempty"`
	// Immutable. The number of machine replicas used at the start of the batch
	// operation. If not set, Vertex AI decides starting number, not greater than
	// [max_replica_count][google.cloud.aiplatform.v1beta1.BatchDedicatedResources.max_replica_count]
	StartingReplicaCount int32 `protobuf:"varint,2,opt,name=starting_replica_count,json=startingReplicaCount,proto3" json:"starting_replica_count,omitempty"`
	// Immutable. The maximum number of machine replicas the batch operation may
	// be scaled to. The default value is 10.
	MaxReplicaCount int32 `protobuf:"varint,3,opt,name=max_replica_count,json=maxReplicaCount,proto3" json:"max_replica_count,omitempty"`
}

func (x *BatchDedicatedResources) Reset() {
	*x = BatchDedicatedResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDedicatedResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDedicatedResources) ProtoMessage() {}

func (x *BatchDedicatedResources) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDedicatedResources.ProtoReflect.Descriptor instead.
func (*BatchDedicatedResources) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{3}
}

func (x *BatchDedicatedResources) GetMachineSpec() *MachineSpec {
	if x != nil {
		return x.MachineSpec
	}
	return nil
}

func (x *BatchDedicatedResources) GetStartingReplicaCount() int32 {
	if x != nil {
		return x.StartingReplicaCount
	}
	return 0
}

func (x *BatchDedicatedResources) GetMaxReplicaCount() int32 {
	if x != nil {
		return x.MaxReplicaCount
	}
	return 0
}

// Statistics information about resource consumption.
type ResourcesConsumed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The number of replica hours used. Note that many replicas may
	// run in parallel, and additionally any given work may be queued for some
	// time. Therefore this value is not strictly related to wall time.
	ReplicaHours float64 `protobuf:"fixed64,1,opt,name=replica_hours,json=replicaHours,proto3" json:"replica_hours,omitempty"`
}

func (x *ResourcesConsumed) Reset() {
	*x = ResourcesConsumed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourcesConsumed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcesConsumed) ProtoMessage() {}

func (x *ResourcesConsumed) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcesConsumed.ProtoReflect.Descriptor instead.
func (*ResourcesConsumed) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{4}
}

func (x *ResourcesConsumed) GetReplicaHours() float64 {
	if x != nil {
		return x.ReplicaHours
	}
	return 0
}

// Represents the spec of disk options.
type DiskSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the boot disk (default is "pd-ssd").
	// Valid values: "pd-ssd" (Persistent Disk Solid State Drive) or
	// "pd-standard" (Persistent Disk Hard Disk Drive).
	BootDiskType string `protobuf:"bytes,1,opt,name=boot_disk_type,json=bootDiskType,proto3" json:"boot_disk_type,omitempty"`
	// Size in GB of the boot disk (default is 100GB).
	BootDiskSizeGb int32 `protobuf:"varint,2,opt,name=boot_disk_size_gb,json=bootDiskSizeGb,proto3" json:"boot_disk_size_gb,omitempty"`
}

func (x *DiskSpec) Reset() {
	*x = DiskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskSpec) ProtoMessage() {}

func (x *DiskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskSpec.ProtoReflect.Descriptor instead.
func (*DiskSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{5}
}

func (x *DiskSpec) GetBootDiskType() string {
	if x != nil {
		return x.BootDiskType
	}
	return ""
}

func (x *DiskSpec) GetBootDiskSizeGb() int32 {
	if x != nil {
		return x.BootDiskSizeGb
	}
	return 0
}

// Represents a mount configuration for Network File System (NFS) to mount.
type NfsMount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. IP address of the NFS server.
	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Required. Source path exported from NFS server.
	// Has to start with '/', and combined with the ip address, it indicates
	// the source mount path in the form of `server:path`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Required. Destination mount path. The NFS will be mounted for the user
	// under /mnt/nfs/<mount_point>
	MountPoint string `protobuf:"bytes,3,opt,name=mount_point,json=mountPoint,proto3" json:"mount_point,omitempty"`
}

func (x *NfsMount) Reset() {
	*x = NfsMount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NfsMount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NfsMount) ProtoMessage() {}

func (x *NfsMount) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NfsMount.ProtoReflect.Descriptor instead.
func (*NfsMount) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{6}
}

func (x *NfsMount) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *NfsMount) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *NfsMount) GetMountPoint() string {
	if x != nil {
		return x.MountPoint
	}
	return ""
}

// The metric specification that defines the target resource utilization
// (CPU utilization, accelerator's duty cycle, and so on) for calculating the
// desired replica count.
type AutoscalingMetricSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource metric name.
	// Supported metrics:
	//
	// * For Online Prediction:
	// * `aiplatform.googleapis.com/prediction/online/accelerator/duty_cycle`
	// * `aiplatform.googleapis.com/prediction/online/cpu/utilization`
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// The target resource utilization in percentage (1% - 100%) for the given
	// metric; once the real usage deviates from the target by a certain
	// percentage, the machine replicas change. The default value is 60
	// (representing 60%) if not provided.
	Target int32 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AutoscalingMetricSpec) Reset() {
	*x = AutoscalingMetricSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoscalingMetricSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoscalingMetricSpec) ProtoMessage() {}

func (x *AutoscalingMetricSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoscalingMetricSpec.ProtoReflect.Descriptor instead.
func (*AutoscalingMetricSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP(), []int{7}
}

func (x *AutoscalingMetricSpec) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *AutoscalingMetricSpec) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

var File_google_cloud_aiplatform_v1beta1_machine_resources_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x26, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0b,
	0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x60, 0x0a, 0x10, 0x61,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x61, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc9, 0x02, 0x0a, 0x12, 0x44,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x32, 0x0a, 0x11, 0x6d, 0x69,
	0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f,
	0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x75, 0x0a, 0x18, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x16,
	0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x70, 0x65, 0x63, 0x73, 0x22, 0x76, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x11,
	0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f, 0x6d,
	0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xde,
	0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x06,
	0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x39, 0x0a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x14, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f,
	0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x3d, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x5b,
	0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x6f,
	0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x67, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x6f, 0x6f,
	0x74, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x62, 0x22, 0x66, 0x0a, 0x08, 0x4e,
	0x66, 0x73, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a,
	0x0b, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x24, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0xec, 0x01, 0x0a, 0x23, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x15, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_aiplatform_v1beta1_machine_resources_proto_goTypes = []interface{}{
	(*MachineSpec)(nil),             // 0: google.cloud.aiplatform.v1beta1.MachineSpec
	(*DedicatedResources)(nil),      // 1: google.cloud.aiplatform.v1beta1.DedicatedResources
	(*AutomaticResources)(nil),      // 2: google.cloud.aiplatform.v1beta1.AutomaticResources
	(*BatchDedicatedResources)(nil), // 3: google.cloud.aiplatform.v1beta1.BatchDedicatedResources
	(*ResourcesConsumed)(nil),       // 4: google.cloud.aiplatform.v1beta1.ResourcesConsumed
	(*DiskSpec)(nil),                // 5: google.cloud.aiplatform.v1beta1.DiskSpec
	(*NfsMount)(nil),                // 6: google.cloud.aiplatform.v1beta1.NfsMount
	(*AutoscalingMetricSpec)(nil),   // 7: google.cloud.aiplatform.v1beta1.AutoscalingMetricSpec
	(AcceleratorType)(0),            // 8: google.cloud.aiplatform.v1beta1.AcceleratorType
}
var file_google_cloud_aiplatform_v1beta1_machine_resources_proto_depIdxs = []int32{
	8, // 0: google.cloud.aiplatform.v1beta1.MachineSpec.accelerator_type:type_name -> google.cloud.aiplatform.v1beta1.AcceleratorType
	0, // 1: google.cloud.aiplatform.v1beta1.DedicatedResources.machine_spec:type_name -> google.cloud.aiplatform.v1beta1.MachineSpec
	7, // 2: google.cloud.aiplatform.v1beta1.DedicatedResources.autoscaling_metric_specs:type_name -> google.cloud.aiplatform.v1beta1.AutoscalingMetricSpec
	0, // 3: google.cloud.aiplatform.v1beta1.BatchDedicatedResources.machine_spec:type_name -> google.cloud.aiplatform.v1beta1.MachineSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_machine_resources_proto_init() }
func file_google_cloud_aiplatform_v1beta1_machine_resources_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_machine_resources_proto != nil {
		return
	}
	file_google_cloud_aiplatform_v1beta1_accelerator_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DedicatedResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutomaticResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDedicatedResources); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourcesConsumed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NfsMount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoscalingMetricSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_machine_resources_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_machine_resources_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_machine_resources_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_machine_resources_proto = out.File
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_machine_resources_proto_depIdxs = nil
}
