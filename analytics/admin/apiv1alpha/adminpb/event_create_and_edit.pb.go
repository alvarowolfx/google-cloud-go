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
// source: google/analytics/admin/v1alpha/event_create_and_edit.proto

package adminpb

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

// Comparison type for matching condition
type MatchingCondition_ComparisonType int32

const (
	// Unknown
	MatchingCondition_COMPARISON_TYPE_UNSPECIFIED MatchingCondition_ComparisonType = 0
	// Equals, case sensitive
	MatchingCondition_EQUALS MatchingCondition_ComparisonType = 1
	// Equals, case insensitive
	MatchingCondition_EQUALS_CASE_INSENSITIVE MatchingCondition_ComparisonType = 2
	// Contains, case sensitive
	MatchingCondition_CONTAINS MatchingCondition_ComparisonType = 3
	// Contains, case insensitive
	MatchingCondition_CONTAINS_CASE_INSENSITIVE MatchingCondition_ComparisonType = 4
	// Starts with, case sensitive
	MatchingCondition_STARTS_WITH MatchingCondition_ComparisonType = 5
	// Starts with, case insensitive
	MatchingCondition_STARTS_WITH_CASE_INSENSITIVE MatchingCondition_ComparisonType = 6
	// Ends with, case sensitive
	MatchingCondition_ENDS_WITH MatchingCondition_ComparisonType = 7
	// Ends with, case insensitive
	MatchingCondition_ENDS_WITH_CASE_INSENSITIVE MatchingCondition_ComparisonType = 8
	// Greater than
	MatchingCondition_GREATER_THAN MatchingCondition_ComparisonType = 9
	// Greater than or equal
	MatchingCondition_GREATER_THAN_OR_EQUAL MatchingCondition_ComparisonType = 10
	// Less than
	MatchingCondition_LESS_THAN MatchingCondition_ComparisonType = 11
	// Less than or equal
	MatchingCondition_LESS_THAN_OR_EQUAL MatchingCondition_ComparisonType = 12
	// regular expression. Only supported for web streams.
	MatchingCondition_REGULAR_EXPRESSION MatchingCondition_ComparisonType = 13
	// regular expression, case insensitive. Only supported for web streams.
	MatchingCondition_REGULAR_EXPRESSION_CASE_INSENSITIVE MatchingCondition_ComparisonType = 14
)

// Enum value maps for MatchingCondition_ComparisonType.
var (
	MatchingCondition_ComparisonType_name = map[int32]string{
		0:  "COMPARISON_TYPE_UNSPECIFIED",
		1:  "EQUALS",
		2:  "EQUALS_CASE_INSENSITIVE",
		3:  "CONTAINS",
		4:  "CONTAINS_CASE_INSENSITIVE",
		5:  "STARTS_WITH",
		6:  "STARTS_WITH_CASE_INSENSITIVE",
		7:  "ENDS_WITH",
		8:  "ENDS_WITH_CASE_INSENSITIVE",
		9:  "GREATER_THAN",
		10: "GREATER_THAN_OR_EQUAL",
		11: "LESS_THAN",
		12: "LESS_THAN_OR_EQUAL",
		13: "REGULAR_EXPRESSION",
		14: "REGULAR_EXPRESSION_CASE_INSENSITIVE",
	}
	MatchingCondition_ComparisonType_value = map[string]int32{
		"COMPARISON_TYPE_UNSPECIFIED":         0,
		"EQUALS":                              1,
		"EQUALS_CASE_INSENSITIVE":             2,
		"CONTAINS":                            3,
		"CONTAINS_CASE_INSENSITIVE":           4,
		"STARTS_WITH":                         5,
		"STARTS_WITH_CASE_INSENSITIVE":        6,
		"ENDS_WITH":                           7,
		"ENDS_WITH_CASE_INSENSITIVE":          8,
		"GREATER_THAN":                        9,
		"GREATER_THAN_OR_EQUAL":               10,
		"LESS_THAN":                           11,
		"LESS_THAN_OR_EQUAL":                  12,
		"REGULAR_EXPRESSION":                  13,
		"REGULAR_EXPRESSION_CASE_INSENSITIVE": 14,
	}
)

func (x MatchingCondition_ComparisonType) Enum() *MatchingCondition_ComparisonType {
	p := new(MatchingCondition_ComparisonType)
	*p = x
	return p
}

func (x MatchingCondition_ComparisonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchingCondition_ComparisonType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_enumTypes[0].Descriptor()
}

func (MatchingCondition_ComparisonType) Type() protoreflect.EnumType {
	return &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_enumTypes[0]
}

func (x MatchingCondition_ComparisonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchingCondition_ComparisonType.Descriptor instead.
func (MatchingCondition_ComparisonType) EnumDescriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescGZIP(), []int{2, 0}
}

