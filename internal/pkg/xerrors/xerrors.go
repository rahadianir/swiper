package xerrors

import (
	"errors"
	"net/http"
)

type ClientError struct {
	Err error
}

func (e ClientError) Error() string {
	return e.Err.Error()
}

type ServerError struct {
	Err error
}

func (e ServerError) Error() string {
	return e.Err.Error()
}

type LogicError struct {
	Err error
}

func (e LogicError) Error() string {
	return e.Err.Error()
}

type AuthError struct {
	Err error
}

func (e AuthError) Error() string {
	return e.Err.Error()
}

func ParseErrorTypeToCodeInt(err error) int {
	switch {
	case errors.As(err, &LogicError{}): //200
		return http.StatusOK
	case errors.As(err, &ClientError{}): //400
		return http.StatusBadRequest
	case errors.As(err, &AuthError{}): //401
		return http.StatusUnauthorized
	case errors.As(err, &ServerError{}): //500
		return http.StatusInternalServerError
	default: //500
		return http.StatusInternalServerError
	}
}
