// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: staff_client/v1/staff.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StaffClient is the client API for Staff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffClient interface {
	Captcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error)
	SignOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SignOutReply, error)
}

type staffClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffClient(cc grpc.ClientConnInterface) StaffClient {
	return &staffClient{cc}
}

func (c *staffClient) Captcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaReply, error) {
	out := new(CaptchaReply)
	err := c.cc.Invoke(ctx, "/api.staff_client.v1.Staff/Captcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/api.staff_client.v1.Staff/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInReply, error) {
	out := new(SignInReply)
	err := c.cc.Invoke(ctx, "/api.staff_client.v1.Staff/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) SignOut(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SignOutReply, error) {
	out := new(SignOutReply)
	err := c.cc.Invoke(ctx, "/api.staff_client.v1.Staff/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServer is the server API for Staff service.
// All implementations must embed UnimplementedStaffServer
// for forward compatibility
type StaffServer interface {
	Captcha(context.Context, *CaptchaRequest) (*CaptchaReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	SignOut(context.Context, *emptypb.Empty) (*SignOutReply, error)
	mustEmbedUnimplementedStaffServer()
}

// UnimplementedStaffServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServer struct {
}

func (UnimplementedStaffServer) Captcha(context.Context, *CaptchaRequest) (*CaptchaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Captcha not implemented")
}
func (UnimplementedStaffServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedStaffServer) SignIn(context.Context, *SignInRequest) (*SignInReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedStaffServer) SignOut(context.Context, *emptypb.Empty) (*SignOutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedStaffServer) mustEmbedUnimplementedStaffServer() {}

// UnsafeStaffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServer will
// result in compilation errors.
type UnsafeStaffServer interface {
	mustEmbedUnimplementedStaffServer()
}

func RegisterStaffServer(s grpc.ServiceRegistrar, srv StaffServer) {
	s.RegisterService(&Staff_ServiceDesc, srv)
}

func _Staff_Captcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).Captcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff_client.v1.Staff/Captcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).Captcha(ctx, req.(*CaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff_client.v1.Staff/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff_client.v1.Staff/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff_client.v1.Staff/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).SignOut(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Staff_ServiceDesc is the grpc.ServiceDesc for Staff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Staff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.staff_client.v1.Staff",
	HandlerType: (*StaffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Captcha",
			Handler:    _Staff_Captcha_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Staff_Register_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Staff_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Staff_SignOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff_client/v1/staff.proto",
}
