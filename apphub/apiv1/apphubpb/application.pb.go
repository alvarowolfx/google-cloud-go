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
// source: google/cloud/apphub/v1/application.proto

package apphubpb

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

// Application state.
type Application_State int32

const (
	// Unspecified state.
	Application_STATE_UNSPECIFIED Application_State = 0
	// The Application is being created.
	Application_CREATING Application_State = 1
	// The Application is ready to register Services and Workloads.
	Application_ACTIVE Application_State = 2
	// The Application is being deleted.
	Application_DELETING Application_State = 3
)

// Enum value maps for Application_State.
var (
	Application_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
	}
	Application_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"CREATING":          1,
		"ACTIVE":            2,
		"DELETING":          3,
	}
)

func (x Application_State) Enum() *Application_State {
	p := new(Application_State)
	*p = x
	return p
}

func (x Application_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Application_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apphub_v1_application_proto_enumTypes[0].Descriptor()
}

func (Application_State) Type() protoreflect.EnumType {
	return &file_google_cloud_apphub_v1_application_proto_enumTypes[0]
}

func (x Application_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Application_State.Descriptor instead.
func (Application_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_application_proto_rawDescGZIP(), []int{0, 0}
}

// Scope Type.
type Scope_Type int32

const (
	// Unspecified type.
	Scope_TYPE_UNSPECIFIED Scope_Type = 0
	// Regional type.
	Scope_REGIONAL Scope_Type = 1
)

// Enum value maps for Scope_Type.
var (
	Scope_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "REGIONAL",
	}
	Scope_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"REGIONAL":         1,
	}
)

func (x Scope_Type) Enum() *Scope_Type {
	p := new(Scope_Type)
	*p = x
	return p
}

func (x Scope_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scope_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_apphub_v1_application_proto_enumTypes[1].Descriptor()
}

func (Scope_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_apphub_v1_application_proto_enumTypes[1]
}

func (x Scope_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scope_Type.Descriptor instead.
func (Scope_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_application_proto_rawDescGZIP(), []int{1, 0}
}

// Application defines the governance boundary for App Hub Entities that
// perform a logical end-to-end business function.
// App Hub supports application level IAM permission to align with governance
// requirements.
type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of an Application. Format:
	// "projects/{host-project-id}/locations/{location}/applications/{application-id}"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. User-defined name for the Application.
	// Can have a maximum length of 63 characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Optional. User-defined description of an Application.
	// Can have a maximum length of 2048 characters.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Consumer provided attributes.
	Attributes *Attributes `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Output only. Create time.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Update time.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Required. Immutable. Defines what data can be included into this
	// Application. Limits which Services and Workloads can be registered.
	Scope *Scope `protobuf:"bytes,9,opt,name=scope,proto3" json:"scope,omitempty"`
	// Output only. A universally unique identifier (in UUID4 format) for the
	// `Application`.
	Uid string `protobuf:"bytes,10,opt,name=uid,proto3" json:"uid,omitempty"`
	// Output only. Application state.
	State Application_State `protobuf:"varint,11,opt,name=state,proto3,enum=google.cloud.apphub.v1.Application_State" json:"state,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_application_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Application) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Application) GetAttributes() *Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Application) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Application) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Application) GetScope() *Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Application) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Application) GetState() Application_State {
	if x != nil {
		return x.State
	}
	return Application_STATE_UNSPECIFIED
}

// Scope of an application.
type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Scope Type.
	Type Scope_Type `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.apphub.v1.Scope_Type" json:"type,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_apphub_v1_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_apphub_v1_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_google_cloud_apphub_v1_application_proto_rawDescGZIP(), []int{1}
}

func (x *Scope) GetType() Scope_Type {
	if x != nil {
		return x.Type
	}
	return Scope_TYPE_UNSPECIFIED
}

var File_google_cloud_apphub_v1_application_proto protoreflect.FileDescriptor

var file_google_cloud_apphub_v1_application_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x05, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xe0, 0x41, 0x03, 0xe2, 0x8c, 0xcf, 0xd7, 0x08, 0x02, 0x08,
	0x01, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x44, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x03, 0x3a, 0x86, 0x01, 0xea, 0x41, 0x82, 0x01, 0x0a, 0x21, 0x61, 0x70, 0x70,
	0x68, 0x75, 0x62, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x32, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x70, 0x0a,
	0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x42,
	0xb2, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x10,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x68, 0x75, 0x62, 0x70, 0x62, 0x3b, 0x61, 0x70,
	0x70, 0x68, 0x75, 0x62, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x48, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41,
	0x70, 0x70, 0x48, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x70, 0x70, 0x48, 0x75, 0x62,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_apphub_v1_application_proto_rawDescOnce sync.Once
	file_google_cloud_apphub_v1_application_proto_rawDescData = file_google_cloud_apphub_v1_application_proto_rawDesc
)

func file_google_cloud_apphub_v1_application_proto_rawDescGZIP() []byte {
	file_google_cloud_apphub_v1_application_proto_rawDescOnce.Do(func() {
		file_google_cloud_apphub_v1_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_apphub_v1_application_proto_rawDescData)
	})
	return file_google_cloud_apphub_v1_application_proto_rawDescData
}

var file_google_cloud_apphub_v1_application_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_apphub_v1_application_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_apphub_v1_application_proto_goTypes = []interface{}{
	(Application_State)(0),        // 0: google.cloud.apphub.v1.Application.State
	(Scope_Type)(0),               // 1: google.cloud.apphub.v1.Scope.Type
	(*Application)(nil),           // 2: google.cloud.apphub.v1.Application
	(*Scope)(nil),                 // 3: google.cloud.apphub.v1.Scope
	(*Attributes)(nil),            // 4: google.cloud.apphub.v1.Attributes
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_google_cloud_apphub_v1_application_proto_depIdxs = []int32{
	4, // 0: google.cloud.apphub.v1.Application.attributes:type_name -> google.cloud.apphub.v1.Attributes
	5, // 1: google.cloud.apphub.v1.Application.create_time:type_name -> google.protobuf.Timestamp
	5, // 2: google.cloud.apphub.v1.Application.update_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.apphub.v1.Application.scope:type_name -> google.cloud.apphub.v1.Scope
	0, // 4: google.cloud.apphub.v1.Application.state:type_name -> google.cloud.apphub.v1.Application.State
	1, // 5: google.cloud.apphub.v1.Scope.type:type_name -> google.cloud.apphub.v1.Scope.Type
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_apphub_v1_application_proto_init() }
func file_google_cloud_apphub_v1_application_proto_init() {
	if File_google_cloud_apphub_v1_application_proto != nil {
		return
	}
	file_google_cloud_apphub_v1_attributes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_apphub_v1_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_google_cloud_apphub_v1_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
			RawDescriptor: file_google_cloud_apphub_v1_application_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_apphub_v1_application_proto_goTypes,
		DependencyIndexes: file_google_cloud_apphub_v1_application_proto_depIdxs,
		EnumInfos:         file_google_cloud_apphub_v1_application_proto_enumTypes,
		MessageInfos:      file_google_cloud_apphub_v1_application_proto_msgTypes,
	}.Build()
	File_google_cloud_apphub_v1_application_proto = out.File
	file_google_cloud_apphub_v1_application_proto_rawDesc = nil
	file_google_cloud_apphub_v1_application_proto_goTypes = nil
	file_google_cloud_apphub_v1_application_proto_depIdxs = nil
}
