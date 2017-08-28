// Code generated by protoc-gen-go. DO NOT EDIT.
// source: environment.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListEnvironmentRequest struct {
	ApplicationId int32 `protobuf:"varint,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
}

func (m *ListEnvironmentRequest) Reset()                    { *m = ListEnvironmentRequest{} }
func (m *ListEnvironmentRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListEnvironmentRequest) ProtoMessage()               {}
func (*ListEnvironmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ListEnvironmentRequest) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

type ListEnvironmentResponse struct {
	Environments []*Environment `protobuf:"bytes,1,rep,name=environments" json:"environments,omitempty"`
}

func (m *ListEnvironmentResponse) Reset()                    { *m = ListEnvironmentResponse{} }
func (m *ListEnvironmentResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListEnvironmentResponse) ProtoMessage()               {}
func (*ListEnvironmentResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListEnvironmentResponse) GetEnvironments() []*Environment {
	if m != nil {
		return m.Environments
	}
	return nil
}

type GetEnvironmentRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetEnvironmentRequest) Reset()                    { *m = GetEnvironmentRequest{} }
func (m *GetEnvironmentRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetEnvironmentRequest) ProtoMessage()               {}
func (*GetEnvironmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetEnvironmentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Environment struct {
	Id            int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ApplicationId int32                      `protobuf:"varint,2,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	Name          string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Slug          string                     `protobuf:"bytes,4,opt,name=slug" json:"slug,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *Environment) Reset()                    { *m = Environment{} }
