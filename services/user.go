package services

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	db "gotaxi/repository"
)

/*
Не буду повторяться, в целом все моменты подсвечены в repository/user.go
Тут все аналогично
*/

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
	u.password, _ = hashPassword(pwd)
	if err != nil {
		log.Println(err) //Как то надо оповестить что хэш не был создан хотя не понятно почему он может не создаться
		// И как будто в последствии эту проблему надо выкидывать наверх
	}

	/*
		Тут 2 стратегии:
		1) прокидывать ошибку наверх. Т.е. метод SetPwd(..) должен возвращать error
		2) забить на ошибку (если прям уверен что она невозможна).

		Для второго либо не возвращаешь ее из hashPassword в принципе,
		либо тут уже делаешь так:
		u.password, _ = hashPassword(pwd)

	*/
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// модель User вынести в свой файл. + методы модели объявлять лучше непосредственно рядом с моделью. Сейчас модель объявлена где-то выше, а метод у нее появляется где-то в середине файла
func (u *User) modelToEntity() *db.User {
	var user db.User
	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.password

	return &user

	/*
		Можно метод реализовать вот так:

		return &db.User{
			Name:     u.Name,
			Email:    u.Email,
			Password: u.password,
		}

		И назвать просто toEntity() - model по мне лишнее.
		Даже лучше будет так: toDbModel() или toDbEntity() - длиннее, но более читаемо
	*/
}

/*
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}*/

func (service *users) Add(user *User) error {
	var err error

	/*
		Лучше 1 раз создать репо и потом его прокидывать, чем инициализировать на каждый вызов Add().
		Т.е. в конструктор создания InitUserSerice надо бы прокиндывать парметр - репозиторий
	*/
	userRepo := db.InitUsersRepo() //

	user.ID, err = userRepo.Insert(user.modelToEntity())

	//Как будто здесь должен быть маппинг из ошибок бд в человекочитаемые

	return err
}
