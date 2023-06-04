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

type CourseReposiotoryTestSuite struct {
	suite.Suite
	appConfig   config.AppConfig
	client      *dao.PostgresClient
	transaction *gorm.DB
	courseRepo  repositories.CourseRepository
}

func (crts *CourseReposiotoryTestSuite) SetupSuite() {
	// config
	os.Setenv("GO_ENV", "testing")
	crts.appConfig = config.InitAppConfig()

	// db
	_, crts.client = dao.NewPostgresClient(crts.appConfig)
	crts.client.Init()

	// repository
	crts.courseRepo = repositories.NewCoursRepository(crts.client)
}

func (crts *CourseReposiotoryTestSuite) SetupTest() {
	crts.transaction = crts.client.Db.Begin()
}

func (crts *CourseReposiotoryTestSuite) TearDownTest() {
	crts.transaction.Rollback()
}

func (crts *CourseReposiotoryTestSuite) TearDownSuite() {

	for _, model := range *models.GetModels() {
		crts.client.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	crts.client.Close(ctx)
}

func (crts *CourseReposiotoryTestSuite) Test_SaveCourse_ShouldReturnNoError_WhenGivenACourse() {
	model := models.Course{Name: "Some Course", Units: 1.0, Code: "ECE530"}
	err := crts.courseRepo.SaveCourse(model)
	crts.Nil(err)
}

func TestCourseRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CourseReposiotoryTestSuite))
}
