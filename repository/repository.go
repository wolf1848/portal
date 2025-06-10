package repository

type UsersRepo interface {
	Insert(user *User) (int32, error)
}
