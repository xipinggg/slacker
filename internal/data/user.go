package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"

	userbiz "slacker/internal/biz/user"
	"slacker/internal/pkg/util/errutil"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) userbiz.Repo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) GetSession(ctx context.Context, code string) (*userbiz.Session, error) {
	session, err := r.data.WXClient.Auth.Session(ctx, code)
	if err != nil {
		err = errors.Unauthorized("BAD_LOGIN_REQUEST", err.Error()).WithCause(err)
		return nil, errutil.WithStack(err)
	}
	if session.ErrCode != 0 {
		return nil, errutil.WithStack(errors.Unauthorized("BAD_LOGIN_CODE", session.ErrMsg))
	}

	return &userbiz.Session{
		OpenID:     session.OpenID,
		UnionID:    session.UnionID,
		SessionKey: session.SessionKey,
	}, nil
}

func (r *UserRepo) GetUser(ctx context.Context, id string) (*userbiz.User, error) {
	user, err := r.data.DBClient.User.Get(ctx, id)
	if err != nil {
		return nil, errutil.Wrap(err, "get by id failed: %s", id)
	}
	return &userbiz.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (r *UserRepo) SaveUser(ctx context.Context, user *userbiz.User) (*userbiz.User, error) {
	u, err := r.data.DBClient.User.
		Create().
		SetID(user.ID).
		SetName(user.Name).
		Save(ctx)
	if err != nil {
		return nil, errutil.WithStack(err)
	}

	return &userbiz.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}
