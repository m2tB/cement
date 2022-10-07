package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "staff_client/api/staff_client/v1"
)

func (s *StaffService) Captcha(ctx context.Context, req *v1.CaptchaRequest) (*v1.CaptchaReply, error) {
	return &v1.CaptchaReply{}, nil
}
func (s *StaffService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return &v1.RegisterReply{}, nil
}
func (s *StaffService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	return &v1.SignInReply{}, nil
}
func (s *StaffService) SignOut(ctx context.Context, r *emptypb.Empty) (*v1.SignOutReply, error) {
	return &v1.SignOutReply{}, nil
}
