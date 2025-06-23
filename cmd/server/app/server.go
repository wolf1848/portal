package app

import (
	"log"

	"github.com/wolf1848/gotaxi/config"
	repo "github.com/wolf1848/gotaxi/repository/server"
)

func Run() {

	conf, err := config.Init()
	if err != nil {
		panic("Not config file!")
	}

	postgres, err := repo.InitDB(conf)
	if err != nil {
		panic("Databse not connection")
	}
	defer repo.Close(postgres)

	repository := repo.InitRepo(postgres)

	//userID, err := repository.User.Insert(&repo.User{Name: "test", Email: "test@mail.ru", Password: "hash"})

	log.Println(err)

	user, err := repository.User.GetUserId(1)

	log.Println(err)

	log.Println(user.Name)

}
