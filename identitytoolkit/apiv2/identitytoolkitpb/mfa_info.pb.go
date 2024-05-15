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
// source: google/cloud/identitytoolkit/v2/mfa_info.proto

package identitytoolkitpb

import (
	reflect "reflect"
	sync "sync"

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

// The information required to auto-retrieve an SMS.
type AutoRetrievalInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Android app's signature hash for Google Play Service's
	// SMS Retriever API.
	AppSignatureHash string `protobuf:"bytes,1,opt,name=app_signature_hash,json=appSignatureHash,proto3" json:"app_signature_hash,omitempty"`
}

func (x *AutoRetrievalInfo) Reset() {
	*x = AutoRetrievalInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoRetrievalInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoRetrievalInfo) ProtoMessage() {}

func (x *AutoRetrievalInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoRetrievalInfo.ProtoReflect.Descriptor instead.
func (*AutoRetrievalInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP(), []int{0}
}

func (x *AutoRetrievalInfo) GetAppSignatureHash() string {
	if x != nil {
		return x.AppSignatureHash
	}
	return ""
}

// App Verification info for a StartMfa request.
type StartMfaPhoneRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required for enrollment. Phone number to be enrolled as MFA.
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// iOS only. Receipt of successful app token validation with APNS.
	IosReceipt string `protobuf:"bytes,2,opt,name=ios_receipt,json=iosReceipt,proto3" json:"ios_receipt,omitempty"`
	// iOS only. Secret delivered to iOS app via APNS.
	IosSecret string `protobuf:"bytes,3,opt,name=ios_secret,json=iosSecret,proto3" json:"ios_secret,omitempty"`
	// Web only. Recaptcha solution.
	RecaptchaToken string `protobuf:"bytes,4,opt,name=recaptcha_token,json=recaptchaToken,proto3" json:"recaptcha_token,omitempty"`
	// Android only. Used by Google Play Services to identify the app for
	// auto-retrieval.
	AutoRetrievalInfo *AutoRetrievalInfo `protobuf:"bytes,5,opt,name=auto_retrieval_info,json=autoRetrievalInfo,proto3" json:"auto_retrieval_info,omitempty"`
	// Android only. Used to assert application identity in place of a
	// recaptcha token. A SafetyNet Token can be generated via the
	// [SafetyNet Android Attestation
	// API](https://developer.android.com/training/safetynet/attestation.html),
	// with the Base64 encoding of the `phone_number` field as the nonce.
	SafetyNetToken string `protobuf:"bytes,6,opt,name=safety_net_token,json=safetyNetToken,proto3" json:"safety_net_token,omitempty"`
}

func (x *StartMfaPhoneRequestInfo) Reset() {
	*x = StartMfaPhoneRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMfaPhoneRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMfaPhoneRequestInfo) ProtoMessage() {}

func (x *StartMfaPhoneRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMfaPhoneRequestInfo.ProtoReflect.Descriptor instead.
func (*StartMfaPhoneRequestInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP(), []int{1}
}

func (x *StartMfaPhoneRequestInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *StartMfaPhoneRequestInfo) GetIosReceipt() string {
	if x != nil {
		return x.IosReceipt
	}
	return ""
}

func (x *StartMfaPhoneRequestInfo) GetIosSecret() string {
	if x != nil {
		return x.IosSecret
	}
	return ""
}

func (x *StartMfaPhoneRequestInfo) GetRecaptchaToken() string {
	if x != nil {
		return x.RecaptchaToken
	}
	return ""
}

func (x *StartMfaPhoneRequestInfo) GetAutoRetrievalInfo() *AutoRetrievalInfo {
	if x != nil {
		return x.AutoRetrievalInfo
	}
	return nil
}

func (x *StartMfaPhoneRequestInfo) GetSafetyNetToken() string {
	if x != nil {
		return x.SafetyNetToken
	}
	return ""
}

// Phone Verification info for a StartMfa response.
type StartMfaPhoneResponseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque string that represents the enrollment session.
	SessionInfo string `protobuf:"bytes,1,opt,name=session_info,json=sessionInfo,proto3" json:"session_info,omitempty"`
}

func (x *StartMfaPhoneResponseInfo) Reset() {
	*x = StartMfaPhoneResponseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartMfaPhoneResponseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartMfaPhoneResponseInfo) ProtoMessage() {}

func (x *StartMfaPhoneResponseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartMfaPhoneResponseInfo.ProtoReflect.Descriptor instead.
func (*StartMfaPhoneResponseInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP(), []int{2}
}

func (x *StartMfaPhoneResponseInfo) GetSessionInfo() string {
	if x != nil {
		return x.SessionInfo
	}
	return ""
}

// Phone Verification info for a FinalizeMfa request.
type FinalizeMfaPhoneRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque string that represents the enrollment session.
	SessionInfo string `protobuf:"bytes,1,opt,name=session_info,json=sessionInfo,proto3" json:"session_info,omitempty"`
	// User-entered verification code.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// Android only. Uses for "instant" phone number verification though GmsCore.
	AndroidVerificationProof string `protobuf:"bytes,3,opt,name=android_verification_proof,json=androidVerificationProof,proto3" json:"android_verification_proof,omitempty"`
	// Required if Android verification proof is presented.
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *FinalizeMfaPhoneRequestInfo) Reset() {
	*x = FinalizeMfaPhoneRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeMfaPhoneRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeMfaPhoneRequestInfo) ProtoMessage() {}

func (x *FinalizeMfaPhoneRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeMfaPhoneRequestInfo.ProtoReflect.Descriptor instead.
func (*FinalizeMfaPhoneRequestInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP(), []int{3}
}

func (x *FinalizeMfaPhoneRequestInfo) GetSessionInfo() string {
	if x != nil {
		return x.SessionInfo
	}
	return ""
}

func (x *FinalizeMfaPhoneRequestInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FinalizeMfaPhoneRequestInfo) GetAndroidVerificationProof() string {
	if x != nil {
		return x.AndroidVerificationProof
	}
	return ""
}

func (x *FinalizeMfaPhoneRequestInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

// Phone Verification info for a FinalizeMfa response.
type FinalizeMfaPhoneResponseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Android only. Long-lived replacement for valid code tied to android device.
	AndroidVerificationProof string `protobuf:"bytes,1,opt,name=android_verification_proof,json=androidVerificationProof,proto3" json:"android_verification_proof,omitempty"`
	// Android only. Expiration time of verification proof in seconds.
	AndroidVerificationProofExpireTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=android_verification_proof_expire_time,json=androidVerificationProofExpireTime,proto3" json:"android_verification_proof_expire_time,omitempty"`
	// For Android verification proof.
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *FinalizeMfaPhoneResponseInfo) Reset() {
	*x = FinalizeMfaPhoneResponseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeMfaPhoneResponseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeMfaPhoneResponseInfo) ProtoMessage() {}

func (x *FinalizeMfaPhoneResponseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeMfaPhoneResponseInfo.ProtoReflect.Descriptor instead.
func (*FinalizeMfaPhoneResponseInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP(), []int{4}
}

func (x *FinalizeMfaPhoneResponseInfo) GetAndroidVerificationProof() string {
	if x != nil {
		return x.AndroidVerificationProof
	}
	return ""
}

func (x *FinalizeMfaPhoneResponseInfo) GetAndroidVerificationProofExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AndroidVerificationProofExpireTime
	}
	return nil
}

func (x *FinalizeMfaPhoneResponseInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

var File_google_cloud_identitytoolkit_v2_mfa_info_proto protoreflect.FileDescriptor

var file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x2f, 0x6d, 0x66, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76,
	0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0xb4, 0x02, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d,
	0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6f, 0x73, 0x5f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6f, 0x73, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6f, 0x73, 0x5f, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6f, 0x73, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x72, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x62,
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x11, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x5f, 0x6e, 0x65, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61,
	0x66, 0x65, 0x74, 0x79, 0x4e, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x19,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x1b, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x1a, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0xef, 0x01, 0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x4d, 0x66, 0x61, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x1a, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x6e, 0x64, 0x72, 0x6f,
	0x69, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x12, 0x6e, 0x0a, 0x26, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x22, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0xdf, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x50, 0x01,
	0x5a, 0x4d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f,
	0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x70, 0x62, 0x3b, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x74, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x70, 0x62, 0xaa,
	0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74, 0x2e, 0x56,
	0x32, 0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x6b, 0x69, 0x74,
	0x5c, 0x56, 0x32, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f,
	0x6c, 0x6b, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescOnce sync.Once
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescData = file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDesc
)

func file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescGZIP() []byte {
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescOnce.Do(func() {
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescData)
	})
	return file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDescData
}

var file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_identitytoolkit_v2_mfa_info_proto_goTypes = []interface{}{
	(*AutoRetrievalInfo)(nil),            // 0: google.cloud.identitytoolkit.v2.AutoRetrievalInfo
	(*StartMfaPhoneRequestInfo)(nil),     // 1: google.cloud.identitytoolkit.v2.StartMfaPhoneRequestInfo
	(*StartMfaPhoneResponseInfo)(nil),    // 2: google.cloud.identitytoolkit.v2.StartMfaPhoneResponseInfo
	(*FinalizeMfaPhoneRequestInfo)(nil),  // 3: google.cloud.identitytoolkit.v2.FinalizeMfaPhoneRequestInfo
	(*FinalizeMfaPhoneResponseInfo)(nil), // 4: google.cloud.identitytoolkit.v2.FinalizeMfaPhoneResponseInfo
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
}
var file_google_cloud_identitytoolkit_v2_mfa_info_proto_depIdxs = []int32{
	0, // 0: google.cloud.identitytoolkit.v2.StartMfaPhoneRequestInfo.auto_retrieval_info:type_name -> google.cloud.identitytoolkit.v2.AutoRetrievalInfo
	5, // 1: google.cloud.identitytoolkit.v2.FinalizeMfaPhoneResponseInfo.android_verification_proof_expire_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_identitytoolkit_v2_mfa_info_proto_init() }
func file_google_cloud_identitytoolkit_v2_mfa_info_proto_init() {
	if File_google_cloud_identitytoolkit_v2_mfa_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoRetrievalInfo); i {
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
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMfaPhoneRequestInfo); i {
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
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartMfaPhoneResponseInfo); i {
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
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeMfaPhoneRequestInfo); i {
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
		file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeMfaPhoneResponseInfo); i {
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
			RawDescriptor: file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_identitytoolkit_v2_mfa_info_proto_goTypes,
		DependencyIndexes: file_google_cloud_identitytoolkit_v2_mfa_info_proto_depIdxs,
		MessageInfos:      file_google_cloud_identitytoolkit_v2_mfa_info_proto_msgTypes,
	}.Build()
	File_google_cloud_identitytoolkit_v2_mfa_info_proto = out.File
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_rawDesc = nil
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_goTypes = nil
	file_google_cloud_identitytoolkit_v2_mfa_info_proto_depIdxs = nil
}
