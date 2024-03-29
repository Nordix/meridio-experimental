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
// source: api/ipam/ipam.proto

package ipam

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

type Subnet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength int32  `protobuf:"varint,2,opt,name=prefixLength,proto3" json:"prefixLength,omitempty"`
}

func (x *Subnet) Reset() {
	*x = Subnet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ipam_ipam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subnet) ProtoMessage() {}

func (x *Subnet) ProtoReflect() protoreflect.Message {
	mi := &file_api_ipam_ipam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subnet.ProtoReflect.Descriptor instead.
func (*Subnet) Descriptor() ([]byte, []int) {
	return file_api_ipam_ipam_proto_rawDescGZIP(), []int{0}
}

func (x *Subnet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Subnet) GetPrefixLength() int32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

type SubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetPool   *Subnet `protobuf:"bytes,1,opt,name=subnetPool,proto3" json:"subnetPool,omitempty"`
	PrefixLength int32   `protobuf:"varint,2,opt,name=prefixLength,proto3" json:"prefixLength,omitempty"`
}

func (x *SubnetRequest) Reset() {
	*x = SubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ipam_ipam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetRequest) ProtoMessage() {}

func (x *SubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ipam_ipam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetRequest.ProtoReflect.Descriptor instead.
func (*SubnetRequest) Descriptor() ([]byte, []int) {
	return file_api_ipam_ipam_proto_rawDescGZIP(), []int{1}
}

func (x *SubnetRequest) GetSubnetPool() *Subnet {
	if x != nil {
		return x.SubnetPool
	}
	return nil
}

func (x *SubnetRequest) GetPrefixLength() int32 {
	if x != nil {
		return x.PrefixLength
	}
	return 0
}

type SubnetRelease struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetPool *Subnet `protobuf:"bytes,1,opt,name=subnetPool,proto3" json:"subnetPool,omitempty"`
	Subnet     *Subnet `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *SubnetRelease) Reset() {
	*x = SubnetRelease{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ipam_ipam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetRelease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetRelease) ProtoMessage() {}

func (x *SubnetRelease) ProtoReflect() protoreflect.Message {
	mi := &file_api_ipam_ipam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetRelease.ProtoReflect.Descriptor instead.
func (*SubnetRelease) Descriptor() ([]byte, []int) {
	return file_api_ipam_ipam_proto_rawDescGZIP(), []int{2}
}

func (x *SubnetRelease) GetSubnetPool() *Subnet {
	if x != nil {
		return x.SubnetPool
	}
	return nil
}

func (x *SubnetRelease) GetSubnet() *Subnet {
	if x != nil {
		return x.Subnet
	}
	return nil
}

var File_api_ipam_ipam_proto protoreflect.FileDescriptor

var file_api_ipam_ipam_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x70, 0x61, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x61, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0x63, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x32, 0x78, 0x0a, 0x0b, 0x49, 0x70, 0x61, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x69, 0x70, 0x61, 0x6d, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x69, 0x78, 0x2f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x70, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ipam_ipam_proto_rawDescOnce sync.Once
	file_api_ipam_ipam_proto_rawDescData = file_api_ipam_ipam_proto_rawDesc
)

func file_api_ipam_ipam_proto_rawDescGZIP() []byte {
	file_api_ipam_ipam_proto_rawDescOnce.Do(func() {
		file_api_ipam_ipam_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ipam_ipam_proto_rawDescData)
	})
	return file_api_ipam_ipam_proto_rawDescData
}

var file_api_ipam_ipam_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_ipam_ipam_proto_goTypes = []interface{}{
	(*Subnet)(nil),        // 0: ipam.Subnet
	(*SubnetRequest)(nil), // 1: ipam.SubnetRequest
	(*SubnetRelease)(nil), // 2: ipam.SubnetRelease
	(*empty.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_api_ipam_ipam_proto_depIdxs = []int32{
	0, // 0: ipam.SubnetRequest.subnetPool:type_name -> ipam.Subnet
	0, // 1: ipam.SubnetRelease.subnetPool:type_name -> ipam.Subnet
	0, // 2: ipam.SubnetRelease.subnet:type_name -> ipam.Subnet
	1, // 3: ipam.IpamService.Allocate:input_type -> ipam.SubnetRequest
	2, // 4: ipam.IpamService.Release:input_type -> ipam.SubnetRelease
	0, // 5: ipam.IpamService.Allocate:output_type -> ipam.Subnet
	3, // 6: ipam.IpamService.Release:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_ipam_ipam_proto_init() }
func file_api_ipam_ipam_proto_init() {
	if File_api_ipam_ipam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ipam_ipam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subnet); i {
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
		file_api_ipam_ipam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetRequest); i {
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
		file_api_ipam_ipam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetRelease); i {
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
			RawDescriptor: file_api_ipam_ipam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ipam_ipam_proto_goTypes,
		DependencyIndexes: file_api_ipam_ipam_proto_depIdxs,
		MessageInfos:      file_api_ipam_ipam_proto_msgTypes,
	}.Build()
	File_api_ipam_ipam_proto = out.File
	file_api_ipam_ipam_proto_rawDesc = nil
	file_api_ipam_ipam_proto_goTypes = nil
	file_api_ipam_ipam_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IpamServiceClient is the client API for IpamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IpamServiceClient interface {
	Allocate(ctx context.Context, in *SubnetRequest, opts ...grpc.CallOption) (*Subnet, error)
	Release(ctx context.Context, in *SubnetRelease, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ipamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIpamServiceClient(cc grpc.ClientConnInterface) IpamServiceClient {
	return &ipamServiceClient{cc}
}

func (c *ipamServiceClient) Allocate(ctx context.Context, in *SubnetRequest, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := c.cc.Invoke(ctx, "/ipam.IpamService/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ipamServiceClient) Release(ctx context.Context, in *SubnetRelease, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ipam.IpamService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IpamServiceServer is the server API for IpamService service.
type IpamServiceServer interface {
	Allocate(context.Context, *SubnetRequest) (*Subnet, error)
	Release(context.Context, *SubnetRelease) (*empty.Empty, error)
}

// UnimplementedIpamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIpamServiceServer struct {
}

func (*UnimplementedIpamServiceServer) Allocate(context.Context, *SubnetRequest) (*Subnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (*UnimplementedIpamServiceServer) Release(context.Context, *SubnetRelease) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}

func RegisterIpamServiceServer(s *grpc.Server, srv IpamServiceServer) {
	s.RegisterService(&_IpamService_serviceDesc, srv)
}

func _IpamService_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamServiceServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipam.IpamService/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamServiceServer).Allocate(ctx, req.(*SubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IpamService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubnetRelease)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IpamServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipam.IpamService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IpamServiceServer).Release(ctx, req.(*SubnetRelease))
	}
	return interceptor(ctx, in, info, handler)
}

var _IpamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ipam.IpamService",
	HandlerType: (*IpamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allocate",
			Handler:    _IpamService_Allocate_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _IpamService_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ipam/ipam.proto",
}
