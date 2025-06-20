package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/gotaxi/config"
)

type UsersRepo interface {
	Insert(user *User) (int32, error)
}

type Database interface{}

type Repo struct {
	db struct {
		pg *pgxpool.Pool
		mysql *sql.DB
	}
}

func InitRepo(c *config.Config) *Repo {
	var err error

	r := &Repo{}

	r.db.pg, err = initPG(c)
	if err != nil {
		panic("AAAAAAA DATABASE")
	}
	defer closePG(r.db.pg)

	r.db.mysql, err = InitMysql(c)

	return r
}
