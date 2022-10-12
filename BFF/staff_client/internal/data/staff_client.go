package data

import (
	"context"
	"fmt"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/biz"
	"staff_client/internal/pkg/captcha"
	"staff_client/internal/pkg/middleware/auth"
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
		r.log.WithContext(ctx).Errorf("[%v] 非法超频请求验证码拦截", req.Mobile)
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
		r.log.WithContext(ctx).Errorf("[%v] staff服务 - staff新增接口错误: %v", req.Mobile, err)
		return false
	}
	return true
}

func (r *staffClientRepo) SignIn(ctx context.Context, req *v1.SignInRequest) *staff_v1.StaffReply {
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
		r.log.WithContext(ctx).Errorf("[%v] staff服务 - staff列表查询接口错误: %v", req.Mobile, err)
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
	return staff
}

func (r *staffClientRepo) SignOut(ctx context.Context) bool {
	// 在上下文 context 中取出 claims 对象
	claim := auth.GetClaimFormContext(ctx)
	if claim["ID"] == nil || claim["Mobile"] == nil {
		r.log.WithContext(ctx).Error("ctx 获取claim数据异常, err: 无法正确提取信息字段")
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

/*--------------------------------------------------------------------------------------------------------------------*/

func (r *staffClientRepo) ReadStaff(ctx context.Context) *staff_v1.ReadStaffReply {
	// 在上下文 context 中取出 claims 对象
	claim := auth.GetClaimFormContext(ctx)
	if claim["ID"] == nil {
		r.log.WithContext(ctx).Error("ctx 获取claim数据异常, err: 无法正确提取信息字段")
		return nil
	}
	staff, err := r.data.sc.ReadStaff(ctx, &staff_v1.ReadStaffRequest{
		Id: claim["ID"].(int64),
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[%v] staff服务 - staff查询接口错误: %v", claim["ID"], err)
		return nil
	}
	return staff
}

func (r *staffClientRepo) UpdateStaff(ctx context.Context, req *v1.UpdateStaffRequest) bool {
	// 在上下文 context 中取出 claims 对象
	claim := auth.GetClaimFormContext(ctx)
	if claim["ID"] == nil || claim["Name"] == nil {
		r.log.WithContext(ctx).Error("ctx 获取claim数据异常, err: 无法正确提取信息字段")
		return false
	}
	exec, err := r.data.sc.UpdateStaff(ctx, &staff_v1.UpdateStaffRequest{
		Id:   claim["ID"].(int64),
		Name: claim["Name"].(string),
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[%v] staff服务 - staff更新接口错误: %v", claim["ID"], err)
		return false
	}
	return exec.Exec
}

/*--------------------------------------------------------------------------------------------------------------------*/

func (r *staffClientRepo) ListStaffTeam(ctx context.Context, req *v1.ListStaffTeamRequest) *staff_v1.ListStaffTeamReply {
	// 在上下文 context 中取出 claims 对象
	claim := auth.GetClaimFormContext(ctx)
	if claim["ID"] == nil {
		r.log.WithContext(ctx).Error("ctx 获取claim数据异常, err: 无法正确提取信息字段")
		return nil
	}
	exec, err := r.data.sc.ListStaffTeam(ctx, &staff_v1.ListStaffTeamRequest{
		Pn:    req.Pn,
		PSize: req.PSize,
		Sid:   claim["ID"].(int64),
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[%v] staff服务 - team与staff关系列表接口错误: %v", claim["ID"], err)
		return nil
	}
	return exec
}
