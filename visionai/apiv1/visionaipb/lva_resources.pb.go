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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/visionai/v1/lva_resources.proto

package visionaipb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message describing the Operator object.
type Operator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The create timestamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The update timestamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Labels as key value pairs.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The definition of the operator.
	OperatorDefinition *OperatorDefinition `protobuf:"bytes,5,opt,name=operator_definition,json=operatorDefinition,proto3" json:"operator_definition,omitempty"`
	// The link to the docker image of the operator.
	DockerImage string `protobuf:"bytes,6,opt,name=docker_image,json=dockerImage,proto3" json:"docker_image,omitempty"`
}

func (x *Operator) Reset() {
	*x = Operator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operator) ProtoMessage() {}

func (x *Operator) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operator.ProtoReflect.Descriptor instead.
func (*Operator) Descriptor() ([]byte, []int) {
	return file_google_cloud_visionai_v1_lva_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Operator) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Operator) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Operator) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Operator) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Operator) GetOperatorDefinition() *OperatorDefinition {
	if x != nil {
		return x.OperatorDefinition
	}
	return nil
}

func (x *Operator) GetDockerImage() string {
	if x != nil {
		return x.DockerImage
	}
	return ""
}

// Message describing the Analysis object.
type Analysis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The create timestamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The update timestamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Labels as key value pairs.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The definition of the analysis.
	AnalysisDefinition *AnalysisDefinition `protobuf:"bytes,5,opt,name=analysis_definition,json=analysisDefinition,proto3" json:"analysis_definition,omitempty"`
	// Map from the input parameter in the definition to the real stream.
	// E.g., suppose you had a stream source operator named "input-0" and you try
	// to receive from the real stream "stream-0". You can add the following
	// mapping: [input-0: stream-0].
	InputStreamsMapping map[string]string `protobuf:"bytes,6,rep,name=input_streams_mapping,json=inputStreamsMapping,proto3" json:"input_streams_mapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map from the output parameter in the definition to the real stream.
	// E.g., suppose you had a stream sink operator named "output-0" and you try
	// to send to the real stream "stream-0". You can add the following
	// mapping: [output-0: stream-0].
	OutputStreamsMapping map[string]string `protobuf:"bytes,7,rep,name=output_streams_mapping,json=outputStreamsMapping,proto3" json:"output_streams_mapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Boolean flag to indicate whether you would like to disable the ability
	// to automatically start a Process when new event happening in the input
	// Stream. If you would like to start a Process manually, the field needs
	// to be set to true.
	DisableEventWatch bool `protobuf:"varint,8,opt,name=disable_event_watch,json=disableEventWatch,proto3" json:"disable_event_watch,omitempty"`
}

func (x *Analysis) Reset() {
	*x = Analysis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Analysis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Analysis) ProtoMessage() {}

func (x *Analysis) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Analysis.ProtoReflect.Descriptor instead.
func (*Analysis) Descriptor() ([]byte, []int) {
	return file_google_cloud_visionai_v1_lva_resources_proto_rawDescGZIP(), []int{1}
}

func (x *Analysis) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Analysis) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Analysis) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Analysis) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Analysis) GetAnalysisDefinition() *AnalysisDefinition {
	if x != nil {
		return x.AnalysisDefinition
	}
	return nil
}

func (x *Analysis) GetInputStreamsMapping() map[string]string {
	if x != nil {
		return x.InputStreamsMapping
	}
	return nil
}

func (x *Analysis) GetOutputStreamsMapping() map[string]string {
	if x != nil {
		return x.OutputStreamsMapping
	}
	return nil
}

func (x *Analysis) GetDisableEventWatch() bool {
	if x != nil {
		return x.DisableEventWatch
	}
	return false
}

