package exceptions

import (
	"net/http"
	"time"
)

type AppError struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"status"`
	TimeStamp  string      `json:"timestamp"`
}

func NewInternalServerError() AppError {
	return AppError{
		Message:    "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}

func NewConflictError(message interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusConflict,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}

func NewBadRequestError(message interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}

func NewUnauthorizedError(message interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}

func NewUnprocessableEntityError(message interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}

func NewNotFoundError(message interface{}) AppError {
	return AppError{
		Message:    message,
		StatusCode: http.StatusNotFound,
		TimeStamp:  time.Now().UTC().Local().String(),
	}
}
