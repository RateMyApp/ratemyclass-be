package models

import (
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name             string
	Address          string
	City             string
	ProvinceOrState  string
	Country          string
	StudentHeadcount int
	StaffHeadcount   int
	Programs         []Program
	Course           []Course
	EmailDomains     []EmailDomain
}
