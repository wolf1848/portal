package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gotaxi/routes/dto"
	"gotaxi/routes/errors"
	"gotaxi/services"
)

/*
Если нужна реализация валидатора, то пусть идет куда-то в отдельный пакет.
Приоритетно приватная структура и публичный конструктор:
NewCustomValidator(validator *validator.Validate) {...}
*/
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

/*
Сервер и хандлеры я б тоже разбил по разным файлам, например:

Вариант 1:
/routes/server/server.go
/routes/handlers/create_user.go
/routes/handlers/update_user.go
/routes/handlers/delete_user.go

Вариант 2:
/routes/server/server.go
/routes/handlers/user/create.go
/routes/handlers/user/update.go
/routes/handlers/user/delete.go

или вариант 3 позволяющий оставить хандлеры приватными:
/routes/server.go
/routes/create_user_handler.go
/routes/update_user_handler.go
/routes/delete_user_handler.go

В целом юзабельны все, я предпочитаю первый и второй. Третий уместен если файлов не много
*/
func ServerInit() {
	e := echo.New()

	v := validator.New()
	e.Validator = &CustomValidator{validator: v}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.POST("/users/create", CreateUserHandler)

	e.Logger.Fatal(e.Start(":8080")) // адрес и порт в конфиг
}

func dtoToModel(d *dto.User) *services.User { // я такие штуки тоже выкидываю в отдельный пакет/файл
	var s services.User
	s.Name = d.Name
	s.Email = d.Email
	s.SetPwd(d.Password)
	return &s
}

func CreateUserHandler(c echo.Context) error {

	var userDto dto.User
	var err error
	if err = c.Bind(&userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(userDto); err != nil {
		/*
			Часть по обработке ошибок можно внести куда-то, потому что во всех других апихах оно тоже пригодится
		*/

		var errorMap = map[string]string{}

		if errs, ok := err.(validator.ValidationErrors); ok {
			// Обрабатываем ошибки валидации
			for _, e := range errs {
				/*
					Могу ошибаться, но кажется что тут потеряются ошибки, в случае если у одного поля будет нарушено сразу несколько правил.
					Пофиксить можно заменив мапу на слайс.

					Еще минус мапы в том, что она не сортирована, при нескольких вызовах апи один и тот же набор ошибок будет выводиться в разном порядке, это не найс
				*/
				errorMap[strings.ToLower(e.Field())] = errors.GetMessage(e)
			}
		}

		return c.JSON(400, errorMap) // 400 поменять на http.StatusBadRequest
	}

	userService := services.InitUserSerice() // Сервис внутри себя должен принять реру,

	userModel := dtoToModel(&userDto)

	err = userService.Add(userModel)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}
