package repositories

import (
	"log"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
)

type CourseRatingRepository interface {
	SaveCourseRating(models.CourseRating) *exceptions.AppError
}

type CourseRatingRepositoryImpl struct {
	client *dao.PostgresClient
}
// save to db
func (cr *CourseRatingRepositoryImpl) SaveCourseRating(courserating models.CourseRating) *exceptions.AppError {
	result := cr.client.Db.Create(&courserating)
	if result.Error != nil {
		log.Print(result.Error)
		ie := exceptions.NewInternalServerError()
		return &ie
	}
	return nil
}
// constructor method
func NewCourseRatingRepository(client *dao.PostgresClient) CourseRatingRepository {
	return &CourseRatingRepositoryImpl{client: client}
}
