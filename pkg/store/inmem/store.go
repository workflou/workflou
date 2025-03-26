package inmem

type Store struct {
	*UserStore
	*SessionStore
}

func New() Store {
	return Store{
		UserStore:    NewUserStore(),
		SessionStore: NewSessionStore(),
	}
}
