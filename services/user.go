package services

import (
	"log"
	"time"

	db "github.com/wolf1848/gotaxi/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	password  string //Делаю приватным для того что бы нельзя было установить не хэшированное значение.
	CreatedAt time.Time
	UpdatedAt time.Time
}

type users struct{}

func InitUserSerice() *users {
	return &users{}
}

func (u *User) SetPwd(pwd string) {
	var err error
	u.password, err = hashPassword(pwd)
	if err != nil {
		log.Println(err) //Как то надо оповестить что хэш не был создан хотя не понятно почему он может не создаться
		// И как будто в последствии эту проблему надо выкидывать наверх
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) modelToEntity() *db.User {
	var user db.User
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.password

	return &user
}

/*
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}*/

func (service *users) Add(user *User) error {
	var err error

	userRepo := db.InitUsersRepo()

	user.ID, err = userRepo.Insert(user.modelToEntity())

	//Как будто здесь должен быть маппинг из ошибок бд в человекочитаемые

	return err
}
