package biz

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/go-kratos/kratos/v2/log"
)

type (
	UserIDType    string // 用户id类型
	LoginCodeType string // 登录凭证类型
)

// User 用户信息
type User struct {
	ID         UserIDType `json:"id,omitempty"`          // 用户id
	Name       string     `json:"name,omitempty"`        // 用户名
	SessionKey string     `json:"session_key,omitempty"` // 用户密钥
}

type UserUseCase struct {
	logger *log.Helper // 日志
}

func NewUserUseCase(l log.Logger) *UserUseCase {
	return &UserUseCase{
		logger: log.NewHelper(l),
	}
}

type VerifyCodeResult struct {
	ID         string
	SessionKey string
}

func (uc *UserUseCase) verifyCode(ctx context.Context, code LoginCodeType) (*VerifyCodeResult, error) {
	// todo 临时写这里测试
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  "wxfc042c932a156bec",   // 小程序app id
		Secret: "[miniprogram_secret]", // 小程序app secret
	})

	if err != nil {
		return nil, fmt.Errorf("%w\n%s", err, debug.Stack())
	}

	session, err := app.Auth.Session(ctx, string(code))
	// todo debug
	uc.logger.WithContext(ctx).Infof("[debug] app.Auth.Session(ctx, string(code)) result: %s", session)
	if err != nil {
		return nil, fmt.Errorf("%w\n%s", err, debug.Stack())
	}
	return &VerifyCodeResult{
		ID:         session.UnionID,
		SessionKey: session.SessionKey,
	}, nil
}

func (uc *UserUseCase) Login(ctx context.Context, code LoginCodeType) (*User, error) {
	result, err := uc.verifyCode(ctx, code)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:         UserIDType(result.ID),
		Name:       "unknown", // todo 不知道怎么获取
		SessionKey: result.SessionKey,
	}, nil
}
