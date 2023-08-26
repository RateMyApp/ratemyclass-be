package repositories

import (
	"log"

	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
)

type ProfessorRepository interface {
	SaveProfessor(models.Professor) *exceptions.AppError
}

type professorRepositoryImpl struct {
	client *dao.PostgresClient
}

func (pr *professorRepositoryImpl) SaveProfessor(professor models.Professor) *exceptions.AppError {
	result := pr.client.Db.Create(&professor)
	if result.Error != nil {
		log.Print(result.Error)
		ie := exceptions.NewInternalServerError()
		return &ie
	}
	return nil
}

func NewProfessorRepository(client *dao.PostgresClient) ProfessorRepository {
	return &professorRepositoryImpl{client: client}
}
