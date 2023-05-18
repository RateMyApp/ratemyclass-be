package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateProfReq struct {
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	DirectoryListing string `json:"directoryListing"`
	Email            string `json:"email"`
	Department       string `json:"department"`
}

func (crp CreateProfReq) Validate() error {
	return validation.ValidateStruct(&crp,
		validation.Field(&crp.Email, validation.Required, is.Email),
		validation.Field(&crp.Firstname, validation.Required),
		validation.Field(&crp.Lastname, validation.Required),
		validation.Field(&crp.Department, validation.Required),
		validation.Field(&crp.DirectoryListing, validation.Required, is.URL),
	)
}
