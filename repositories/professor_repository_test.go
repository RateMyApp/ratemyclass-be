package repositories_test

import (
	"context"
	"os"
	"testing"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ProfessorRepositoryTestSuite struct {
	suite.Suite
	appConfig     config.AppConfig
	client        *dao.PostgresClient
	transaction   *gorm.DB
	professorRepo repositories.ProfessorRepository
}

func (prts *ProfessorRepositoryTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	prts.appConfig = config.InitAppConfig()
	_, postgresClient := dao.NewPostgresClient(prts.appConfig)
	prts.client = postgresClient
	prts.client.Init()
	prts.professorRepo = repositories.NewProfessorRepository(postgresClient)
}

func (prts *ProfessorRepositoryTestSuite) SetupTest() {
	prts.transaction = prts.client.Db.Begin()
}

func (prts *ProfessorRepositoryTestSuite) TearDownTest() {
	prts.transaction.Rollback()
}

func (prts *ProfessorRepositoryTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		prts.client.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	prts.client.Close(ctx)
}

func (prts *ProfessorRepositoryTestSuite) Test_SaveProfessor_ShouldReturnNoError_WhenGivenAValidUser() {
	newProf := models.Professor{Email: "prof@gmail.com", Firstname: "Test", Lastname: "Prof", DirectoryListing: "https://prof.com", Department: "English"}
	result := prts.professorRepo.SaveProfessor(newProf)
	prts.Nil(result)
}

func TestProfessorRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProfessorRepositoryTestSuite))
}
