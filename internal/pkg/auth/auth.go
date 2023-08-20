package auth

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"slacker/internal/conf"
	"slacker/internal/pkg/util/errutil"
	"slacker/internal/pkg/util/strutil"
)

type TokenClaims struct {
	jwtv4.RegisteredClaims // 标准字段

	User *User `json:"x-claim-custom"` // 自定义字段
}

func CreateToken(c *conf.Auth, user *User, now time.Time, expires time.Duration) (string, error) {
	claims := &TokenClaims{
		// 标准字段
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(now.Add(expires)), // 过期时间
			IssuedAt:  jwtv4.NewNumericDate(now),              // 签发时间
			NotBefore: jwtv4.NewNumericDate(now),              // 生效时间
		},
		// 自定义字段
		User: user,
	}

	token, err := jwtv4.NewWithClaims(SigningMethod(c.GetMethod()), claims).SignedString(strutil.Bytes(c.GetSecretKey()))
	if err != nil {
		return "", errutil.WithStack(err)
	}

	return token, nil
}

// FromContext returns the authorized user meta info, read only.
func FromContext(ctx context.Context) (*User, error) {
	claims, ok := jwt.FromContext(ctx)
	if !ok {
		return nil, errutil.WithStack(errors.Unauthorized("UNAUTHORIZED", "claims not found from context"))
	}

	customClaims, ok := claims.(*TokenClaims)
	if !ok {
		return nil, errutil.WithStack(errors.Unauthorized("UNAUTHORIZED", "custom claims type is not supported"))
	}

	return customClaims.User, nil
}

func WithClaims() jwtv4.Claims {
	return &TokenClaims{}
}
