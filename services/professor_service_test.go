package services_test

import (
	"context"
	"os"
	"testing"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/mocks"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ProfessorServiceTestSuite struct {
	suite.Suite
	appConfig         config.AppConfig
	client            *dao.PostgresClient
	transaction       *gorm.DB
	professorRepo     repositories.ProfessorRepository
	professorServ     services.ProfessorService
	professorRepoMock *mocks.ProfessorRepository
	professorServMock services.ProfessorService
}

func (psts *ProfessorServiceTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	appConfig = config.InitAppConfig()
	_, psts.client = dao.NewPostgresClient(appConfig)
	psts.client.Init()
	// normal setup
	psts.professorRepo = repositories.NewProfessorRepository(psts.client)
	psts.professorServ = services.NewProfessorService(psts.professorRepo)

	// mock setup
	psts.professorRepoMock = new(mocks.ProfessorRepository)
	psts.professorServMock = services.NewProfessorService(psts.professorRepoMock)
}

func (psts *ProfessorServiceTestSuite) SetupTest() {
	psts.transaction = psts.client.Db.Begin()
}

func (psts *ProfessorServiceTestSuite) TearDownTest() {
	psts.transaction.Rollback()
}

func (psts *ProfessorServiceTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		psts.client.Db.Unscoped().Where("1 = 1 ").Delete(model)
	}
	ctx := context.Background()
	psts.client.Close(ctx)
}

func (psts *ProfessorServiceTestSuite) Test_CreateProfessor_ShouldNotReturnError_WhenGivenValidCommand() {
	command := services.CreateProfessorCommand{Lastname: "Prof", Firstname: "Test", Department: "English", DirectoryListing: "https://test.com", Email: "test@gmail.com"}
	result := psts.professorServ.CreateProfessor(command)
	psts.Nil(result)
}

func (psts *ProfessorServiceTestSuite) Test_CreateProfessor_ShouldReturnError_WhenGivenValidCommand() {
	command := services.CreateProfessorCommand{Lastname: "Prof", Firstname: "Test", Department: "English", DirectoryListing: "https://test.com", Email: "test@gmail.com"}
	internalErr := exceptions.NewInternalServerError()
	psts.professorRepoMock.On("SaveProfessor", mock.Anything).Return(&internalErr)
	result := psts.professorServMock.CreateProfessor(command)
	psts.NotNil(result)
	psts.Equal(internalErr.Message, result.Message)
	psts.Equal(internalErr.TimeStamp, result.TimeStamp)
	psts.Equal(internalErr.StatusCode, result.StatusCode)
}

func TestProfessorServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProfessorServiceTestSuite))
}
