package remote

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
)

func GetRemoteAddressFromContext(ctx context.Context) string {
	if tp, ok := transport.FromServerContext(ctx); ok {
		return tp.RequestHeader().Get("Remote-Address")
	}
	return ""
}
