package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wolf1848/gotaxi/routes/dto"
	"github.com/wolf1848/gotaxi/routes/errors"
	"github.com/wolf1848/gotaxi/services"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func ServerInit() {
	e := echo.New()

	v := validator.New()
	e.Validator = &CustomValidator{validator: v}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Роуты
	e.POST("/users/create", CreateUserHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

func dtoToModel(d *dto.User) *services.User {
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

		var errorMap = map[string]string{}

		if errs, ok := err.(validator.ValidationErrors); ok {
			// Обрабатываем ошибки валидации
			for _, e := range errs {
				errorMap[strings.ToLower(e.Field())] = errors.GetMessage(e)
			}
		}

		return c.JSON(400, errorMap)
	}

	userService := services.InitUserSerice()

	userModel := dtoToModel(&userDto)

	err = userService.Add(userModel)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}
