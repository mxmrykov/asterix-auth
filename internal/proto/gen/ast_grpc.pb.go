// protoc --go_out=./internal/proto/gen --go-grpc_out=./internal/proto/gen internal/proto/ast.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: internal/proto/ast.proto

package __

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
	Ast_GetIAID_FullMethodName = "/Ast/GetIAID"
)

// AstClient is the client API for Ast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AstClient interface {
	GetIAID(ctx context.Context, in *GetIAIDRequest, opts ...grpc.CallOption) (*GetIAIDResponse, error)
}

type astClient struct {
	cc grpc.ClientConnInterface
}

func NewAstClient(cc grpc.ClientConnInterface) AstClient {
	return &astClient{cc}
}

func (c *astClient) GetIAID(ctx context.Context, in *GetIAIDRequest, opts ...grpc.CallOption) (*GetIAIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIAIDResponse)
	err := c.cc.Invoke(ctx, Ast_GetIAID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AstServer is the server API for Ast service.
// All implementations must embed UnimplementedAstServer
// for forward compatibility.
type AstServer interface {
	GetIAID(context.Context, *GetIAIDRequest) (*GetIAIDResponse, error)
	mustEmbedUnimplementedAstServer()
}

// UnimplementedAstServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAstServer struct{}

func (UnimplementedAstServer) GetIAID(context.Context, *GetIAIDRequest) (*GetIAIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIAID not implemented")
}
func (UnimplementedAstServer) mustEmbedUnimplementedAstServer() {}
func (UnimplementedAstServer) testEmbeddedByValue()             {}

// UnsafeAstServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AstServer will
// result in compilation errors.
type UnsafeAstServer interface {
	mustEmbedUnimplementedAstServer()
}

func RegisterAstServer(s grpc.ServiceRegistrar, srv AstServer) {
	// If the following call pancis, it indicates UnimplementedAstServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ast_ServiceDesc, srv)
}

func _Ast_GetIAID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIAIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AstServer).GetIAID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ast_GetIAID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AstServer).GetIAID(ctx, req.(*GetIAIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ast_ServiceDesc is the grpc.ServiceDesc for Ast service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ast_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ast",
	HandlerType: (*AstServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIAID",
			Handler:    _Ast_GetIAID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/ast.proto",
}