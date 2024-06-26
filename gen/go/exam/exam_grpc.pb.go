// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: exam/exam.proto

package exam

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExamService_CreateExam_FullMethodName = "/exam.ExamService/CreateExam"
	ExamService_GetExam_FullMethodName    = "/exam.ExamService/GetExam"
	ExamService_UpdateExam_FullMethodName = "/exam.ExamService/UpdateExam"
	ExamService_DeleteExam_FullMethodName = "/exam.ExamService/DeleteExam"
)

// ExamServiceClient is the client API for ExamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExamServiceClient interface {
	CreateExam(ctx context.Context, in *CreateExamRequest, opts ...grpc.CallOption) (*Exam, error)
	GetExam(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Exam, error)
	UpdateExam(ctx context.Context, in *UpdateExamRequest, opts ...grpc.CallOption) (*Exam, error)
	DeleteExam(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Exam, error)
}

type examServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExamServiceClient(cc grpc.ClientConnInterface) ExamServiceClient {
	return &examServiceClient{cc}
}

func (c *examServiceClient) CreateExam(ctx context.Context, in *CreateExamRequest, opts ...grpc.CallOption) (*Exam, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Exam)
	err := c.cc.Invoke(ctx, ExamService_CreateExam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examServiceClient) GetExam(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Exam, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Exam)
	err := c.cc.Invoke(ctx, ExamService_GetExam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examServiceClient) UpdateExam(ctx context.Context, in *UpdateExamRequest, opts ...grpc.CallOption) (*Exam, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Exam)
	err := c.cc.Invoke(ctx, ExamService_UpdateExam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *examServiceClient) DeleteExam(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Exam, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Exam)
	err := c.cc.Invoke(ctx, ExamService_DeleteExam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExamServiceServer is the server API for ExamService service.
// All implementations must embed UnimplementedExamServiceServer
// for forward compatibility
type ExamServiceServer interface {
	CreateExam(context.Context, *CreateExamRequest) (*Exam, error)
	GetExam(context.Context, *GetByIdRequest) (*Exam, error)
	UpdateExam(context.Context, *UpdateExamRequest) (*Exam, error)
	DeleteExam(context.Context, *GetByIdRequest) (*Exam, error)
	mustEmbedUnimplementedExamServiceServer()
}

// UnimplementedExamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExamServiceServer struct {
}

func (UnimplementedExamServiceServer) CreateExam(context.Context, *CreateExamRequest) (*Exam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExam not implemented")
}
func (UnimplementedExamServiceServer) GetExam(context.Context, *GetByIdRequest) (*Exam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExam not implemented")
}
func (UnimplementedExamServiceServer) UpdateExam(context.Context, *UpdateExamRequest) (*Exam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExam not implemented")
}
func (UnimplementedExamServiceServer) DeleteExam(context.Context, *GetByIdRequest) (*Exam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExam not implemented")
}
func (UnimplementedExamServiceServer) mustEmbedUnimplementedExamServiceServer() {}

// UnsafeExamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExamServiceServer will
// result in compilation errors.
type UnsafeExamServiceServer interface {
	mustEmbedUnimplementedExamServiceServer()
}

func RegisterExamServiceServer(s grpc.ServiceRegistrar, srv ExamServiceServer) {
	s.RegisterService(&ExamService_ServiceDesc, srv)
}

func _ExamService_CreateExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).CreateExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExamService_CreateExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).CreateExam(ctx, req.(*CreateExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamService_GetExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).GetExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExamService_GetExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).GetExam(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamService_UpdateExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).UpdateExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExamService_UpdateExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).UpdateExam(ctx, req.(*UpdateExamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExamService_DeleteExam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExamServiceServer).DeleteExam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExamService_DeleteExam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExamServiceServer).DeleteExam(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExamService_ServiceDesc is the grpc.ServiceDesc for ExamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exam.ExamService",
	HandlerType: (*ExamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExam",
			Handler:    _ExamService_CreateExam_Handler,
		},
		{
			MethodName: "GetExam",
			Handler:    _ExamService_GetExam_Handler,
		},
		{
			MethodName: "UpdateExam",
			Handler:    _ExamService_UpdateExam_Handler,
		},
		{
			MethodName: "DeleteExam",
			Handler:    _ExamService_DeleteExam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exam/exam.proto",
}
