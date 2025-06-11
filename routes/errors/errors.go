package errors

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// сделать приватным
var RulesMessage = map[string]string{
	"required": "Поле \"{field}\" обязательно для заполнения", // чтоб не экранировать ничего удобно юзать апостроф: `Поле "{field}" должно содержать валидный email`,
	"email":    `Поле "{field}" должно содержать валидный email`,
	"min":      "Минимальная длина поля \"{field}\" {params} символов",
}

// сделать приватным
var FieldTranslate = map[string]string{
	"Name":     "Имя",
	"Email":    "Email",
	"Password": "Пароль",
}

func GetMessage(e validator.FieldError) string {
	return strings.NewReplacer(
		"{field}", FieldTranslate[e.Field()], // "{field}" и "{params}" я б закинул в приватные константы
		"{params}", e.Param(),
	).Replace(RulesMessage[e.Tag()])
}
