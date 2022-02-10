package user

type UserCreatedUpdatedNotifier interface {
	NotifyUserCreatedUpdated(user Data) error
}

type UserDeletedNotifier interface {
	NotifyUserDeleted(userID string) error
}
