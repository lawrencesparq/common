// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: api/upload/upload.proto

package upload

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UploadService_UploadFile_FullMethodName = "/api.UploadService/UploadFile"
	UploadService_DeleteFile_FullMethodName = "/api.UploadService/DeleteFile"
	UploadService_EditFile_FullMethodName   = "/api.UploadService/EditFile"
)

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	DeleteFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
	EditFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, UploadService_UploadFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadServiceClient) DeleteFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, UploadService_DeleteFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadServiceClient) EditFile(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, UploadService_EditFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility.
type UploadServiceServer interface {
	UploadFile(context.Context, *UploadRequest) (*UploadResponse, error)
	DeleteFile(context.Context, *UploadRequest) (*UploadResponse, error)
	EditFile(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUploadServiceServer struct{}

func (UnimplementedUploadServiceServer) UploadFile(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedUploadServiceServer) DeleteFile(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedUploadServiceServer) EditFile(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditFile not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}
func (UnimplementedUploadServiceServer) testEmbeddedByValue()                       {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	// If the following call pancis, it indicates UnimplementedUploadServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).UploadFile(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).DeleteFile(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadService_EditFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).EditFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_EditFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).EditFile(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFile",
			Handler:    _UploadService_UploadFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _UploadService_DeleteFile_Handler,
		},
		{
			MethodName: "EditFile",
			Handler:    _UploadService_EditFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/upload/upload.proto",
}