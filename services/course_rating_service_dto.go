package services

type CreateCourseRatingCommand struct {
	ProfessorID      uint    
	ExperienceRating float32 
	DifficultyRating float32 
	Review           string  
	CourseID         uint   
	UserID           uint    
	IsAnonymous      bool   
}