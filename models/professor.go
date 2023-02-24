package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Name         string
	Email        string
	Department   string
	Courses      []*Course `gorm:"many2many:professor_courses"`
	CourseRating []CourseRating
	ProgramID    uint
}
