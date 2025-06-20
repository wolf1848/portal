package app

import (
	"log"

	"github.com/wolf1848/gotaxi/config"
	"github.com/wolf1848/gotaxi/repository"
)

func Run() {

	conf, err := config.Init()
	if err != nil {
		panic("AAAAAA")
	}

	repository.InitRepo(conf)

	log.Print(conf.Database.Mysql.Password)

}
