package models

import "gorm.io/gorm"

type Program struct {
	gorm.Model
	CourseCode string
	CourseName string
	SchoolID   uint
	Professor  []Professor
}
