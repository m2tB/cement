// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.20.0
// source: staff_client/v1/staff.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationStaffCaptcha = "/api.staff_client.v1.Staff/Captcha"
const OperationStaffRegister = "/api.staff_client.v1.Staff/Register"
const OperationStaffSignIn = "/api.staff_client.v1.Staff/SignIn"
const OperationStaffSignOut = "/api.staff_client.v1.Staff/SignOut"

type StaffHTTPServer interface {
	Captcha(context.Context, *CaptchaRequest) (*CaptchaReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	SignOut(context.Context, *emptypb.Empty) (*SignOutReply, error)
}

func RegisterStaffHTTPServer(s *http.Server, srv StaffHTTPServer) {
	r := s.Route("/")
	r.GET("/api/staff/captcha", _Staff_Captcha0_HTTP_Handler(srv))
	r.POST("/api/staff/register", _Staff_Register0_HTTP_Handler(srv))
	r.POST("/api/staff/signIn", _Staff_SignIn0_HTTP_Handler(srv))
	r.POST("/api/staff/signOut", _Staff_SignOut0_HTTP_Handler(srv))
}

func _Staff_Captcha0_HTTP_Handler(srv StaffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffCaptcha)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Captcha(ctx, req.(*CaptchaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CaptchaReply)
		return ctx.Result(200, reply)
	}
}

func _Staff_Register0_HTTP_Handler(srv StaffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Staff_SignIn0_HTTP_Handler(srv StaffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffSignIn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignIn(ctx, req.(*SignInRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignInReply)
		return ctx.Result(200, reply)
	}
}

func _Staff_SignOut0_HTTP_Handler(srv StaffHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffSignOut)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignOut(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignOutReply)
		return ctx.Result(200, reply)
	}
}

type StaffHTTPClient interface {
	Captcha(ctx context.Context, req *CaptchaRequest, opts ...http.CallOption) (rsp *CaptchaReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	SignIn(ctx context.Context, req *SignInRequest, opts ...http.CallOption) (rsp *SignInReply, err error)
	SignOut(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SignOutReply, err error)
}

type StaffHTTPClientImpl struct {
	cc *http.Client
}

func NewStaffHTTPClient(client *http.Client) StaffHTTPClient {
	return &StaffHTTPClientImpl{client}
}

func (c *StaffHTTPClientImpl) Captcha(ctx context.Context, in *CaptchaRequest, opts ...http.CallOption) (*CaptchaReply, error) {
	var out CaptchaReply
	pattern := "/api/staff/captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStaffCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/staff/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffHTTPClientImpl) SignIn(ctx context.Context, in *SignInRequest, opts ...http.CallOption) (*SignInReply, error) {
	var out SignInReply
	pattern := "/api/staff/signIn"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffSignIn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffHTTPClientImpl) SignOut(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SignOutReply, error) {
	var out SignOutReply
	pattern := "/api/staff/signOut"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffSignOut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
