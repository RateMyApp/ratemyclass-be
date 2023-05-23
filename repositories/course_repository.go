package repositories

import (
	"log"

	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
)

type CourseRepository interface {
	SaveCourse(models.Course) *exceptions.AppError
}

type courseRepositoryImpl struct {
	client *dao.PostgresClient
}

func (self *courseRepositoryImpl) SaveCourse(course models.Course) *exceptions.AppError {
	result := self.client.Db.Create(&course)
	if result.Error != nil {
		log.Print(result.Error)
		ie := exceptions.NewInternalServerError()
		return &ie
	}

	return nil
}

func NewCoursRepository(client *dao.PostgresClient) CourseRepository {
	return &courseRepositoryImpl{client: client}
}
