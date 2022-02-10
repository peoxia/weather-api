package user

type Deleter interface {
	DeleteUser(id string) error
}
