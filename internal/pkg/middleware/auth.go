package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"slacker/internal/conf"
	"slacker/internal/pkg/auth"
	"slacker/internal/pkg/util/strutil"
)

func Auth(c *conf.Auth) middleware.Middleware {
	return jwt.Server(
		func(token *jwtv4.Token) (any, error) {
			return strutil.Bytes(c.GetSecretKey()), nil
		},
		jwt.WithSigningMethod(auth.SigningMethod(c.GetMethod())),
		jwt.WithClaims(auth.WithClaims),
	)
}
