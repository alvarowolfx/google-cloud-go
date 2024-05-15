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
// source: google/cloud/datacatalog/v1/bigquery.proto

package datacatalogpb

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

// The type of the BigQuery connection.
type BigQueryConnectionSpec_ConnectionType int32

const (
	// Unspecified type.
	BigQueryConnectionSpec_CONNECTION_TYPE_UNSPECIFIED BigQueryConnectionSpec_ConnectionType = 0
	// Cloud SQL connection.
	BigQueryConnectionSpec_CLOUD_SQL BigQueryConnectionSpec_ConnectionType = 1
)

// Enum value maps for BigQueryConnectionSpec_ConnectionType.
var (
	BigQueryConnectionSpec_ConnectionType_name = map[int32]string{
		0: "CONNECTION_TYPE_UNSPECIFIED",
		1: "CLOUD_SQL",
	}
	BigQueryConnectionSpec_ConnectionType_value = map[string]int32{
		"CONNECTION_TYPE_UNSPECIFIED": 0,
		"CLOUD_SQL":                   1,
	}
)

func (x BigQueryConnectionSpec_ConnectionType) Enum() *BigQueryConnectionSpec_ConnectionType {
	p := new(BigQueryConnectionSpec_ConnectionType)
	*p = x
	return p
}

func (x BigQueryConnectionSpec_ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BigQueryConnectionSpec_ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes[0].Descriptor()
}

func (BigQueryConnectionSpec_ConnectionType) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes[0]
}

func (x BigQueryConnectionSpec_ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BigQueryConnectionSpec_ConnectionType.Descriptor instead.
func (BigQueryConnectionSpec_ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP(), []int{0, 0}
}

// Supported Cloud SQL database types.
type CloudSqlBigQueryConnectionSpec_DatabaseType int32

const (
	// Unspecified database type.
	CloudSqlBigQueryConnectionSpec_DATABASE_TYPE_UNSPECIFIED CloudSqlBigQueryConnectionSpec_DatabaseType = 0
	// Cloud SQL for PostgreSQL.
	CloudSqlBigQueryConnectionSpec_POSTGRES CloudSqlBigQueryConnectionSpec_DatabaseType = 1
	// Cloud SQL for MySQL.
	CloudSqlBigQueryConnectionSpec_MYSQL CloudSqlBigQueryConnectionSpec_DatabaseType = 2
)

// Enum value maps for CloudSqlBigQueryConnectionSpec_DatabaseType.
var (
	CloudSqlBigQueryConnectionSpec_DatabaseType_name = map[int32]string{
		0: "DATABASE_TYPE_UNSPECIFIED",
		1: "POSTGRES",
		2: "MYSQL",
	}
	CloudSqlBigQueryConnectionSpec_DatabaseType_value = map[string]int32{
		"DATABASE_TYPE_UNSPECIFIED": 0,
		"POSTGRES":                  1,
		"MYSQL":                     2,
	}
)

func (x CloudSqlBigQueryConnectionSpec_DatabaseType) Enum() *CloudSqlBigQueryConnectionSpec_DatabaseType {
	p := new(CloudSqlBigQueryConnectionSpec_DatabaseType)
	*p = x
	return p
}

func (x CloudSqlBigQueryConnectionSpec_DatabaseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudSqlBigQueryConnectionSpec_DatabaseType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes[1].Descriptor()
}

func (CloudSqlBigQueryConnectionSpec_DatabaseType) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes[1]
}

func (x CloudSqlBigQueryConnectionSpec_DatabaseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudSqlBigQueryConnectionSpec_DatabaseType.Descriptor instead.
func (CloudSqlBigQueryConnectionSpec_DatabaseType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP(), []int{1, 0}
}

// Specification for the BigQuery connection.
type BigQueryConnectionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the BigQuery connection.
	ConnectionType BigQueryConnectionSpec_ConnectionType `protobuf:"varint,1,opt,name=connection_type,json=connectionType,proto3,enum=google.cloud.datacatalog.v1.BigQueryConnectionSpec_ConnectionType" json:"connection_type,omitempty"`
	// Types that are assignable to ConnectionSpec:
	//
	//	*BigQueryConnectionSpec_CloudSql
	ConnectionSpec isBigQueryConnectionSpec_ConnectionSpec `protobuf_oneof:"connection_spec"`
	// True if there are credentials attached to the BigQuery connection; false
	// otherwise.
	HasCredential bool `protobuf:"varint,3,opt,name=has_credential,json=hasCredential,proto3" json:"has_credential,omitempty"`
}

func (x *BigQueryConnectionSpec) Reset() {
	*x = BigQueryConnectionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryConnectionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryConnectionSpec) ProtoMessage() {}

