package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type signUpDto struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (sd signUpDto) Validate() error {
	return validation.ValidateStruct(&sd,
		validation.Field(&sd.Email, validation.Required, is.Email),
		validation.Field(&sd.Password, validation.Required, validation.Length(8, 0)),
		validation.Field(&sd.Firstname, validation.Required, validation.Length(2, 0)),
		validation.Field(&sd.Lastname, validation.Required, validation.Length(2, 0)),
	)
}
