package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"slacker/internal/biz"
	"slacker/internal/pkg/util/errutil"

	"github.com/go-kratos/kratos/v2/log"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *UserRepo) GetSession(ctx context.Context, code biz.LoginCodeType) (*biz.Session, error) {
	session, err := r.data.WXClient.Auth.Session(ctx, string(code))
	if err != nil {
		err = errors.Unauthorized("BAD_LOGIN_REQUEST", err.Error()).WithCause(err)
		return nil, errutil.WithStack(err)
	}
	if session.ErrCode != 0 {
		return nil, errutil.WithStack(errors.Unauthorized("BAD_LOGIN_CODE", session.ErrMSG))
	}

	return &biz.Session{
		OpenID:     session.OpenID,
		UnionID:    session.UnionID,
		SessionKey: session.SessionKey,
	}, nil
}
