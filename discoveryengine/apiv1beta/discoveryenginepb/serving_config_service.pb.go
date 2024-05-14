// Copyright 2022 Google LLC
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
// source: google/cloud/discoveryengine/v1beta/serving_config_service.proto

package discoveryenginepb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for UpdateServingConfig method.
type UpdateServingConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ServingConfig to update.
	ServingConfig *ServingConfig `protobuf:"bytes,1,opt,name=serving_config,json=servingConfig,proto3" json:"serving_config,omitempty"`
	// Indicates which fields in the provided
	// [ServingConfig][google.cloud.discoveryengine.v1beta.ServingConfig] to
	// update. The following are NOT supported:
	//
	// * [ServingConfig.name][google.cloud.discoveryengine.v1beta.ServingConfig.name]
	//
	// If not set, all supported fields are updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateServingConfigRequest) Reset() {
	*x = UpdateServingConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServingConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServingConfigRequest) ProtoMessage() {}

func (x *UpdateServingConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServingConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateServingConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateServingConfigRequest) GetServingConfig() *ServingConfig {
	if x != nil {
		return x.ServingConfig
	}
	return nil
}

func (x *UpdateServingConfigRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request for GetServingConfig method.
type GetServingConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the ServingConfig to get. Format:
	// `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}/servingConfigs/{serving_config_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetServingConfigRequest) Reset() {
	*x = GetServingConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServingConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServingConfigRequest) ProtoMessage() {}

