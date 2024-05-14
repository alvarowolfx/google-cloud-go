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
// source: google/cloud/aiplatform/v1beta1/model_monitoring_alert.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Monitoring alert triggered condition.
type ModelMonitoringAlertCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alert triggered condition.
	//
	// Types that are assignable to Condition:
	//
	//	*ModelMonitoringAlertCondition_Threshold
	Condition isModelMonitoringAlertCondition_Condition `protobuf_oneof:"condition"`
}

func (x *ModelMonitoringAlertCondition) Reset() {
	*x = ModelMonitoringAlertCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMonitoringAlertCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMonitoringAlertCondition) ProtoMessage() {}

func (x *ModelMonitoringAlertCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMonitoringAlertCondition.ProtoReflect.Descriptor instead.
func (*ModelMonitoringAlertCondition) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescGZIP(), []int{0}
}

func (m *ModelMonitoringAlertCondition) GetCondition() isModelMonitoringAlertCondition_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (x *ModelMonitoringAlertCondition) GetThreshold() float64 {
	if x, ok := x.GetCondition().(*ModelMonitoringAlertCondition_Threshold); ok {
		return x.Threshold
	}
	return 0
}

type isModelMonitoringAlertCondition_Condition interface {
	isModelMonitoringAlertCondition_Condition()
}

type ModelMonitoringAlertCondition_Threshold struct {
	// A condition that compares a stats value against a threshold. Alert will
	// be triggered if value above the threshold.
	Threshold float64 `protobuf:"fixed64,1,opt,name=threshold,proto3,oneof"`
}

func (*ModelMonitoringAlertCondition_Threshold) isModelMonitoringAlertCondition_Condition() {}

// Represents a single model monitoring anomaly.
type ModelMonitoringAnomaly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Anomaly:
	//
	//	*ModelMonitoringAnomaly_TabularAnomaly_
	Anomaly isModelMonitoringAnomaly_Anomaly `protobuf_oneof:"anomaly"`
	// Model monitoring job resource name.
	ModelMonitoringJob string `protobuf:"bytes,2,opt,name=model_monitoring_job,json=modelMonitoringJob,proto3" json:"model_monitoring_job,omitempty"`
	// Algorithm used to calculated the metrics, eg: jensen_shannon_divergence,
	// l_infinity.
	Algorithm string `protobuf:"bytes,3,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
}

func (x *ModelMonitoringAnomaly) Reset() {
	*x = ModelMonitoringAnomaly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMonitoringAnomaly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMonitoringAnomaly) ProtoMessage() {}

func (x *ModelMonitoringAnomaly) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMonitoringAnomaly.ProtoReflect.Descriptor instead.
func (*ModelMonitoringAnomaly) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescGZIP(), []int{1}
}

func (m *ModelMonitoringAnomaly) GetAnomaly() isModelMonitoringAnomaly_Anomaly {
	if m != nil {
		return m.Anomaly
	}
	return nil
}

func (x *ModelMonitoringAnomaly) GetTabularAnomaly() *ModelMonitoringAnomaly_TabularAnomaly {
	if x, ok := x.GetAnomaly().(*ModelMonitoringAnomaly_TabularAnomaly_); ok {
		return x.TabularAnomaly
	}
	return nil
}

func (x *ModelMonitoringAnomaly) GetModelMonitoringJob() string {
	if x != nil {
		return x.ModelMonitoringJob
	}
	return ""
}

func (x *ModelMonitoringAnomaly) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

type isModelMonitoringAnomaly_Anomaly interface {
	isModelMonitoringAnomaly_Anomaly()
}

type ModelMonitoringAnomaly_TabularAnomaly_ struct {
	// Tabular anomaly.
	TabularAnomaly *ModelMonitoringAnomaly_TabularAnomaly `protobuf:"bytes,1,opt,name=tabular_anomaly,json=tabularAnomaly,proto3,oneof"`
}

func (*ModelMonitoringAnomaly_TabularAnomaly_) isModelMonitoringAnomaly_Anomaly() {}

// Represents a single monitoring alert. This is currently used in the
// SearchModelMonitoringAlerts api, thus the alert wrapped in this message
// belongs to the resource asked in the request.
type ModelMonitoringAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stats name.
	StatsName string `protobuf:"bytes,1,opt,name=stats_name,json=statsName,proto3" json:"stats_name,omitempty"`
	// One of the supported monitoring objectives:
	// `raw-feature-drift`
	// `prediction-output-drift`
	// `feature-attribution`
	ObjectiveType string `protobuf:"bytes,2,opt,name=objective_type,json=objectiveType,proto3" json:"objective_type,omitempty"`
	// Alert creation time.
	AlertTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=alert_time,json=alertTime,proto3" json:"alert_time,omitempty"`
	// Anomaly details.
	Anomaly *ModelMonitoringAnomaly `protobuf:"bytes,4,opt,name=anomaly,proto3" json:"anomaly,omitempty"`
}

func (x *ModelMonitoringAlert) Reset() {
	*x = ModelMonitoringAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMonitoringAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMonitoringAlert) ProtoMessage() {}

