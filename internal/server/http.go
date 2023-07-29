package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"slacker/api/slacker/v1"
	"slacker/internal/conf"
	"slacker/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return operation != v1.OperationUserLogin
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(logger log.Logger, c *conf.Server, user *service.UserService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (any, error) {
					return []byte("666"), nil
				},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
					jwt.WithClaims(func() jwtv4.Claims {
						return &jwtv4.MapClaims{}
					}))).Match(NewWhiteListMatcher()).Build(),
			validate.Validator(),
		),
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
	// 注册应用层服务
	{
		v1.RegisterUserHTTPServer(srv, user)
	}

	return srv
}
