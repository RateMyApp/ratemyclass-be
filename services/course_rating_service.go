package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
)

type CourseRatingService interface{
	CreateCourseRating(CreateCourseRatingCommand) *exceptions.AppError
}

type CourseRatingServiceImpl struct{
	repo repositories.CourseRatingRepository
}
// create course rating
func (crs * CourseRatingServiceImpl) CreateCourseRating(command CreateCourseRatingCommand) * exceptions.AppError{
	course_rating:=models.CourseRating{ProfessorID: command.ProfessorID, ExperienceRating: command.ExperienceRating, DifficultyRating: command.DifficultyRating,Review: command.Review, CourseID: command.CourseID,UserID: command.UserID, IsAnonymous: command.IsAnonymous}
	return crs.repo.SaveCourseRating(course_rating)
}

// construtor method
func NewCourseRatingService(repo repositories.CourseRatingRepository) CourseRatingService {
	return &CourseRatingServiceImpl{repo: repo}
}
