package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	netHttp "net/http"
	v1 "staff_client/api/staff_client/v1"
	"staff_client/internal/conf"
	"staff_client/internal/service"
	"strings"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, auth *conf.Auth, service *service.StaffClientService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			// 接口访问的参数校验
			validate.Validator(),
			tracing.Server(),
			// jwt 验证
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(auth.JwtKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
			).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
		// 浏览器跨域设置
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		), LocalHttpRequestFilter()),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterStaffClientHTTPServer(srv, service)
	return srv
}

/*--------------------------------------------------------------------------------------------------------------------*/

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.staff_client.v1.StaffClient/Captcha"] = struct{}{}
	whiteList["/api.staff_client.v1.StaffClient/SignIn"] = struct{}{}
	whiteList["/api.staff_client.v1.StaffClient/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

/*--------------------------------------------------------------------------------------------------------------------*/

// LocalHttpRequestFilter 获取http请求的真实IP并填充到Header中
func LocalHttpRequestFilter() http.FilterFunc {
	return func(next netHttp.Handler) netHttp.Handler {
		return netHttp.HandlerFunc(func(writer netHttp.ResponseWriter, req *netHttp.Request) {
			remoteAddr := RemoteIp(req)
			req.Header.Add("Remote-Address", remoteAddr)
			next.ServeHTTP(writer, req)
		})
	}
}

// RemoteIp 获取请求中的IP
func RemoteIp(req *netHttp.Request) string {
	var remoteAddr string
	// RemoteAddr
	remoteAddr = req.RemoteAddr
	if remoteAddr != "" {
		return strings.Split(remoteAddr, ":")[0]
	}
	// ipv4
	remoteAddr = req.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Forwarded-For
	remoteAddr = req.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Real-Ip
	remoteAddr = req.Header.Get("X-Real-Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
