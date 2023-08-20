package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	pb "slacker/api/slacker/v1"
	"slacker/internal/biz"
	"slacker/internal/conf"
	"slacker/internal/pkg/auth"
)

type UserService struct {
	pb.UnimplementedUserServer

	conf   *conf.Auth
	logger *log.Helper
	uc     *biz.UserUseCase
}

func NewUserService(l log.Logger, uc *biz.UserUseCase, confAuth *conf.Auth) *UserService {
	return &UserService{
		conf:   confAuth,
		logger: log.NewHelper(l),
		uc:     uc,
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	user, err := s.uc.Login(ctx, biz.LoginCodeType(req.Code))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expires := s.conf.GetExpiresTime().AsDuration()

	u := &auth.User{
		ID:         string(user.ID),
		SessionKey: user.SessionKey,
	}
	token, err := auth.CreateToken(s.conf, u, now, expires)
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
