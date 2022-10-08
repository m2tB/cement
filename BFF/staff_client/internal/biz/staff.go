package biz

import (
	"context"
	"fmt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"regexp"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/conf"
	"staff_client/internal/pkg/captcha"
	"staff_client/internal/pkg/middleware/auth"
	"time"

	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

// 定义错误信息
var (
	ErrCaptchaInvalid      = errors.New("verification code error")
	ErrMobileInvalid       = errors.New("mobile invalid")
	ErrNameInvalid         = errors.New("name invalid")
	ErrStaffNotFound       = errors.New("staff not found")
	ErrSignInFailed        = errors.New("sign in failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")
)

// Staff is a Staff model.
type Staff struct {
	ID        int64
	Mobile    string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// StaffRepo is a Staff repo.
type StaffRepo interface {
	Save(context.Context, *Staff) (bool, error)
	Read(context.Context, string) (*Staff, error)
}

// StaffUsecase is a Staff usecase.
type StaffUsecase struct {
	repo    StaffRepo
	log     *log.Helper
	signKey string // 这里是为了生存 token 的时候可以直接取配置文件里面的配置
}

// NewStaffUsecase new a Staff usecase.
func NewStaffUsecase(repo StaffRepo, logger log.Logger, conf *conf.Auth) *StaffUsecase {
	return &StaffUsecase{repo: repo, log: log.NewHelper(logger), signKey: conf.JwtKey}
}

// Captcha send a Captcha code used by sms
func (uc *StaffUsecase) Captcha(ctx context.Context, req *v1.CaptchaRequest) (*v1.CaptchaReply, error) {
	// TODO - 判断redis是否对该号码做限制
	code, err := captcha.SendCaptcha(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	// TODO - 存入redis做数据绑定
	fmt.Printf("code - %v", code)
	return &v1.CaptchaReply{
		Exec: true,
	}, nil
}

func (uc *StaffUsecase) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.RegisterReply, error) {
	// 验证入参名称校验判断
	reg := regexp.MustCompile("[\u4e00-\u9fa5a-zA-Z\\d]")
	if !reg.MatchString(req.Name) {
		return nil, ErrNameInvalid
	}
	exec, err := uc.repo.Save(ctx, &Staff{
		Mobile: req.Mobile,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterReply{
		Exec: exec,
	}, nil
}

func (uc *StaffUsecase) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	// TODO - 提取redis数据进行校验
	fmt.Printf("captcha code - %v", req.Captcha)
	staff, err := uc.repo.Read(ctx, req.Mobile)
	if err != nil {
		return nil, ErrStaffNotFound
	}
	claims := auth.CusClaims{
		Mobile: staff.Mobile,
		Name:   staff.Name,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),
			ExpiresAt: jwt2.NewNumericDate(time.Now().AddDate(0, 1, 0)),
			Issuer:    "Staff_Client",
		},
	}
	token, _ := auth.CreateToken(claims, uc.signKey)
	if err != nil {
		return nil, ErrGenerateTokenFailed
	}
	return &v1.SignInReply{
		Id:        staff.ID,
		Name:      staff.Name,
		Mobile:    staff.Mobile,
		CreatedAt: staff.CreatedAt,
		Token:     token,
	}, nil
}