func (m *Environment) String() string            { return proto1.CompactTextString(m) }
func (*Environment) ProtoMessage()               {}
func (*Environment) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Environment) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Environment) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Environment) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Environment) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type DestroyEnvironmentRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DestroyEnvironmentRequest) Reset()                    { *m = DestroyEnvironmentRequest{} }
func (m *DestroyEnvironmentRequest) String() string            { return proto1.CompactTextString(m) }
func (*DestroyEnvironmentRequest) ProtoMessage()               {}
func (*DestroyEnvironmentRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DestroyEnvironmentRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto1.RegisterType((*ListEnvironmentRequest)(nil), "soapbox.ListEnvironmentRequest")
	proto1.RegisterType((*ListEnvironmentResponse)(nil), "soapbox.ListEnvironmentResponse")
	proto1.RegisterType((*GetEnvironmentRequest)(nil), "soapbox.GetEnvironmentRequest")
	proto1.RegisterType((*Environment)(nil), "soapbox.Environment")
	proto1.RegisterType((*DestroyEnvironmentRequest)(nil), "soapbox.DestroyEnvironmentRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Environments service

type EnvironmentsClient interface {
	ListEnvironments(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error)
	GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*Environment, error)
	CreateEnvironment(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error)
	DestroyEnvironment(ctx context.Context, in *DestroyEnvironmentRequest, opts ...grpc.CallOption) (*Empty, error)
}

type environmentsClient struct {
	cc *grpc.ClientConn
}

func NewEnvironmentsClient(cc *grpc.ClientConn) EnvironmentsClient {
	return &environmentsClient{cc}
}

func (c *environmentsClient) ListEnvironments(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error) {
	out := new(ListEnvironmentResponse)
	err := grpc.Invoke(ctx, "/soapbox.Environments/ListEnvironments", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) GetEnvironment(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := grpc.Invoke(ctx, "/soapbox.Environments/GetEnvironment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) CreateEnvironment(ctx context.Context, in *Environment, opts ...grpc.CallOption) (*Environment, error) {
	out := new(Environment)
	err := grpc.Invoke(ctx, "/soapbox.Environments/CreateEnvironment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentsClient) DestroyEnvironment(ctx context.Context, in *DestroyEnvironmentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/soapbox.Environments/DestroyEnvironment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Environments service

type EnvironmentsServer interface {
	ListEnvironments(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error)
	GetEnvironment(context.Context, *GetEnvironmentRequest) (*Environment, error)
	CreateEnvironment(context.Context, *Environment) (*Environment, error)
	DestroyEnvironment(context.Context, *DestroyEnvironmentRequest) (*Empty, error)
}

func RegisterEnvironmentsServer(s *grpc.Server, srv EnvironmentsServer) {
	s.RegisterService(&_Environments_serviceDesc, srv)
}

func _Environments_ListEnvironments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Environments/ListEnvironments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).ListEnvironments(ctx, req.(*ListEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_GetEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).GetEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Environments/GetEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).GetEnvironment(ctx, req.(*GetEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_CreateEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Environment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).CreateEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Environments/CreateEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).CreateEnvironment(ctx, req.(*Environment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Environments_DestroyEnvironment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentsServer).DestroyEnvironment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Environments/DestroyEnvironment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentsServer).DestroyEnvironment(ctx, req.(*DestroyEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Environments_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soapbox.Environments",
	HandlerType: (*EnvironmentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEnvironments",
			Handler:    _Environments_ListEnvironments_Handler,
		},
		{
			MethodName: "GetEnvironment",
			Handler:    _Environments_GetEnvironment_Handler,
		},
		{
			MethodName: "CreateEnvironment",
			Handler:    _Environments_CreateEnvironment_Handler,
		},
		{
			MethodName: "DestroyEnvironment",
			Handler:    _Environments_DestroyEnvironment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "environment.proto",
}

func init() { proto1.RegisterFile("environment.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcb, 0x6a, 0xfa, 0x40,
	0x14, 0xc6, 0x4d, 0xbc, 0xe1, 0x51, 0xc3, 0xdf, 0xf9, 0xf7, 0x92, 0x66, 0x51, 0xc3, 0x40, 0x69,
	0xa0, 0x10, 0xc1, 0x6e, 0xda, 0x55, 0xb1, 0x17, 0x7a, 0xa1, 0xab, 0xb4, 0x50, 0xe8, 0x46, 0xa2,
	0x19, 0xc3, 0x80, 0xc9, 0xa4, 0x99, 0x49, 0xa9, 0x4f, 0xd4, 0xd7, 0xe8, 0xa3, 0x15, 0x13, 0x8d,
	0x51, 0x47, 0xba, 0xca, 0xcc, 0x39, 0xdf, 0x7c, 0x39, 0xe7, 0xf7, 0x41, 0x87, 0x84, 0x9f, 0x34,
	0x66, 0x61, 0x40, 0x42, 0x61, 0x47, 0x31, 0x13, 0x0c, 0xd5, 0x39, 0x73, 0xa3, 0x11, 0xfb, 0x32,
	0xfe, 0x8f, 0x59, 0x38, 0xa1, 0x7e, 0x12, 0xbb, 0x82, 0xb2, 0x30, 0xeb, 0x1a, 0xed, 0x45, 0x77,
	0x71, 0xed, 0xfa, 0x8c, 0xf9, 0x53, 0xd2, 0x4b, 0x6f, 0xa3, 0x64, 0xd2, 0x13, 0x34, 0x20, 0x5c,
	0xb8, 0x41, 0x94, 0x09, 0xf0, 0x15, 0x1c, 0x3c, 0x53, 0x2e, 0xee, 0x56, 0xbf, 0x71, 0xc8, 0x47,
	0x42, 0xb8, 0x40, 0x27, 0xa0, 0xb9, 0x51, 0x34, 0xa5, 0xe3, 0xd4, 0x7e, 0x48, 0x3d, 0x5d, 0x31,
	0x15, 0xab, 0xea, 0xb4, 0x0b, 0xd5, 0x47, 0x0f, 0xbf, 0xc0, 0xe1, 0x96, 0x01, 0x8f, 0x58, 0xc8,
	0x09, 0xba, 0x80, 0x56, 0x61, 0x7c, 0xae, 0x2b, 0x66, 0xd9, 0x6a, 0xf6, 0xf7, 0xec, 0xe5, 0x88,
	0xc5, 0x37, 0x6b, 0x4a, 0x7c, 0x0a, 0xfb, 0xf7, 0x44, 0x36, 0x94, 0x06, 0x6a, 0x3e, 0x88, 0x4a,
	0x3d, 0xfc, 0xad, 0x40, 0xb3, 0x20, 0xdb, 0xec, 0x4b, 0x96, 0x50, 0x25, 0x4b, 0x20, 0x04, 0x95,
	0xd0, 0x0d, 0x88, 0x5e, 0x36, 0x15, 0xab, 0xe1, 0xa4, 0xe7, 0x79, 0x8d, 0x4f, 0x13, 0x5f, 0xaf,
	0x64, 0xb5, 0xf9, 0x19, 0x5d, 0x02, 0x8c, 0x63, 0xe2, 0x0a, 0xe2, 0x0d, 0x5d, 0xa1, 0xd7, 0x4c,
	0xc5, 0x6a, 0xf6, 0x0d, 0x3b, 0x63, 0x6c, 0x2f, 0x19, 0xdb, 0xaf, 0x4b, 0xc6, 0x4e, 0x63, 0xa1,
	0x1e, 0x08, 0x7c, 0x06, 0x47, 0xb7, 0x84, 0x8b, 0x98, 0xcd, 0xfe, 0x5e, 0xab, 0xff, 0xa3, 0x42,
	0xab, 0x20, 0xe3, 0xe8, 0x0d, 0xfe, 0x6d, 0x50, 0xe6, 0xa8, 0x9b, 0x83, 0x94, 0x27, 0x68, 0x98,
	0xbb, 0x05, 0x59, 0x42, 0xb8, 0x84, 0x1e, 0x40, 0x5b, 0x27, 0x8d, 0x8e, 0xf3, 0x57, 0xd2, 0x08,
	0x0c, 0x69, 0x7e, 0xb8, 0x84, 0x06, 0xd0, 0xb9, 0x49, 0xb7, 0x2d, 0x9a, 0x49, 0xc5, 0x3b, 0x2d,
	0x9e, 0x00, 0x6d, 0x33, 0x42, 0x38, 0x57, 0xef, 0x04, 0x68, 0x68, 0x2b, 0xc7, 0x20, 0x12, 0x33,
	0x5c, 0xba, 0xae, 0xbf, 0x57, 0xb3, 0x40, 0x6a, 0xe9, 0xe7, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x17, 0xed, 0x90, 0x64, 0x4b, 0x03, 0x00, 0x00,
}
