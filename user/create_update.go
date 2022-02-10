package user

type CreatorUpdater interface {
	CreateOrUpdateUser(updatedData Data) (string, error)
}