func (x *BigQueryConnectionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryConnectionSpec.ProtoReflect.Descriptor instead.
func (*BigQueryConnectionSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP(), []int{0}
}

func (x *BigQueryConnectionSpec) GetConnectionType() BigQueryConnectionSpec_ConnectionType {
	if x != nil {
		return x.ConnectionType
	}
	return BigQueryConnectionSpec_CONNECTION_TYPE_UNSPECIFIED
}

func (m *BigQueryConnectionSpec) GetConnectionSpec() isBigQueryConnectionSpec_ConnectionSpec {
	if m != nil {
		return m.ConnectionSpec
	}
	return nil
}

func (x *BigQueryConnectionSpec) GetCloudSql() *CloudSqlBigQueryConnectionSpec {
	if x, ok := x.GetConnectionSpec().(*BigQueryConnectionSpec_CloudSql); ok {
		return x.CloudSql
	}
	return nil
}

func (x *BigQueryConnectionSpec) GetHasCredential() bool {
	if x != nil {
		return x.HasCredential
	}
	return false
}

type isBigQueryConnectionSpec_ConnectionSpec interface {
	isBigQueryConnectionSpec_ConnectionSpec()
}

type BigQueryConnectionSpec_CloudSql struct {
	// Specification for the BigQuery connection to a Cloud SQL instance.
	CloudSql *CloudSqlBigQueryConnectionSpec `protobuf:"bytes,2,opt,name=cloud_sql,json=cloudSql,proto3,oneof"`
}

func (*BigQueryConnectionSpec_CloudSql) isBigQueryConnectionSpec_ConnectionSpec() {}

// Specification for the BigQuery connection to a Cloud SQL instance.
type CloudSqlBigQueryConnectionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud SQL instance ID in the format of `project:location:instance`.
	InstanceId string `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Database name.
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// Type of the Cloud SQL database.
	Type CloudSqlBigQueryConnectionSpec_DatabaseType `protobuf:"varint,3,opt,name=type,proto3,enum=google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec_DatabaseType" json:"type,omitempty"`
}

func (x *CloudSqlBigQueryConnectionSpec) Reset() {
	*x = CloudSqlBigQueryConnectionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSqlBigQueryConnectionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSqlBigQueryConnectionSpec) ProtoMessage() {}

func (x *CloudSqlBigQueryConnectionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSqlBigQueryConnectionSpec.ProtoReflect.Descriptor instead.
func (*CloudSqlBigQueryConnectionSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP(), []int{1}
}

func (x *CloudSqlBigQueryConnectionSpec) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CloudSqlBigQueryConnectionSpec) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *CloudSqlBigQueryConnectionSpec) GetType() CloudSqlBigQueryConnectionSpec_DatabaseType {
	if x != nil {
		return x.Type
	}
	return CloudSqlBigQueryConnectionSpec_DATABASE_TYPE_UNSPECIFIED
}

// Fields specific for BigQuery routines.
type BigQueryRoutineSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Paths of the imported libraries.
	ImportedLibraries []string `protobuf:"bytes,1,rep,name=imported_libraries,json=importedLibraries,proto3" json:"imported_libraries,omitempty"`
}

func (x *BigQueryRoutineSpec) Reset() {
	*x = BigQueryRoutineSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryRoutineSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryRoutineSpec) ProtoMessage() {}

func (x *BigQueryRoutineSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryRoutineSpec.ProtoReflect.Descriptor instead.
func (*BigQueryRoutineSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP(), []int{2}
}

func (x *BigQueryRoutineSpec) GetImportedLibraries() []string {
	if x != nil {
		return x.ImportedLibraries
	}
	return nil
}

var File_google_cloud_datacatalog_v1_bigquery_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_bigquery_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69,
	0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x02, 0x0a, 0x16, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x6b, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x67,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x71, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70,
	0x65, 0x63, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x68, 0x61, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x40, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x4f, 0x55,
	0x44, 0x5f, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x22, 0x83, 0x02, 0x0a, 0x1e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x71, 0x6c, 0x42,
	0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x41, 0x54, 0x41,
	0x42, 0x41, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4f, 0x53, 0x54, 0x47,
	0x52, 0x45, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x59, 0x53, 0x51, 0x4c, 0x10, 0x02,
	0x22, 0x44, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x42, 0xd5, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x42, 0x69, 0x67, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62,
	0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0xf8, 0x01,
	0x01, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_bigquery_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_bigquery_proto_rawDescData = file_google_cloud_datacatalog_v1_bigquery_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_bigquery_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_bigquery_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_bigquery_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_bigquery_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_bigquery_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_datacatalog_v1_bigquery_proto_goTypes = []interface{}{
	(BigQueryConnectionSpec_ConnectionType)(0),       // 0: google.cloud.datacatalog.v1.BigQueryConnectionSpec.ConnectionType
	(CloudSqlBigQueryConnectionSpec_DatabaseType)(0), // 1: google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.DatabaseType
	(*BigQueryConnectionSpec)(nil),                   // 2: google.cloud.datacatalog.v1.BigQueryConnectionSpec
	(*CloudSqlBigQueryConnectionSpec)(nil),           // 3: google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec
	(*BigQueryRoutineSpec)(nil),                      // 4: google.cloud.datacatalog.v1.BigQueryRoutineSpec
}
var file_google_cloud_datacatalog_v1_bigquery_proto_depIdxs = []int32{
	0, // 0: google.cloud.datacatalog.v1.BigQueryConnectionSpec.connection_type:type_name -> google.cloud.datacatalog.v1.BigQueryConnectionSpec.ConnectionType
	3, // 1: google.cloud.datacatalog.v1.BigQueryConnectionSpec.cloud_sql:type_name -> google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec
	1, // 2: google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.type:type_name -> google.cloud.datacatalog.v1.CloudSqlBigQueryConnectionSpec.DatabaseType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_bigquery_proto_init() }
func file_google_cloud_datacatalog_v1_bigquery_proto_init() {
	if File_google_cloud_datacatalog_v1_bigquery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryConnectionSpec); i {
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
		file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudSqlBigQueryConnectionSpec); i {
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
		file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryRoutineSpec); i {
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
	file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BigQueryConnectionSpec_CloudSql)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_datacatalog_v1_bigquery_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_bigquery_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_bigquery_proto_depIdxs,
		EnumInfos:         file_google_cloud_datacatalog_v1_bigquery_proto_enumTypes,
		MessageInfos:      file_google_cloud_datacatalog_v1_bigquery_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_bigquery_proto = out.File
	file_google_cloud_datacatalog_v1_bigquery_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_bigquery_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_bigquery_proto_depIdxs = nil
}
