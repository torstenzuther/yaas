package dal

type Stores struct {
	GrantStore   GrantStore
	SessionStore SessionStore
	ClientStore  ClientStore
	UserStore    UserStore
}