func (x *GetServingConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServingConfigRequest.ProtoReflect.Descriptor instead.
func (*GetServingConfigRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetServingConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request for ListServingConfigs method.
type ListServingConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of the parent resource. Format:
	// `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Maximum number of results to return. If unspecified, defaults
	// to 100. If a value greater than 100 is provided, at most 100 results are
	// returned.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A page token, received from a previous `ListServingConfigs` call.
	// Provide this to retrieve the subsequent page.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListServingConfigsRequest) Reset() {
	*x = ListServingConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServingConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServingConfigsRequest) ProtoMessage() {}

func (x *ListServingConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServingConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListServingConfigsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListServingConfigsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListServingConfigsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListServingConfigsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response for ListServingConfigs method.
type ListServingConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the ServingConfigs for a given dataStore.
	ServingConfigs []*ServingConfig `protobuf:"bytes,1,rep,name=serving_configs,json=servingConfigs,proto3" json:"serving_configs,omitempty"`
	// Pagination token, if not returned indicates the last page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListServingConfigsResponse) Reset() {
	*x = ListServingConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServingConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServingConfigsResponse) ProtoMessage() {}

func (x *ListServingConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServingConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListServingConfigsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListServingConfigsResponse) GetServingConfigs() []*ServingConfig {
	if x != nil {
		return x.ServingConfigs
	}
	return nil
}

func (x *ListServingConfigsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_google_cloud_discoveryengine_v1beta_serving_config_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x63, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2e, 0x0a, 0x2c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x34, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2e, 0x12, 0x2c, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa1, 0x01, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x82, 0x0b,
	0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xfa, 0x03, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0xed, 0x02, 0xda, 0x41, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xc9, 0x02, 0x3a, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5a, 0x72, 0x3a, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x60, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x5a, 0x6f,
	0x3a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0x5d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x32,
	0x52, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f,
	0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0x81, 0x03, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xfa, 0x01, 0xda, 0x41, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0xec, 0x01, 0x5a, 0x53, 0x12, 0x51, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d,
	0x5a, 0x50, 0x12, 0x4e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x2a,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f,
	0x2a, 0x7d, 0x12, 0x43, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x94, 0x03, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xfc, 0x01, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0xec, 0x01, 0x5a, 0x53, 0x12, 0x51, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x5a, 0x50, 0x12, 0x4e, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x43, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x52,
	0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x42, 0xa0, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x19,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02,
	0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45,
	0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x26, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_goTypes = []interface{}{
	(*UpdateServingConfigRequest)(nil), // 0: google.cloud.discoveryengine.v1beta.UpdateServingConfigRequest
	(*GetServingConfigRequest)(nil),    // 1: google.cloud.discoveryengine.v1beta.GetServingConfigRequest
	(*ListServingConfigsRequest)(nil),  // 2: google.cloud.discoveryengine.v1beta.ListServingConfigsRequest
	(*ListServingConfigsResponse)(nil), // 3: google.cloud.discoveryengine.v1beta.ListServingConfigsResponse
	(*ServingConfig)(nil),              // 4: google.cloud.discoveryengine.v1beta.ServingConfig
	(*fieldmaskpb.FieldMask)(nil),      // 5: google.protobuf.FieldMask
}
var file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.discoveryengine.v1beta.UpdateServingConfigRequest.serving_config:type_name -> google.cloud.discoveryengine.v1beta.ServingConfig
	5, // 1: google.cloud.discoveryengine.v1beta.UpdateServingConfigRequest.update_mask:type_name -> google.protobuf.FieldMask
	4, // 2: google.cloud.discoveryengine.v1beta.ListServingConfigsResponse.serving_configs:type_name -> google.cloud.discoveryengine.v1beta.ServingConfig
	0, // 3: google.cloud.discoveryengine.v1beta.ServingConfigService.UpdateServingConfig:input_type -> google.cloud.discoveryengine.v1beta.UpdateServingConfigRequest
	1, // 4: google.cloud.discoveryengine.v1beta.ServingConfigService.GetServingConfig:input_type -> google.cloud.discoveryengine.v1beta.GetServingConfigRequest
	2, // 5: google.cloud.discoveryengine.v1beta.ServingConfigService.ListServingConfigs:input_type -> google.cloud.discoveryengine.v1beta.ListServingConfigsRequest
	4, // 6: google.cloud.discoveryengine.v1beta.ServingConfigService.UpdateServingConfig:output_type -> google.cloud.discoveryengine.v1beta.ServingConfig
	4, // 7: google.cloud.discoveryengine.v1beta.ServingConfigService.GetServingConfig:output_type -> google.cloud.discoveryengine.v1beta.ServingConfig
	3, // 8: google.cloud.discoveryengine.v1beta.ServingConfigService.ListServingConfigs:output_type -> google.cloud.discoveryengine.v1beta.ListServingConfigsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_init() }
func file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_serving_config_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1beta_serving_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateServingConfigRequest); i {
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
		file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServingConfigRequest); i {
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
		file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServingConfigsRequest); i {
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
		file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServingConfigsResponse); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_serving_config_service_proto = out.File
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_serving_config_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServingConfigServiceClient is the client API for ServingConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServingConfigServiceClient interface {
	// Updates a ServingConfig.
	//
	// Returns a NOT_FOUND error if the ServingConfig does not exist.
	UpdateServingConfig(ctx context.Context, in *UpdateServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error)
	// Gets a ServingConfig.
	//
	// Returns a NotFound error if the ServingConfig does not exist.
	GetServingConfig(ctx context.Context, in *GetServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error)
	// Lists all ServingConfigs linked to this dataStore.
	ListServingConfigs(ctx context.Context, in *ListServingConfigsRequest, opts ...grpc.CallOption) (*ListServingConfigsResponse, error)
}

type servingConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServingConfigServiceClient(cc grpc.ClientConnInterface) ServingConfigServiceClient {
	return &servingConfigServiceClient{cc}
}

func (c *servingConfigServiceClient) UpdateServingConfig(ctx context.Context, in *UpdateServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error) {
	out := new(ServingConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.ServingConfigService/UpdateServingConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servingConfigServiceClient) GetServingConfig(ctx context.Context, in *GetServingConfigRequest, opts ...grpc.CallOption) (*ServingConfig, error) {
	out := new(ServingConfig)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.ServingConfigService/GetServingConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servingConfigServiceClient) ListServingConfigs(ctx context.Context, in *ListServingConfigsRequest, opts ...grpc.CallOption) (*ListServingConfigsResponse, error) {
	out := new(ListServingConfigsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1beta.ServingConfigService/ListServingConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServingConfigServiceServer is the server API for ServingConfigService service.
type ServingConfigServiceServer interface {
	// Updates a ServingConfig.
	//
	// Returns a NOT_FOUND error if the ServingConfig does not exist.
	UpdateServingConfig(context.Context, *UpdateServingConfigRequest) (*ServingConfig, error)
	// Gets a ServingConfig.
	//
	// Returns a NotFound error if the ServingConfig does not exist.
	GetServingConfig(context.Context, *GetServingConfigRequest) (*ServingConfig, error)
	// Lists all ServingConfigs linked to this dataStore.
	ListServingConfigs(context.Context, *ListServingConfigsRequest) (*ListServingConfigsResponse, error)
}

// UnimplementedServingConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServingConfigServiceServer struct {
}

func (*UnimplementedServingConfigServiceServer) UpdateServingConfig(context.Context, *UpdateServingConfigRequest) (*ServingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServingConfig not implemented")
}
func (*UnimplementedServingConfigServiceServer) GetServingConfig(context.Context, *GetServingConfigRequest) (*ServingConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServingConfig not implemented")
}
func (*UnimplementedServingConfigServiceServer) ListServingConfigs(context.Context, *ListServingConfigsRequest) (*ListServingConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServingConfigs not implemented")
}

func RegisterServingConfigServiceServer(s *grpc.Server, srv ServingConfigServiceServer) {
	s.RegisterService(&_ServingConfigService_serviceDesc, srv)
}

func _ServingConfigService_UpdateServingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).UpdateServingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.ServingConfigService/UpdateServingConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).UpdateServingConfig(ctx, req.(*UpdateServingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServingConfigService_GetServingConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServingConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).GetServingConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.ServingConfigService/GetServingConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).GetServingConfig(ctx, req.(*GetServingConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServingConfigService_ListServingConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServingConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServingConfigServiceServer).ListServingConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1beta.ServingConfigService/ListServingConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServingConfigServiceServer).ListServingConfigs(ctx, req.(*ListServingConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServingConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1beta.ServingConfigService",
	HandlerType: (*ServingConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateServingConfig",
			Handler:    _ServingConfigService_UpdateServingConfig_Handler,
		},
		{
			MethodName: "GetServingConfig",
			Handler:    _ServingConfigService_GetServingConfig_Handler,
		},
		{
			MethodName: "ListServingConfigs",
			Handler:    _ServingConfigService_ListServingConfigs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1beta/serving_config_service.proto",
}
