// Copyright 2021 Google LLC
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
// source: google/cloud/automl/v1/data_items.proto

package automlpb

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

// Unit of the document dimension.
type DocumentDimensions_DocumentDimensionUnit int32

const (
	// Should not be used.
	DocumentDimensions_DOCUMENT_DIMENSION_UNIT_UNSPECIFIED DocumentDimensions_DocumentDimensionUnit = 0
	// Document dimension is measured in inches.
	DocumentDimensions_INCH DocumentDimensions_DocumentDimensionUnit = 1
	// Document dimension is measured in centimeters.
	DocumentDimensions_CENTIMETER DocumentDimensions_DocumentDimensionUnit = 2
	// Document dimension is measured in points. 72 points = 1 inch.
	DocumentDimensions_POINT DocumentDimensions_DocumentDimensionUnit = 3
)

// Enum value maps for DocumentDimensions_DocumentDimensionUnit.
var (
	DocumentDimensions_DocumentDimensionUnit_name = map[int32]string{
		0: "DOCUMENT_DIMENSION_UNIT_UNSPECIFIED",
		1: "INCH",
		2: "CENTIMETER",
		3: "POINT",
	}
	DocumentDimensions_DocumentDimensionUnit_value = map[string]int32{
		"DOCUMENT_DIMENSION_UNIT_UNSPECIFIED": 0,
		"INCH":                                1,
		"CENTIMETER":                          2,
		"POINT":                               3,
	}
)

func (x DocumentDimensions_DocumentDimensionUnit) Enum() *DocumentDimensions_DocumentDimensionUnit {
	p := new(DocumentDimensions_DocumentDimensionUnit)
	*p = x
	return p
}

func (x DocumentDimensions_DocumentDimensionUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentDimensions_DocumentDimensionUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_automl_v1_data_items_proto_enumTypes[0].Descriptor()
}

func (DocumentDimensions_DocumentDimensionUnit) Type() protoreflect.EnumType {
	return &file_google_cloud_automl_v1_data_items_proto_enumTypes[0]
}

func (x DocumentDimensions_DocumentDimensionUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentDimensions_DocumentDimensionUnit.Descriptor instead.
func (DocumentDimensions_DocumentDimensionUnit) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{2, 0}
}

// The type of TextSegment in the context of the original document.
type Document_Layout_TextSegmentType int32

const (
	// Should not be used.
	Document_Layout_TEXT_SEGMENT_TYPE_UNSPECIFIED Document_Layout_TextSegmentType = 0
	// The text segment is a token. e.g. word.
	Document_Layout_TOKEN Document_Layout_TextSegmentType = 1
	// The text segment is a paragraph.
	Document_Layout_PARAGRAPH Document_Layout_TextSegmentType = 2
	// The text segment is a form field.
	Document_Layout_FORM_FIELD Document_Layout_TextSegmentType = 3
	// The text segment is the name part of a form field. It will be treated
	// as child of another FORM_FIELD TextSegment if its span is subspan of
	// another TextSegment with type FORM_FIELD.
	Document_Layout_FORM_FIELD_NAME Document_Layout_TextSegmentType = 4
	// The text segment is the text content part of a form field. It will be
	// treated as child of another FORM_FIELD TextSegment if its span is
	// subspan of another TextSegment with type FORM_FIELD.
	Document_Layout_FORM_FIELD_CONTENTS Document_Layout_TextSegmentType = 5
	// The text segment is a whole table, including headers, and all rows.
	Document_Layout_TABLE Document_Layout_TextSegmentType = 6
	// The text segment is a table's headers. It will be treated as child of
	// another TABLE TextSegment if its span is subspan of another TextSegment
	// with type TABLE.
	Document_Layout_TABLE_HEADER Document_Layout_TextSegmentType = 7
	// The text segment is a row in table. It will be treated as child of
	// another TABLE TextSegment if its span is subspan of another TextSegment
	// with type TABLE.
	Document_Layout_TABLE_ROW Document_Layout_TextSegmentType = 8
	// The text segment is a cell in table. It will be treated as child of
	// another TABLE_ROW TextSegment if its span is subspan of another
	// TextSegment with type TABLE_ROW.
	Document_Layout_TABLE_CELL Document_Layout_TextSegmentType = 9
)

