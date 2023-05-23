package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
)

type CourseService interface {
	// Saves Course within Database
	CreateCourse(CreateCourseCommand) *exceptions.AppError
}

type courseServiceImpl struct {
	courseRepo repositories.CourseRepository
}

func (self *courseServiceImpl) CreateCourse(command CreateCourseCommand) *exceptions.AppError {
	newCourse := models.Course{Name: command.Name, Units: float32(command.Units), Code: command.Code}
	err := self.courseRepo.SaveCourse(newCourse)
	if err != nil {
		return err
	}
	return nil
}

func NewCourseService() CourseService {
	return &courseServiceImpl{}
}
