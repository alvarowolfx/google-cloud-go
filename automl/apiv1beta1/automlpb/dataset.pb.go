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
// source: google/cloud/automl/v1beta1/dataset.proto

package automlpb

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

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are assignable to DatasetMetadata:
	//
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	//	*Dataset_ImageObjectDetectionDatasetMetadata
	//	*Dataset_VideoClassificationDatasetMetadata
	//	*Dataset_VideoObjectTrackingDatasetMetadata
	//	*Dataset_TextExtractionDatasetMetadata
	//	*Dataset_TextSentimentDatasetMetadata
	//	*Dataset_TablesDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User-provided description of the dataset. The description can be up to
	// 25000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,17,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Dataset) Reset() {
	*x = Dataset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1beta1_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dataset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dataset) ProtoMessage() {}

func (x *Dataset) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1beta1_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dataset.ProtoReflect.Descriptor instead.
func (*Dataset) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1beta1_dataset_proto_rawDescGZIP(), []int{0}
}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (x *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetImageObjectDetectionDatasetMetadata() *ImageObjectDetectionDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_ImageObjectDetectionDatasetMetadata); ok {
		return x.ImageObjectDetectionDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetVideoClassificationDatasetMetadata() *VideoClassificationDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_VideoClassificationDatasetMetadata); ok {
		return x.VideoClassificationDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetVideoObjectTrackingDatasetMetadata() *VideoObjectTrackingDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_VideoObjectTrackingDatasetMetadata); ok {
		return x.VideoObjectTrackingDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetTextExtractionDatasetMetadata() *TextExtractionDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_TextExtractionDatasetMetadata); ok {
		return x.TextExtractionDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetTextSentimentDatasetMetadata() *TextSentimentDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_TextSentimentDatasetMetadata); ok {
		return x.TextSentimentDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetTablesDatasetMetadata() *TablesDatasetMetadata {
	if x, ok := x.GetDatasetMetadata().(*Dataset_TablesDatasetMetadata); ok {
		return x.TablesDatasetMetadata
	}
	return nil
}

func (x *Dataset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dataset) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Dataset) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dataset) GetExampleCount() int32 {
	if x != nil {
		return x.ExampleCount
	}
	return 0
}

func (x *Dataset) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Dataset) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	// Metadata for a dataset used for translation.
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	// Metadata for a dataset used for image classification.
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	// Metadata for a dataset used for text classification.
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageObjectDetectionDatasetMetadata struct {
	// Metadata for a dataset used for image object detection.
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `protobuf:"bytes,26,opt,name=image_object_detection_dataset_metadata,json=imageObjectDetectionDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoClassificationDatasetMetadata struct {
	// Metadata for a dataset used for video classification.
	VideoClassificationDatasetMetadata *VideoClassificationDatasetMetadata `protobuf:"bytes,31,opt,name=video_classification_dataset_metadata,json=videoClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoObjectTrackingDatasetMetadata struct {
	// Metadata for a dataset used for video object tracking.
	VideoObjectTrackingDatasetMetadata *VideoObjectTrackingDatasetMetadata `protobuf:"bytes,29,opt,name=video_object_tracking_dataset_metadata,json=videoObjectTrackingDatasetMetadata,proto3,oneof"`
}

type Dataset_TextExtractionDatasetMetadata struct {
	// Metadata for a dataset used for text extraction.
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `protobuf:"bytes,28,opt,name=text_extraction_dataset_metadata,json=textExtractionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextSentimentDatasetMetadata struct {
	// Metadata for a dataset used for text sentiment.
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `protobuf:"bytes,30,opt,name=text_sentiment_dataset_metadata,json=textSentimentDatasetMetadata,proto3,oneof"`
}

type Dataset_TablesDatasetMetadata struct {
	// Metadata for a dataset used for Tables.
	TablesDatasetMetadata *TablesDatasetMetadata `protobuf:"bytes,33,opt,name=tables_dataset_metadata,json=tablesDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageObjectDetectionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoObjectTrackingDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextExtractionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextSentimentDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TablesDatasetMetadata) isDataset_DatasetMetadata() {}

var File_google_cloud_automl_v1beta1_dataset_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1beta1_dataset_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x0c, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x12, 0x7b, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x94, 0x01, 0x0a, 0x25, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x22, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x91, 0x01, 0x0a, 0x24, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x21, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x98, 0x01, 0x0a, 0x27,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x23, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x94, 0x01, 0x0a, 0x25, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x22, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x95, 0x01,
	0x0a, 0x26, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x22, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x85, 0x01, 0x0a, 0x20, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x1d,
	0x74, 0x65, 0x78, 0x74, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x82, 0x01,
	0x0a, 0x1f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x1c, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x6c, 0x0a, 0x17, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x21, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x15, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65,
	0x74, 0x61, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x3a,
	0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x1d, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x7d, 0x42,
	0x12, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x9b, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x70, 0x62, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1beta1_dataset_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1beta1_dataset_proto_rawDescData = file_google_cloud_automl_v1beta1_dataset_proto_rawDesc
)

