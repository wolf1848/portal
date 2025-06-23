package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	User *userRepo
}

func InitRepo(postgres *pgxpool.Pool) *Repo {

	return &Repo{
		initUsersRepo(postgres),
	}

}