func (x *ModelMonitoringAlert) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMonitoringAlert.ProtoReflect.Descriptor instead.
func (*ModelMonitoringAlert) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescGZIP(), []int{2}
}

func (x *ModelMonitoringAlert) GetStatsName() string {
	if x != nil {
		return x.StatsName
	}
	return ""
}

func (x *ModelMonitoringAlert) GetObjectiveType() string {
	if x != nil {
		return x.ObjectiveType
	}
	return ""
}

func (x *ModelMonitoringAlert) GetAlertTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AlertTime
	}
	return nil
}

func (x *ModelMonitoringAlert) GetAnomaly() *ModelMonitoringAnomaly {
	if x != nil {
		return x.Anomaly
	}
	return nil
}

// Tabular anomaly details.
type ModelMonitoringAnomaly_TabularAnomaly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Additional anomaly information. e.g. Google Cloud Storage uri.
	AnomalyUri string `protobuf:"bytes,1,opt,name=anomaly_uri,json=anomalyUri,proto3" json:"anomaly_uri,omitempty"`
	// Overview of this anomaly.
	Summary string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	// Anomaly body.
	Anomaly *structpb.Value `protobuf:"bytes,3,opt,name=anomaly,proto3" json:"anomaly,omitempty"`
	// The time the anomaly was triggered.
	TriggerTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	// The alert condition associated with this anomaly.
	Condition *ModelMonitoringAlertCondition `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) Reset() {
	*x = ModelMonitoringAnomaly_TabularAnomaly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelMonitoringAnomaly_TabularAnomaly) ProtoMessage() {}

func (x *ModelMonitoringAnomaly_TabularAnomaly) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelMonitoringAnomaly_TabularAnomaly.ProtoReflect.Descriptor instead.
func (*ModelMonitoringAnomaly_TabularAnomaly) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) GetAnomalyUri() string {
	if x != nil {
		return x.AnomalyUri
	}
	return ""
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) GetAnomaly() *structpb.Value {
	if x != nil {
		return x.Anomaly
	}
	return nil
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) GetTriggerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TriggerTime
	}
	return nil
}

func (x *ModelMonitoringAnomaly_TabularAnomaly) GetCondition() *ModelMonitoringAlertCondition {
	if x != nil {
		return x.Condition
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c,
	0x0a, 0x1d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x83, 0x04, 0x0a,
	0x16, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12, 0x71, 0x0a, 0x0f, 0x74, 0x61, 0x62, 0x75, 0x6c,
	0x61, 0x72, 0x5f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x2e, 0x54, 0x61, 0x62, 0x75, 0x6c, 0x61,
	0x72, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x61, 0x62, 0x75,
	0x6c, 0x61, 0x72, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6a,
	0x6f, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x4a, 0x6f, 0x62, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x1a, 0x9a, 0x02, 0x0a, 0x0e, 0x54,
	0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x55, 0x72, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x6e, 0x6f, 0x6d,
	0x61, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x61, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x79, 0x22, 0xea, 0x01, 0x0a, 0x14, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x07,
	0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x41,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x42,
	0xf0, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x19, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41,
	0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_goTypes = []interface{}{
	(*ModelMonitoringAlertCondition)(nil),         // 0: google.cloud.aiplatform.v1beta1.ModelMonitoringAlertCondition
	(*ModelMonitoringAnomaly)(nil),                // 1: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly
	(*ModelMonitoringAlert)(nil),                  // 2: google.cloud.aiplatform.v1beta1.ModelMonitoringAlert
	(*ModelMonitoringAnomaly_TabularAnomaly)(nil), // 3: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.TabularAnomaly
	(*timestamppb.Timestamp)(nil),                 // 4: google.protobuf.Timestamp
	(*structpb.Value)(nil),                        // 5: google.protobuf.Value
}
var file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_depIdxs = []int32{
	3, // 0: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.tabular_anomaly:type_name -> google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.TabularAnomaly
	4, // 1: google.cloud.aiplatform.v1beta1.ModelMonitoringAlert.alert_time:type_name -> google.protobuf.Timestamp
	1, // 2: google.cloud.aiplatform.v1beta1.ModelMonitoringAlert.anomaly:type_name -> google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly
	5, // 3: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.TabularAnomaly.anomaly:type_name -> google.protobuf.Value
	4, // 4: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.TabularAnomaly.trigger_time:type_name -> google.protobuf.Timestamp
	0, // 5: google.cloud.aiplatform.v1beta1.ModelMonitoringAnomaly.TabularAnomaly.condition:type_name -> google.cloud.aiplatform.v1beta1.ModelMonitoringAlertCondition
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_init() }
func file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMonitoringAlertCondition); i {
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
		file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMonitoringAnomaly); i {
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
		file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMonitoringAlert); i {
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
		file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelMonitoringAnomaly_TabularAnomaly); i {
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
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ModelMonitoringAlertCondition_Threshold)(nil),
	}
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ModelMonitoringAnomaly_TabularAnomaly_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto = out.File
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_model_monitoring_alert_proto_depIdxs = nil
}
