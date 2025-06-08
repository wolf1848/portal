package repository

import (
	"context"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func (u *User) Create() (int32, error) {
	var userID int32

	err := Pool.QueryRow(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		u.Name, u.Email, u.Password).Scan(&userID)

	return userID, err
}
