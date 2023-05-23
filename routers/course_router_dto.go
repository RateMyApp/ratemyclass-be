package routers

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	appvalidation "github.com/ratemyapp/validation"
)

type createCourseReq struct {
	Code  string
	Name  string
	Units float32
}

func (self createCourseReq) Validate() error {
	return validation.ValidateStruct(&self,
		validation.Field(&self.Code, validation.Required, is.UpperCase),
		validation.Field(&self.Name, validation.Required),
		validation.Field(&self.Units, validation.Required, validation.By(appvalidation.MaxDecimalPlaces(2))),
	)
}
