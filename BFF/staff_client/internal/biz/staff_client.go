package biz

import (
	"context"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"regexp"
	staff_v1 "staff_client/api/staff/v1"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/conf"
	"staff_client/internal/pkg/middleware/auth"
	"time"

	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义错误信息
var (
	ErrCaptchaInvalid      = errors.New("verification code error")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")

	ErrNameInvalid = errors.New("name invalid")

	ErrRegisterInvalid = errors.New("register invalid")
	ErrSignInFailed    = errors.New("sign in failed")

	ErrStaffNotFound     = errors.New("staff not found")
	ErrStaffUpdateFailed = errors.New("staff update failed")
	ErrProgramExec       = errors.New("program exec error")
)

// StaffClientRepo is a StaffClient repo.
type StaffClientRepo interface {
	Captcha(context.Context, *v1.CaptchaRequest) bool
	Register(context.Context, *v1.RegisterRequest) bool
	SignIn(context.Context, *v1.SignInRequest) *staff_v1.StaffReply
	SignOut(context.Context) bool

	ReadStaff(context.Context) *staff_v1.ReadStaffReply
	UpdateStaff(context.Context, *v1.UpdateStaffRequest) bool

	ListStaffTeam(context.Context, *v1.ListStaffTeamRequest) *staff_v1.ListStaffTeamReply
}

// StaffClientUsecase is a StaffClient usecase.
type StaffClientUsecase struct {
	repo    StaffClientRepo
	log     *log.Helper
	signKey string // 这里是为了生存 token 的时候可以直接取配置文件里面的配置
}

// NewStaffClientUsecase new a StaffClient usecase.
func NewStaffClientUsecase(repo StaffClientRepo, logger log.Logger, conf *conf.Auth) *StaffClientUsecase {
	return &StaffClientUsecase{repo: repo, log: log.NewHelper(logger), signKey: conf.JwtKey}
}

/*--------------------------------------------------------------------------------------------------------------------*/

// Captcha send a Captcha code used by sms
func (uc *StaffClientUsecase) Captcha(ctx context.Context, req *v1.CaptchaRequest) (*v1.CaptchaReply, error) {
	exec := uc.repo.Captcha(ctx, req)
	if !exec {
		return nil, ErrCaptchaInvalid
	}
	return &v1.CaptchaReply{
		Exec: true,
	}, nil
}

func (uc *StaffClientUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	// 验证入参名称校验判断
	reg := regexp.MustCompile("[\u4e00-\u9fa5a-zA-Z\\d]")
	if !reg.MatchString(req.Name) {
		return nil, ErrNameInvalid
	}
	exec := uc.repo.Register(ctx, req)
	if !exec {
		return nil, ErrRegisterInvalid
	}
	return &v1.RegisterReply{
		Exec: true,
	}, nil
}

func (uc *StaffClientUsecase) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	staff := uc.repo.SignIn(ctx, req)
	if staff == nil {
		return nil, ErrSignInFailed
	}
	// 生成token令牌
	claims := auth.CusClaims{
		ID:     staff.Id,
		Mobile: staff.Mobile,
		Name:   staff.Name,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			Issuer:    "Staff_Client",
		},
	}
	token, err := auth.GenerateToken(claims, uc.signKey)
	if err != nil {
		return nil, ErrGenerateTokenFailed
	}
	return &v1.SignInReply{
		Id:        staff.Id,
		Name:      staff.Name,
		Mobile:    staff.Mobile,
		CreatedAt: staff.CreatedAt,
		Token:     token,
	}, nil
}

func (uc *StaffClientUsecase) SignOut(ctx context.Context) (*v1.SignOutReply, error) {
	exec := uc.repo.SignOut(ctx)
	if !exec {
		return nil, ErrAuthFailed
	}
	return &v1.SignOutReply{
		Exec: true,
	}, nil
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (uc *StaffClientUsecase) ReadStaff(ctx context.Context) (*v1.ReadStaffReply, error) {
	staff := uc.repo.ReadStaff(ctx)
	if staff == nil {
		return nil, ErrStaffNotFound
	}
	return &v1.ReadStaffReply{
		Id:        staff.Staff.Id,
		Name:      staff.Staff.Name,
		Mobile:    staff.Staff.Mobile,
		CreatedAt: staff.Staff.CreatedAt,
		UpdatedAt: staff.Staff.UpdatedAt,
	}, nil
}

func (uc *StaffClientUsecase) UpdateStaff(ctx context.Context, req *v1.UpdateStaffRequest) (*v1.UpdateStaffReply, error) {
	exec := uc.repo.UpdateStaff(ctx, req)
	if !exec {
		return nil, ErrStaffUpdateFailed
	}
	return &v1.UpdateStaffReply{
		Exec: true,
	}, nil
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (uc *StaffClientUsecase) ListStaffTeam(ctx context.Context, req *v1.ListStaffTeamRequest) (*v1.ListStaffTeamReply, error) {
	exec := uc.repo.ListStaffTeam(ctx, req)
	if exec == nil {
		return nil, ErrProgramExec
	}
	var data []*v1.StaffTeamReply
	if exec.Total > 0 {
		for _, v := range exec.Data {
			data = append(data, &v1.StaffTeamReply{
				Tid:       v.Tid,
				TName:     v.TName,
				CreatedAt: v.CreatedAt,
			})
		}
	}
	return &v1.ListStaffTeamReply{
		Total: exec.Total,
		Data:  data,
	}, nil
}