// Defines an event parameter to mutate.
type ParameterMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the parameter to mutate.
	// This value must:
	// * be less than 40 characters.
	// * be unique across across all mutations within the rule
	// * consist only of letters, digits or _ (underscores)
	// For event edit rules, the name may also be set to 'event_name' to modify
	// the event_name in place.
	Parameter string `protobuf:"bytes,1,opt,name=parameter,proto3" json:"parameter,omitempty"`
	// Required. The value mutation to perform.
	// * Must be less than 100 characters.
	// * To specify a constant value for the param, use the value's string.
	// * To copy value from another parameter, use syntax like
	// "[[other_parameter]]" For more details, see this [help center
	// article](https://support.google.com/analytics/answer/10085872#modify-an-event&zippy=%2Cin-this-article%2Cmodify-parameters).
	ParameterValue string `protobuf:"bytes,2,opt,name=parameter_value,json=parameterValue,proto3" json:"parameter_value,omitempty"`
}

func (x *ParameterMutation) Reset() {
	*x = ParameterMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterMutation) ProtoMessage() {}

func (x *ParameterMutation) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterMutation.ProtoReflect.Descriptor instead.
func (*ParameterMutation) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescGZIP(), []int{0}
}

func (x *ParameterMutation) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *ParameterMutation) GetParameterValue() string {
	if x != nil {
		return x.ParameterValue
	}
	return ""
}

// An Event Create Rule defines conditions that will trigger the creation
// of an entirely new event based upon matched criteria of a source event.
// Additional mutations of the parameters from the source event can be defined.
//
// Unlike Event Edit rules, Event Creation Rules have no defined order.  They
// will all be run independently.
//
// Event Edit and Event Create rules can't be used to modify an event created
// from an Event Create rule.
type EventCreateRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name for this EventCreateRule resource.
	// Format:
	// properties/{property}/dataStreams/{data_stream}/eventCreateRules/{event_create_rule}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the new event to be created.
	//
	// This value must:
	// * be less than 40 characters
	// * consist only of letters, digits or _ (underscores)
	// * start with a letter
	DestinationEvent string `protobuf:"bytes,2,opt,name=destination_event,json=destinationEvent,proto3" json:"destination_event,omitempty"`
	// Required. Must have at least one condition, and can have up to 10 max.
	// Conditions on the source event must match for this rule to be applied.
	EventConditions []*MatchingCondition `protobuf:"bytes,3,rep,name=event_conditions,json=eventConditions,proto3" json:"event_conditions,omitempty"`
	// If true, the source parameters are copied to the new event.
	// If false, or unset, all non-internal parameters are not copied from the
	// source event. Parameter mutations are applied after the parameters have
	// been copied.
	SourceCopyParameters bool `protobuf:"varint,4,opt,name=source_copy_parameters,json=sourceCopyParameters,proto3" json:"source_copy_parameters,omitempty"`
	// Parameter mutations define parameter behavior on the new event, and
	// are applied in order.
	// A maximum of 20 mutations can be applied.
	ParameterMutations []*ParameterMutation `protobuf:"bytes,5,rep,name=parameter_mutations,json=parameterMutations,proto3" json:"parameter_mutations,omitempty"`
}

func (x *EventCreateRule) Reset() {
	*x = EventCreateRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateRule) ProtoMessage() {}

func (x *EventCreateRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCreateRule.ProtoReflect.Descriptor instead.
func (*EventCreateRule) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescGZIP(), []int{1}
}

func (x *EventCreateRule) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EventCreateRule) GetDestinationEvent() string {
	if x != nil {
		return x.DestinationEvent
	}
	return ""
}

func (x *EventCreateRule) GetEventConditions() []*MatchingCondition {
	if x != nil {
		return x.EventConditions
	}
	return nil
}

func (x *EventCreateRule) GetSourceCopyParameters() bool {
	if x != nil {
		return x.SourceCopyParameters
	}
	return false
}

func (x *EventCreateRule) GetParameterMutations() []*ParameterMutation {
	if x != nil {
		return x.ParameterMutations
	}
	return nil
}

// Defines a condition for when an Event Edit or Event Creation rule applies to
// an event.
type MatchingCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the field that is compared against for the condition.
	// If 'event_name' is specified this condition will apply to the name of the
	// event.  Otherwise the condition will apply to a parameter with the
	// specified name.
	//
	// This value cannot contain spaces.
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// Required. The type of comparison to be applied to the value.
	ComparisonType MatchingCondition_ComparisonType `protobuf:"varint,2,opt,name=comparison_type,json=comparisonType,proto3,enum=google.analytics.admin.v1alpha.MatchingCondition_ComparisonType" json:"comparison_type,omitempty"`
	// Required. The value being compared against for this condition.  The runtime
	// implementation may perform type coercion of this value to evaluate this
	// condition based on the type of the parameter value.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Whether or not the result of the comparison should be negated. For example,
	// if `negated` is true, then 'equals' comparisons would function as 'not
	// equals'.
	Negated bool `protobuf:"varint,4,opt,name=negated,proto3" json:"negated,omitempty"`
}

