package user

// Root 聚合根
type Root struct {
	user    *User
	session *Session
}

func (r *Root) ID() string {
	return r.user.ID
}

func (r *Root) SessionKey() string {
	return r.session.SessionKey
}
