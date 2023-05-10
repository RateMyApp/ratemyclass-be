package repositories

import (
	"errors"
	"log"

	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) (*models.User, *exceptions.AppError)
	SaveUser(models.User) *exceptions.AppError
}

type UserRepositoryPostgresImpl struct {
	postgresClient *dao.PostgresClient
}

func (ur *UserRepositoryPostgresImpl) FindUserByEmail(email string) (*models.User, *exceptions.AppError) {
	var userModel models.User
	result := ur.postgresClient.Db.First(&userModel, "email= ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Print(result.Error)
		appError := exceptions.NewInternalServerError()
		return nil, &appError
	}
	return &userModel, nil
}

func (ur *UserRepositoryPostgresImpl) SaveUser(user models.User) *exceptions.AppError {
	result := ur.postgresClient.Db.Create(&user)

	if result.Error != nil {
		log.Print(result.Error)
		ie := exceptions.NewInternalServerError()
		return &ie
	}

	return nil
}


func NewUserRepository(postgresClient *dao.PostgresClient) UserRepository {
	return &UserRepositoryPostgresImpl{postgresClient: postgresClient}
}
