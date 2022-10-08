package captcha

import (
	"context"
	"fmt"
	"time"
)

func SendCaptcha(_ context.Context, mobile string) (int, error) {
	code := time.Now().Nanosecond() / 1e3
	fmt.Printf("[%s] 生成验证码 -> %v", mobile, code)
	return code, nil
}
