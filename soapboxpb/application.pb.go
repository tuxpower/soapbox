// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application.proto

package soapboxpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ApplicationType int32

const (
	ApplicationType_SERVER  ApplicationType = 0
	ApplicationType_CRONJOB ApplicationType = 1
)

var ApplicationType_name = map[int32]string{
	0: "SERVER",
	1: "CRONJOB",
}
var ApplicationType_value = map[string]int32{
	"SERVER":  0,
	"CRONJOB": 1,
}

func (x ApplicationType) String() string {
	return proto.EnumName(ApplicationType_name, int32(x))
}
func (ApplicationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type CreationState int32

const (
	CreationState_CREATE_INFRASTRUCTURE_WAIT CreationState = 0
	CreationState_SUCCEEDED                  CreationState = 1
	CreationState_FAILED                     CreationState = 2
)

var CreationState_name = map[int32]string{
	0: "CREATE_INFRASTRUCTURE_WAIT",
	1: "SUCCEEDED",
	2: "FAILED",
}
var CreationState_value = map[string]int32{
	"CREATE_INFRASTRUCTURE_WAIT": 0,
	"SUCCEEDED":                  1,
	"FAILED":                     2,
}

func (x CreationState) String() string {
	return proto.EnumName(CreationState_name, int32(x))
}
func (CreationState) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type Application struct {
	Id                 int32           `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name               string          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description        string          `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	ExternalDns        string          `protobuf:"bytes,4,opt,name=external_dns,json=externalDns" json:"external_dns,omitempty"`
	GithubRepoUrl      string          `protobuf:"bytes,5,opt,name=github_repo_url,json=githubRepoUrl" json:"github_repo_url,omitempty"`
	DockerfilePath     string          `protobuf:"bytes,6,opt,name=dockerfile_path,json=dockerfilePath" json:"dockerfile_path,omitempty"`
	EntrypointOverride string          `protobuf:"bytes,7,opt,name=entrypoint_override,json=entrypointOverride" json:"entrypoint_override,omitempty"`
	Type               ApplicationType `protobuf:"varint,8,opt,name=type,enum=soapbox.ApplicationType" json:"type,omitempty"`
	CreatedAt          string          `protobuf:"bytes,9,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Slug               string          `protobuf:"bytes,10,opt,name=slug" json:"slug,omitempty"`
	InternalDns        string          `protobuf:"bytes,11,opt,name=internal_dns,json=internalDns" json:"internal_dns,omitempty"`
	CreationState      CreationState   `protobuf:"varint,12,opt,name=creation_state,json=creationState,enum=soapbox.CreationState" json:"creation_state,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Application) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Application) GetExternalDns() string {
	if m != nil {
		return m.ExternalDns
	}
	return ""
}

func (m *Application) GetGithubRepoUrl() string {
	if m != nil {
		return m.GithubRepoUrl
	}
	return ""
}

func (m *Application) GetDockerfilePath() string {
	if m != nil {
		return m.DockerfilePath
	}
	return ""
}

func (m *Application) GetEntrypointOverride() string {
	if m != nil {
		return m.EntrypointOverride
	}
	return ""
}

func (m *Application) GetType() ApplicationType {
	if m != nil {
		return m.Type
	}
	return ApplicationType_SERVER
}

func (m *Application) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Application) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Application) GetInternalDns() string {
	if m != nil {
		return m.InternalDns
	}
	return ""
}

func (m *Application) GetCreationState() CreationState {
	if m != nil {
		return m.CreationState
	}
	return CreationState_CREATE_INFRASTRUCTURE_WAIT
}

type ListApplicationResponse struct {
	Applications []*Application `protobuf:"bytes,1,rep,name=applications" json:"applications,omitempty"`
}

func (m *ListApplicationResponse) Reset()                    { *m = ListApplicationResponse{} }
func (m *ListApplicationResponse) String() string            { return proto.CompactTextString(m) }
func (*ListApplicationResponse) ProtoMessage()               {}
func (*ListApplicationResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListApplicationResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

type GetApplicationRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetApplicationRequest) Reset()                    { *m = GetApplicationRequest{} }
func (m *GetApplicationRequest) String() string            { return proto.CompactTextString(m) }
func (*GetApplicationRequest) ProtoMessage()               {}
func (*GetApplicationRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetApplicationRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Application)(nil), "soapbox.Application")
	proto.RegisterType((*ListApplicationResponse)(nil), "soapbox.ListApplicationResponse")
	proto.RegisterType((*GetApplicationRequest)(nil), "soapbox.GetApplicationRequest")
	proto.RegisterEnum("soapbox.ApplicationType", ApplicationType_name, ApplicationType_value)
	proto.RegisterEnum("soapbox.CreationState", CreationState_name, CreationState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Applications service

type ApplicationsClient interface {
	ListApplications(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListApplicationResponse, error)
	CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error)
	GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error)
}

