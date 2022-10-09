// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: staff/v1/staff.proto

package v1

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

// StaffClient is the client API for Staff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffClient interface {
	// staff相关接口
	CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffReply, error)
	UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffReply, error)
	DeleteStaff(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*DeleteStaffReply, error)
	ReadStaff(ctx context.Context, in *ReadStaffRequest, opts ...grpc.CallOption) (*ReadStaffReply, error)
	RecoveryStaff(ctx context.Context, in *RecoveryStaffRequest, opts ...grpc.CallOption) (*RecoveryStaffReply, error)
	ListStaff(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffReply, error)
	// team相关接口
	CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamReply, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamReply, error)
	DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamReply, error)
	ReadTeam(ctx context.Context, in *ReadTeamRequest, opts ...grpc.CallOption) (*ReadTeamReply, error)
	ListTeam(ctx context.Context, in *ListTeamRequest, opts ...grpc.CallOption) (*ListTeamReply, error)
	// team 绑定 staff 相关接口
	InviteStaff(ctx context.Context, in *InviteStaffRequest, opts ...grpc.CallOption) (*InviteStaffReply, error)
	ExpelStaff(ctx context.Context, in *ExpelStaffRequest, opts ...grpc.CallOption) (*ExpelStaffReply, error)
	ListTeamStaff(ctx context.Context, in *ListTeamStaffRequest, opts ...grpc.CallOption) (*ListTeamStaffReply, error)
	ListStaffTeam(ctx context.Context, in *ListStaffTeamRequest, opts ...grpc.CallOption) (*ListStaffTeamReply, error)
	// team 绑定 team 相关接口
	CreateSubTeam(ctx context.Context, in *CreateSubTeamRequest, opts ...grpc.CallOption) (*CreateSubTeamReply, error)
	UpdateSubTeam(ctx context.Context, in *UpdateSubTeamRequest, opts ...grpc.CallOption) (*UpdateSubTeamReply, error)
	DeleteSubTeam(ctx context.Context, in *DeleteSubTeamRequest, opts ...grpc.CallOption) (*DeleteSubTeamReply, error)
	ListTeamSubTeam(ctx context.Context, in *ListTeamSubTeamRequest, opts ...grpc.CallOption) (*ListTeamSubTeamReply, error)
}

type staffClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffClient(cc grpc.ClientConnInterface) StaffClient {
	return &staffClient{cc}
}

