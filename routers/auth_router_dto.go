package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type registerUserReq struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (self registerUserReq) Validate() error {
	return validation.ValidateStruct(&self,
		validation.Field(&self.Email, validation.Required, is.Email),
		validation.Field(&self.Password, validation.Required, validation.Length(8, 0)),
		validation.Field(&self.Firstname, validation.Required, validation.Length(2, 0)),
		validation.Field(&self.Lastname, validation.Required, validation.Length(2, 0)),
	)
}

type loginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginUserResp struct {
	Firstname   string
	Lastname    string
	Email       string
	AccessToken string
}
