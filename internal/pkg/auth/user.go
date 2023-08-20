package auth

// User is the authorized meta info, include user id and other metadata.
type User struct {
	ID         string            // 用户id
	SessionKey string            // 用户密钥
	Metadata   map[string]string // 其他元数据，留作扩展
}

func (u *User) Value(key string) string {
	return u.Metadata[key]
}

func (u *User) SetValue(key, value string) {
	u.Metadata[key] = value
}
