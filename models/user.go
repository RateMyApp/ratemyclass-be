package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname    string
	Lastname     string
	Password     string
	Address      string
	Program      string
	Email        string
	Phone        string
	CourseRating []CourseRating
	Course       []*Course `gorm:"many2many:course_students"`
}
