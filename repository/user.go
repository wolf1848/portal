package repository

import (
	"context"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type users struct{}

func InitUsersRepo() *users {
	return &users{}
}

func (u *users) Insert(user *User) (int32, error) {
	var userID int32

	err := Pool.QueryRow(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password).Scan(&userID)

	return userID, err
}
