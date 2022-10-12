// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.1
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

const OperationStaffClientCaptcha = "/api.staff_client.v1.StaffClient/Captcha"
const OperationStaffClientListStaffTeam = "/api.staff_client.v1.StaffClient/ListStaffTeam"
const OperationStaffClientReadStaff = "/api.staff_client.v1.StaffClient/ReadStaff"
const OperationStaffClientRegister = "/api.staff_client.v1.StaffClient/Register"
const OperationStaffClientSignIn = "/api.staff_client.v1.StaffClient/SignIn"
const OperationStaffClientSignOut = "/api.staff_client.v1.StaffClient/SignOut"
const OperationStaffClientUpdateStaff = "/api.staff_client.v1.StaffClient/UpdateStaff"

type StaffClientHTTPServer interface {
	Captcha(context.Context, *CaptchaRequest) (*CaptchaReply, error)
	ListStaffTeam(context.Context, *ListStaffTeamRequest) (*ListStaffTeamReply, error)
	ReadStaff(context.Context, *emptypb.Empty) (*ReadStaffReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	SignOut(context.Context, *emptypb.Empty) (*SignOutReply, error)
	UpdateStaff(context.Context, *UpdateStaffRequest) (*UpdateStaffReply, error)
}

func RegisterStaffClientHTTPServer(s *http.Server, srv StaffClientHTTPServer) {
	r := s.Route("/")
	r.POST("/api/staff_client/captcha", _StaffClient_Captcha0_HTTP_Handler(srv))
	r.POST("/api/staff_client/register", _StaffClient_Register0_HTTP_Handler(srv))
	r.POST("/api/staff_client/signIn", _StaffClient_SignIn0_HTTP_Handler(srv))
	r.POST("/api/staff_client/signOut", _StaffClient_SignOut0_HTTP_Handler(srv))
	r.POST("/api/staff_client/readStaff", _StaffClient_ReadStaff0_HTTP_Handler(srv))
	r.POST("/api/staff_client/updateStaff", _StaffClient_UpdateStaff0_HTTP_Handler(srv))
	r.POST("/api/staff_client/listStaffTeam", _StaffClient_ListStaffTeam0_HTTP_Handler(srv))
}

func _StaffClient_Captcha0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CaptchaRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientCaptcha)
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

func _StaffClient_Register0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientRegister)
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

func _StaffClient_SignIn0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientSignIn)
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

func _StaffClient_SignOut0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientSignOut)
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

func _StaffClient_ReadStaff0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientReadStaff)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ReadStaff(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReadStaffReply)
		return ctx.Result(200, reply)
	}
}

func _StaffClient_UpdateStaff0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateStaffRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientUpdateStaff)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateStaff(ctx, req.(*UpdateStaffRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateStaffReply)
		return ctx.Result(200, reply)
	}
}

func _StaffClient_ListStaffTeam0_HTTP_Handler(srv StaffClientHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListStaffTeamRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStaffClientListStaffTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListStaffTeam(ctx, req.(*ListStaffTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListStaffTeamReply)
		return ctx.Result(200, reply)
	}
}

type StaffClientHTTPClient interface {
	Captcha(ctx context.Context, req *CaptchaRequest, opts ...http.CallOption) (rsp *CaptchaReply, err error)
	ListStaffTeam(ctx context.Context, req *ListStaffTeamRequest, opts ...http.CallOption) (rsp *ListStaffTeamReply, err error)
	ReadStaff(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *ReadStaffReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	SignIn(ctx context.Context, req *SignInRequest, opts ...http.CallOption) (rsp *SignInReply, err error)
	SignOut(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *SignOutReply, err error)
	UpdateStaff(ctx context.Context, req *UpdateStaffRequest, opts ...http.CallOption) (rsp *UpdateStaffReply, err error)
}

type StaffClientHTTPClientImpl struct {
	cc *http.Client
}

func NewStaffClientHTTPClient(client *http.Client) StaffClientHTTPClient {
	return &StaffClientHTTPClientImpl{client}
}

func (c *StaffClientHTTPClientImpl) Captcha(ctx context.Context, in *CaptchaRequest, opts ...http.CallOption) (*CaptchaReply, error) {
	var out CaptchaReply
	pattern := "/api/staff_client/captcha"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientCaptcha))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) ListStaffTeam(ctx context.Context, in *ListStaffTeamRequest, opts ...http.CallOption) (*ListStaffTeamReply, error) {
	var out ListStaffTeamReply
	pattern := "/api/staff_client/listStaffTeam"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientListStaffTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) ReadStaff(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*ReadStaffReply, error) {
	var out ReadStaffReply
	pattern := "/api/staff_client/readStaff"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientReadStaff))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/staff_client/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) SignIn(ctx context.Context, in *SignInRequest, opts ...http.CallOption) (*SignInReply, error) {
	var out SignInReply
	pattern := "/api/staff_client/signIn"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientSignIn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) SignOut(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*SignOutReply, error) {
	var out SignOutReply
	pattern := "/api/staff_client/signOut"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientSignOut))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StaffClientHTTPClientImpl) UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...http.CallOption) (*UpdateStaffReply, error) {
	var out UpdateStaffReply
	pattern := "/api/staff_client/updateStaff"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStaffClientUpdateStaff))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
