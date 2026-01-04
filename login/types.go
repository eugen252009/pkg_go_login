package login

import "net/http"

var ErrorParseForm = &LoginError{Code: http.StatusUnprocessableEntity, Message: "Invalid Formdata"}

const (
	ContextAuthKey   = "login/auth/data"
	LogLoginLogin    = "Login/Login"
	LogLoginRegister = "Login/Register"
)

type AuthInfo struct {
	Email    string
	Password string
	ParseErr *error
}

type LoginError struct {
	Code    int
	Message string
}
