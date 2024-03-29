//
//Copyright (c) 2021 Nordix Foundation
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: api/target/target.proto

package target

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConduitEventStatus int32

const (
	ConduitEventStatus_Connect    ConduitEventStatus = 0
	ConduitEventStatus_Disconnect ConduitEventStatus = 1
)

// Enum value maps for ConduitEventStatus.
var (
	ConduitEventStatus_name = map[int32]string{
		0: "Connect",
		1: "Disconnect",
	}
	ConduitEventStatus_value = map[string]int32{
		"Connect":    0,
		"Disconnect": 1,
	}
)

func (x ConduitEventStatus) Enum() *ConduitEventStatus {
	p := new(ConduitEventStatus)
	*p = x
	return p
}

func (x ConduitEventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConduitEventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_target_target_proto_enumTypes[0].Descriptor()
}

func (ConduitEventStatus) Type() protoreflect.EnumType {
	return &file_api_target_target_proto_enumTypes[0]
}

func (x ConduitEventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConduitEventStatus.Descriptor instead.
func (ConduitEventStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{0}
}

type StreamEventStatus int32

const (
	StreamEventStatus_Request StreamEventStatus = 0
	StreamEventStatus_Close   StreamEventStatus = 1
)

// Enum value maps for StreamEventStatus.
var (
	StreamEventStatus_name = map[int32]string{
		0: "Request",
		1: "Close",
	}
	StreamEventStatus_value = map[string]int32{
		"Request": 0,
		"Close":   1,
	}
)

func (x StreamEventStatus) Enum() *StreamEventStatus {
	p := new(StreamEventStatus)
	*p = x
	return p
}

func (x StreamEventStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamEventStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_target_target_proto_enumTypes[1].Descriptor()
}

func (StreamEventStatus) Type() protoreflect.EnumType {
	return &file_api_target_target_proto_enumTypes[1]
}

func (x StreamEventStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamEventStatus.Descriptor instead.
func (StreamEventStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{1}
}

type Trench struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the trench
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Trench) Reset() {
	*x = Trench{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trench) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trench) ProtoMessage() {}

func (x *Trench) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trench.ProtoReflect.Descriptor instead.
func (*Trench) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{0}
}

func (x *Trench) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The message contains the properties of a conduit
type Conduit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the service
	// The name will be compiled using the proxy and the trench name
	// e.g.: load-balancer with the proxy used in trench-a will
	// become proxy.load-balancer.trench-a
	NetworkServiceName string `protobuf:"bytes,1,opt,name=networkServiceName,proto3" json:"networkServiceName,omitempty"`
	// Trench the conduit belongs to
	Trench *Trench `protobuf:"bytes,2,opt,name=trench,proto3" json:"trench,omitempty"`
}

func (x *Conduit) Reset() {
	*x = Conduit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conduit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conduit) ProtoMessage() {}

func (x *Conduit) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conduit.ProtoReflect.Descriptor instead.
func (*Conduit) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{1}
}

func (x *Conduit) GetNetworkServiceName() string {
	if x != nil {
		return x.NetworkServiceName
	}
	return ""
}

func (x *Conduit) GetTrench() *Trench {
	if x != nil {
		return x.Trench
	}
	return nil
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Conduit the stream belongs to
	Conduit *Conduit `protobuf:"bytes,2,opt,name=conduit,proto3" json:"conduit,omitempty"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{2}
}

func (x *Stream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stream) GetConduit() *Conduit {
	if x != nil {
		return x.Conduit
	}
	return nil
}

type ConduitEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conduit            *Conduit           `protobuf:"bytes,1,opt,name=conduit,proto3" json:"conduit,omitempty"`
	ConduitEventStatus ConduitEventStatus `protobuf:"varint,2,opt,name=conduitEventStatus,proto3,enum=target.ConduitEventStatus" json:"conduitEventStatus,omitempty"`
}

func (x *ConduitEvent) Reset() {
	*x = ConduitEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConduitEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConduitEvent) ProtoMessage() {}

func (x *ConduitEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConduitEvent.ProtoReflect.Descriptor instead.
func (*ConduitEvent) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{3}
}

func (x *ConduitEvent) GetConduit() *Conduit {
	if x != nil {
		return x.Conduit
	}
	return nil
}

func (x *ConduitEvent) GetConduitEventStatus() ConduitEventStatus {
	if x != nil {
		return x.ConduitEventStatus
	}
	return ConduitEventStatus_Connect
}

type StreamEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream            *Stream           `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	StreamEventStatus StreamEventStatus `protobuf:"varint,2,opt,name=streamEventStatus,proto3,enum=target.StreamEventStatus" json:"streamEventStatus,omitempty"`
}