func (x *MatchingCondition) Reset() {
	*x = MatchingCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchingCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchingCondition) ProtoMessage() {}

func (x *MatchingCondition) ProtoReflect() protoreflect.Message {
	mi := &file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchingCondition.ProtoReflect.Descriptor instead.
func (*MatchingCondition) Descriptor() ([]byte, []int) {
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescGZIP(), []int{2}
}

func (x *MatchingCondition) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *MatchingCondition) GetComparisonType() MatchingCondition_ComparisonType {
	if x != nil {
		return x.ComparisonType
	}
	return MatchingCondition_COMPARISON_TYPE_UNSPECIFIED
}

func (x *MatchingCondition) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MatchingCondition) GetNegated() bool {
	if x != nil {
		return x.Negated
	}
	return false
}

var File_google_analytics_admin_v1alpha_event_create_and_edit_proto protoreflect.FileDescriptor

var file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x6e,
	0x64, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe5,
	0x03, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x11, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x10, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a,
	0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x34, 0x0a, 0x16, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x70, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x62, 0x0a, 0x13, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x89, 0x01, 0xea, 0x41, 0x85,
	0x01, 0x0a, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x54, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x7d, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x7d, 0x22, 0xd4, 0x04, 0x0a, 0x11, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x05,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x6e, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69,
	0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x22, 0xfe, 0x02, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x52, 0x49, 0x53, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45,
	0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e,
	0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x54, 0x41,
	0x49, 0x4e, 0x53, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x4e, 0x53, 0x49,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x52, 0x54, 0x53,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45,
	0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x44,
	0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x4e, 0x44, 0x53,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x4e,
	0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x10, 0x09, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x48, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51,
	0x55, 0x41, 0x4c, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48,
	0x41, 0x4e, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x54, 0x48, 0x41,
	0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x10, 0x0d, 0x12, 0x27, 0x0a, 0x23, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f,
	0x45, 0x58, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x41, 0x53, 0x45, 0x5f,
	0x49, 0x4e, 0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x0e, 0x42, 0x66, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x62, 0x3b, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescOnce sync.Once
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescData = file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDesc
)

func file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescGZIP() []byte {
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescOnce.Do(func() {
		file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescData)
	})
	return file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDescData
}

var file_google_analytics_admin_v1alpha_event_create_and_edit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_analytics_admin_v1alpha_event_create_and_edit_proto_goTypes = []interface{}{
	(MatchingCondition_ComparisonType)(0), // 0: google.analytics.admin.v1alpha.MatchingCondition.ComparisonType
	(*ParameterMutation)(nil),             // 1: google.analytics.admin.v1alpha.ParameterMutation
	(*EventCreateRule)(nil),               // 2: google.analytics.admin.v1alpha.EventCreateRule
	(*MatchingCondition)(nil),             // 3: google.analytics.admin.v1alpha.MatchingCondition
}
var file_google_analytics_admin_v1alpha_event_create_and_edit_proto_depIdxs = []int32{
	3, // 0: google.analytics.admin.v1alpha.EventCreateRule.event_conditions:type_name -> google.analytics.admin.v1alpha.MatchingCondition
	1, // 1: google.analytics.admin.v1alpha.EventCreateRule.parameter_mutations:type_name -> google.analytics.admin.v1alpha.ParameterMutation
	0, // 2: google.analytics.admin.v1alpha.MatchingCondition.comparison_type:type_name -> google.analytics.admin.v1alpha.MatchingCondition.ComparisonType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_analytics_admin_v1alpha_event_create_and_edit_proto_init() }
func file_google_analytics_admin_v1alpha_event_create_and_edit_proto_init() {
	if File_google_analytics_admin_v1alpha_event_create_and_edit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterMutation); i {
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
		file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateRule); i {
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
		file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchingCondition); i {
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
			RawDescriptor: file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_analytics_admin_v1alpha_event_create_and_edit_proto_goTypes,
		DependencyIndexes: file_google_analytics_admin_v1alpha_event_create_and_edit_proto_depIdxs,
		EnumInfos:         file_google_analytics_admin_v1alpha_event_create_and_edit_proto_enumTypes,
		MessageInfos:      file_google_analytics_admin_v1alpha_event_create_and_edit_proto_msgTypes,
	}.Build()
	File_google_analytics_admin_v1alpha_event_create_and_edit_proto = out.File
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_rawDesc = nil
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_goTypes = nil
	file_google_analytics_admin_v1alpha_event_create_and_edit_proto_depIdxs = nil
}
