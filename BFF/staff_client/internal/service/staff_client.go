package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "staff_client/api/staff_client/v1"
)

func (s *StaffClientService) Captcha(ctx context.Context, req *v1.CaptchaRequest) (*v1.CaptchaReply, error) {
	return s.uc.Captcha(ctx, req)
}
func (s *StaffClientService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	return s.uc.Register(ctx, req)
}
func (s *StaffClientService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	return s.uc.SignIn(ctx, req)
}
func (s *StaffClientService) SignOut(ctx context.Context, _ *emptypb.Empty) (*v1.SignOutReply, error) {
	return s.uc.SignOut(ctx)
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (s *StaffClientService) ReadStaff(ctx context.Context, _ *emptypb.Empty) (*v1.ReadStaffReply, error) {
	return s.uc.ReadStaff(ctx)
}

func (s *StaffClientService) UpdateStaff(ctx context.Context, req *v1.UpdateStaffRequest) (*v1.UpdateStaffReply, error) {
	return s.uc.UpdateStaff(ctx, req)
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (s *StaffClientService) ListStaffTeam(ctx context.Context, req *v1.ListStaffTeamRequest) (*v1.ListStaffTeamReply, error) {
	return s.uc.ListStaffTeam(ctx, req)
}
