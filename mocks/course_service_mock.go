package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
)

type CourseServiceMock struct {
	mock.Mock
}

func (csm *CourseServiceMock) CreateCourse(command services.CreateCourseCommand) *exceptions.AppError{
	args := csm.Called(command)

	var args0 *exceptions.AppError = nil
	
	if val,ok := args.Get(0).(*exceptions.AppError); ok {
		args0 = val
	}
	return args0
}


var _ services.CourseService = (*CourseServiceMock)(nil)