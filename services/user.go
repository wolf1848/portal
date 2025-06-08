package services

import (
	"log"
	"time"

	db "github.com/wolf1848/gotaxi/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

/*
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}*/

func (u *User) Create(password string) (int32, error) {

	pwd, err := hashPassword(password)
	if err != nil {
		log.Println(err) //Ту надо прервать исполнение создания пользователя
	}

	userRepo := db.User{Name: u.Name, Email: u.Email, Password: pwd}

	userId, errDb := userRepo.Create()

	//Как будто здесь должен быть маппинг из ошибок бд в человекочитаемые

	return userId, errDb
}
