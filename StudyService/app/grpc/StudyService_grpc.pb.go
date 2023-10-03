// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: StudyService.proto

package GrpcStudyService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StudyServiceClient is the client API for StudyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudyServiceClient interface {
	GetUserRecord(ctx context.Context, in *GetUserRecordRequest, opts ...grpc.CallOption) (*GetUserRecordResponse, error)
	CreateUserRecord(ctx context.Context, in *CreateUserRecordRequest, opts ...grpc.CallOption) (*CreateUserRecordResponse, error)
}

type studyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudyServiceClient(cc grpc.ClientConnInterface) StudyServiceClient {
	return &studyServiceClient{cc}
}

func (c *studyServiceClient) GetUserRecord(ctx context.Context, in *GetUserRecordRequest, opts ...grpc.CallOption) (*GetUserRecordResponse, error) {
	out := new(GetUserRecordResponse)
	err := c.cc.Invoke(ctx, "/StudyService/GetUserRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceClient) CreateUserRecord(ctx context.Context, in *CreateUserRecordRequest, opts ...grpc.CallOption) (*CreateUserRecordResponse, error) {
	out := new(CreateUserRecordResponse)
	err := c.cc.Invoke(ctx, "/StudyService/CreateUserRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudyServiceServer is the server API for StudyService service.
// All implementations must embed UnimplementedStudyServiceServer
// for forward compatibility
type StudyServiceServer interface {
	GetUserRecord(context.Context, *GetUserRecordRequest) (*GetUserRecordResponse, error)
	CreateUserRecord(context.Context, *CreateUserRecordRequest) (*CreateUserRecordResponse, error)
	mustEmbedUnimplementedStudyServiceServer()
}

// UnimplementedStudyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudyServiceServer struct {
}

func (UnimplementedStudyServiceServer) GetUserRecord(context.Context, *GetUserRecordRequest) (*GetUserRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecord not implemented")
}
func (UnimplementedStudyServiceServer) CreateUserRecord(context.Context, *CreateUserRecordRequest) (*CreateUserRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserRecord not implemented")
}
func (UnimplementedStudyServiceServer) mustEmbedUnimplementedStudyServiceServer() {}

// UnsafeStudyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudyServiceServer will
// result in compilation errors.
type UnsafeStudyServiceServer interface {
	mustEmbedUnimplementedStudyServiceServer()
}

func RegisterStudyServiceServer(s grpc.ServiceRegistrar, srv StudyServiceServer) {
	s.RegisterService(&StudyService_ServiceDesc, srv)
}

func _StudyService_GetUserRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceServer).GetUserRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StudyService/GetUserRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceServer).GetUserRecord(ctx, req.(*GetUserRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyService_CreateUserRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceServer).CreateUserRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StudyService/CreateUserRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceServer).CreateUserRecord(ctx, req.(*CreateUserRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudyService_ServiceDesc is the grpc.ServiceDesc for StudyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StudyService",
	HandlerType: (*StudyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserRecord",
			Handler:    _StudyService_GetUserRecord_Handler,
		},
		{
			MethodName: "CreateUserRecord",
			Handler:    _StudyService_CreateUserRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "StudyService.proto",
}