// Message describing the Process object.
type Process struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The create timestamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The update timestamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. Reference to an existing Analysis resource.
	Analysis string `protobuf:"bytes,4,opt,name=analysis,proto3" json:"analysis,omitempty"`
	// Optional. Attribute overrides of the Analyzers.
	// Format for each single override item:
	// "{analyzer_name}:{attribute_key}={value}"
	AttributeOverrides []string `protobuf:"bytes,5,rep,name=attribute_overrides,json=attributeOverrides,proto3" json:"attribute_overrides,omitempty"`
	// Optional. Status of the Process.
	RunStatus *RunStatus `protobuf:"bytes,6,opt,name=run_status,json=runStatus,proto3" json:"run_status,omitempty"`
	// Optional. Run mode of the Process.
	RunMode RunMode `protobuf:"varint,7,opt,name=run_mode,json=runMode,proto3,enum=google.cloud.visionai.v1.RunMode" json:"run_mode,omitempty"`
	// Optional. Event ID of the input/output streams.
	// This is useful when you have a StreamSource/StreamSink operator in the
	// Analysis, and you want to manually specify the Event to read from/write to.
	EventId string `protobuf:"bytes,8,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Optional. Optional: Batch ID of the Process.
	BatchId string `protobuf:"bytes,9,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// Optional. Optional: The number of retries for a process in submission mode
	// the system should try before declaring failure. By default, no retry will
	// be performed.
	RetryCount int32 `protobuf:"varint,10,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
}

func (x *Process) Reset() {
	*x = Process{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Process) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Process) ProtoMessage() {}

func (x *Process) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Process.ProtoReflect.Descriptor instead.
func (*Process) Descriptor() ([]byte, []int) {
	return file_google_cloud_visionai_v1_lva_resources_proto_rawDescGZIP(), []int{2}
}

func (x *Process) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Process) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Process) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Process) GetAnalysis() string {
	if x != nil {
		return x.Analysis
	}
	return ""
}

func (x *Process) GetAttributeOverrides() []string {
	if x != nil {
		return x.AttributeOverrides
	}
	return nil
}

func (x *Process) GetRunStatus() *RunStatus {
	if x != nil {
		return x.RunStatus
	}
	return nil
}

func (x *Process) GetRunMode() RunMode {
	if x != nil {
		return x.RunMode
	}
	return RunMode_RUN_MODE_UNSPECIFIED
}

func (x *Process) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Process) GetBatchId() string {
	if x != nil {
		return x.BatchId
	}
	return ""
}

func (x *Process) GetRetryCount() int32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

var File_google_cloud_visionai_v1_lva_resources_proto protoreflect.FileDescriptor

var file_google_cloud_visionai_v1_lva_resources_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x76, 0x61, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c,
	0x76, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x04, 0x0a, 0x08, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5d, 0x0a, 0x13, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x63, 0xea, 0x41, 0x60, 0x0a, 0x20, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x7d, 0x22, 0xa1, 0x07, 0x0a, 0x08, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x5d, 0x0a, 0x13, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6f, 0x0a, 0x15, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x13, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x72, 0x0a, 0x16, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x14, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x46, 0x0a, 0x18, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x47, 0x0a, 0x19,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x75, 0xea, 0x41, 0x72, 0x0a, 0x20, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x4e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65,
	0x73, 0x2f, 0x7b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x7d, 0x22, 0x85, 0x05, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x44, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x08, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x12, 0x34, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0a,
	0x72, 0x75, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x72, 0x75, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x07, 0x72, 0x75, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x74,
	0xea, 0x41, 0x71, 0x0a, 0x1f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x4e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x7d,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x7d, 0x42, 0xc1, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x4c, 0x76, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x61, 0x69, 0x70, 0x62, 0x3b, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x61, 0x69, 0x70, 0x62, 0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x49, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x56,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x49, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x56, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x41, 0x49, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_visionai_v1_lva_resources_proto_rawDescOnce sync.Once
	file_google_cloud_visionai_v1_lva_resources_proto_rawDescData = file_google_cloud_visionai_v1_lva_resources_proto_rawDesc
)

func file_google_cloud_visionai_v1_lva_resources_proto_rawDescGZIP() []byte {
	file_google_cloud_visionai_v1_lva_resources_proto_rawDescOnce.Do(func() {
		file_google_cloud_visionai_v1_lva_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_visionai_v1_lva_resources_proto_rawDescData)
	})
	return file_google_cloud_visionai_v1_lva_resources_proto_rawDescData
}

var file_google_cloud_visionai_v1_lva_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_cloud_visionai_v1_lva_resources_proto_goTypes = []interface{}{
	(*Operator)(nil),              // 0: google.cloud.visionai.v1.Operator
	(*Analysis)(nil),              // 1: google.cloud.visionai.v1.Analysis
	(*Process)(nil),               // 2: google.cloud.visionai.v1.Process
	nil,                           // 3: google.cloud.visionai.v1.Operator.LabelsEntry
	nil,                           // 4: google.cloud.visionai.v1.Analysis.LabelsEntry
	nil,                           // 5: google.cloud.visionai.v1.Analysis.InputStreamsMappingEntry
	nil,                           // 6: google.cloud.visionai.v1.Analysis.OutputStreamsMappingEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*OperatorDefinition)(nil),    // 8: google.cloud.visionai.v1.OperatorDefinition
	(*AnalysisDefinition)(nil),    // 9: google.cloud.visionai.v1.AnalysisDefinition
	(*RunStatus)(nil),             // 10: google.cloud.visionai.v1.RunStatus
	(RunMode)(0),                  // 11: google.cloud.visionai.v1.RunMode
}
var file_google_cloud_visionai_v1_lva_resources_proto_depIdxs = []int32{
	7,  // 0: google.cloud.visionai.v1.Operator.create_time:type_name -> google.protobuf.Timestamp
	7,  // 1: google.cloud.visionai.v1.Operator.update_time:type_name -> google.protobuf.Timestamp
	3,  // 2: google.cloud.visionai.v1.Operator.labels:type_name -> google.cloud.visionai.v1.Operator.LabelsEntry
	8,  // 3: google.cloud.visionai.v1.Operator.operator_definition:type_name -> google.cloud.visionai.v1.OperatorDefinition
	7,  // 4: google.cloud.visionai.v1.Analysis.create_time:type_name -> google.protobuf.Timestamp
	7,  // 5: google.cloud.visionai.v1.Analysis.update_time:type_name -> google.protobuf.Timestamp
	4,  // 6: google.cloud.visionai.v1.Analysis.labels:type_name -> google.cloud.visionai.v1.Analysis.LabelsEntry
	9,  // 7: google.cloud.visionai.v1.Analysis.analysis_definition:type_name -> google.cloud.visionai.v1.AnalysisDefinition
	5,  // 8: google.cloud.visionai.v1.Analysis.input_streams_mapping:type_name -> google.cloud.visionai.v1.Analysis.InputStreamsMappingEntry
	6,  // 9: google.cloud.visionai.v1.Analysis.output_streams_mapping:type_name -> google.cloud.visionai.v1.Analysis.OutputStreamsMappingEntry
	7,  // 10: google.cloud.visionai.v1.Process.create_time:type_name -> google.protobuf.Timestamp
	7,  // 11: google.cloud.visionai.v1.Process.update_time:type_name -> google.protobuf.Timestamp
	10, // 12: google.cloud.visionai.v1.Process.run_status:type_name -> google.cloud.visionai.v1.RunStatus
	11, // 13: google.cloud.visionai.v1.Process.run_mode:type_name -> google.cloud.visionai.v1.RunMode
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_google_cloud_visionai_v1_lva_resources_proto_init() }
func file_google_cloud_visionai_v1_lva_resources_proto_init() {
	if File_google_cloud_visionai_v1_lva_resources_proto != nil {
		return
	}
	file_google_cloud_visionai_v1_lva_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operator); i {
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
		file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Analysis); i {
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
		file_google_cloud_visionai_v1_lva_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Process); i {
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
			RawDescriptor: file_google_cloud_visionai_v1_lva_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_visionai_v1_lva_resources_proto_goTypes,
		DependencyIndexes: file_google_cloud_visionai_v1_lva_resources_proto_depIdxs,
		MessageInfos:      file_google_cloud_visionai_v1_lva_resources_proto_msgTypes,
	}.Build()
	File_google_cloud_visionai_v1_lva_resources_proto = out.File
	file_google_cloud_visionai_v1_lva_resources_proto_rawDesc = nil
	file_google_cloud_visionai_v1_lva_resources_proto_goTypes = nil
	file_google_cloud_visionai_v1_lva_resources_proto_depIdxs = nil
}
