package captcha

import (
	"context"
	"time"
)

func GenerateCaptcha(_ context.Context) int {
	code := time.Now().Nanosecond() / 1e3
	return code
}
