package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
)

type CourseRatingServiceMock struct {
	mock.Mock
}

func (crsm *CourseRatingServiceMock) CreateCourseRating(command services.CreateCourseRatingCommand) *exceptions.AppError {
	args := crsm.Called(command)
	
	var return0 *exceptions.AppError = nil
	
	if val, ok := args.Get(0).(*exceptions.AppError) ; ok{
			return0 = val
	}
	return return0
}
//type checking
var _ services.CourseRatingService = (*CourseRatingServiceMock)(nil)