package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
)

type CourseService interface {
	// Saves Course within a Database
	CreateCourse(CreateCourseCommand) *exceptions.AppError
}

type courseServiceImpl struct {
	courseRepo repositories.CourseRepository
}


// Creates a course using the CreateCourseCommand
func (csi *courseServiceImpl) CreateCourse(command CreateCourseCommand) *exceptions.AppError {
	newCourse := models.Course{Name: command.Name, Units: command.Units, Code: command.Code}
	err := csi.courseRepo.SaveCourse(newCourse)
	if err != nil {
		return err
	}
	return nil
}


// Initialises a new Course Service Impl
func NewCourseService(courseRepo repositories.CourseRepository) CourseService {
	return &courseServiceImpl{courseRepo: courseRepo}
}
