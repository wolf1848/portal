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
	"github.com/wolf1848/gotaxi/config"

	"database/sql"
    _ "github.com/go-sql-driver/mysql"
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
	e.GET("/test", TestMysql)

	e.Logger.Fatal(e.Start(":8080"))
}

func TestMysql(c echo.Context) error {

	db, err := sql.Open("mysql", config.GetDBConnectionString("mysql"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM b_user LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s\n", id, name)
	}

	return c.JSON(http.StatusCreated, map[string]string{"test": "route test"})
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