func (c *staffClient) CreateStaff(ctx context.Context, in *CreateStaffRequest, opts ...grpc.CallOption) (*CreateStaffReply, error) {
	out := new(CreateStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/CreateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) UpdateStaff(ctx context.Context, in *UpdateStaffRequest, opts ...grpc.CallOption) (*UpdateStaffReply, error) {
	out := new(UpdateStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/UpdateStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) DeleteStaff(ctx context.Context, in *DeleteStaffRequest, opts ...grpc.CallOption) (*DeleteStaffReply, error) {
	out := new(DeleteStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/DeleteStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ReadStaff(ctx context.Context, in *ReadStaffRequest, opts ...grpc.CallOption) (*ReadStaffReply, error) {
	out := new(ReadStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ReadStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) RecoveryStaff(ctx context.Context, in *RecoveryStaffRequest, opts ...grpc.CallOption) (*RecoveryStaffReply, error) {
	out := new(RecoveryStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/RecoveryStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ListStaff(ctx context.Context, in *ListStaffRequest, opts ...grpc.CallOption) (*ListStaffReply, error) {
	out := new(ListStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ListStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...grpc.CallOption) (*CreateTeamReply, error) {
	out := new(CreateTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*UpdateTeamReply, error) {
	out := new(UpdateTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) DeleteTeam(ctx context.Context, in *DeleteTeamRequest, opts ...grpc.CallOption) (*DeleteTeamReply, error) {
	out := new(DeleteTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ReadTeam(ctx context.Context, in *ReadTeamRequest, opts ...grpc.CallOption) (*ReadTeamReply, error) {
	out := new(ReadTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ReadTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ListTeam(ctx context.Context, in *ListTeamRequest, opts ...grpc.CallOption) (*ListTeamReply, error) {
	out := new(ListTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ListTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) InviteStaff(ctx context.Context, in *InviteStaffRequest, opts ...grpc.CallOption) (*InviteStaffReply, error) {
	out := new(InviteStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/InviteStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ExpelStaff(ctx context.Context, in *ExpelStaffRequest, opts ...grpc.CallOption) (*ExpelStaffReply, error) {
	out := new(ExpelStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ExpelStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ListTeamStaff(ctx context.Context, in *ListTeamStaffRequest, opts ...grpc.CallOption) (*ListTeamStaffReply, error) {
	out := new(ListTeamStaffReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ListTeamStaff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ListStaffTeam(ctx context.Context, in *ListStaffTeamRequest, opts ...grpc.CallOption) (*ListStaffTeamReply, error) {
	out := new(ListStaffTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ListStaffTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) CreateSubTeam(ctx context.Context, in *CreateSubTeamRequest, opts ...grpc.CallOption) (*CreateSubTeamReply, error) {
	out := new(CreateSubTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/CreateSubTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) UpdateSubTeam(ctx context.Context, in *UpdateSubTeamRequest, opts ...grpc.CallOption) (*UpdateSubTeamReply, error) {
	out := new(UpdateSubTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/UpdateSubTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) DeleteSubTeam(ctx context.Context, in *DeleteSubTeamRequest, opts ...grpc.CallOption) (*DeleteSubTeamReply, error) {
	out := new(DeleteSubTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/DeleteSubTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffClient) ListTeamSubTeam(ctx context.Context, in *ListTeamSubTeamRequest, opts ...grpc.CallOption) (*ListTeamSubTeamReply, error) {
	out := new(ListTeamSubTeamReply)
	err := c.cc.Invoke(ctx, "/api.staff.v1.Staff/ListTeamSubTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServer is the server API for Staff service.
// All implementations must embed UnimplementedStaffServer
// for forward compatibility
type StaffServer interface {
	// staff相关接口
	CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffReply, error)
	UpdateStaff(context.Context, *UpdateStaffRequest) (*UpdateStaffReply, error)
	DeleteStaff(context.Context, *DeleteStaffRequest) (*DeleteStaffReply, error)
	ReadStaff(context.Context, *ReadStaffRequest) (*ReadStaffReply, error)
	RecoveryStaff(context.Context, *RecoveryStaffRequest) (*RecoveryStaffReply, error)
	ListStaff(context.Context, *ListStaffRequest) (*ListStaffReply, error)
	// team相关接口
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamReply, error)
	DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamReply, error)
	ReadTeam(context.Context, *ReadTeamRequest) (*ReadTeamReply, error)
	ListTeam(context.Context, *ListTeamRequest) (*ListTeamReply, error)
	// team 绑定 staff 相关接口
	InviteStaff(context.Context, *InviteStaffRequest) (*InviteStaffReply, error)
	ExpelStaff(context.Context, *ExpelStaffRequest) (*ExpelStaffReply, error)
	ListTeamStaff(context.Context, *ListTeamStaffRequest) (*ListTeamStaffReply, error)
	ListStaffTeam(context.Context, *ListStaffTeamRequest) (*ListStaffTeamReply, error)
	// team 绑定 team 相关接口
	CreateSubTeam(context.Context, *CreateSubTeamRequest) (*CreateSubTeamReply, error)
	UpdateSubTeam(context.Context, *UpdateSubTeamRequest) (*UpdateSubTeamReply, error)
	DeleteSubTeam(context.Context, *DeleteSubTeamRequest) (*DeleteSubTeamReply, error)
	ListTeamSubTeam(context.Context, *ListTeamSubTeamRequest) (*ListTeamSubTeamReply, error)
	mustEmbedUnimplementedStaffServer()
}

// UnimplementedStaffServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServer struct {
}

func (UnimplementedStaffServer) CreateStaff(context.Context, *CreateStaffRequest) (*CreateStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStaff not implemented")
}
func (UnimplementedStaffServer) UpdateStaff(context.Context, *UpdateStaffRequest) (*UpdateStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStaff not implemented")
}
func (UnimplementedStaffServer) DeleteStaff(context.Context, *DeleteStaffRequest) (*DeleteStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStaff not implemented")
}
func (UnimplementedStaffServer) ReadStaff(context.Context, *ReadStaffRequest) (*ReadStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadStaff not implemented")
}
func (UnimplementedStaffServer) RecoveryStaff(context.Context, *RecoveryStaffRequest) (*RecoveryStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoveryStaff not implemented")
}
func (UnimplementedStaffServer) ListStaff(context.Context, *ListStaffRequest) (*ListStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaff not implemented")
}
func (UnimplementedStaffServer) CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (UnimplementedStaffServer) UpdateTeam(context.Context, *UpdateTeamRequest) (*UpdateTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (UnimplementedStaffServer) DeleteTeam(context.Context, *DeleteTeamRequest) (*DeleteTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (UnimplementedStaffServer) ReadTeam(context.Context, *ReadTeamRequest) (*ReadTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTeam not implemented")
}
func (UnimplementedStaffServer) ListTeam(context.Context, *ListTeamRequest) (*ListTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeam not implemented")
}
func (UnimplementedStaffServer) InviteStaff(context.Context, *InviteStaffRequest) (*InviteStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteStaff not implemented")
}
func (UnimplementedStaffServer) ExpelStaff(context.Context, *ExpelStaffRequest) (*ExpelStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpelStaff not implemented")
}
func (UnimplementedStaffServer) ListTeamStaff(context.Context, *ListTeamStaffRequest) (*ListTeamStaffReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeamStaff not implemented")
}
func (UnimplementedStaffServer) ListStaffTeam(context.Context, *ListStaffTeamRequest) (*ListStaffTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStaffTeam not implemented")
}
func (UnimplementedStaffServer) CreateSubTeam(context.Context, *CreateSubTeamRequest) (*CreateSubTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubTeam not implemented")
}
func (UnimplementedStaffServer) UpdateSubTeam(context.Context, *UpdateSubTeamRequest) (*UpdateSubTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubTeam not implemented")
}
func (UnimplementedStaffServer) DeleteSubTeam(context.Context, *DeleteSubTeamRequest) (*DeleteSubTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubTeam not implemented")
}
func (UnimplementedStaffServer) ListTeamSubTeam(context.Context, *ListTeamSubTeamRequest) (*ListTeamSubTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeamSubTeam not implemented")
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

func _Staff_CreateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).CreateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/CreateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).CreateStaff(ctx, req.(*CreateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_UpdateStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).UpdateStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/UpdateStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).UpdateStaff(ctx, req.(*UpdateStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_DeleteStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).DeleteStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/DeleteStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).DeleteStaff(ctx, req.(*DeleteStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ReadStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ReadStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ReadStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ReadStaff(ctx, req.(*ReadStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_RecoveryStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoveryStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).RecoveryStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/RecoveryStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).RecoveryStaff(ctx, req.(*RecoveryStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ListStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ListStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ListStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ListStaff(ctx, req.(*ListStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).CreateTeam(ctx, req.(*CreateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).DeleteTeam(ctx, req.(*DeleteTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ReadTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ReadTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ReadTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ReadTeam(ctx, req.(*ReadTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ListTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ListTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ListTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ListTeam(ctx, req.(*ListTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_InviteStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).InviteStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/InviteStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).InviteStaff(ctx, req.(*InviteStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ExpelStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpelStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ExpelStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ExpelStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ExpelStaff(ctx, req.(*ExpelStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ListTeamStaff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ListTeamStaff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ListTeamStaff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ListTeamStaff(ctx, req.(*ListTeamStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ListStaffTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStaffTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ListStaffTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ListStaffTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ListStaffTeam(ctx, req.(*ListStaffTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_CreateSubTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).CreateSubTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/CreateSubTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).CreateSubTeam(ctx, req.(*CreateSubTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_UpdateSubTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSubTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).UpdateSubTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/UpdateSubTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).UpdateSubTeam(ctx, req.(*UpdateSubTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_DeleteSubTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).DeleteSubTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/DeleteSubTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).DeleteSubTeam(ctx, req.(*DeleteSubTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Staff_ListTeamSubTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeamSubTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServer).ListTeamSubTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.staff.v1.Staff/ListTeamSubTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServer).ListTeamSubTeam(ctx, req.(*ListTeamSubTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Staff_ServiceDesc is the grpc.ServiceDesc for Staff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Staff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.staff.v1.Staff",
	HandlerType: (*StaffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStaff",
			Handler:    _Staff_CreateStaff_Handler,
		},
		{
			MethodName: "UpdateStaff",
			Handler:    _Staff_UpdateStaff_Handler,
		},
		{
			MethodName: "DeleteStaff",
			Handler:    _Staff_DeleteStaff_Handler,
		},
		{
			MethodName: "ReadStaff",
			Handler:    _Staff_ReadStaff_Handler,
		},
		{
			MethodName: "RecoveryStaff",
			Handler:    _Staff_RecoveryStaff_Handler,
		},
		{
			MethodName: "ListStaff",
			Handler:    _Staff_ListStaff_Handler,
		},
		{
			MethodName: "CreateTeam",
			Handler:    _Staff_CreateTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Staff_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Staff_DeleteTeam_Handler,
		},
		{
			MethodName: "ReadTeam",
			Handler:    _Staff_ReadTeam_Handler,
		},
		{
			MethodName: "ListTeam",
			Handler:    _Staff_ListTeam_Handler,
		},
		{
			MethodName: "InviteStaff",
			Handler:    _Staff_InviteStaff_Handler,
		},
		{
			MethodName: "ExpelStaff",
			Handler:    _Staff_ExpelStaff_Handler,
		},
		{
			MethodName: "ListTeamStaff",
			Handler:    _Staff_ListTeamStaff_Handler,
		},
		{
			MethodName: "ListStaffTeam",
			Handler:    _Staff_ListStaffTeam_Handler,
		},
		{
			MethodName: "CreateSubTeam",
			Handler:    _Staff_CreateSubTeam_Handler,
		},
		{
			MethodName: "UpdateSubTeam",
			Handler:    _Staff_UpdateSubTeam_Handler,
		},
		{
			MethodName: "DeleteSubTeam",
			Handler:    _Staff_DeleteSubTeam_Handler,
		},
		{
			MethodName: "ListTeamSubTeam",
			Handler:    _Staff_ListTeamSubTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff/v1/staff.proto",
}