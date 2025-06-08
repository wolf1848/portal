package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	db "github.com/wolf1848/gotaxi/repository"
	"github.com/wolf1848/gotaxi/services"
)

// CustomValidator - кастомная структура валидатора
type CustomValidator struct {
	validator *validator.Validate
}

// Validate - реализация интерфейса валидатора
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Инициализация БД
	if err := db.InitDB(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.CloseDB()

	user := services.User{Name: "Тестович", Email: "test1@mail.ru", Password: "123321"}

	_, err := user.Create()

	log.Println(err)

	/*e := echo.New()

	// Регистрация валидатора
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.GET("/", handlers.HomeHandler)
	e.GET("/users", handlers.GetUsersHandler)
	e.POST("/users", handlers.CreateUserHandler)

	e.Logger.Fatal(e.Start(":8080"))*/
}
