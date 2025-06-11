package repository

import (
	"context"
)

// Я предпочитаю хранить все в отдельныз файлах. Этот вынес бы как то так: /db/repositories/user/models/user.go
type User struct {
	Name     string
	Email    string
	Password string
}

/*
Я бы все таки выделял под каждую сущность свой пакет, например /db/repositories/user/repository.go
Структуру users тогда переименовать в repository, а конструктор выглядел бы так:

	func NewRepository(pool *pgxpool.Pool) *repository {
		return &repository{
			pool: pool,
		}
	}

при этом *pgxpool.Pool - я бы сразу заменил на интерфейс. (пока тебе нужен метод QueryRow)

Интерфейсы объявляются по месту применения, т.е. прям в этом файле.
Но в данном случае, т.к. количество репозиториев будет расти и подмена *pgxpool.Pool нужна будет много где,
то допустимо вынести это куда-то отдельно, я бы вынес как-то так:
/db/database/pool.go, можно убрать прослойку database, но так структура проекта поприятней на мой взгляд
*/
type users struct{}

func InitUsersRepo() *users {
	return &users{}
}

func (u *users) Insert(user *User) (int32, error) {
	var userID int32

	err := Pool.QueryRow(context.Background(), // для Pool стоит выделить поле в структуре users
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password).Scan(&userID)

	return userID, err
}
