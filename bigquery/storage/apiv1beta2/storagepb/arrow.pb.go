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
// source: google/cloud/bigquery/storage/v1beta2/arrow.proto

package storagepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The IPC format to use when serializing Arrow streams.
type ArrowSerializationOptions_Format int32

const (
	// If unspecied the IPC format as of 0.15 release will be used.
	ArrowSerializationOptions_FORMAT_UNSPECIFIED ArrowSerializationOptions_Format = 0
	// Use the legacy IPC message format as of Apache Arrow Release 0.14.
	ArrowSerializationOptions_ARROW_0_14 ArrowSerializationOptions_Format = 1
	// Use the message format as of Apache Arrow Release 0.15.
	ArrowSerializationOptions_ARROW_0_15 ArrowSerializationOptions_Format = 2
)

// Enum value maps for ArrowSerializationOptions_Format.
var (
	ArrowSerializationOptions_Format_name = map[int32]string{
		0: "FORMAT_UNSPECIFIED",
		1: "ARROW_0_14",
		2: "ARROW_0_15",
	}
	ArrowSerializationOptions_Format_value = map[string]int32{
		"FORMAT_UNSPECIFIED": 0,
		"ARROW_0_14":         1,
		"ARROW_0_15":         2,
	}
)

func (x ArrowSerializationOptions_Format) Enum() *ArrowSerializationOptions_Format {
	p := new(ArrowSerializationOptions_Format)
	*p = x
	return p
}

func (x ArrowSerializationOptions_Format) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArrowSerializationOptions_Format) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_enumTypes[0].Descriptor()
}

func (ArrowSerializationOptions_Format) Type() protoreflect.EnumType {
	return &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_enumTypes[0]
}

func (x ArrowSerializationOptions_Format) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArrowSerializationOptions_Format.Descriptor instead.
func (ArrowSerializationOptions_Format) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescGZIP(), []int{2, 0}
}

// Arrow schema as specified in
// https://arrow.apache.org/docs/python/api/datatypes.html
// and serialized to bytes using IPC:
// https://arrow.apache.org/docs/format/Columnar.html#serialization-and-interprocess-communication-ipc
//
// See code samples on how this message can be deserialized.
type ArrowSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IPC serialized Arrow schema.
	SerializedSchema []byte `protobuf:"bytes,1,opt,name=serialized_schema,json=serializedSchema,proto3" json:"serialized_schema,omitempty"`
}

func (x *ArrowSchema) Reset() {
	*x = ArrowSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowSchema) ProtoMessage() {}

func (x *ArrowSchema) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowSchema.ProtoReflect.Descriptor instead.
func (*ArrowSchema) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescGZIP(), []int{0}
}

func (x *ArrowSchema) GetSerializedSchema() []byte {
	if x != nil {
		return x.SerializedSchema
	}
	return nil
}

// Arrow RecordBatch.
type ArrowRecordBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IPC-serialized Arrow RecordBatch.
	SerializedRecordBatch []byte `protobuf:"bytes,1,opt,name=serialized_record_batch,json=serializedRecordBatch,proto3" json:"serialized_record_batch,omitempty"`
}

func (x *ArrowRecordBatch) Reset() {
	*x = ArrowRecordBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowRecordBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowRecordBatch) ProtoMessage() {}

func (x *ArrowRecordBatch) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowRecordBatch.ProtoReflect.Descriptor instead.
func (*ArrowRecordBatch) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescGZIP(), []int{1}
}

func (x *ArrowRecordBatch) GetSerializedRecordBatch() []byte {
	if x != nil {
		return x.SerializedRecordBatch
	}
	return nil
}

// Contains options specific to Arrow Serialization.
type ArrowSerializationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Arrow IPC format to use.
	Format ArrowSerializationOptions_Format `protobuf:"varint,1,opt,name=format,proto3,enum=google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions_Format" json:"format,omitempty"`
}

func (x *ArrowSerializationOptions) Reset() {
	*x = ArrowSerializationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrowSerializationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrowSerializationOptions) ProtoMessage() {}

func (x *ArrowSerializationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrowSerializationOptions.ProtoReflect.Descriptor instead.
func (*ArrowSerializationOptions) Descriptor() ([]byte, []int) {
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescGZIP(), []int{2}
}

func (x *ArrowSerializationOptions) GetFormat() ArrowSerializationOptions_Format {
	if x != nil {
		return x.Format
	}
	return ArrowSerializationOptions_FORMAT_UNSPECIFIED
}

var File_google_cloud_bigquery_storage_v1beta2_arrow_proto protoreflect.FileDescriptor

var file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x22, 0x3a, 0x0a, 0x0b, 0x41, 0x72,
	0x72, 0x6f, 0x77, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x15, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x5f, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x22, 0x40, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x52, 0x52, 0x4f, 0x57, 0x5f, 0x30, 0x5f, 0x31,
	0x34, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x52, 0x52, 0x4f, 0x57, 0x5f, 0x30, 0x5f, 0x31,
	0x35, 0x10, 0x02, 0x42, 0x7e, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32,
	0x42, 0x0a, 0x41, 0x72, 0x72, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescOnce sync.Once
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescData = file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDesc
)

func file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescGZIP() []byte {
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescOnce.Do(func() {
		file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescData)
	})
	return file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDescData
}

var file_google_cloud_bigquery_storage_v1beta2_arrow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_bigquery_storage_v1beta2_arrow_proto_goTypes = []interface{}{
	(ArrowSerializationOptions_Format)(0), // 0: google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions.Format
	(*ArrowSchema)(nil),                   // 1: google.cloud.bigquery.storage.v1beta2.ArrowSchema
	(*ArrowRecordBatch)(nil),              // 2: google.cloud.bigquery.storage.v1beta2.ArrowRecordBatch
	(*ArrowSerializationOptions)(nil),     // 3: google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions
}
var file_google_cloud_bigquery_storage_v1beta2_arrow_proto_depIdxs = []int32{
	0, // 0: google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions.format:type_name -> google.cloud.bigquery.storage.v1beta2.ArrowSerializationOptions.Format
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_bigquery_storage_v1beta2_arrow_proto_init() }
func file_google_cloud_bigquery_storage_v1beta2_arrow_proto_init() {
	if File_google_cloud_bigquery_storage_v1beta2_arrow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowSchema); i {
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
		file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowRecordBatch); i {
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
		file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrowSerializationOptions); i {
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
			RawDescriptor: file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_bigquery_storage_v1beta2_arrow_proto_goTypes,
		DependencyIndexes: file_google_cloud_bigquery_storage_v1beta2_arrow_proto_depIdxs,
		EnumInfos:         file_google_cloud_bigquery_storage_v1beta2_arrow_proto_enumTypes,
		MessageInfos:      file_google_cloud_bigquery_storage_v1beta2_arrow_proto_msgTypes,
	}.Build()
	File_google_cloud_bigquery_storage_v1beta2_arrow_proto = out.File
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_rawDesc = nil
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_goTypes = nil
	file_google_cloud_bigquery_storage_v1beta2_arrow_proto_depIdxs = nil
}
