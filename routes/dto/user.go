package dto

type User struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}
