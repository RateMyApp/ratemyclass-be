package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string
	Address        	 string
	Program          string
	Email            string
	Phone            string
	CourseRating []CourseRating
	Course       []*Course `gorm:"many2many:course_students"`
}
