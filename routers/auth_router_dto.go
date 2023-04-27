package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type signUpDto struct {
	Email     string
	Password  string
	Firstname string
	LastName  string
}

func (sd *signUpDto) Validation() error {
	return validation.ValidateStruct(sd,
		validation.Field(sd.Email, validation.Required, is.Email),
		validation.Field(sd.Password, validation.Required, validation.Min(8)),
		validation.Field(sd.Firstname, validation.Required, validation.Min(2)),
		validation.Field(sd.LastName, validation.Required, validation.Min(2)),
	)
}
