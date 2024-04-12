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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/video/stitcher/v1/slates.proto

package stitcherpb

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

// Slate object
type Slate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of the slate, in the form of
	// `projects/{project_number}/locations/{location}/slates/{id}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The URI to fetch the source content for the slate. This URI must return an
	// MP4 video with at least one audio track.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// gam_slate has all the GAM-related attributes of slates.
	GamSlate *Slate_GamSlate `protobuf:"bytes,3,opt,name=gam_slate,json=gamSlate,proto3" json:"gam_slate,omitempty"`
}

func (x *Slate) Reset() {
	*x = Slate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slate) ProtoMessage() {}

func (x *Slate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slate.ProtoReflect.Descriptor instead.
func (*Slate) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_slates_proto_rawDescGZIP(), []int{0}
}

func (x *Slate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Slate) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Slate) GetGamSlate() *Slate_GamSlate {
	if x != nil {
		return x.GamSlate
	}
	return nil
}

// GamSlate object has Google Ad Manager (GAM) related properties for the
// slate.
type Slate_GamSlate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Ad Manager network code to associate with the live config.
	NetworkCode string `protobuf:"bytes,1,opt,name=network_code,json=networkCode,proto3" json:"network_code,omitempty"`
	// Output only. The identifier generated for the slate by GAM.
	GamSlateId int64 `protobuf:"varint,2,opt,name=gam_slate_id,json=gamSlateId,proto3" json:"gam_slate_id,omitempty"`
}

func (x *Slate_GamSlate) Reset() {
	*x = Slate_GamSlate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slate_GamSlate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slate_GamSlate) ProtoMessage() {}

func (x *Slate_GamSlate) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slate_GamSlate.ProtoReflect.Descriptor instead.
func (*Slate_GamSlate) Descriptor() ([]byte, []int) {
	return file_google_cloud_video_stitcher_v1_slates_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Slate_GamSlate) GetNetworkCode() string {
	if x != nil {
		return x.NetworkCode
	}
	return ""
}

func (x *Slate_GamSlate) GetGamSlateId() int64 {
	if x != nil {
		return x.GamSlateId
	}
	return 0
}

var File_google_cloud_video_stitcher_v1_slates_proto protoreflect.FileDescriptor

var file_google_cloud_video_stitcher_v1_slates_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x05, 0x53, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x4b,
	0x0a, 0x09, 0x67, 0x61, 0x6d, 0x5f, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x53, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x59, 0x0a, 0x08, 0x47,
	0x61, 0x6d, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x25, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x5f, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x53,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x3a, 0x5f, 0xea, 0x41, 0x5c, 0x0a, 0x22, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x7d, 0x42, 0x73, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70,
	0x62, 0x3b, 0x73, 0x74, 0x69, 0x74, 0x63, 0x68, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_video_stitcher_v1_slates_proto_rawDescOnce sync.Once
	file_google_cloud_video_stitcher_v1_slates_proto_rawDescData = file_google_cloud_video_stitcher_v1_slates_proto_rawDesc
)

func file_google_cloud_video_stitcher_v1_slates_proto_rawDescGZIP() []byte {
	file_google_cloud_video_stitcher_v1_slates_proto_rawDescOnce.Do(func() {
		file_google_cloud_video_stitcher_v1_slates_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_video_stitcher_v1_slates_proto_rawDescData)
	})
	return file_google_cloud_video_stitcher_v1_slates_proto_rawDescData
}

var file_google_cloud_video_stitcher_v1_slates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_video_stitcher_v1_slates_proto_goTypes = []interface{}{
	(*Slate)(nil),          // 0: google.cloud.video.stitcher.v1.Slate
	(*Slate_GamSlate)(nil), // 1: google.cloud.video.stitcher.v1.Slate.GamSlate
}
var file_google_cloud_video_stitcher_v1_slates_proto_depIdxs = []int32{
	1, // 0: google.cloud.video.stitcher.v1.Slate.gam_slate:type_name -> google.cloud.video.stitcher.v1.Slate.GamSlate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_video_stitcher_v1_slates_proto_init() }
func file_google_cloud_video_stitcher_v1_slates_proto_init() {
	if File_google_cloud_video_stitcher_v1_slates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slate); i {
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
		file_google_cloud_video_stitcher_v1_slates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slate_GamSlate); i {
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
			RawDescriptor: file_google_cloud_video_stitcher_v1_slates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_video_stitcher_v1_slates_proto_goTypes,
		DependencyIndexes: file_google_cloud_video_stitcher_v1_slates_proto_depIdxs,
		MessageInfos:      file_google_cloud_video_stitcher_v1_slates_proto_msgTypes,
	}.Build()
	File_google_cloud_video_stitcher_v1_slates_proto = out.File
	file_google_cloud_video_stitcher_v1_slates_proto_rawDesc = nil
	file_google_cloud_video_stitcher_v1_slates_proto_goTypes = nil
	file_google_cloud_video_stitcher_v1_slates_proto_depIdxs = nil
}
