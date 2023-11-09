package user

import (
	"time"
)

// User 用户信息实体
type User struct {
	ID        string // 用户id
	Name      string // 用户名
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// Session 会话信息实体
type Session struct {
	OpenID     string // 关联UserID
	UnionID    string
	SessionKey string // 用户密钥
}
