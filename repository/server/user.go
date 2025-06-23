package server

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type userRepo struct {
	db *pgxpool.Pool
}

func (u *userRepo) Insert(user *User) (int32, error) {
	var userID int32

	err := u.db.QueryRow(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password).Scan(&userID)

	return userID, err
}

func (u *userRepo) GetUserId(id int32) (*User, error) {
	var user User

	err := u.db.QueryRow(context.Background(),
		"SELECT id, name, email, password FROM users WHERE id=$1 ",
		id).Scan(&user)

	return &user, err
}

func initUsersRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db,
	}
}
