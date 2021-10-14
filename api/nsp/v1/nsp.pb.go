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
// source: api/nsp/v1/nsp.proto

package v1

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

type Target_Status int32

const (
	Target_Enabled  Target_Status = 0
	Target_Disabled Target_Status = 1
)

// Enum value maps for Target_Status.
var (
	Target_Status_name = map[int32]string{
		0: "Enabled",
		1: "Disabled",
	}
	Target_Status_value = map[string]int32{
		"Enabled":  0,
		"Disabled": 1,
	}
)

func (x Target_Status) Enum() *Target_Status {
	p := new(Target_Status)
	*p = x
	return p
}

func (x Target_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Target_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_nsp_v1_nsp_proto_enumTypes[0].Descriptor()
}

func (Target_Status) Type() protoreflect.EnumType {
	return &file_api_nsp_v1_nsp_proto_enumTypes[0]
}

func (x Target_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Target_Status.Descriptor instead.
func (Target_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{1, 0}
}

type Target_Type int32

const (
	Target_DEFAULT  Target_Type = 0
	Target_FRONTEND Target_Type = 1
)

// Enum value maps for Target_Type.
var (
	Target_Type_name = map[int32]string{
		0: "DEFAULT",
		1: "FRONTEND",
	}
	Target_Type_value = map[string]int32{
		"DEFAULT":  0,
		"FRONTEND": 1,
	}
)

func (x Target_Type) Enum() *Target_Type {
	p := new(Target_Type)
	*p = x
	return p
}

func (x Target_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Target_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_api_nsp_v1_nsp_proto_enumTypes[1].Descriptor()
}

func (Target_Type) Type() protoreflect.EnumType {
	return &file_api_nsp_v1_nsp_proto_enumTypes[1]
}

func (x Target_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Target_Type.Descriptor instead.
func (Target_Type) EnumDescriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{1, 1}
}

type TargetEvent_Status int32

const (
	TargetEvent_Register   TargetEvent_Status = 0
	TargetEvent_Unregister TargetEvent_Status = 1
	TargetEvent_Updated    TargetEvent_Status = 2
)

// Enum value maps for TargetEvent_Status.
var (
	TargetEvent_Status_name = map[int32]string{
		0: "Register",
		1: "Unregister",
		2: "Updated",
	}
	TargetEvent_Status_value = map[string]int32{
		"Register":   0,
		"Unregister": 1,
		"Updated":    2,
	}
)

func (x TargetEvent_Status) Enum() *TargetEvent_Status {
	p := new(TargetEvent_Status)
	*p = x
	return p
}

func (x TargetEvent_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetEvent_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_nsp_v1_nsp_proto_enumTypes[2].Descriptor()
}

func (TargetEvent_Status) Type() protoreflect.EnumType {
	return &file_api_nsp_v1_nsp_proto_enumTypes[2]
}

func (x TargetEvent_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetEvent_Status.Descriptor instead.
func (TargetEvent_Status) EnumDescriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{3, 0}
}

type GetTargetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*Target `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *GetTargetsResponse) Reset() {
	*x = GetTargetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nsp_v1_nsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetsResponse) ProtoMessage() {}

func (x *GetTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_nsp_v1_nsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetsResponse.ProtoReflect.Descriptor instead.
func (*GetTargetsResponse) Descriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetTargetsResponse) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips     []string          `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
	Type    Target_Type       `protobuf:"varint,2,opt,name=type,proto3,enum=nsp.v1.Target_Type" json:"type,omitempty"`
	Context map[string]string `protobuf:"bytes,3,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Status  Target_Status     `protobuf:"varint,4,opt,name=status,proto3,enum=nsp.v1.Target_Status" json:"status,omitempty"`
	Stream  *Stream           `protobuf:"bytes,5,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nsp_v1_nsp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_api_nsp_v1_nsp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{1}
}

func (x *Target) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *Target) GetType() Target_Type {
	if x != nil {
		return x.Type
	}
	return Target_DEFAULT
}