type applicationsClient struct {
	cc *grpc.ClientConn
}

func NewApplicationsClient(cc *grpc.ClientConn) ApplicationsClient {
	return &applicationsClient{cc}
}

func (c *applicationsClient) ListApplications(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListApplicationResponse, error) {
	out := new(ListApplicationResponse)
	err := grpc.Invoke(ctx, "/soapbox.Applications/ListApplications", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/soapbox.Applications/CreateApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationsClient) GetApplication(ctx context.Context, in *GetApplicationRequest, opts ...grpc.CallOption) (*Application, error) {
	out := new(Application)
	err := grpc.Invoke(ctx, "/soapbox.Applications/GetApplication", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Applications service

type ApplicationsServer interface {
	ListApplications(context.Context, *Empty) (*ListApplicationResponse, error)
	CreateApplication(context.Context, *Application) (*Application, error)
	GetApplication(context.Context, *GetApplicationRequest) (*Application, error)
}

func RegisterApplicationsServer(s *grpc.Server, srv ApplicationsServer) {
	s.RegisterService(&_Applications_serviceDesc, srv)
}

func _Applications_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Applications/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).ListApplications(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Applications/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).CreateApplication(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _Applications_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationsServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soapbox.Applications/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationsServer).GetApplication(ctx, req.(*GetApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Applications_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soapbox.Applications",
	HandlerType: (*ApplicationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListApplications",
			Handler:    _Applications_ListApplications_Handler,
		},
		{
			MethodName: "CreateApplication",
			Handler:    _Applications_CreateApplication_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _Applications_GetApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "application.proto",
}

func init() { proto.RegisterFile("application.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5d, 0x6f, 0xda, 0x3e,
	0x14, 0xc6, 0x09, 0x50, 0xf8, 0x73, 0x02, 0x21, 0xf5, 0x7f, 0x2f, 0x16, 0xd2, 0xaa, 0x8c, 0x8b,
	0x15, 0xa1, 0xa9, 0x93, 0xd8, 0xcd, 0x6e, 0x76, 0x91, 0x86, 0xb0, 0x51, 0x55, 0x65, 0x32, 0xb0,
	0x49, 0xbb, 0x89, 0x02, 0xf1, 0x8a, 0xb5, 0x34, 0xf6, 0x62, 0x33, 0x95, 0xcf, 0xb6, 0x8f, 0xb3,
	0x2f, 0x32, 0xc5, 0xbc, 0x05, 0x44, 0xef, 0x8e, 0x9f, 0xe7, 0xc9, 0xc9, 0x39, 0x3f, 0xd9, 0x70,
	0x1e, 0x0a, 0x11, 0xb3, 0x79, 0xa8, 0x18, 0x4f, 0xae, 0x44, 0xca, 0x15, 0x47, 0x55, 0xc9, 0x43,
	0x31, 0xe3, 0x8f, 0xad, 0xc6, 0xa6, 0x58, 0xeb, 0xed, 0x3f, 0x25, 0x30, 0xdd, 0x7d, 0x1a, 0x59,
	0x50, 0x64, 0x11, 0x36, 0x1c, 0xa3, 0x73, 0x46, 0x8a, 0x2c, 0x42, 0x08, 0xca, 0x49, 0xf8, 0x40,
	0x71, 0xd1, 0x31, 0x3a, 0x35, 0xa2, 0x6b, 0xe4, 0x80, 0x19, 0x51, 0x39, 0x4f, 0x99, 0xc8, 0x3e,
	0xc1, 0x25, 0x6d, 0xe5, 0x25, 0xf4, 0x1a, 0xea, 0xf4, 0x51, 0xd1, 0x34, 0x09, 0xe3, 0x20, 0x4a,
	0x24, 0x2e, 0xaf, 0x23, 0x5b, 0xad, 0x9f, 0x48, 0xf4, 0x06, 0x9a, 0xf7, 0x4c, 0x2d, 0x96, 0xb3,
	0x20, 0xa5, 0x82, 0x07, 0xcb, 0x34, 0xc6, 0x67, 0x3a, 0xd5, 0x58, 0xcb, 0x84, 0x0a, 0x3e, 0x4d,
	0x63, 0x74, 0x09, 0xcd, 0x88, 0xcf, 0x7f, 0xd2, 0xf4, 0x07, 0x8b, 0x69, 0x20, 0x42, 0xb5, 0xc0,
	0x15, 0x9d, 0xb3, 0xf6, 0xf2, 0x97, 0x50, 0x2d, 0xd0, 0x3b, 0xf8, 0x9f, 0x26, 0x2a, 0x5d, 0x09,
	0xce, 0x12, 0x15, 0xf0, 0xdf, 0x34, 0x4d, 0x59, 0x44, 0x71, 0x55, 0x87, 0xd1, 0xde, 0x1a, 0x6d,
	0x1c, 0xf4, 0x16, 0xca, 0x6a, 0x25, 0x28, 0xfe, 0xcf, 0x31, 0x3a, 0x56, 0x0f, 0x5f, 0x6d, 0xc1,
	0xe4, 0x70, 0x4c, 0x56, 0x82, 0x12, 0x9d, 0x42, 0xaf, 0x00, 0xe6, 0x29, 0x0d, 0x15, 0x8d, 0x82,
	0x50, 0xe1, 0x9a, 0xee, 0x5a, 0xdb, 0x28, 0xae, 0xca, 0x38, 0xc9, 0x78, 0x79, 0x8f, 0x61, 0xcd,
	0x29, 0xab, 0x33, 0x0a, 0x2c, 0xc9, 0x51, 0x30, 0xd7, 0x14, 0xb6, 0x5a, 0x46, 0xe1, 0x23, 0x58,
	0xba, 0x07, 0xe3, 0x49, 0x20, 0x55, 0xa8, 0x28, 0xae, 0xeb, 0x69, 0x5e, 0xec, 0xa6, 0xf1, 0x36,
	0xf6, 0x38, 0x73, 0x49, 0x63, 0x9e, 0x3f, 0xb6, 0xc7, 0xf0, 0xf2, 0x96, 0x49, 0x95, 0x9b, 0x98,
	0x50, 0x29, 0x78, 0x22, 0x29, 0xfa, 0x00, 0xf5, 0xdc, 0x2d, 0x90, 0xd8, 0x70, 0x4a, 0x1d, 0xb3,
	0xf7, 0xec, 0xd4, 0x96, 0xe4, 0x20, 0xd9, 0xbe, 0x84, 0xe7, 0x9f, 0xe8, 0x61, 0xcf, 0x5f, 0x4b,
	0x2a, 0xd5, 0xf1, 0xdd, 0xe8, 0x76, 0xa1, 0x79, 0xc4, 0x0a, 0x01, 0x54, 0xc6, 0x3e, 0xf9, 0xea,
	0x13, 0xbb, 0x80, 0x4c, 0xa8, 0x7a, 0x64, 0x74, 0x77, 0x33, 0xba, 0xb6, 0x8d, 0xee, 0x0d, 0x34,
	0x0e, 0x36, 0x41, 0x17, 0xd0, 0xf2, 0x88, 0xef, 0x4e, 0xfc, 0x60, 0x78, 0x37, 0x20, 0xee, 0x78,
	0x42, 0xa6, 0xde, 0x64, 0x4a, 0xfc, 0xe0, 0x9b, 0x3b, 0x9c, 0xd8, 0x05, 0xd4, 0x80, 0xda, 0x78,
	0xea, 0x79, 0xbe, 0xdf, 0xf7, 0xfb, 0xb6, 0x91, 0x35, 0x1e, 0xb8, 0xc3, 0x5b, 0xbf, 0x6f, 0x17,
	0x7b, 0x7f, 0x0d, 0xa8, 0xe7, 0x7e, 0x2c, 0xd1, 0x00, 0xec, 0x23, 0x0c, 0x12, 0x59, 0xbb, 0x4d,
	0xfd, 0x07, 0xa1, 0x56, 0x2d, 0x67, 0x77, 0x7e, 0x82, 0x58, 0xbb, 0x80, 0x5c, 0x38, 0xd7, 0x43,
	0xd2, 0xfc, 0x8b, 0x38, 0x89, 0xac, 0x75, 0x52, 0x6d, 0x17, 0xd0, 0x67, 0xb0, 0x0e, 0xe1, 0xa1,
	0x8b, 0x5d, 0xf2, 0x24, 0xd5, 0xa7, 0x3a, 0x5d, 0x9b, 0xdf, 0x6b, 0x1b, 0x43, 0xcc, 0x66, 0x15,
	0xfd, 0x5a, 0xdf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x18, 0x33, 0xc5, 0xda, 0x03, 0x00,
	0x00,
}
