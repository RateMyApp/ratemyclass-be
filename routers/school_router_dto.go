package routers

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CreateSchoolReq struct {
	Address         string   `json:"address"`
	City            string   `json:"city"`
	PronviceOrState string   `json:"provinceOrState"`
	Country         string   `json:"country"`
	Name            string   `json:"name"`
	EmailDomains    []string `json:"domains"`
}

func (csr *CreateSchoolReq) Validate() error {

	csr.Name = strings.TrimSpace(csr.Name)
	csr.Address = strings.TrimSpace(csr.Address)
	csr.City = strings.TrimSpace(csr.City)
	csr.PronviceOrState = strings.TrimSpace(csr.PronviceOrState)
	csr.Country = strings.TrimSpace(csr.Country)

	return validation.ValidateStruct(csr,
		validation.Field(&csr.Name, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&csr.Address, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&csr.Country, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&csr.PronviceOrState, validation.Required, validation.NilOrNotEmpty),
		validation.Field(&csr.EmailDomains, validation.Required, validation.Length(1, 0), validation.Each(is.Email)),
	)
}

type SearchForSchoolInfoRes struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
