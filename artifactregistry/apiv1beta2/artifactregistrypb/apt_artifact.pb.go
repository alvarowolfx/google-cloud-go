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
// source: google/devtools/artifactregistry/v1beta2/apt_artifact.proto

package artifactregistrypb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Package type is either binary or source.
type AptArtifact_PackageType int32

const (
	// Package type is not specified.
	AptArtifact_PACKAGE_TYPE_UNSPECIFIED AptArtifact_PackageType = 0
	// Binary package.
	AptArtifact_BINARY AptArtifact_PackageType = 1
	// Source package.
	AptArtifact_SOURCE AptArtifact_PackageType = 2
)

// Enum value maps for AptArtifact_PackageType.
var (
	AptArtifact_PackageType_name = map[int32]string{
		0: "PACKAGE_TYPE_UNSPECIFIED",
		1: "BINARY",
		2: "SOURCE",
	}
	AptArtifact_PackageType_value = map[string]int32{
		"PACKAGE_TYPE_UNSPECIFIED": 0,
		"BINARY":                   1,
		"SOURCE":                   2,
	}
)

func (x AptArtifact_PackageType) Enum() *AptArtifact_PackageType {
	p := new(AptArtifact_PackageType)
	*p = x
	return p
}

func (x AptArtifact_PackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AptArtifact_PackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_enumTypes[0].Descriptor()
}

func (AptArtifact_PackageType) Type() protoreflect.EnumType {
	return &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_enumTypes[0]
}