// Enum value maps for Document_Layout_TextSegmentType.
var (
	Document_Layout_TextSegmentType_name = map[int32]string{
		0: "TEXT_SEGMENT_TYPE_UNSPECIFIED",
		1: "TOKEN",
		2: "PARAGRAPH",
		3: "FORM_FIELD",
		4: "FORM_FIELD_NAME",
		5: "FORM_FIELD_CONTENTS",
		6: "TABLE",
		7: "TABLE_HEADER",
		8: "TABLE_ROW",
		9: "TABLE_CELL",
	}
	Document_Layout_TextSegmentType_value = map[string]int32{
		"TEXT_SEGMENT_TYPE_UNSPECIFIED": 0,
		"TOKEN":                         1,
		"PARAGRAPH":                     2,
		"FORM_FIELD":                    3,
		"FORM_FIELD_NAME":               4,
		"FORM_FIELD_CONTENTS":           5,
		"TABLE":                         6,
		"TABLE_HEADER":                  7,
		"TABLE_ROW":                     8,
		"TABLE_CELL":                    9,
	}
)

func (x Document_Layout_TextSegmentType) Enum() *Document_Layout_TextSegmentType {
	p := new(Document_Layout_TextSegmentType)
	*p = x
	return p
}

func (x Document_Layout_TextSegmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Document_Layout_TextSegmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_automl_v1_data_items_proto_enumTypes[1].Descriptor()
}

func (Document_Layout_TextSegmentType) Type() protoreflect.EnumType {
	return &file_google_cloud_automl_v1_data_items_proto_enumTypes[1]
}

func (x Document_Layout_TextSegmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Document_Layout_TextSegmentType.Descriptor instead.
func (Document_Layout_TextSegmentType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{3, 0, 0}
}

// A representation of an image.
// Only images up to 30MB in size are supported.
type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input only. The data representing the image.
	// For Predict calls [image_bytes][google.cloud.automl.v1.Image.image_bytes] must be set .
	//
	// Types that are assignable to Data:
	//	*Image_ImageBytes
	Data isImage_Data `protobuf_oneof:"data"`
	// Output only. HTTP URI to the thumbnail image.
	ThumbnailUri string `protobuf:"bytes,4,opt,name=thumbnail_uri,json=thumbnailUri,proto3" json:"thumbnail_uri,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{0}
}

func (m *Image) GetData() isImage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Image) GetImageBytes() []byte {
	if x, ok := x.GetData().(*Image_ImageBytes); ok {
		return x.ImageBytes
	}
	return nil
}

func (x *Image) GetThumbnailUri() string {
	if x != nil {
		return x.ThumbnailUri
	}
	return ""
}

type isImage_Data interface {
	isImage_Data()
}

type Image_ImageBytes struct {
	// Image content represented as a stream of bytes.
	// Note: As with all `bytes` fields, protobuffers use a pure binary
	// representation, whereas JSON representations use base64.
	ImageBytes []byte `protobuf:"bytes,1,opt,name=image_bytes,json=imageBytes,proto3,oneof"`
}

func (*Image_ImageBytes) isImage_Data() {}

// A representation of a text snippet.
type TextSnippet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The content of the text snippet as a string. Up to 250000
	// characters long.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Optional. The format of [content][google.cloud.automl.v1.TextSnippet.content]. Currently the only two allowed
	// values are "text/html" and "text/plain". If left blank, the format is
	// automatically determined from the type of the uploaded [content][google.cloud.automl.v1.TextSnippet.content].
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Output only. HTTP URI where you can download the content.
	ContentUri string `protobuf:"bytes,4,opt,name=content_uri,json=contentUri,proto3" json:"content_uri,omitempty"`
}

func (x *TextSnippet) Reset() {
	*x = TextSnippet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextSnippet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextSnippet) ProtoMessage() {}

func (x *TextSnippet) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextSnippet.ProtoReflect.Descriptor instead.
func (*TextSnippet) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{1}
}

func (x *TextSnippet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextSnippet) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *TextSnippet) GetContentUri() string {
	if x != nil {
		return x.ContentUri
	}
	return ""
}

// Message that describes dimension of a document.
type DocumentDimensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unit of the dimension.
	Unit DocumentDimensions_DocumentDimensionUnit `protobuf:"varint,1,opt,name=unit,proto3,enum=google.cloud.automl.v1.DocumentDimensions_DocumentDimensionUnit" json:"unit,omitempty"`
	// Width value of the document, works together with the unit.
	Width float32 `protobuf:"fixed32,2,opt,name=width,proto3" json:"width,omitempty"`
	// Height value of the document, works together with the unit.
	Height float32 `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *DocumentDimensions) Reset() {
	*x = DocumentDimensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentDimensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentDimensions) ProtoMessage() {}