func (x *StreamEvent) Reset() {
	*x = StreamEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_target_target_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEvent) ProtoMessage() {}

func (x *StreamEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_target_target_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEvent.ProtoReflect.Descriptor instead.
func (*StreamEvent) Descriptor() ([]byte, []int) {
	return file_api_target_target_proto_rawDescGZIP(), []int{4}
}

func (x *StreamEvent) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *StreamEvent) GetStreamEventStatus() StreamEventStatus {
	if x != nil {
		return x.StreamEventStatus
	}
	return StreamEventStatus_Request
}

var File_api_target_target_proto protoreflect.FileDescriptor

var file_api_target_target_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c,
	0x0a, 0x06, 0x54, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x72, 0x65, 0x6e, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x2e, 0x54, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x74, 0x72, 0x65, 0x6e, 0x63, 0x68, 0x22,
	0x47, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x64, 0x75, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x64, 0x75, 0x69, 0x74, 0x12, 0x4a, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x63, 0x6f,
	0x6e, 0x64, 0x75, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x7e, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x47, 0x0a, 0x11, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x31, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x10, 0x01, 0x2a, 0x2b, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x10, 0x01,
	0x32, 0xe7, 0x02, 0x0a, 0x0a, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x12,
	0x34, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x75, 0x69, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x75, 0x69, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x33, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x2e, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12,
	0x0e, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x13, 0x2e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x69, 0x78, 0x2f,
	0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_target_target_proto_rawDescOnce sync.Once
	file_api_target_target_proto_rawDescData = file_api_target_target_proto_rawDesc
)

func file_api_target_target_proto_rawDescGZIP() []byte {
	file_api_target_target_proto_rawDescOnce.Do(func() {
		file_api_target_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_target_target_proto_rawDescData)
	})
	return file_api_target_target_proto_rawDescData
}

