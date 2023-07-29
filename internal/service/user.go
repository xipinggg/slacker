package service

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	pb "slacker/api/slacker/v1"
	"slacker/internal/biz"
)

var methods = map[string]jwt.SigningMethod{
	"256": jwt.SigningMethodHS256,
	"384": jwt.SigningMethodHS384,
	"512": jwt.SigningMethodHS512,
}

func SignMethod(method string) jwt.SigningMethod {
	return methods[method]
}

type User biz.User

type TokenClaims struct {
	jwt.RegisteredClaims // 标准字段

	User *User `json:"x-claim-user"` // 自定义字段
}

type UserService struct {
	pb.UnimplementedUserServer

	logger *log.Helper
	uc     *biz.UserUseCase
}

func NewUserService(l log.Logger, uc *biz.UserUseCase) *UserService {
	return &UserService{
		logger: log.NewHelper(l),
		uc:     uc,
	}
}

func (s *UserService) makeToken(_ context.Context, user *User, now time.Time, expires time.Duration) (string, error) {
	claims := TokenClaims{
		// 标准字段
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expires)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(now),              // 签发时间
			NotBefore: jwt.NewNumericDate(now),              // 生效时间
		},
		// 自定义字段
		User: user,
	}

	// todo
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("666"))
	if err != nil {
		return "", fmt.Errorf("%w\n%s", err, debug.Stack())
	}
	return token, nil
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.uc.Login(ctx, biz.LoginCodeType(req.Code))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expires := 24 * time.Hour //todo
	token, err := s.makeToken(ctx, (*User)(user), now, expires)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.UserInfo{
			Id:   string(user.ID),
			Name: user.Name,
		},
		Token: &pb.TokenInfo{
			Value:     token,
			ExpiresAt: now.Add(expires).Unix(),
		},
	}, nil
}
