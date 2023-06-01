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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: google/cloud/talent/v4/event.proto

package talentpb

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

// An enumeration of an event attributed to the behavior of the end user,
// such as a job seeker.
type JobEvent_JobEventType int32

const (
	// The event is unspecified by other provided values.
	JobEvent_JOB_EVENT_TYPE_UNSPECIFIED JobEvent_JobEventType = 0
	// The job seeker or other entity interacting with the service has
	// had a job rendered in their view, such as in a list of search results in
	// a compressed or clipped format. This event is typically associated with
	// the viewing of a jobs list on a single page by a job seeker.
	JobEvent_IMPRESSION JobEvent_JobEventType = 1
	// The job seeker, or other entity interacting with the service, has
	// viewed the details of a job, including the full description. This
	// event doesn't apply to the viewing a snippet of a job appearing as a
	// part of the job search results. Viewing a snippet is associated with an
	// [impression][google.cloud.talent.v4.JobEvent.JobEventType.IMPRESSION]).
	JobEvent_VIEW JobEvent_JobEventType = 2
	// The job seeker or other entity interacting with the service
	// performed an action to view a job and was redirected to a different
	// website for job.
	JobEvent_VIEW_REDIRECT JobEvent_JobEventType = 3
	// The job seeker or other entity interacting with the service
	// began the process or demonstrated the intention of applying for a job.
	JobEvent_APPLICATION_START JobEvent_JobEventType = 4
	// The job seeker or other entity interacting with the service
	// submitted an application for a job.
	JobEvent_APPLICATION_FINISH JobEvent_JobEventType = 5
	// The job seeker or other entity interacting with the service
	// submitted an application for a job with a single click without
	// entering information. If a job seeker performs this action, send only
	// this event to the service. Do not also send
	// [JobEventType.APPLICATION_START][google.cloud.talent.v4.JobEvent.JobEventType.APPLICATION_START]
	// or
	// [JobEventType.APPLICATION_FINISH][google.cloud.talent.v4.JobEvent.JobEventType.APPLICATION_FINISH]
	// events.
	JobEvent_APPLICATION_QUICK_SUBMISSION JobEvent_JobEventType = 6
	// The job seeker or other entity interacting with the service
	// performed an action to apply to a job and was redirected to a different
	// website to complete the application.
	JobEvent_APPLICATION_REDIRECT JobEvent_JobEventType = 7
	// The job seeker or other entity interacting with the service began the
	// process or demonstrated the intention of applying for a job from the
	// search results page without viewing the details of the job posting.
	// If sending this event, JobEventType.VIEW event shouldn't be sent.
	JobEvent_APPLICATION_START_FROM_SEARCH JobEvent_JobEventType = 8
	// The job seeker, or other entity interacting with the service, performs an
	// action with a single click from the search results page to apply to a job
	// (without viewing the details of the job posting), and is redirected
	// to a different website to complete the application. If a candidate
	// performs this action, send only this event to the service. Do not also
	// send
	// [JobEventType.APPLICATION_START][google.cloud.talent.v4.JobEvent.JobEventType.APPLICATION_START],
	// [JobEventType.APPLICATION_FINISH][google.cloud.talent.v4.JobEvent.JobEventType.APPLICATION_FINISH]
	// or [JobEventType.VIEW][google.cloud.talent.v4.JobEvent.JobEventType.VIEW]
	// events.
	JobEvent_APPLICATION_REDIRECT_FROM_SEARCH JobEvent_JobEventType = 9
	// This event should be used when a company submits an application
	// on behalf of a job seeker. This event is intended for use by staffing
	// agencies attempting to place candidates.
	JobEvent_APPLICATION_COMPANY_SUBMIT JobEvent_JobEventType = 10
	// The job seeker or other entity interacting with the service demonstrated
	// an interest in a job by bookmarking or saving it.
	JobEvent_BOOKMARK JobEvent_JobEventType = 11
	// The job seeker or other entity interacting with the service was
	// sent a notification, such as an email alert or device notification,
	// containing one or more jobs listings generated by the service.
	JobEvent_NOTIFICATION JobEvent_JobEventType = 12
	// The job seeker or other entity interacting with the service was
	// employed by the hiring entity (employer). Send this event
	// only if the job seeker was hired through an application that was
	// initiated by a search conducted through the Cloud Talent Solution
	// service.
	JobEvent_HIRED JobEvent_JobEventType = 13
	// A recruiter or staffing agency submitted an application on behalf of the
	// candidate after interacting with the service to identify a suitable job
	// posting.
	JobEvent_SENT_CV JobEvent_JobEventType = 14
	// The entity interacting with the service (for example, the job seeker),
	// was granted an initial interview by the hiring entity (employer). This
	// event should only be sent if the job seeker was granted an interview as
	// part of an application that was initiated by a search conducted through /
	// recommendation provided by the Cloud Talent Solution service.
	JobEvent_INTERVIEW_GRANTED JobEvent_JobEventType = 15
)