var file_api_target_target_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_target_target_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_target_target_proto_goTypes = []interface{}{
	(ConduitEventStatus)(0), // 0: target.ConduitEventStatus
	(StreamEventStatus)(0),  // 1: target.StreamEventStatus
	(*Trench)(nil),          // 2: target.Trench
	(*Conduit)(nil),         // 3: target.Conduit
	(*Stream)(nil),          // 4: target.Stream
	(*ConduitEvent)(nil),    // 5: target.ConduitEvent
	(*StreamEvent)(nil),     // 6: target.StreamEvent
	(*empty.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_api_target_target_proto_depIdxs = []int32{
	2,  // 0: target.Conduit.trench:type_name -> target.Trench
	3,  // 1: target.Stream.conduit:type_name -> target.Conduit
	3,  // 2: target.ConduitEvent.conduit:type_name -> target.Conduit
	0,  // 3: target.ConduitEvent.conduitEventStatus:type_name -> target.ConduitEventStatus
	4,  // 4: target.StreamEvent.stream:type_name -> target.Stream
	1,  // 5: target.StreamEvent.streamEventStatus:type_name -> target.StreamEventStatus
	3,  // 6: target.Ambassador.Connect:input_type -> target.Conduit
	3,  // 7: target.Ambassador.Disconnect:input_type -> target.Conduit
	7,  // 8: target.Ambassador.WatchConduits:input_type -> google.protobuf.Empty
	4,  // 9: target.Ambassador.Request:input_type -> target.Stream
	4,  // 10: target.Ambassador.Close:input_type -> target.Stream
	7,  // 11: target.Ambassador.WatchStreams:input_type -> google.protobuf.Empty
	7,  // 12: target.Ambassador.Connect:output_type -> google.protobuf.Empty
	7,  // 13: target.Ambassador.Disconnect:output_type -> google.protobuf.Empty
	5,  // 14: target.Ambassador.WatchConduits:output_type -> target.ConduitEvent
	7,  // 15: target.Ambassador.Request:output_type -> google.protobuf.Empty
	7,  // 16: target.Ambassador.Close:output_type -> google.protobuf.Empty
	6,  // 17: target.Ambassador.WatchStreams:output_type -> target.StreamEvent
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_target_target_proto_init() }
func file_api_target_target_proto_init() {
	if File_api_target_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_target_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trench); i {
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
		file_api_target_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conduit); i {
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
		file_api_target_target_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_api_target_target_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConduitEvent); i {
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
		file_api_target_target_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEvent); i {
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
			RawDescriptor: file_api_target_target_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_target_target_proto_goTypes,
		DependencyIndexes: file_api_target_target_proto_depIdxs,
		EnumInfos:         file_api_target_target_proto_enumTypes,
		MessageInfos:      file_api_target_target_proto_msgTypes,
	}.Build()
	File_api_target_target_proto = out.File
	file_api_target_target_proto_rawDesc = nil
	file_api_target_target_proto_goTypes = nil
	file_api_target_target_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AmbassadorClient is the client API for Ambassador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AmbassadorClient interface {
	// Connect to a conduit, so a new interface will be created.
	// The Ambassador will also connect to the trench the
	// conduit belongs to. And, the VIPs will be added to
	// the loopback interface.
	Connect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error)
	// Disconnect from a conduit, so the interface will be removed.
	// The Ambassador will also close the streams which are opened
	// in the conduit. It will disconnect the target from the trench.
	// And, the VIPs will be removed from the loopback interface.
	Disconnect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error)
	// WatchConduits will, first, returns all conduits connected with
	// "Connect" as event status, and then, will send an event on each
	// conduit connection/disconnection.
	WatchConduits(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Ambassador_WatchConduitsClient, error)
	// Request to a stream, so the identifier will be registered
	// in the NSP service, the LBs will start load-balancing the
	// traffic to the target.
	Request(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error)
	// Close a stream, so the identifier will be unregistered
	// in the NSP service, the LBs will stop load-balancing the
	// traffic to the target.
	Close(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error)
	// WatchStreams will, first, returns all streams opened with "Request"
	// as event status, and then, will send an event on each stream
	// request/close.
	WatchStreams(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Ambassador_WatchStreamsClient, error)
}

type ambassadorClient struct {
	cc grpc.ClientConnInterface
}

func NewAmbassadorClient(cc grpc.ClientConnInterface) AmbassadorClient {
	return &ambassadorClient{cc}
}

func (c *ambassadorClient) Connect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) Disconnect(ctx context.Context, in *Conduit, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) WatchConduits(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Ambassador_WatchConduitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ambassador_serviceDesc.Streams[0], "/target.Ambassador/WatchConduits", opts...)
	if err != nil {
		return nil, err
	}
	x := &ambassadorWatchConduitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ambassador_WatchConduitsClient interface {
	Recv() (*ConduitEvent, error)
	grpc.ClientStream
}

type ambassadorWatchConduitsClient struct {
	grpc.ClientStream
}