func (x *DocumentDimensions) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentDimensions.ProtoReflect.Descriptor instead.
func (*DocumentDimensions) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{2}
}

func (x *DocumentDimensions) GetUnit() DocumentDimensions_DocumentDimensionUnit {
	if x != nil {
		return x.Unit
	}
	return DocumentDimensions_DOCUMENT_DIMENSION_UNIT_UNSPECIFIED
}

func (x *DocumentDimensions) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *DocumentDimensions) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

// A structured text document e.g. a PDF.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An input config specifying the content of the document.
	InputConfig *DocumentInputConfig `protobuf:"bytes,1,opt,name=input_config,json=inputConfig,proto3" json:"input_config,omitempty"`
	// The plain text version of this document.
	DocumentText *TextSnippet `protobuf:"bytes,2,opt,name=document_text,json=documentText,proto3" json:"document_text,omitempty"`
	// Describes the layout of the document.
	// Sorted by [page_number][].
	Layout []*Document_Layout `protobuf:"bytes,3,rep,name=layout,proto3" json:"layout,omitempty"`
	// The dimensions of the page in the document.
	DocumentDimensions *DocumentDimensions `protobuf:"bytes,4,opt,name=document_dimensions,json=documentDimensions,proto3" json:"document_dimensions,omitempty"`
	// Number of pages in the document.
	PageCount int32 `protobuf:"varint,5,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{3}
}

func (x *Document) GetInputConfig() *DocumentInputConfig {
	if x != nil {
		return x.InputConfig
	}
	return nil
}

func (x *Document) GetDocumentText() *TextSnippet {
	if x != nil {
		return x.DocumentText
	}
	return nil
}

func (x *Document) GetLayout() []*Document_Layout {
	if x != nil {
		return x.Layout
	}
	return nil
}

func (x *Document) GetDocumentDimensions() *DocumentDimensions {
	if x != nil {
		return x.DocumentDimensions
	}
	return nil
}

func (x *Document) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

// Example data used for training or prediction.
type ExamplePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The example data.
	//
	// Types that are assignable to Payload:
	//	*ExamplePayload_Image
	//	*ExamplePayload_TextSnippet
	//	*ExamplePayload_Document
	Payload isExamplePayload_Payload `protobuf_oneof:"payload"`
}

func (x *ExamplePayload) Reset() {
	*x = ExamplePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamplePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamplePayload) ProtoMessage() {}

func (x *ExamplePayload) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamplePayload.ProtoReflect.Descriptor instead.
func (*ExamplePayload) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{4}
}

func (m *ExamplePayload) GetPayload() isExamplePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ExamplePayload) GetImage() *Image {
	if x, ok := x.GetPayload().(*ExamplePayload_Image); ok {
		return x.Image
	}
	return nil
}

func (x *ExamplePayload) GetTextSnippet() *TextSnippet {
	if x, ok := x.GetPayload().(*ExamplePayload_TextSnippet); ok {
		return x.TextSnippet
	}
	return nil
}

func (x *ExamplePayload) GetDocument() *Document {
	if x, ok := x.GetPayload().(*ExamplePayload_Document); ok {
		return x.Document
	}
	return nil
}

type isExamplePayload_Payload interface {
	isExamplePayload_Payload()
}

type ExamplePayload_Image struct {
	// Example image.
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3,oneof"`
}

type ExamplePayload_TextSnippet struct {
	// Example text.
	TextSnippet *TextSnippet `protobuf:"bytes,2,opt,name=text_snippet,json=textSnippet,proto3,oneof"`
}

type ExamplePayload_Document struct {
	// Example document.
	Document *Document `protobuf:"bytes,4,opt,name=document,proto3,oneof"`
}

func (*ExamplePayload_Image) isExamplePayload_Payload() {}

func (*ExamplePayload_TextSnippet) isExamplePayload_Payload() {}

func (*ExamplePayload_Document) isExamplePayload_Payload() {}

