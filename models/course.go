package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	Units        float32
	Code         string
	Professor    []*Professor `gorm:"many2many:professor_courses"`
	CourseRating []CourseRating
	SchoolID     uint
}