func (x *Target) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Target) GetStatus() Target_Status {
	if x != nil {
		return x.Status
	}
	return Target_Enabled
}

func (x *Target) GetStream() *Stream {
	if x != nil {
		return x.Stream
	}
	return nil
}

type TargetType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Target_Type `protobuf:"varint,1,opt,name=type,proto3,enum=nsp.v1.Target_Type" json:"type,omitempty"`
}

func (x *TargetType) Reset() {
	*x = TargetType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nsp_v1_nsp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetType) ProtoMessage() {}

func (x *TargetType) ProtoReflect() protoreflect.Message {
	mi := &file_api_nsp_v1_nsp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetType.ProtoReflect.Descriptor instead.
func (*TargetType) Descriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{2}
}

func (x *TargetType) GetType() Target_Type {
	if x != nil {
		return x.Type
	}
	return Target_DEFAULT
}

type TargetEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *Target            `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Status TargetEvent_Status `protobuf:"varint,2,opt,name=status,proto3,enum=nsp.v1.TargetEvent_Status" json:"status,omitempty"`
}

func (x *TargetEvent) Reset() {
	*x = TargetEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_nsp_v1_nsp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetEvent) ProtoMessage() {}

func (x *TargetEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_nsp_v1_nsp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetEvent.ProtoReflect.Descriptor instead.
func (*TargetEvent) Descriptor() ([]byte, []int) {
	return file_api_nsp_v1_nsp_proto_rawDescGZIP(), []int{3}
}

func (x *TargetEvent) GetTarget() *Target {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *TargetEvent) GetStatus() TargetEvent_Status {
	if x != nil {
		return x.Status
	}
	return TargetEvent_Register
}

var File_api_nsp_v1_nsp_proto protoreflect.FileDescriptor

var file_api_nsp_v1_nsp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69,
	0x2f, 0x6e, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73,
	0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x26, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x22, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x46, 0x52, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x22, 0x35, 0x0a, 0x0a, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x10, 0x02, 0x32, 0xba, 0x02, 0x0a, 0x1e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0a,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x6e, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x13, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x6e, 0x73, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x1a,
	0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x6e, 0x73, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e,
	0x6f, 0x72, 0x64, 0x69, 0x78, 0x2f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x73, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_nsp_v1_nsp_proto_rawDescOnce sync.Once
	file_api_nsp_v1_nsp_proto_rawDescData = file_api_nsp_v1_nsp_proto_rawDesc
)

func file_api_nsp_v1_nsp_proto_rawDescGZIP() []byte {
	file_api_nsp_v1_nsp_proto_rawDescOnce.Do(func() {
		file_api_nsp_v1_nsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_nsp_v1_nsp_proto_rawDescData)
	})
	return file_api_nsp_v1_nsp_proto_rawDescData
}

var file_api_nsp_v1_nsp_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_nsp_v1_nsp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_nsp_v1_nsp_proto_goTypes = []interface{}{
	(Target_Status)(0),         // 0: nsp.v1.Target.Status
	(Target_Type)(0),           // 1: nsp.v1.Target.Type
	(TargetEvent_Status)(0),    // 2: nsp.v1.TargetEvent.Status
	(*GetTargetsResponse)(nil), // 3: nsp.v1.GetTargetsResponse
	(*Target)(nil),             // 4: nsp.v1.Target
	(*TargetType)(nil),         // 5: nsp.v1.TargetType
	(*TargetEvent)(nil),        // 6: nsp.v1.TargetEvent
	nil,                        // 7: nsp.v1.Target.ContextEntry
	(*Stream)(nil),             // 8: nsp.v1.Stream
	(*empty.Empty)(nil),        // 9: google.protobuf.Empty
}
var file_api_nsp_v1_nsp_proto_depIdxs = []int32{
	4,  // 0: nsp.v1.GetTargetsResponse.targets:type_name -> nsp.v1.Target
	1,  // 1: nsp.v1.Target.type:type_name -> nsp.v1.Target.Type
	7,  // 2: nsp.v1.Target.context:type_name -> nsp.v1.Target.ContextEntry
	0,  // 3: nsp.v1.Target.status:type_name -> nsp.v1.Target.Status
	8,  // 4: nsp.v1.Target.stream:type_name -> nsp.v1.Stream
	1,  // 5: nsp.v1.TargetType.type:type_name -> nsp.v1.Target.Type
	4,  // 6: nsp.v1.TargetEvent.target:type_name -> nsp.v1.Target
	2,  // 7: nsp.v1.TargetEvent.status:type_name -> nsp.v1.TargetEvent.Status
	4,  // 8: nsp.v1.NetworkServicePlateformService.Register:input_type -> nsp.v1.Target
	4,  // 9: nsp.v1.NetworkServicePlateformService.Unregister:input_type -> nsp.v1.Target
	5,  // 10: nsp.v1.NetworkServicePlateformService.Monitor:input_type -> nsp.v1.TargetType
	5,  // 11: nsp.v1.NetworkServicePlateformService.GetTargets:input_type -> nsp.v1.TargetType
	4,  // 12: nsp.v1.NetworkServicePlateformService.Update:input_type -> nsp.v1.Target
	9,  // 13: nsp.v1.NetworkServicePlateformService.Register:output_type -> google.protobuf.Empty
	9,  // 14: nsp.v1.NetworkServicePlateformService.Unregister:output_type -> google.protobuf.Empty
	6,  // 15: nsp.v1.NetworkServicePlateformService.Monitor:output_type -> nsp.v1.TargetEvent
	3,  // 16: nsp.v1.NetworkServicePlateformService.GetTargets:output_type -> nsp.v1.GetTargetsResponse
	9,  // 17: nsp.v1.NetworkServicePlateformService.Update:output_type -> google.protobuf.Empty
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_nsp_v1_nsp_proto_init() }
func file_api_nsp_v1_nsp_proto_init() {
	if File_api_nsp_v1_nsp_proto != nil {
		return
	}
	file_api_nsp_v1_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_nsp_v1_nsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTargetsResponse); i {
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
		file_api_nsp_v1_nsp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
		file_api_nsp_v1_nsp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetType); i {
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
		file_api_nsp_v1_nsp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetEvent); i {
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
			RawDescriptor: file_api_nsp_v1_nsp_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_nsp_v1_nsp_proto_goTypes,
		DependencyIndexes: file_api_nsp_v1_nsp_proto_depIdxs,
		EnumInfos:         file_api_nsp_v1_nsp_proto_enumTypes,
		MessageInfos:      file_api_nsp_v1_nsp_proto_msgTypes,
	}.Build()
	File_api_nsp_v1_nsp_proto = out.File
	file_api_nsp_v1_nsp_proto_rawDesc = nil
	file_api_nsp_v1_nsp_proto_goTypes = nil
	file_api_nsp_v1_nsp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NetworkServicePlateformServiceClient is the client API for NetworkServicePlateformService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServicePlateformServiceClient interface {
	Register(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error)
	Unregister(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error)
	Monitor(ctx context.Context, in *TargetType, opts ...grpc.CallOption) (NetworkServicePlateformService_MonitorClient, error)
	GetTargets(ctx context.Context, in *TargetType, opts ...grpc.CallOption) (*GetTargetsResponse, error)
	Update(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error)
}

type networkServicePlateformServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServicePlateformServiceClient(cc grpc.ClientConnInterface) NetworkServicePlateformServiceClient {
	return &networkServicePlateformServiceClient{cc}
}

func (c *networkServicePlateformServiceClient) Register(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nsp.v1.NetworkServicePlateformService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServicePlateformServiceClient) Unregister(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nsp.v1.NetworkServicePlateformService/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServicePlateformServiceClient) Monitor(ctx context.Context, in *TargetType, opts ...grpc.CallOption) (NetworkServicePlateformService_MonitorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkServicePlateformService_serviceDesc.Streams[0], "/nsp.v1.NetworkServicePlateformService/Monitor", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkServicePlateformServiceMonitorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkServicePlateformService_MonitorClient interface {
	Recv() (*TargetEvent, error)
	grpc.ClientStream
}

type networkServicePlateformServiceMonitorClient struct {
	grpc.ClientStream
}

func (x *networkServicePlateformServiceMonitorClient) Recv() (*TargetEvent, error) {
	m := new(TargetEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkServicePlateformServiceClient) GetTargets(ctx context.Context, in *TargetType, opts ...grpc.CallOption) (*GetTargetsResponse, error) {
	out := new(GetTargetsResponse)
	err := c.cc.Invoke(ctx, "/nsp.v1.NetworkServicePlateformService/GetTargets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServicePlateformServiceClient) Update(ctx context.Context, in *Target, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/nsp.v1.NetworkServicePlateformService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServicePlateformServiceServer is the server API for NetworkServicePlateformService service.
type NetworkServicePlateformServiceServer interface {
	Register(context.Context, *Target) (*empty.Empty, error)
	Unregister(context.Context, *Target) (*empty.Empty, error)
	Monitor(*TargetType, NetworkServicePlateformService_MonitorServer) error
	GetTargets(context.Context, *TargetType) (*GetTargetsResponse, error)
	Update(context.Context, *Target) (*empty.Empty, error)
}

// UnimplementedNetworkServicePlateformServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServicePlateformServiceServer struct {
}

func (*UnimplementedNetworkServicePlateformServiceServer) Register(context.Context, *Target) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedNetworkServicePlateformServiceServer) Unregister(context.Context, *Target) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}
func (*UnimplementedNetworkServicePlateformServiceServer) Monitor(*TargetType, NetworkServicePlateformService_MonitorServer) error {
	return status.Errorf(codes.Unimplemented, "method Monitor not implemented")
}
func (*UnimplementedNetworkServicePlateformServiceServer) GetTargets(context.Context, *TargetType) (*GetTargetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTargets not implemented")
}
func (*UnimplementedNetworkServicePlateformServiceServer) Update(context.Context, *Target) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterNetworkServicePlateformServiceServer(s *grpc.Server, srv NetworkServicePlateformServiceServer) {
	s.RegisterService(&_NetworkServicePlateformService_serviceDesc, srv)
}

func _NetworkServicePlateformService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServicePlateformServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.NetworkServicePlateformService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServicePlateformServiceServer).Register(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServicePlateformService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServicePlateformServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.NetworkServicePlateformService/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServicePlateformServiceServer).Unregister(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServicePlateformService_Monitor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TargetType)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkServicePlateformServiceServer).Monitor(m, &networkServicePlateformServiceMonitorServer{stream})
}

type NetworkServicePlateformService_MonitorServer interface {
	Send(*TargetEvent) error
	grpc.ServerStream
}

type networkServicePlateformServiceMonitorServer struct {
	grpc.ServerStream
}

func (x *networkServicePlateformServiceMonitorServer) Send(m *TargetEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _NetworkServicePlateformService_GetTargets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TargetType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServicePlateformServiceServer).GetTargets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.NetworkServicePlateformService/GetTargets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServicePlateformServiceServer).GetTargets(ctx, req.(*TargetType))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServicePlateformService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServicePlateformServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsp.v1.NetworkServicePlateformService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServicePlateformServiceServer).Update(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServicePlateformService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nsp.v1.NetworkServicePlateformService",
	HandlerType: (*NetworkServicePlateformServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _NetworkServicePlateformService_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _NetworkServicePlateformService_Unregister_Handler,
		},
		{
			MethodName: "GetTargets",
			Handler:    _NetworkServicePlateformService_GetTargets_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkServicePlateformService_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Monitor",
			Handler:       _NetworkServicePlateformService_Monitor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/nsp/v1/nsp.proto",
}