// Enum value maps for JobEvent_JobEventType.
var (
	JobEvent_JobEventType_name = map[int32]string{
		0:  "JOB_EVENT_TYPE_UNSPECIFIED",
		1:  "IMPRESSION",
		2:  "VIEW",
		3:  "VIEW_REDIRECT",
		4:  "APPLICATION_START",
		5:  "APPLICATION_FINISH",
		6:  "APPLICATION_QUICK_SUBMISSION",
		7:  "APPLICATION_REDIRECT",
		8:  "APPLICATION_START_FROM_SEARCH",
		9:  "APPLICATION_REDIRECT_FROM_SEARCH",
		10: "APPLICATION_COMPANY_SUBMIT",
		11: "BOOKMARK",
		12: "NOTIFICATION",
		13: "HIRED",
		14: "SENT_CV",
		15: "INTERVIEW_GRANTED",
	}
	JobEvent_JobEventType_value = map[string]int32{
		"JOB_EVENT_TYPE_UNSPECIFIED":       0,
		"IMPRESSION":                       1,
		"VIEW":                             2,
		"VIEW_REDIRECT":                    3,
		"APPLICATION_START":                4,
		"APPLICATION_FINISH":               5,
		"APPLICATION_QUICK_SUBMISSION":     6,
		"APPLICATION_REDIRECT":             7,
		"APPLICATION_START_FROM_SEARCH":    8,
		"APPLICATION_REDIRECT_FROM_SEARCH": 9,
		"APPLICATION_COMPANY_SUBMIT":       10,
		"BOOKMARK":                         11,
		"NOTIFICATION":                     12,
		"HIRED":                            13,
		"SENT_CV":                          14,
		"INTERVIEW_GRANTED":                15,
	}
)

func (x JobEvent_JobEventType) Enum() *JobEvent_JobEventType {
	p := new(JobEvent_JobEventType)
	*p = x
	return p
}

func (x JobEvent_JobEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobEvent_JobEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_talent_v4_event_proto_enumTypes[0].Descriptor()
}

func (JobEvent_JobEventType) Type() protoreflect.EnumType {
	return &file_google_cloud_talent_v4_event_proto_enumTypes[0]
}

func (x JobEvent_JobEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobEvent_JobEventType.Descriptor instead.
func (JobEvent_JobEventType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_event_proto_rawDescGZIP(), []int{1, 0}
}

// An event issued when an end user interacts with the application that
// implements Cloud Talent Solution. Providing this information improves the
// quality of results for the API clients, enabling the
// service to perform optimally. The number of events sent must be consistent
// with other calls, such as job searches, issued to the service by the client.
type ClientEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Strongly recommended for the best service experience.
	//
	// A unique ID generated in the API responses. It can be found in
	// [ResponseMetadata.request_id][google.cloud.talent.v4.ResponseMetadata.request_id].
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Required. A unique identifier, generated by the client application.
	EventId string `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Required. The timestamp of the event.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Required.
	//
	// The detail information of a specific event type.
	//
	// Types that are assignable to Event:
	//	*ClientEvent_JobEvent
	Event isClientEvent_Event `protobuf_oneof:"event"`
	// Notes about the event provided by recruiters or other users, for example,
	// feedback on why a job was bookmarked.
	EventNotes string `protobuf:"bytes,9,opt,name=event_notes,json=eventNotes,proto3" json:"event_notes,omitempty"`
}

func (x *ClientEvent) Reset() {
	*x = ClientEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientEvent) ProtoMessage() {}

func (x *ClientEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientEvent.ProtoReflect.Descriptor instead.
func (*ClientEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_event_proto_rawDescGZIP(), []int{0}
}

func (x *ClientEvent) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ClientEvent) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *ClientEvent) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (m *ClientEvent) GetEvent() isClientEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *ClientEvent) GetJobEvent() *JobEvent {
	if x, ok := x.GetEvent().(*ClientEvent_JobEvent); ok {
		return x.JobEvent
	}
	return nil
}

func (x *ClientEvent) GetEventNotes() string {
	if x != nil {
		return x.EventNotes
	}
	return ""
}

type isClientEvent_Event interface {
	isClientEvent_Event()
}

type ClientEvent_JobEvent struct {
	// An event issued when a job seeker interacts with the application that
	// implements Cloud Talent Solution.
	JobEvent *JobEvent `protobuf:"bytes,5,opt,name=job_event,json=jobEvent,proto3,oneof"`
}

func (*ClientEvent_JobEvent) isClientEvent_Event() {}

// An event issued when a job seeker interacts with the application that
// implements Cloud Talent Solution.
type JobEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The type of the event (see
	// [JobEventType][google.cloud.talent.v4.JobEvent.JobEventType]).
	Type JobEvent_JobEventType `protobuf:"varint,1,opt,name=type,proto3,enum=google.cloud.talent.v4.JobEvent_JobEventType" json:"type,omitempty"`
	// Required. The [job name(s)][google.cloud.talent.v4.Job.name] associated
	// with this event. For example, if this is an
	// [impression][google.cloud.talent.v4.JobEvent.JobEventType.IMPRESSION]
	// event, this field contains the identifiers of all jobs shown to the job
	// seeker. If this was a
	// [view][google.cloud.talent.v4.JobEvent.JobEventType.VIEW] event, this field
	// contains the identifier of the viewed job.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/jobs/{job_id}", for
	// example, "projects/foo/tenants/bar/jobs/baz".
	Jobs []string `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *JobEvent) Reset() {
	*x = JobEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobEvent) ProtoMessage() {}

