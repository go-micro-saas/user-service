// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/user-service/v1/services/user_auth.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/user-service/api/user-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvUserAuthV1_Ping_FullMethodName         = "/saas.api.user.servicev1.SrvUserAuthV1/Ping"
	SrvUserAuthV1_RefreshToken_FullMethodName = "/saas.api.user.servicev1.SrvUserAuthV1/RefreshToken"
	SrvUserAuthV1_LoginByEmail_FullMethodName = "/saas.api.user.servicev1.SrvUserAuthV1/LoginByEmail"
	SrvUserAuthV1_OpenApiLogin_FullMethodName = "/saas.api.user.servicev1.SrvUserAuthV1/OpenApiLogin"
)

// SrvUserAuthV1Client is the client API for SrvUserAuthV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvUserAuthV1Client interface {
	// 身份验证-Ping测试
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
	// 身份验证-刷新Token
	RefreshToken(ctx context.Context, in *resources.RefreshTokenReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-Email登录
	LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
	// 身份验证-开发平台登录
	OpenApiLogin(ctx context.Context, in *resources.OpenApiLoginReq, opts ...grpc.CallOption) (*resources.LoginResp, error)
}

type srvUserAuthV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSrvUserAuthV1Client(cc grpc.ClientConnInterface) SrvUserAuthV1Client {
	return &srvUserAuthV1Client{cc}
}

func (c *srvUserAuthV1Client) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) RefreshToken(ctx context.Context, in *resources.RefreshTokenReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) LoginByEmail(ctx context.Context, in *resources.LoginByEmailReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_LoginByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUserAuthV1Client) OpenApiLogin(ctx context.Context, in *resources.OpenApiLoginReq, opts ...grpc.CallOption) (*resources.LoginResp, error) {
	out := new(resources.LoginResp)
	err := c.cc.Invoke(ctx, SrvUserAuthV1_OpenApiLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvUserAuthV1Server is the server API for SrvUserAuthV1 service.
// All implementations must embed UnimplementedSrvUserAuthV1Server
// for forward compatibility
type SrvUserAuthV1Server interface {
	// 身份验证-Ping测试
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	// 身份验证-刷新Token
	RefreshToken(context.Context, *resources.RefreshTokenReq) (*resources.LoginResp, error)
	// 身份验证-Email登录
	LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error)
	// 身份验证-开发平台登录
	OpenApiLogin(context.Context, *resources.OpenApiLoginReq) (*resources.LoginResp, error)
	mustEmbedUnimplementedSrvUserAuthV1Server()
}

// UnimplementedSrvUserAuthV1Server must be embedded to have forward compatible implementations.
type UnimplementedSrvUserAuthV1Server struct {
}

func (UnimplementedSrvUserAuthV1Server) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvUserAuthV1Server) RefreshToken(context.Context, *resources.RefreshTokenReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedSrvUserAuthV1Server) LoginByEmail(context.Context, *resources.LoginByEmailReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByEmail not implemented")
}
func (UnimplementedSrvUserAuthV1Server) OpenApiLogin(context.Context, *resources.OpenApiLoginReq) (*resources.LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenApiLogin not implemented")
}
func (UnimplementedSrvUserAuthV1Server) mustEmbedUnimplementedSrvUserAuthV1Server() {}

// UnsafeSrvUserAuthV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvUserAuthV1Server will
// result in compilation errors.
type UnsafeSrvUserAuthV1Server interface {
	mustEmbedUnimplementedSrvUserAuthV1Server()
}

func RegisterSrvUserAuthV1Server(s grpc.ServiceRegistrar, srv SrvUserAuthV1Server) {
	s.RegisterService(&SrvUserAuthV1_ServiceDesc, srv)
}

func _SrvUserAuthV1_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).RefreshToken(ctx, req.(*resources.RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_LoginByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.LoginByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).LoginByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_LoginByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).LoginByEmail(ctx, req.(*resources.LoginByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SrvUserAuthV1_OpenApiLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.OpenApiLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvUserAuthV1Server).OpenApiLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvUserAuthV1_OpenApiLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvUserAuthV1Server).OpenApiLogin(ctx, req.(*resources.OpenApiLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvUserAuthV1_ServiceDesc is the grpc.ServiceDesc for SrvUserAuthV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvUserAuthV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.user.servicev1.SrvUserAuthV1",
	HandlerType: (*SrvUserAuthV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvUserAuthV1_Ping_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SrvUserAuthV1_RefreshToken_Handler,
		},
		{
			MethodName: "LoginByEmail",
			Handler:    _SrvUserAuthV1_LoginByEmail_Handler,
		},
		{
			MethodName: "OpenApiLogin",
			Handler:    _SrvUserAuthV1_OpenApiLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user-service/v1/services/user_auth.service.v1.proto",
}