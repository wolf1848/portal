package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/wolf1848/gotaxi/repository"
)

func HomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Главная страница")
}

func GetUsersHandler(c echo.Context) error {
	ctx := c.Request().Context()

	rows, err := db.Pool.Query(ctx, "SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var u models.User
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Валидация с использованием зарегистрированного валидатора
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := db.Pool.Exec(ctx,
		"INSERT INTO users (name, email) VALUES ($1, $2)",
		u.Name, u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}