// Describes the layout information of a [text_segment][google.cloud.automl.v1.Document.Layout.text_segment] in the document.
type Document_Layout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text Segment that represents a segment in
	// [document_text][google.cloud.automl.v1p1beta.Document.document_text].
	TextSegment *TextSegment `protobuf:"bytes,1,opt,name=text_segment,json=textSegment,proto3" json:"text_segment,omitempty"`
	// Page number of the [text_segment][google.cloud.automl.v1.Document.Layout.text_segment] in the original document, starts
	// from 1.
	PageNumber int32 `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	// The position of the [text_segment][google.cloud.automl.v1.Document.Layout.text_segment] in the page.
	// Contains exactly 4
	// [normalized_vertices][google.cloud.automl.v1p1beta.BoundingPoly.normalized_vertices]
	// and they are connected by edges in the order provided, which will
	// represent a rectangle parallel to the frame. The
	// [NormalizedVertex-s][google.cloud.automl.v1p1beta.NormalizedVertex] are
	// relative to the page.
	// Coordinates are based on top-left as point (0,0).
	BoundingPoly *BoundingPoly `protobuf:"bytes,3,opt,name=bounding_poly,json=boundingPoly,proto3" json:"bounding_poly,omitempty"`
	// The type of the [text_segment][google.cloud.automl.v1.Document.Layout.text_segment] in document.
	TextSegmentType Document_Layout_TextSegmentType `protobuf:"varint,4,opt,name=text_segment_type,json=textSegmentType,proto3,enum=google.cloud.automl.v1.Document_Layout_TextSegmentType" json:"text_segment_type,omitempty"`
}

func (x *Document_Layout) Reset() {
	*x = Document_Layout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document_Layout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document_Layout) ProtoMessage() {}

func (x *Document_Layout) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_data_items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document_Layout.ProtoReflect.Descriptor instead.
func (*Document_Layout) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_data_items_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Document_Layout) GetTextSegment() *TextSegment {
	if x != nil {
		return x.TextSegment
	}
	return nil
}

func (x *Document_Layout) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *Document_Layout) GetBoundingPoly() *BoundingPoly {
	if x != nil {
		return x.BoundingPoly
	}
	return nil
}

func (x *Document_Layout) GetTextSegmentType() Document_Layout_TextSegmentType {
	if x != nil {
		return x.TextSegmentType
	}
	return Document_Layout_TEXT_SEGMENT_TYPE_UNSPECIFIED
}

var File_google_cloud_automl_v1_data_items_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1_data_items_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x75, 0x72,
	0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x55, 0x72, 0x69, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x65, 0x0a,
	0x0b, 0x54, 0x65, 0x78, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x55, 0x72, 0x69, 0x22, 0xff, 0x01, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x65, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x4f, 0x43, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x49, 0x4d, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x43, 0x48, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43,
	0x45, 0x4e, 0x54, 0x49, 0x4d, 0x45, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x22, 0xd0, 0x06, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x52,
	0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x3f, 0x0a,
	0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x5b,
	0x0a, 0x13, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x6d,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x12, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0xec, 0x03, 0x0a, 0x06, 0x4c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x49,
	0x0a, 0x0d, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x52, 0x0c, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x79, 0x12, 0x63, 0x0a, 0x11, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x65,
	0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x74,
	0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xc8,
	0x01, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x53, 0x45, 0x47, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x50, 0x41, 0x52, 0x41, 0x47, 0x52, 0x41, 0x50, 0x48, 0x10, 0x02, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x46, 0x49, 0x45,
	0x4c, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x05, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x41, 0x42, 0x4c,
	0x45, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x41,
	0x42, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x57, 0x10, 0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x5f, 0x43, 0x45, 0x4c, 0x4c, 0x10, 0x09, 0x22, 0xdc, 0x01, 0x0a, 0x0e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x35, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x6e, 0x69, 0x70,
	0x70, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x48, 0x00,
	0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x3e, 0x0a,
	0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0xa0, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0xaa, 0x02, 0x16,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x4d, 0x4c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1_data_items_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1_data_items_proto_rawDescData = file_google_cloud_automl_v1_data_items_proto_rawDesc
)

func file_google_cloud_automl_v1_data_items_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1_data_items_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1_data_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1_data_items_proto_rawDescData)
	})
	return file_google_cloud_automl_v1_data_items_proto_rawDescData
}

var file_google_cloud_automl_v1_data_items_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_automl_v1_data_items_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_automl_v1_data_items_proto_goTypes = []interface{}{
	(DocumentDimensions_DocumentDimensionUnit)(0), // 0: google.cloud.automl.v1.DocumentDimensions.DocumentDimensionUnit
	(Document_Layout_TextSegmentType)(0),          // 1: google.cloud.automl.v1.Document.Layout.TextSegmentType
	(*Image)(nil),                                 // 2: google.cloud.automl.v1.Image
	(*TextSnippet)(nil),                           // 3: google.cloud.automl.v1.TextSnippet
	(*DocumentDimensions)(nil),                    // 4: google.cloud.automl.v1.DocumentDimensions
	(*Document)(nil),                              // 5: google.cloud.automl.v1.Document
	(*ExamplePayload)(nil),                        // 6: google.cloud.automl.v1.ExamplePayload
	(*Document_Layout)(nil),                       // 7: google.cloud.automl.v1.Document.Layout
	(*DocumentInputConfig)(nil),                   // 8: google.cloud.automl.v1.DocumentInputConfig
	(*TextSegment)(nil),                           // 9: google.cloud.automl.v1.TextSegment
	(*BoundingPoly)(nil),                          // 10: google.cloud.automl.v1.BoundingPoly
}
var file_google_cloud_automl_v1_data_items_proto_depIdxs = []int32{
	0,  // 0: google.cloud.automl.v1.DocumentDimensions.unit:type_name -> google.cloud.automl.v1.DocumentDimensions.DocumentDimensionUnit
	8,  // 1: google.cloud.automl.v1.Document.input_config:type_name -> google.cloud.automl.v1.DocumentInputConfig
	3,  // 2: google.cloud.automl.v1.Document.document_text:type_name -> google.cloud.automl.v1.TextSnippet
	7,  // 3: google.cloud.automl.v1.Document.layout:type_name -> google.cloud.automl.v1.Document.Layout
	4,  // 4: google.cloud.automl.v1.Document.document_dimensions:type_name -> google.cloud.automl.v1.DocumentDimensions
	2,  // 5: google.cloud.automl.v1.ExamplePayload.image:type_name -> google.cloud.automl.v1.Image
	3,  // 6: google.cloud.automl.v1.ExamplePayload.text_snippet:type_name -> google.cloud.automl.v1.TextSnippet
	5,  // 7: google.cloud.automl.v1.ExamplePayload.document:type_name -> google.cloud.automl.v1.Document
	9,  // 8: google.cloud.automl.v1.Document.Layout.text_segment:type_name -> google.cloud.automl.v1.TextSegment
	10, // 9: google.cloud.automl.v1.Document.Layout.bounding_poly:type_name -> google.cloud.automl.v1.BoundingPoly
	1,  // 10: google.cloud.automl.v1.Document.Layout.text_segment_type:type_name -> google.cloud.automl.v1.Document.Layout.TextSegmentType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1_data_items_proto_init() }
func file_google_cloud_automl_v1_data_items_proto_init() {
	if File_google_cloud_automl_v1_data_items_proto != nil {
		return
	}
	file_google_cloud_automl_v1_geometry_proto_init()
	file_google_cloud_automl_v1_io_proto_init()
	file_google_cloud_automl_v1_text_segment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1_data_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_google_cloud_automl_v1_data_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextSnippet); i {
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
		file_google_cloud_automl_v1_data_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentDimensions); i {
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
		file_google_cloud_automl_v1_data_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_google_cloud_automl_v1_data_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamplePayload); i {
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
		file_google_cloud_automl_v1_data_items_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document_Layout); i {
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
	file_google_cloud_automl_v1_data_items_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Image_ImageBytes)(nil),
	}
	file_google_cloud_automl_v1_data_items_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ExamplePayload_Image)(nil),
		(*ExamplePayload_TextSnippet)(nil),
		(*ExamplePayload_Document)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_automl_v1_data_items_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1_data_items_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1_data_items_proto_depIdxs,
		EnumInfos:         file_google_cloud_automl_v1_data_items_proto_enumTypes,
		MessageInfos:      file_google_cloud_automl_v1_data_items_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1_data_items_proto = out.File
	file_google_cloud_automl_v1_data_items_proto_rawDesc = nil
	file_google_cloud_automl_v1_data_items_proto_goTypes = nil
	file_google_cloud_automl_v1_data_items_proto_depIdxs = nil
}
