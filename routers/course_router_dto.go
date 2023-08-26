package routers

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	appvalidation "github.com/ratemyapp/validation"
)

type createCourseReq struct {
	Code  string  `code:"code"`
	Name  string  `json:"name"`
	Units float32 `json:"units"`
}

func (ccr *createCourseReq) Validate() error {
	ccr.Code = strings.TrimSpace(ccr.Code)
	ccr.Name = strings.TrimSpace(ccr.Name)
	return validation.ValidateStruct(ccr,
		validation.Field(&ccr.Code, validation.Required, is.UpperCase),
		validation.Field(&ccr.Name, validation.Required),
		validation.Field(&ccr.Units, validation.Required, validation.By(appvalidation.MaxDecimalPlaces(2))),
	)
}
