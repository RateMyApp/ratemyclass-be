package models

import "gorm.io/gorm"

type CourseRating struct {
	gorm.Model
	ProfessorID      uint
	ExperienceRating float32
	DifficultyRating float32
	Review           string
	CourseID         uint
	// StudentID uint
}