func (x AptArtifact_PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AptArtifact_PackageType.Descriptor instead.
func (AptArtifact_PackageType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{0, 0}
}

// A detailed representation of an Apt artifact. Information in the record
// is derived from the archive's control file.
// See https://www.debian.org/doc/debian-policy/ch-controlfields.html
type AptArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Artifact Registry resource name of the artifact.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The Apt package name of the artifact.
	PackageName string `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// Output only. An artifact is a binary or source package.
	PackageType AptArtifact_PackageType `protobuf:"varint,3,opt,name=package_type,json=packageType,proto3,enum=google.devtools.artifactregistry.v1beta2.AptArtifact_PackageType" json:"package_type,omitempty"`
	// Output only. Operating system architecture of the artifact.
	Architecture string `protobuf:"bytes,4,opt,name=architecture,proto3" json:"architecture,omitempty"`
	// Output only. Repository component of the artifact.
	Component string `protobuf:"bytes,5,opt,name=component,proto3" json:"component,omitempty"`
	// Output only. Contents of the artifact's control metadata file.
	ControlFile []byte `protobuf:"bytes,6,opt,name=control_file,json=controlFile,proto3" json:"control_file,omitempty"`
}

func (x *AptArtifact) Reset() {
	*x = AptArtifact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AptArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AptArtifact) ProtoMessage() {}

func (x *AptArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AptArtifact.ProtoReflect.Descriptor instead.
func (*AptArtifact) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *AptArtifact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AptArtifact) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *AptArtifact) GetPackageType() AptArtifact_PackageType {
	if x != nil {
		return x.PackageType
	}
	return AptArtifact_PACKAGE_TYPE_UNSPECIFIED
}

func (x *AptArtifact) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *AptArtifact) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *AptArtifact) GetControlFile() []byte {
	if x != nil {
		return x.ControlFile
	}
	return nil
}

// Google Cloud Storage location where the artifacts currently reside.
type ImportAptArtifactsGcsSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud Storage paths URI (e.g., gs://my_bucket//my_object).
	Uris []string `protobuf:"bytes,1,rep,name=uris,proto3" json:"uris,omitempty"`
	// Supports URI wildcards for matching multiple objects from a single URI.
	UseWildcards bool `protobuf:"varint,2,opt,name=use_wildcards,json=useWildcards,proto3" json:"use_wildcards,omitempty"`
}

func (x *ImportAptArtifactsGcsSource) Reset() {
	*x = ImportAptArtifactsGcsSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAptArtifactsGcsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAptArtifactsGcsSource) ProtoMessage() {}

func (x *ImportAptArtifactsGcsSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAptArtifactsGcsSource.ProtoReflect.Descriptor instead.
func (*ImportAptArtifactsGcsSource) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{1}
}

func (x *ImportAptArtifactsGcsSource) GetUris() []string {
	if x != nil {
		return x.Uris
	}
	return nil
}

func (x *ImportAptArtifactsGcsSource) GetUseWildcards() bool {
	if x != nil {
		return x.UseWildcards
	}
	return false
}

// The request to import new apt artifacts.
type ImportAptArtifactsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source location of the package binaries.
	//
	// Types that are assignable to Source:
	//
	//	*ImportAptArtifactsRequest_GcsSource
	Source isImportAptArtifactsRequest_Source `protobuf_oneof:"source"`
	// The name of the parent resource where the artifacts will be imported.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ImportAptArtifactsRequest) Reset() {
	*x = ImportAptArtifactsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAptArtifactsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAptArtifactsRequest) ProtoMessage() {}

func (x *ImportAptArtifactsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAptArtifactsRequest.ProtoReflect.Descriptor instead.
func (*ImportAptArtifactsRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{2}
}

func (m *ImportAptArtifactsRequest) GetSource() isImportAptArtifactsRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ImportAptArtifactsRequest) GetGcsSource() *ImportAptArtifactsGcsSource {
	if x, ok := x.GetSource().(*ImportAptArtifactsRequest_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (x *ImportAptArtifactsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

type isImportAptArtifactsRequest_Source interface {
	isImportAptArtifactsRequest_Source()
}

type ImportAptArtifactsRequest_GcsSource struct {
	// Google Cloud Storage location where input content is located.
	GcsSource *ImportAptArtifactsGcsSource `protobuf:"bytes,2,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*ImportAptArtifactsRequest_GcsSource) isImportAptArtifactsRequest_Source() {}

// Error information explaining why a package was not imported.
type ImportAptArtifactsErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source that was not imported.
	//
	// Types that are assignable to Source:
	//
	//	*ImportAptArtifactsErrorInfo_GcsSource
	Source isImportAptArtifactsErrorInfo_Source `protobuf_oneof:"source"`
	// The detailed error status.
	Error *status.Status `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ImportAptArtifactsErrorInfo) Reset() {
	*x = ImportAptArtifactsErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAptArtifactsErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAptArtifactsErrorInfo) ProtoMessage() {}

func (x *ImportAptArtifactsErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAptArtifactsErrorInfo.ProtoReflect.Descriptor instead.
func (*ImportAptArtifactsErrorInfo) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{3}
}

func (m *ImportAptArtifactsErrorInfo) GetSource() isImportAptArtifactsErrorInfo_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *ImportAptArtifactsErrorInfo) GetGcsSource() *ImportAptArtifactsGcsSource {
	if x, ok := x.GetSource().(*ImportAptArtifactsErrorInfo_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (x *ImportAptArtifactsErrorInfo) GetError() *status.Status {
	if x != nil {
		return x.Error
	}
	return nil
}

type isImportAptArtifactsErrorInfo_Source interface {
	isImportAptArtifactsErrorInfo_Source()
}

type ImportAptArtifactsErrorInfo_GcsSource struct {
	// Google Cloud Storage location requested.
	GcsSource *ImportAptArtifactsGcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*ImportAptArtifactsErrorInfo_GcsSource) isImportAptArtifactsErrorInfo_Source() {}

// The response message from importing APT artifacts.
type ImportAptArtifactsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Apt artifacts imported.
	AptArtifacts []*AptArtifact `protobuf:"bytes,1,rep,name=apt_artifacts,json=aptArtifacts,proto3" json:"apt_artifacts,omitempty"`
	// Detailed error info for artifacts that were not imported.
	Errors []*ImportAptArtifactsErrorInfo `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ImportAptArtifactsResponse) Reset() {
	*x = ImportAptArtifactsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAptArtifactsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAptArtifactsResponse) ProtoMessage() {}

func (x *ImportAptArtifactsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAptArtifactsResponse.ProtoReflect.Descriptor instead.
func (*ImportAptArtifactsResponse) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{4}
}

func (x *ImportAptArtifactsResponse) GetAptArtifacts() []*AptArtifact {
	if x != nil {
		return x.AptArtifacts
	}
	return nil
}

func (x *ImportAptArtifactsResponse) GetErrors() []*ImportAptArtifactsErrorInfo {
	if x != nil {
		return x.Errors
	}
	return nil
}

// The operation metadata for importing artifacts.
type ImportAptArtifactsMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportAptArtifactsMetadata) Reset() {
	*x = ImportAptArtifactsMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportAptArtifactsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportAptArtifactsMetadata) ProtoMessage() {}