func (x *JobEvent) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobEvent.ProtoReflect.Descriptor instead.
func (*JobEvent) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_event_proto_rawDescGZIP(), []int{1}
}

func (x *JobEvent) GetType() JobEvent_JobEventType {
	if x != nil {
		return x.Type
	}
	return JobEvent_JOB_EVENT_TYPE_UNSPECIFIED
}

func (x *JobEvent) GetJobs() []string {
	if x != nil {
		return x.Jobs
	}
	return nil
}

var File_google_cloud_talent_v4_event_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4_event_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9,
	0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x3f, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x4a, 0x6f, 0x62, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x74, 0x65,
	0x73, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xec, 0x03, 0x0a, 0x08, 0x4a,
	0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x4a,
	0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xfe, 0x02, 0x0a, 0x0c, 0x4a, 0x6f, 0x62,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x4a, 0x4f, 0x42,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4d, 0x50,
	0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x52, 0x45, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x04, 0x12, 0x16, 0x0a,
	0x12, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49, 0x4e,
	0x49, 0x53, 0x48, 0x10, 0x05, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x51, 0x55, 0x49, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x10,
	0x07, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x53, 0x45, 0x41, 0x52,
	0x43, 0x48, 0x10, 0x08, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x52, 0x4f,
	0x4d, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x09, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x50,
	0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x4e,
	0x59, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x4f,
	0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x56,
	0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x56, 0x49, 0x45, 0x57, 0x5f,
	0x47, 0x52, 0x41, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0f, 0x42, 0x64, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x53, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_talent_v4_event_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4_event_proto_rawDescData = file_google_cloud_talent_v4_event_proto_rawDesc
)

func file_google_cloud_talent_v4_event_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4_event_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4_event_proto_rawDescData)
	})
	return file_google_cloud_talent_v4_event_proto_rawDescData
}

var file_google_cloud_talent_v4_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_talent_v4_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_talent_v4_event_proto_goTypes = []interface{}{
	(JobEvent_JobEventType)(0),    // 0: google.cloud.talent.v4.JobEvent.JobEventType
	(*ClientEvent)(nil),           // 1: google.cloud.talent.v4.ClientEvent
	(*JobEvent)(nil),              // 2: google.cloud.talent.v4.JobEvent
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_talent_v4_event_proto_depIdxs = []int32{
	3, // 0: google.cloud.talent.v4.ClientEvent.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: google.cloud.talent.v4.ClientEvent.job_event:type_name -> google.cloud.talent.v4.JobEvent
	0, // 2: google.cloud.talent.v4.JobEvent.type:type_name -> google.cloud.talent.v4.JobEvent.JobEventType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4_event_proto_init() }
func file_google_cloud_talent_v4_event_proto_init() {
	if File_google_cloud_talent_v4_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_talent_v4_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientEvent); i {
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
		file_google_cloud_talent_v4_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobEvent); i {
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
	file_google_cloud_talent_v4_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ClientEvent_JobEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_talent_v4_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_talent_v4_event_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4_event_proto_depIdxs,
		EnumInfos:         file_google_cloud_talent_v4_event_proto_enumTypes,
		MessageInfos:      file_google_cloud_talent_v4_event_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4_event_proto = out.File
	file_google_cloud_talent_v4_event_proto_rawDesc = nil
	file_google_cloud_talent_v4_event_proto_goTypes = nil
	file_google_cloud_talent_v4_event_proto_depIdxs = nil
}
