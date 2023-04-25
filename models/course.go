package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	Units        float32
	Professor    []Professor `gorm:"many2many:professor_courses"`
	User         []*User `gorm:"many2many:course_students"`
	CourseRating []CourseRating
}
