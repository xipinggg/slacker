package biz

//
//import (
//	"context"
//	"time"
//
//	"github.com/go-kratos/kratos/v2/log"
//	"slacker/internal/data/ent"
//)
//
//type UserRepo interface {
//	GetSession(ctx context.Context, code string) (*Session, error)
//	GetUser(ctx context.Context, id string) (*User, error)
//	SaveUser(ctx context.Context, user *User) error
//}
//
//type UserRoot struct {
//	user    *User
//	session *Session
//}
//
//func (r *UserRoot) ID() string {
//	return r.user.ID
//}
//
//func (r *UserRoot) SessionKey() string {
//	return r.session.SessionKey
//}
//
//// User 用户信息
//type User struct {
//	ID        string // 用户id
//	Name      string // 用户名
//	CreatedAt time.Time
//	UpdatedAt *time.Time
//}
//
//type Session struct {
//	OpenID     string
//	UnionID    string
//	SessionKey string // 用户密钥
//}
//
//type UserUseCase struct {
//	logger *log.Helper // 日志
//	repo   UserRepo    // 数据访问层
//}
//
//func NewUserUseCase(l log.Logger, repo UserRepo) *UserUseCase {
//	return &UserUseCase{
//		logger: log.NewHelper(l),
//		repo:   repo,
//	}
//}
//
//func (uc *UserUseCase) Login(ctx context.Context, code string) (*UserRoot, error) {
//	session, err := uc.repo.GetSession(ctx, code)
//	if err != nil {
//		return nil, err
//	}
//
//	userID := session.OpenID
//	user, err := uc.repo.GetUser(ctx, userID)
//	if err != nil {
//		if !ent.IsNotFound(err) {
//			return nil, err
//		}
//		// 用户不存在则创建一条记录
//		user = &User{
//			ID:   userID,
//			Name: time.Now().String(), // todo 临时名称
//		}
//		if err := uc.repo.SaveUser(ctx, user); err != nil {
//			return nil, err
//		}
//	}
//
//	return &UserRoot{
//		user:    user,
//		session: session,
//	}, nil
//}
