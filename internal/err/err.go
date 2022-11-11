package err

import (
	"net/http"
)

type CustomErr struct {
	HttpStatus int
	Message    string
}

func (e CustomErr) Error() string {
	return e.Message
}

func ErrorByStatus(status int) CustomErr {
	e := CustomErr{}
	switch status {
	case 200:
		e.HttpStatus = http.StatusOK
	case 400:
		e.HttpStatus = http.StatusBadRequest
	case 401:
		e.HttpStatus = http.StatusUnauthorized
	case 408:
		e.HttpStatus = http.StatusRequestTimeout
	default:
		e.HttpStatus = http.StatusInternalServerError
	}
	e.Message = http.StatusText(e.HttpStatus)
	return e
}
