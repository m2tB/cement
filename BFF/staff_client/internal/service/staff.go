package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "staff_client/api/staff_client/v1"
)

func (s *StaffService) Captcha(ctx context.Context, req *v1.CaptchaRequest) (*v1.CaptchaReply, error) {
	return s.uc.Captcha(ctx, req)
}
func (s *StaffService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return s.uc.Register(ctx, req)
}
func (s *StaffService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	return s.uc.SignIn(ctx, req)
}
func (s *StaffService) SignOut(ctx context.Context, _ *emptypb.Empty) (*v1.SignOutReply, error) {
	return s.uc.SignOut(ctx)
}
