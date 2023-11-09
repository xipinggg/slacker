package user

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"slacker/internal/data/ent"
)

type Repo interface {
	GetSession(ctx context.Context, code string) (*Session, error)
	GetUser(ctx context.Context, id string) (*User, error)
	SaveUser(ctx context.Context, user *User) (*User, error)
}

type UseCase struct {
	logger *log.Helper // 日志
	repo   Repo        // 资源库
}

func NewUseCase(l log.Logger, repo Repo) *UseCase {
	return &UseCase{
		logger: log.NewHelper(l),
		repo:   repo,
	}
}

func (uc *UseCase) Login(ctx context.Context, code string) (*Root, error) {
	session, err := uc.repo.GetSession(ctx, code)
	if err != nil {
		return nil, err
	}

	userID := session.OpenID
	user, err := uc.repo.GetUser(ctx, userID)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}
		var err error
		// 用户不存在则创建一条记录
		user, err = uc.repo.SaveUser(ctx, &User{
			ID:   userID,
			Name: time.Now().String(), // todo 临时名称
		})
		if err != nil {
			return nil, err
		}
	}

	return &Root{
		user:    user,
		session: session,
	}, nil
}
