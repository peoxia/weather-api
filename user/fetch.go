package user

type Getter interface {
	GetUser(id string) (*Data, error)
}

type ListGetter interface {
	GetUsers(pageID int, country string) ([]Data, error)
}
