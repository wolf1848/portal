package services

type UsersService interface {
	Add(user *User) error
}