func (x *ImportAptArtifactsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportAptArtifactsMetadata.ProtoReflect.Descriptor instead.
func (*ImportAptArtifactsMetadata) Descriptor() ([]byte, []int) {
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP(), []int{5}
}

var File_google_devtools_artifactregistry_v1beta2_apt_artifact_proto protoreflect.FileDescriptor

var file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x61, 0x70, 0x74, 0x5f, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x04, 0x0a,
	0x0b, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x69, 0x0a,
	0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x41,
	0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x0b,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x50,
	0x41, 0x43, 0x4b, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e,
	0x41, 0x52, 0x59, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10,
	0x02, 0x3a, 0x90, 0x01, 0xea, 0x41, 0x8c, 0x01, 0x0a, 0x2b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x12, 0x5d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x7d, 0x2f, 0x61, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x70, 0x74, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x7d, 0x22, 0x56, 0x0a, 0x1b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x72, 0x69, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x5f, 0x77,
	0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x75, 0x73, 0x65, 0x57, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x73, 0x22, 0xa5, 0x01, 0x0a,
	0x19, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x66, 0x0a, 0x0a, 0x67, 0x63,
	0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x41, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x47, 0x63, 0x73, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x1b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41,
	0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x66, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48,
	0x00, 0x52, 0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x22, 0xd7, 0x01, 0x0a, 0x1a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x0d, 0x61, 0x70, 0x74, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x32, 0x2e, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x0c, 0x61,
	0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x5d, 0x0a, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x94, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x42, 0x10, 0x41, 0x70, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x70,
	0x62, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x70, 0x62, 0xaa, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x32, 0xca, 0x02, 0x25, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x32, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescOnce sync.Once
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescData = file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDesc
)

func file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescGZIP() []byte {
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescOnce.Do(func() {
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescData)
	})
	return file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDescData
}

var file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_goTypes = []interface{}{
	(AptArtifact_PackageType)(0),        // 0: google.devtools.artifactregistry.v1beta2.AptArtifact.PackageType
	(*AptArtifact)(nil),                 // 1: google.devtools.artifactregistry.v1beta2.AptArtifact
	(*ImportAptArtifactsGcsSource)(nil), // 2: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsGcsSource
	(*ImportAptArtifactsRequest)(nil),   // 3: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsRequest
	(*ImportAptArtifactsErrorInfo)(nil), // 4: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsErrorInfo
	(*ImportAptArtifactsResponse)(nil),  // 5: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsResponse
	(*ImportAptArtifactsMetadata)(nil),  // 6: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsMetadata
	(*status.Status)(nil),               // 7: google.rpc.Status
}
var file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_depIdxs = []int32{
	0, // 0: google.devtools.artifactregistry.v1beta2.AptArtifact.package_type:type_name -> google.devtools.artifactregistry.v1beta2.AptArtifact.PackageType
	2, // 1: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsRequest.gcs_source:type_name -> google.devtools.artifactregistry.v1beta2.ImportAptArtifactsGcsSource
	2, // 2: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsErrorInfo.gcs_source:type_name -> google.devtools.artifactregistry.v1beta2.ImportAptArtifactsGcsSource
	7, // 3: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsErrorInfo.error:type_name -> google.rpc.Status
	1, // 4: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsResponse.apt_artifacts:type_name -> google.devtools.artifactregistry.v1beta2.AptArtifact
	4, // 5: google.devtools.artifactregistry.v1beta2.ImportAptArtifactsResponse.errors:type_name -> google.devtools.artifactregistry.v1beta2.ImportAptArtifactsErrorInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_init() }
func file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_init() {
	if File_google_devtools_artifactregistry_v1beta2_apt_artifact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AptArtifact); i {
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
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAptArtifactsGcsSource); i {
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
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAptArtifactsRequest); i {
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
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAptArtifactsErrorInfo); i {
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
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAptArtifactsResponse); i {
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
		file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportAptArtifactsMetadata); i {
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
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ImportAptArtifactsRequest_GcsSource)(nil),
	}
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ImportAptArtifactsErrorInfo_GcsSource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_goTypes,
		DependencyIndexes: file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_depIdxs,
		EnumInfos:         file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_enumTypes,
		MessageInfos:      file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_msgTypes,
	}.Build()
	File_google_devtools_artifactregistry_v1beta2_apt_artifact_proto = out.File
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_rawDesc = nil
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_goTypes = nil
	file_google_devtools_artifactregistry_v1beta2_apt_artifact_proto_depIdxs = nil
}
