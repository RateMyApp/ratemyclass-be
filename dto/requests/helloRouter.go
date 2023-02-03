package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SayHelloRequest struct {
	Name string `json:"name"`
}

func (sr SayHelloRequest) Validate() error {
	return validation.ValidateStruct(&sr, validation.Field(&sr.Name, validation.Required.Error("is required")))
}
