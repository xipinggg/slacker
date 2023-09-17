package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo interface {
	GetSession(ctx context.Context, code LoginCodeType) (*Session, error)
}

type (
	UserIDType    string // 用户id类型
	LoginCodeType string // 登录凭证类型
)

// User 用户信息
type User struct {
	ID         UserIDType // 用户id
	SessionKey string     // 用户密钥
}

type Session struct {
	OpenID     string
	UnionID    string
	SessionKey string
}

type verifyCodeResult struct {
	ID         string
	SessionKey string
}

type UserUseCase struct {
	logger *log.Helper // 日志
	repo   UserRepo    // 数据访问层
}

func NewUserUseCase(l log.Logger, repo UserRepo) *UserUseCase {
	return &UserUseCase{
		logger: log.NewHelper(l),
		repo:   repo,
	}
}

func (uc *UserUseCase) Login(ctx context.Context, code LoginCodeType) (*User, error) {
	result, err := uc.verifyCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:         UserIDType(result.ID),
		SessionKey: result.SessionKey,
	}, nil
}

func (uc *UserUseCase) verifyCode(ctx context.Context, code LoginCodeType) (*verifyCodeResult, error) {
	session, err := uc.repo.GetSession(ctx, code)
	if err != nil {
		return nil, err
	}

	return &verifyCodeResult{
		ID:         session.OpenID,
		SessionKey: session.SessionKey,
	}, nil
}
