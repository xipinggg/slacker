package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"slacker/api/slacker/v1"
	"slacker/internal/conf"
	"slacker/internal/pkg/middleware"
	"slacker/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return operation != v1.OperationUserLogin
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	logger log.Logger,
	confServer *conf.Server,
	confAuth *conf.Auth,
	user *service.UserService,
	record *service.RecordService,
) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(),
			logging.Server(logger),
			validate.Validator(),
			selector.Server(middleware.Auth(confAuth)).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if confServer.Http.Network != "" {
		opts = append(opts, http.Network(confServer.Http.Network))
	}
	if confServer.Http.Addr != "" {
		opts = append(opts, http.Address(confServer.Http.Addr))
	}
	if confServer.Http.Timeout != nil {
		opts = append(opts, http.Timeout(confServer.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)

	// 注册框架接口路由
	{
		// swagger ui, path: /q/swagger-ui
		srv.HandlePrefix("/q/", openapiv2.NewHandler())
		// healthz
		srv.Route("/").GET("/healthz", func(c http.Context) error {
			return c.Result(200, nil)
		})
	}

	// 注册应用层服务
	{
		// 用户管理
		v1.RegisterUserHTTPServer(srv, user)
		// 打卡管理
		v1.RegisterRecordHTTPServer(srv, record)
	}

	return srv
}
