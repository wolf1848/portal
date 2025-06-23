package app

import (
	"log"

	"github.com/wolf1848/gotaxi/config"
	"github.com/wolf1848/gotaxi/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() {

	conf, err := config.Init()
	if err != nil {
		panic("AAAAAA")
	}

	postgres, err := repository.InitPG(conf)
	if err != nil {
		panic("AAAAAAA DATABASE")
	}
	defer repository.ClosePG(postgres)

	mysql, err := repository.InitMysql(c)

	repository.InitRepo(conf)

	log.Print(conf.Database.Mysql.Password)

}