func file_google_cloud_automl_v1beta1_dataset_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1beta1_dataset_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1beta1_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1beta1_dataset_proto_rawDescData)
	})
	return file_google_cloud_automl_v1beta1_dataset_proto_rawDescData
}

var file_google_cloud_automl_v1beta1_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_automl_v1beta1_dataset_proto_goTypes = []interface{}{
	(*Dataset)(nil),                             // 0: google.cloud.automl.v1beta1.Dataset
	(*TranslationDatasetMetadata)(nil),          // 1: google.cloud.automl.v1beta1.TranslationDatasetMetadata
	(*ImageClassificationDatasetMetadata)(nil),  // 2: google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata
	(*TextClassificationDatasetMetadata)(nil),   // 3: google.cloud.automl.v1beta1.TextClassificationDatasetMetadata
	(*ImageObjectDetectionDatasetMetadata)(nil), // 4: google.cloud.automl.v1beta1.ImageObjectDetectionDatasetMetadata
	(*VideoClassificationDatasetMetadata)(nil),  // 5: google.cloud.automl.v1beta1.VideoClassificationDatasetMetadata
	(*VideoObjectTrackingDatasetMetadata)(nil),  // 6: google.cloud.automl.v1beta1.VideoObjectTrackingDatasetMetadata
	(*TextExtractionDatasetMetadata)(nil),       // 7: google.cloud.automl.v1beta1.TextExtractionDatasetMetadata
	(*TextSentimentDatasetMetadata)(nil),        // 8: google.cloud.automl.v1beta1.TextSentimentDatasetMetadata
	(*TablesDatasetMetadata)(nil),               // 9: google.cloud.automl.v1beta1.TablesDatasetMetadata
	(*timestamppb.Timestamp)(nil),               // 10: google.protobuf.Timestamp
}
var file_google_cloud_automl_v1beta1_dataset_proto_depIdxs = []int32{
	1,  // 0: google.cloud.automl.v1beta1.Dataset.translation_dataset_metadata:type_name -> google.cloud.automl.v1beta1.TranslationDatasetMetadata
	2,  // 1: google.cloud.automl.v1beta1.Dataset.image_classification_dataset_metadata:type_name -> google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata
	3,  // 2: google.cloud.automl.v1beta1.Dataset.text_classification_dataset_metadata:type_name -> google.cloud.automl.v1beta1.TextClassificationDatasetMetadata
	4,  // 3: google.cloud.automl.v1beta1.Dataset.image_object_detection_dataset_metadata:type_name -> google.cloud.automl.v1beta1.ImageObjectDetectionDatasetMetadata
	5,  // 4: google.cloud.automl.v1beta1.Dataset.video_classification_dataset_metadata:type_name -> google.cloud.automl.v1beta1.VideoClassificationDatasetMetadata
	6,  // 5: google.cloud.automl.v1beta1.Dataset.video_object_tracking_dataset_metadata:type_name -> google.cloud.automl.v1beta1.VideoObjectTrackingDatasetMetadata
	7,  // 6: google.cloud.automl.v1beta1.Dataset.text_extraction_dataset_metadata:type_name -> google.cloud.automl.v1beta1.TextExtractionDatasetMetadata
	8,  // 7: google.cloud.automl.v1beta1.Dataset.text_sentiment_dataset_metadata:type_name -> google.cloud.automl.v1beta1.TextSentimentDatasetMetadata
	9,  // 8: google.cloud.automl.v1beta1.Dataset.tables_dataset_metadata:type_name -> google.cloud.automl.v1beta1.TablesDatasetMetadata
	10, // 9: google.cloud.automl.v1beta1.Dataset.create_time:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1beta1_dataset_proto_init() }
func file_google_cloud_automl_v1beta1_dataset_proto_init() {
	if File_google_cloud_automl_v1beta1_dataset_proto != nil {
		return
	}
	file_google_cloud_automl_v1beta1_image_proto_init()
	file_google_cloud_automl_v1beta1_tables_proto_init()
	file_google_cloud_automl_v1beta1_text_proto_init()
	file_google_cloud_automl_v1beta1_translation_proto_init()
	file_google_cloud_automl_v1beta1_video_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1beta1_dataset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dataset); i {
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
	file_google_cloud_automl_v1beta1_dataset_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
		(*Dataset_ImageObjectDetectionDatasetMetadata)(nil),
		(*Dataset_VideoClassificationDatasetMetadata)(nil),
		(*Dataset_VideoObjectTrackingDatasetMetadata)(nil),
		(*Dataset_TextExtractionDatasetMetadata)(nil),
		(*Dataset_TextSentimentDatasetMetadata)(nil),
		(*Dataset_TablesDatasetMetadata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1beta1_dataset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1beta1_dataset_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1beta1_dataset_proto_depIdxs,
		MessageInfos:      file_google_cloud_automl_v1beta1_dataset_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1beta1_dataset_proto = out.File
	file_google_cloud_automl_v1beta1_dataset_proto_rawDesc = nil
	file_google_cloud_automl_v1beta1_dataset_proto_goTypes = nil
	file_google_cloud_automl_v1beta1_dataset_proto_depIdxs = nil
}
