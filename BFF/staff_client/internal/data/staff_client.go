package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/biz"
	"staff_client/internal/pkg/captcha"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	staff_v1 "staff_client/api/staff/v1"
)

type staffClientRepo struct {
	data *Data
	log  *log.Helper
}

var (
	LimitCaptchaPre = "limit_captcha_"
	SignMarkPre     = "sign_mark_"
)

// NewStaffClientRepo .
func NewStaffClientRepo(data *Data, logger log.Logger) biz.StaffClientRepo {
	return &staffClientRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/staff_client")),
	}
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (r *staffClientRepo) Captcha(ctx context.Context, req *v1.CaptchaRequest) bool {
	exec, err := r.data.rdb.Exists(ctx, LimitCaptchaPre+req.Mobile).Result()
	if err != nil || exec > 0 {
		r.log.WithContext(ctx).Errorf("%v 非法超频请求验证码拦截", req.Mobile)
		return false
	}
	code := captcha.GenerateCaptcha(ctx)
	r.log.WithContext(ctx).Infof("[%v] 生成验证码 -> %v", req.Mobile, code)
	err = r.data.rdb.SetNX(ctx, LimitCaptchaPre+req.Mobile, code, time.Minute).Err()
	if err != nil {
		r.log.WithContext(ctx).Errorf("redis 设置 [%v] 验证码 [%v] 信息失败: %v", req.Mobile, code, err)
		return false
	}
	// TODO - sms发送code到mobile
	r.log.WithContext(ctx).Infof("sms 发送验证码 [%v] 到 mobile [%v]", code, req.Mobile)
	return true
}

func (r *staffClientRepo) Register(ctx context.Context, req *v1.RegisterRequest) bool {
	_, err := r.data.sc.CreateStaff(ctx, &staff_v1.CreateStaffRequest{
		Mobile: req.Mobile,
		Name:   req.Name,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[%v] staff服务注册接口错误: %v", req.Mobile, err)
		return false
	}
	return true
}

func (r *staffClientRepo) SignIn(ctx context.Context, req *v1.SignInRequest) *biz.StaffClient {
	exec, err := r.data.rdb.Get(ctx, LimitCaptchaPre+req.Mobile).Result()
	if err != nil {
		r.log.WithContext(ctx).Errorf("redis 获取 [%v] 验证码信息失败: %v", req.Mobile, err)
		return nil
	}
	if exec != req.Captcha {
		r.log.WithContext(ctx).Errorf("[%v] 验证码匹配错误", req.Mobile)
		return nil
	}
	staffs, err := r.data.sc.ListStaff(ctx, &staff_v1.ListStaffRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[%v] staff服务列表查询接口错误: %v", req.Mobile, err)
		return nil
	}
	if staffs.GetTotal() <= 0 {
		return nil
	}
	staff := staffs.GetData()[0]
	// 获取请求真实ip
	//remoteAddress := remote.GetRemoteAddressFromContext(ctx)
	err = r.data.rdb.SetNX(ctx, fmt.Sprintf("%v%v", SignMarkPre, staff.Id), req.Mobile, time.Hour*24).Err()
	if err != nil {
		r.log.WithContext(ctx).Errorf("redis 设置 [%v] staff 登录信息失败: %v", req.Mobile, err)
		return nil
	}
	return &biz.StaffClient{
		ID:        staff.Id,
		Mobile:    staff.Mobile,
		Name:      staff.Name,
		CreatedAt: staff.CreatedAt,
		UpdatedAt: staff.UpdatedAt,
	}
}

func (r *staffClientRepo) SignOut(ctx context.Context) bool {
	// 在上下文 context 中取出 claims 对象
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return false
	}
	claim := claims.(jwt2.MapClaims)
	if claim["ID"] == nil || claim["Mobile"] == nil {
		return false
	}
	exec, err := r.data.rdb.ExpireAt(ctx, fmt.Sprintf("%v%v", SignMarkPre, claim["ID"]), time.Now()).Result()
	if err != nil {
		r.log.WithContext(ctx).Errorf("redis 设置 [%v] 登录信息过期失败: %v", claim["Mobile"], err)
		return false
	}
	if !exec {
		return false
	}
	return true
}
