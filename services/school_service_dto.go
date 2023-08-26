package services

type CreateSchoolCommand struct {
	Name            string
	Address         string
	City            string
	ProvinceOrState string
	Country         string
	EmailDomains    []string
}

type SearchSchoolInfoResult struct {
	Id       uint
	Name     string
	Location string
}
