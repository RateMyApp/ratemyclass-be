package models

type CourseRating struct {
	gorm.Model
	Professor        Professor
	ExperienceRating float32
	DifficultyRating float32
	Review           string
	Course           Course
	// Student Student
}
