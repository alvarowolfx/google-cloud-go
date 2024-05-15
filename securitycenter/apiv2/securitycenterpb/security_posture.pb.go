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
// source: google/cloud/securitycenter/v2/security_posture.proto

package securitycenterpb

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

// Represents a posture that is deployed on Google Cloud by the
// Security Command Center Posture Management service.
// A posture contains one or more policy sets. A policy set is a
// group of policies that enforce a set of security rules on Google
// Cloud.
type SecurityPosture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the posture, for example, `CIS-Posture`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The version of the posture, for example, `c7cfa2a8`.
	RevisionId string `protobuf:"bytes,2,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// The project, folder, or organization on which the posture is deployed,
	// for example, `projects/{project_number}`.
	PostureDeploymentResource string `protobuf:"bytes,3,opt,name=posture_deployment_resource,json=postureDeploymentResource,proto3" json:"posture_deployment_resource,omitempty"`
	// The name of the posture deployment, for example,
	// `organizations/{org_id}/posturedeployments/{posture_deployment_id}`.
	PostureDeployment string `protobuf:"bytes,4,opt,name=posture_deployment,json=postureDeployment,proto3" json:"posture_deployment,omitempty"`
	// The name of the updated policy, for example,
	// `projects/{project_id}/policies/{constraint_name}`.
	ChangedPolicy string `protobuf:"bytes,5,opt,name=changed_policy,json=changedPolicy,proto3" json:"changed_policy,omitempty"`
	// The name of the updated policy set, for example, `cis-policyset`.
	PolicySet string `protobuf:"bytes,6,opt,name=policy_set,json=policySet,proto3" json:"policy_set,omitempty"`
	// The ID of the updated policy, for example, `compute-policy-1`.
	Policy string `protobuf:"bytes,7,opt,name=policy,proto3" json:"policy,omitempty"`
	// The details about a change in an updated policy that violates the deployed
	// posture.
	PolicyDriftDetails []*SecurityPosture_PolicyDriftDetails `protobuf:"bytes,8,rep,name=policy_drift_details,json=policyDriftDetails,proto3" json:"policy_drift_details,omitempty"`
}

func (x *SecurityPosture) Reset() {
	*x = SecurityPosture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityPosture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityPosture) ProtoMessage() {}

func (x *SecurityPosture) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityPosture.ProtoReflect.Descriptor instead.
func (*SecurityPosture) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_security_posture_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityPosture) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityPosture) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *SecurityPosture) GetPostureDeploymentResource() string {
	if x != nil {
		return x.PostureDeploymentResource
	}
	return ""
}

func (x *SecurityPosture) GetPostureDeployment() string {
	if x != nil {
		return x.PostureDeployment
	}
	return ""
}

func (x *SecurityPosture) GetChangedPolicy() string {
	if x != nil {
		return x.ChangedPolicy
	}
	return ""
}

func (x *SecurityPosture) GetPolicySet() string {
	if x != nil {
		return x.PolicySet
	}
	return ""
}

func (x *SecurityPosture) GetPolicy() string {
	if x != nil {
		return x.Policy
	}
	return ""
}

func (x *SecurityPosture) GetPolicyDriftDetails() []*SecurityPosture_PolicyDriftDetails {
	if x != nil {
		return x.PolicyDriftDetails
	}
	return nil
}

// The policy field that violates the deployed posture and its expected and
// detected values.
type SecurityPosture_PolicyDriftDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the updated field, for example
	// constraint.implementation.policy_rules[0].enforce
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The value of this field that was configured in a posture, for example,
	// `true` or `allowed_values={"projects/29831892"}`.
	ExpectedValue string `protobuf:"bytes,2,opt,name=expected_value,json=expectedValue,proto3" json:"expected_value,omitempty"`
	// The detected value that violates the deployed posture, for example,
	// `false` or `allowed_values={"projects/22831892"}`.
	DetectedValue string `protobuf:"bytes,3,opt,name=detected_value,json=detectedValue,proto3" json:"detected_value,omitempty"`
}

func (x *SecurityPosture_PolicyDriftDetails) Reset() {
	*x = SecurityPosture_PolicyDriftDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityPosture_PolicyDriftDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityPosture_PolicyDriftDetails) ProtoMessage() {}

func (x *SecurityPosture_PolicyDriftDetails) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityPosture_PolicyDriftDetails.ProtoReflect.Descriptor instead.
func (*SecurityPosture_PolicyDriftDetails) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v2_security_posture_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SecurityPosture_PolicyDriftDetails) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *SecurityPosture_PolicyDriftDetails) GetExpectedValue() string {
	if x != nil {
		return x.ExpectedValue
	}
	return ""
}

func (x *SecurityPosture_PolicyDriftDetails) GetDetectedValue() string {
	if x != nil {
		return x.DetectedValue
	}
	return ""
}

var File_google_cloud_securitycenter_v2_security_posture_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v2_security_posture_proto_rawDesc = []byte{
	0x0a, 0x35, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x22, 0x83, 0x04, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x3e, 0x0a, 0x1b, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6f,
	0x73, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x53, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x74, 0x0a,
	0x14, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x64, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x44, 0x72, 0x69, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x72, 0x69, 0x66, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x78, 0x0a, 0x12, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x72, 0x69,
	0x66, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xee, 0x01,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x42, 0x14, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x50, 0x6f,
	0x73, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v2_security_posture_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v2_security_posture_proto_rawDescData = file_google_cloud_securitycenter_v2_security_posture_proto_rawDesc
)

func file_google_cloud_securitycenter_v2_security_posture_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v2_security_posture_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v2_security_posture_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v2_security_posture_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v2_security_posture_proto_rawDescData
}

var file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v2_security_posture_proto_goTypes = []interface{}{
	(*SecurityPosture)(nil),                    // 0: google.cloud.securitycenter.v2.SecurityPosture
	(*SecurityPosture_PolicyDriftDetails)(nil), // 1: google.cloud.securitycenter.v2.SecurityPosture.PolicyDriftDetails
}
var file_google_cloud_securitycenter_v2_security_posture_proto_depIdxs = []int32{
	1, // 0: google.cloud.securitycenter.v2.SecurityPosture.policy_drift_details:type_name -> google.cloud.securitycenter.v2.SecurityPosture.PolicyDriftDetails
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v2_security_posture_proto_init() }
func file_google_cloud_securitycenter_v2_security_posture_proto_init() {
	if File_google_cloud_securitycenter_v2_security_posture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityPosture); i {
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
		file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityPosture_PolicyDriftDetails); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v2_security_posture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v2_security_posture_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v2_security_posture_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v2_security_posture_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v2_security_posture_proto = out.File
	file_google_cloud_securitycenter_v2_security_posture_proto_rawDesc = nil
	file_google_cloud_securitycenter_v2_security_posture_proto_goTypes = nil
	file_google_cloud_securitycenter_v2_security_posture_proto_depIdxs = nil
}
