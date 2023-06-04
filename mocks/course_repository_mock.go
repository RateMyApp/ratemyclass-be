package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/stretchr/testify/mock"
)

type CourseRepositoryMock struct {
	mock.Mock
}

func (crm *CourseRepositoryMock) SaveCourse(course models.Course) *exceptions.AppError {
	args := crm.Called(course)

	var arg0 *exceptions.AppError = nil

	if val, ok := args.Get(0).(*exceptions.AppError); ok {
		arg0 = val
	}

	return arg0
}

var _ repositories.CourseRepository = (*CourseRepositoryMock)(nil)
