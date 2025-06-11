package repository

// не вижу применения интерфейсу
type UsersRepo interface {
	Insert(user *User) (int32, error)
}
