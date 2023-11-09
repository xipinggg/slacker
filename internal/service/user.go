package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "slacker/api/slacker/v1"
	userbiz "slacker/internal/biz/user"
	"slacker/internal/conf"
	"slacker/internal/pkg/auth"
)

type UserService struct {
	pb.UnimplementedUserServer

	conf   *conf.Auth
	logger *log.Helper
	uc     *userbiz.UseCase
}

func NewUserService(l log.Logger, uc *userbiz.UseCase, confAuth *conf.Auth) *UserService {
	return &UserService{
		conf:   confAuth,
		logger: log.NewHelper(l),
		uc:     uc,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	root, err := s.uc.Login(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expires := s.conf.GetExpiresTime().AsDuration()

	u := &auth.User{
		ID:         root.ID(),
		SessionKey: root.SessionKey(),
	}
	token, err := auth.CreateToken(s.conf, u, now, expires)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		User: &pb.UserInfo{
			Id: root.ID(),
		},
		Token: &pb.TokenInfo{
			Value:     token,
			ExpiresAt: now.Add(expires).Unix(),
		},
	}, nil
}
