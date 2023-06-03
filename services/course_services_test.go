package services_test

import (
	"context"
	"net/http"
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

type CourseServicesTestSuite struct {
	suite.Suite
	config            config.AppConfig
	transaction       *gorm.DB
	dbclient          *dao.PostgresClient
	courseRepo        repositories.CourseRepository
	courseRepoMock    mocks.CourseRepositoryMock
	courseService     services.CourseService
	courseServiceMock services.CourseService
}

func (csts *CourseServicesTestSuite) SetupSuite() {
	// config
	os.Setenv("GO_ENV", "testing")
	csts.config = config.InitAppConfig()

	// db
	_, csts.dbclient = dao.NewPostgresClient(csts.config)
	csts.dbclient.Init()

	// repository
	csts.courseRepo = repositories.NewCoursRepository(csts.dbclient)
	csts.courseRepoMock = mocks.CourseRepositoryMock{}

	// services
	csts.courseService = services.NewCourseService(csts.courseRepo)
	csts.courseServiceMock = services.NewCourseService(&csts.courseRepoMock)
}

func (csts *CourseServicesTestSuite) SetupTest() {
	csts.transaction = csts.dbclient.Db.Begin()
}

func (csts *CourseServicesTestSuite) TearDownTest() {
	csts.transaction.Rollback()
}

func (csts *CourseServicesTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		csts.dbclient.Db.Unscoped().Where("1 = 1 ").Delete(model)
	}
	ctx := context.Background()
	csts.dbclient.Close(ctx)
}

func (csts *CourseServicesTestSuite) Test_CreateCourse_ShouldReturnNoError_WhenGivenCommand() {
	command := services.CreateCourseCommand{Code: "ECE650", Name: "New Course", Units: 0.3}

	err := csts.courseService.CreateCourse(command)

	csts.Nil(err)
}

func (csts *CourseServicesTestSuite) Test_CreateCourse_ShouldReturnErr_WhenGivenCommand() {
	command := services.CreateCourseCommand{Code: "ECE650", Name: "New Course", Units: 0.3}
	internalServerErr := exceptions.NewInternalServerError()
	csts.courseRepoMock.On("SaveCourse", mock.Anything).Return(&internalServerErr)
	err := csts.courseServiceMock.CreateCourse(command)
	csts.NotNil(err)
	csts.Equal(err.StatusCode, http.StatusInternalServerError)
}

func TestCourseServicesTestSuite(t *testing.T) {
	suite.Run(t, new(CourseServicesTestSuite))
}
