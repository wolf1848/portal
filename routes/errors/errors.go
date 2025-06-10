package errors

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var RulesMessage = map[string]string{
	"required": "Поле \"{field}\" обязательно для заполнения",
	"email":    "Поле \"{field}\" должно содержать валидный email",
	"min":      "Минимальная длина поля \"{field}\" {params} символов",
}

var FieldTranslate = map[string]string{
	"Name":     "Имя",
	"Email":    "Email",
	"Password": "Пароль",
}

func GetMessage(e validator.FieldError) string {
	return strings.NewReplacer(
		"{field}", FieldTranslate[e.Field()],
		"{params}", e.Param(),
	).Replace(RulesMessage[e.Tag()])
}
