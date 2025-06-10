package errors

type UserValidate struct {
	err     error
	Message string
	Rules   string
}

func (e *UserValidate) Error() string {
	return e.err.Error()
}
