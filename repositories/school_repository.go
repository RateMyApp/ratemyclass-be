package repositories

import (
	"log"
	"strings"

	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
)

type SchoolRepository interface {
	SearchSchoolsByName(string) ([]models.School, *exceptions.AppError)
	SaveSchool(models.School) *exceptions.AppError
}

type schoolRepositoryImpl struct {
	client *dao.PostgresClient
}

func (sr *schoolRepositoryImpl) SaveSchool(school models.School) *exceptions.AppError {
	result := sr.client.Db.Create(&school)
	if result.Error != nil {
		log.Println(result.Error)
		ie := exceptions.NewInternalServerError()
		return &ie
	}

	return nil
}

func (sr *schoolRepositoryImpl) SearchSchoolsByName(search string) ([]models.School, *exceptions.AppError) {
	var schools []models.School
	search = "%" + strings.ToLower(search) + "%"
	result := sr.client.Db.Where("LOWER(name) LIKE ? OR LOWER(city) LIKE ? OR LOWER(province_or_state) LIKE ? OR LOWER(country) LIKE ?", search, search, search, search).Find(&schools)

	if result.Error != nil {

		log.Println(result.Error)
		ie := exceptions.NewInternalServerError()
		return schools, &ie
	}

	return schools, nil
}

func NewSchoolRepository(client *dao.PostgresClient) SchoolRepository {
	return &schoolRepositoryImpl{client: client}
}
