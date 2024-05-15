// Copyright 2024 Google LLC
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
// source: google/cloud/securitycenter/v2/mute_config.proto

package securitycenterpb

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

// The type of MuteConfig.
type MuteConfig_MuteConfigType int32

const (
	// Unused.
	MuteConfig_MUTE_CONFIG_TYPE_UNSPECIFIED MuteConfig_MuteConfigType = 0
	// A static mute config, which sets the static mute state of future matching
	// findings to muted. Once the static mute state has been set, finding or
	// config modifications will not affect the state.
	MuteConfig_STATIC MuteConfig_MuteConfigType = 1
)

// Enum value maps for MuteConfig_MuteConfigType.
var (
	MuteConfig_MuteConfigType_name = map[int32]string{
		0: "MUTE_CONFIG_TYPE_UNSPECIFIED",
		1: "STATIC",
	}
	MuteConfig_MuteConfigType_value = map[string]int32{
		"MUTE_CONFIG_TYPE_UNSPECIFIED": 0,
		"STATIC":                       1,
	}
)

func (x MuteConfig_MuteConfigType) Enum() *MuteConfig_MuteConfigType {
	p := new(MuteConfig_MuteConfigType)
	*p = x
	return p
}

func (x MuteConfig_MuteConfigType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MuteConfig_MuteConfigType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v2_mute_config_proto_enumTypes[0].Descriptor()
}

func (MuteConfig_MuteConfigType) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v2_mute_config_proto_enumTypes[0]
}

func (x MuteConfig_MuteConfigType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MuteConfig_MuteConfigType.Descriptor instead.
func (MuteConfig_MuteConfigType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_mute_config_proto_rawDescGZIP(), []int{0, 0}
}

// A mute config is a Cloud SCC resource that contains the configuration
// to mute create/update events of findings.
type MuteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field will be ignored if provided on config creation. The following
	// list shows some examples of the format:
	//
	// + `organizations/{organization}/muteConfigs/{mute_config}`
	// +
	// `organizations/{organization}locations/{location}//muteConfigs/{mute_config}`
	// + `folders/{folder}/muteConfigs/{mute_config}`
	// + `folders/{folder}/locations/{location}/muteConfigs/{mute_config}`
	// + `projects/{project}/muteConfigs/{mute_config}`
	// + `projects/{project}/locations/{location}/muteConfigs/{mute_config}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of the mute config.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Required. An expression that defines the filter to apply across
	// create/update events of findings. While creating a filter string, be
	// mindful of the scope in which the mute configuration is being created.
	// E.g., If a filter contains project = X but is created under the project = Y
	// scope, it might not match any findings.
	//
	// The following field and operator combinations are supported:
	//
	// * severity: `=`, `:`
	// * category: `=`, `:`
	// * resource.name: `=`, `:`
	// * resource.project_name: `=`, `:`
	// * resource.project_display_name: `=`, `:`
	// * resource.folders.resource_folder: `=`, `:`
	// * resource.parent_name: `=`, `:`
	// * resource.parent_display_name: `=`, `:`
	// * resource.type: `=`, `:`
	// * finding_class: `=`, `:`
	// * indicator.ip_addresses: `=`, `:`
	// * indicator.domains: `=`, `:`
	Filter string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	// Output only. The time at which the mute config was created.
	// This field is set by the server and will be ignored if provided on config
	// creation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The most recent time at which the mute config was updated.
	// This field is set by the server and will be ignored if provided on config
	// creation or update.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Email address of the user who last edited the mute config.
	// This field is set by the server and will be ignored if provided on config
	// creation or update.
	MostRecentEditor string `protobuf:"bytes,6,opt,name=most_recent_editor,json=mostRecentEditor,proto3" json:"most_recent_editor,omitempty"`
	// Required. The type of the mute config, which determines what type of mute
	// state the config affects. Immutable after creation.
	Type MuteConfig_MuteConfigType `protobuf:"varint,8,opt,name=type,proto3,enum=google.cloud.securitycenter.v2.MuteConfig_MuteConfigType" json:"type,omitempty"`
}

func (x *MuteConfig) Reset() {
	*x = MuteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_mute_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuteConfig) ProtoMessage() {}

func (x *MuteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_mute_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuteConfig.ProtoReflect.Descriptor instead.
func (*MuteConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_mute_config_proto_rawDescGZIP(), []int{0}
}

func (x *MuteConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MuteConfig) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MuteConfig) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *MuteConfig) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *MuteConfig) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *MuteConfig) GetMostRecentEditor() string {
	if x != nil {
		return x.MostRecentEditor
	}
	return ""
}

func (x *MuteConfig) GetType() MuteConfig_MuteConfigType {
	if x != nil {
		return x.Type
	}
	return MuteConfig_MUTE_CONFIG_TYPE_UNSPECIFIED
}

var File_google_cloud_securitycenter_v2_mute_config_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_mute_config_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd7, 0x06, 0x0a, 0x0a, 0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x12, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x10, 0x6d, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x52, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3e, 0x0a, 0x0e,
	0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x4d, 0x55, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x49, 0x43, 0x10, 0x01, 0x3a, 0xaa, 0x03, 0xea,
	0x41, 0xa6, 0x03, 0x0a, 0x28, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x36, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0x4b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x7d, 0x12, 0x2a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2f, 0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0x3f,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x7d,
	0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x2f, 0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12,
	0x2c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f,
	0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x12, 0x41, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2f, 0x7b, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d,
	0x2a, 0x0b, 0x6d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x32, 0x0a, 0x6d,
	0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xe9, 0x01, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x42, 0x0f, 0x4d, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_mute_config_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_mute_config_proto_rawDescData = file_google_cloud_securitycenter_v2_mute_config_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_mute_config_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_mute_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_mute_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_mute_config_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_mute_config_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_mute_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v2_mute_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v2_mute_config_proto_goTypes = []interface{}{
	(MuteConfig_MuteConfigType)(0), // 0: google.cloud.securitycenter.v2.MuteConfig.MuteConfigType
	(*MuteConfig)(nil),             // 1: google.cloud.securitycenter.v2.MuteConfig
	(*timestamppb.Timestamp)(nil),  // 2: google.protobuf.Timestamp
}
var file_google_cloud_securitycenter_v2_mute_config_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v2.MuteConfig.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.securitycenter.v2.MuteConfig.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.cloud.securitycenter.v2.MuteConfig.type:type_name -> google.cloud.securitycenter.v2.MuteConfig.MuteConfigType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_mute_config_proto_init() }
func file_google_cloud_securitycenter_v2_mute_config_proto_init() {
	if File_google_cloud_securitycenter_v2_mute_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v2_mute_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuteConfig); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v2_mute_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_mute_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_mute_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v2_mute_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v2_mute_config_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_mute_config_proto = out.File
	file_google_cloud_securitycenter_v2_mute_config_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_mute_config_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_mute_config_proto_depIdxs = nil
}
