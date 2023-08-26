package models

import "gorm.io/gorm"

type EmailDomain struct {
	gorm.Model
	SchoolID uint
	Domain   string
}
