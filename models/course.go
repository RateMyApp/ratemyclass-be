package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	Units        float32
	Code         string
	SchoolID     *uint
	Professor    []*Professor `gorm:"many2many:professor_courses"`
	CourseRating []CourseRating
}
