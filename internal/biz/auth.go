package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"slacker/internal/conf"
)

var methods = map[string]jwt.SigningMethod{
	"256": jwt.SigningMethodHS256,
	"384": jwt.SigningMethodHS384,
	"512": jwt.SigningMethodHS512,
}

func SignMethod(method string) jwt.SigningMethod {
	return methods[method]
}

type TokenClaims struct {
	jwt.RegisteredClaims // 标准字段

	User *User `json:"x-claim-user"` // 自定义字段
}

type AuthUsecase struct {
	logger *log.Helper // 日志

	secretKey   []byte            // 签名密钥
	method      jwt.SigningMethod // 签名方法
	expiresTime time.Duration     // token过期时间
}

func NewAuthUsecase(l log.Logger, c *conf.Auth) *AuthUsecase {
	return &AuthUsecase{
		logger:      log.NewHelper(l),
		secretKey:   []byte(c.SecretKey),
		method:      SignMethod(c.Method),
		expiresTime: c.ExpiresTime.AsDuration(),
	}
}

func (uc *AuthUsecase) CreateToken(_ context.Context, user *User) (string, error) {
	now := time.Now()
	claims := TokenClaims{
		// 标准字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(uc.expiresTime)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),                     // 签发时间
			NotBefore: jwt.NewNumericDate(now),                     // 生效时间
		},
		// 自定义字段
		User: user,
	}

	return jwt.NewWithClaims(uc.method, claims).SignedString(uc.secretKey)
}