func (x *ambassadorWatchConduitsClient) Recv() (*ConduitEvent, error) {
	m := new(ConduitEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ambassadorClient) Request(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) Close(ctx context.Context, in *Stream, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/target.Ambassador/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ambassadorClient) WatchStreams(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Ambassador_WatchStreamsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Ambassador_serviceDesc.Streams[1], "/target.Ambassador/WatchStreams", opts...)
	if err != nil {
		return nil, err
	}
	x := &ambassadorWatchStreamsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ambassador_WatchStreamsClient interface {
	Recv() (*StreamEvent, error)
	grpc.ClientStream
}

type ambassadorWatchStreamsClient struct {
	grpc.ClientStream
}

func (x *ambassadorWatchStreamsClient) Recv() (*StreamEvent, error) {
	m := new(StreamEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AmbassadorServer is the server API for Ambassador service.
type AmbassadorServer interface {
	// Connect to a conduit, so a new interface will be created.
	// The Ambassador will also connect to the trench the
	// conduit belongs to. And, the VIPs will be added to
	// the loopback interface.
	Connect(context.Context, *Conduit) (*empty.Empty, error)
	// Disconnect from a conduit, so the interface will be removed.
	// The Ambassador will also close the streams which are opened
	// in the conduit. It will disconnect the target from the trench.
	// And, the VIPs will be removed from the loopback interface.
	Disconnect(context.Context, *Conduit) (*empty.Empty, error)
	// WatchConduits will, first, returns all conduits connected with
	// "Connect" as event status, and then, will send an event on each
	// conduit connection/disconnection.
	WatchConduits(*empty.Empty, Ambassador_WatchConduitsServer) error
	// Request to a stream, so the identifier will be registered
	// in the NSP service, the LBs will start load-balancing the
	// traffic to the target.
	Request(context.Context, *Stream) (*empty.Empty, error)
	// Close a stream, so the identifier will be unregistered
	// in the NSP service, the LBs will stop load-balancing the
	// traffic to the target.
	Close(context.Context, *Stream) (*empty.Empty, error)
	// WatchStreams will, first, returns all streams opened with "Request"
	// as event status, and then, will send an event on each stream
	// request/close.
	WatchStreams(*empty.Empty, Ambassador_WatchStreamsServer) error
}

// UnimplementedAmbassadorServer can be embedded to have forward compatible implementations.
type UnimplementedAmbassadorServer struct {
}

func (*UnimplementedAmbassadorServer) Connect(context.Context, *Conduit) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedAmbassadorServer) Disconnect(context.Context, *Conduit) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (*UnimplementedAmbassadorServer) WatchConduits(*empty.Empty, Ambassador_WatchConduitsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchConduits not implemented")
}
func (*UnimplementedAmbassadorServer) Request(context.Context, *Stream) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (*UnimplementedAmbassadorServer) Close(context.Context, *Stream) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (*UnimplementedAmbassadorServer) WatchStreams(*empty.Empty, Ambassador_WatchStreamsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchStreams not implemented")
}

func RegisterAmbassadorServer(s *grpc.Server, srv AmbassadorServer) {
	s.RegisterService(&_Ambassador_serviceDesc, srv)
}

func _Ambassador_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conduit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Connect(ctx, req.(*Conduit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Conduit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Disconnect(ctx, req.(*Conduit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_WatchConduits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AmbassadorServer).WatchConduits(m, &ambassadorWatchConduitsServer{stream})
}

type Ambassador_WatchConduitsServer interface {
	Send(*ConduitEvent) error
	grpc.ServerStream
}

type ambassadorWatchConduitsServer struct {
	grpc.ServerStream
}

func (x *ambassadorWatchConduitsServer) Send(m *ConduitEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Ambassador_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Request(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stream)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/target.Ambassador/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).Close(ctx, req.(*Stream))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ambassador_WatchStreams_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AmbassadorServer).WatchStreams(m, &ambassadorWatchStreamsServer{stream})
}

type Ambassador_WatchStreamsServer interface {
	Send(*StreamEvent) error
	grpc.ServerStream
}

type ambassadorWatchStreamsServer struct {
	grpc.ServerStream
}

func (x *ambassadorWatchStreamsServer) Send(m *StreamEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _Ambassador_serviceDesc = grpc.ServiceDesc{
	ServiceName: "target.Ambassador",
	HandlerType: (*AmbassadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Ambassador_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Ambassador_Disconnect_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _Ambassador_Request_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Ambassador_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchConduits",
			Handler:       _Ambassador_WatchConduits_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchStreams",
			Handler:       _Ambassador_WatchStreams_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/target/target.proto",
}
