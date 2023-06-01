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
// source: google/maps/routing/v2/waypoint.proto

package routingpb

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

// Encapsulates a waypoint. Waypoints mark both the beginning and end of a
// route, and include intermediate stops along the route.
type Waypoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Different ways to represent a location.
	//
	// Types that are assignable to LocationType:
	//	*Waypoint_Location
	//	*Waypoint_PlaceId
	//	*Waypoint_Address
	LocationType isWaypoint_LocationType `protobuf_oneof:"location_type"`
	// Marks this waypoint as a milestone rather a stopping point. For
	// each non-via waypoint in the request, the response appends an entry to the
	// [legs][google.maps.routing.v2.Route.legs]
	// array to provide the details for stopovers on that leg of the trip. Set
	// this value to true when you want the route to pass through this waypoint
	// without stopping over. Via waypoints don't cause an entry to be added to
	// the `legs` array, but they do route the journey through the waypoint. You
	// can only set this value on waypoints that are intermediates. The request
	// fails if you set this field on terminal waypoints. If
	// `ComputeRoutesRequest.optimize_waypoint_order` is set to true then this
	// field cannot be set to true; otherwise, the request fails.
	Via bool `protobuf:"varint,3,opt,name=via,proto3" json:"via,omitempty"`
	// Indicates that the waypoint is meant for vehicles to stop at, where the
	// intention is to either pickup or drop-off. When you set this value, the
	// calculated route won't include non-`via` waypoints on roads that are
	// unsuitable for pickup and drop-off. This option works only for `DRIVE` and
	// `TWO_WHEELER` travel modes, and when the `location_type` is
	// [Location][google.maps.routing.v2.Location].
	VehicleStopover bool `protobuf:"varint,4,opt,name=vehicle_stopover,json=vehicleStopover,proto3" json:"vehicle_stopover,omitempty"`
	// Indicates that the location of this waypoint is meant to have a preference
	// for the vehicle to stop at a particular side of road. When you set this
	// value, the route will pass through the location so that the vehicle can
	// stop at the side of road that the location is biased towards from the
	// center of the road. This option works only for 'DRIVE' and 'TWO_WHEELER'
	// [RouteTravelMode][google.maps.routing.v2.RouteTravelMode].
	SideOfRoad bool `protobuf:"varint,5,opt,name=side_of_road,json=sideOfRoad,proto3" json:"side_of_road,omitempty"`
}

func (x *Waypoint) Reset() {
	*x = Waypoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_waypoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Waypoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Waypoint) ProtoMessage() {}

func (x *Waypoint) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_waypoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Waypoint.ProtoReflect.Descriptor instead.
func (*Waypoint) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_waypoint_proto_rawDescGZIP(), []int{0}
}

func (m *Waypoint) GetLocationType() isWaypoint_LocationType {
	if m != nil {
		return m.LocationType
	}
	return nil
}

func (x *Waypoint) GetLocation() *Location {
	if x, ok := x.GetLocationType().(*Waypoint_Location); ok {
		return x.Location
	}
	return nil
}

func (x *Waypoint) GetPlaceId() string {
	if x, ok := x.GetLocationType().(*Waypoint_PlaceId); ok {
		return x.PlaceId
	}
	return ""
}

func (x *Waypoint) GetAddress() string {
	if x, ok := x.GetLocationType().(*Waypoint_Address); ok {
		return x.Address
	}
	return ""
}

func (x *Waypoint) GetVia() bool {
	if x != nil {
		return x.Via
	}
	return false
}

func (x *Waypoint) GetVehicleStopover() bool {
	if x != nil {
		return x.VehicleStopover
	}
	return false
}

func (x *Waypoint) GetSideOfRoad() bool {
	if x != nil {
		return x.SideOfRoad
	}
	return false
}

type isWaypoint_LocationType interface {
	isWaypoint_LocationType()
}

type Waypoint_Location struct {
	// A point specified using geographic coordinates, including an optional
	// heading.
	Location *Location `protobuf:"bytes,1,opt,name=location,proto3,oneof"`
}

type Waypoint_PlaceId struct {
	// The POI Place ID associated with the waypoint.
	PlaceId string `protobuf:"bytes,2,opt,name=place_id,json=placeId,proto3,oneof"`
}

type Waypoint_Address struct {
	// Human readable address or a plus code.
	// See https://plus.codes for details.
	Address string `protobuf:"bytes,7,opt,name=address,proto3,oneof"`
}

func (*Waypoint_Location) isWaypoint_LocationType() {}

func (*Waypoint_PlaceId) isWaypoint_LocationType() {}

func (*Waypoint_Address) isWaypoint_LocationType() {}

var File_google_maps_routing_v2_waypoint_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_waypoint_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x1a,
	0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x08, 0x57, 0x61, 0x79, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x69, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x76, 0x69, 0x61, 0x12, 0x29,
	0x0a, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x53, 0x74, 0x6f, 0x70, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x69, 0x64,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x72, 0x6f, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x73, 0x69, 0x64, 0x65, 0x4f, 0x66, 0x52, 0x6f, 0x61, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0xc2, 0x01, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x57, 0x61, 0x79,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x05, 0x47, 0x4d,
	0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70,
	0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_waypoint_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_waypoint_proto_rawDescData = file_google_maps_routing_v2_waypoint_proto_rawDesc
)

func file_google_maps_routing_v2_waypoint_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_waypoint_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_waypoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_waypoint_proto_rawDescData)
	})
	return file_google_maps_routing_v2_waypoint_proto_rawDescData
}

var file_google_maps_routing_v2_waypoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routing_v2_waypoint_proto_goTypes = []interface{}{
	(*Waypoint)(nil), // 0: google.maps.routing.v2.Waypoint
	(*Location)(nil), // 1: google.maps.routing.v2.Location
}
var file_google_maps_routing_v2_waypoint_proto_depIdxs = []int32{
	1, // 0: google.maps.routing.v2.Waypoint.location:type_name -> google.maps.routing.v2.Location
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_waypoint_proto_init() }
func file_google_maps_routing_v2_waypoint_proto_init() {
	if File_google_maps_routing_v2_waypoint_proto != nil {
		return
	}
	file_google_maps_routing_v2_location_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routing_v2_waypoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Waypoint); i {
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
	file_google_maps_routing_v2_waypoint_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Waypoint_Location)(nil),
		(*Waypoint_PlaceId)(nil),
		(*Waypoint_Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_waypoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_waypoint_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_waypoint_proto_depIdxs,
		MessageInfos:      file_google_maps_routing_v2_waypoint_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_waypoint_proto = out.File
	file_google_maps_routing_v2_waypoint_proto_rawDesc = nil
	file_google_maps_routing_v2_waypoint_proto_goTypes = nil
	file_google_maps_routing_v2_waypoint_proto_depIdxs = nil
}
