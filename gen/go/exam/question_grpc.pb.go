// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: exam/question.proto

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
	QuestionService_CreateQuestion_FullMethodName = "/exam.QuestionService/CreateQuestion"
	QuestionService_GetQuestion_FullMethodName    = "/exam.QuestionService/GetQuestion"
	QuestionService_UpdateQuestion_FullMethodName = "/exam.QuestionService/UpdateQuestion"
	QuestionService_DeleteQuestion_FullMethodName = "/exam.QuestionService/DeleteQuestion"
)

// QuestionServiceClient is the client API for QuestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuestionServiceClient interface {
	CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*Question, error)
	GetQuestion(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Question, error)
	UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*Question, error)
	DeleteQuestion(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Question, error)
}

type questionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuestionServiceClient(cc grpc.ClientConnInterface) QuestionServiceClient {
	return &questionServiceClient{cc}
}

func (c *questionServiceClient) CreateQuestion(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*Question, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Question)
	err := c.cc.Invoke(ctx, QuestionService_CreateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) GetQuestion(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Question, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Question)
	err := c.cc.Invoke(ctx, QuestionService_GetQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) UpdateQuestion(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*Question, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Question)
	err := c.cc.Invoke(ctx, QuestionService_UpdateQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionServiceClient) DeleteQuestion(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*Question, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Question)
	err := c.cc.Invoke(ctx, QuestionService_DeleteQuestion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionServiceServer is the server API for QuestionService service.
// All implementations must embed UnimplementedQuestionServiceServer
// for forward compatibility
type QuestionServiceServer interface {
	CreateQuestion(context.Context, *CreateQuestionRequest) (*Question, error)
	GetQuestion(context.Context, *GetByIdRequest) (*Question, error)
	UpdateQuestion(context.Context, *UpdateQuestionRequest) (*Question, error)
	DeleteQuestion(context.Context, *GetByIdRequest) (*Question, error)
	mustEmbedUnimplementedQuestionServiceServer()
}

// UnimplementedQuestionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuestionServiceServer struct {
}

func (UnimplementedQuestionServiceServer) CreateQuestion(context.Context, *CreateQuestionRequest) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) GetQuestion(context.Context, *GetByIdRequest) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) UpdateQuestion(context.Context, *UpdateQuestionRequest) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) DeleteQuestion(context.Context, *GetByIdRequest) (*Question, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestion not implemented")
}
func (UnimplementedQuestionServiceServer) mustEmbedUnimplementedQuestionServiceServer() {}

// UnsafeQuestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuestionServiceServer will
// result in compilation errors.
type UnsafeQuestionServiceServer interface {
	mustEmbedUnimplementedQuestionServiceServer()
}

func RegisterQuestionServiceServer(s grpc.ServiceRegistrar, srv QuestionServiceServer) {
	s.RegisterService(&QuestionService_ServiceDesc, srv)
}

func _QuestionService_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_CreateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).CreateQuestion(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_GetQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).GetQuestion(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_UpdateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_UpdateQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).UpdateQuestion(ctx, req.(*UpdateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionService_DeleteQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionServiceServer).DeleteQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuestionService_DeleteQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionServiceServer).DeleteQuestion(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuestionService_ServiceDesc is the grpc.ServiceDesc for QuestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exam.QuestionService",
	HandlerType: (*QuestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _QuestionService_CreateQuestion_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _QuestionService_GetQuestion_Handler,
		},
		{
			MethodName: "UpdateQuestion",
			Handler:    _QuestionService_UpdateQuestion_Handler,
		},
		{
			MethodName: "DeleteQuestion",
			Handler:    _QuestionService_DeleteQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exam/question.proto",
}
