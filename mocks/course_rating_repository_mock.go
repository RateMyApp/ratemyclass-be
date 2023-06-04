package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/stretchr/testify/mock"
)

type CourseRatingRepository struct {
	mock.Mock
}

func (cr *CourseRatingRepository) SaveCourseRating(courserating models.CourseRating) *exceptions.AppError {
	args := cr.Called(courserating)

	var arg0 *exceptions.AppError

	if val, ok := args.Get(0).(*exceptions.AppError); ok {
		arg0 = val
	} else {
		arg0 = nil
	}

	return arg0
}
