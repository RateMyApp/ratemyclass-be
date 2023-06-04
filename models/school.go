package models

import (
	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	Name             string
	Location         string
	StudentHeadcount int
	StaffHeadcount   int
	Programs         []Program
	Course           []Course
}
