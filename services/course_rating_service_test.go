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

type CourseRatingServiceTestSuite struct {
	suite.Suite
	appConfig         config.AppConfig
	client            *dao.PostgresClient
	transaction       *gorm.DB
	courseratingrepo     repositories.CourseRatingRepository
	courseratingserv     services.CourseRatingService
	coueseratingRepoMock *mocks.CourseRatingRepository
	courseratingServMock services.CourseRatingService
}

func (crst *CourseRatingServiceTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	crst.appConfig = config.InitAppConfig()
	_, crst.client = dao.NewPostgresClient(crst.appConfig)
	crst.client.Init()
	crst.courseratingrepo = repositories.NewCourseRatingRepository(crst.client)
	crst.courseratingserv = services.NewCourseRatingService(crst.courseratingrepo)
	crst.coueseratingRepoMock = new(mocks.CourseRatingRepository)
	crst.courseratingServMock = services.NewCourseRatingService(crst.coueseratingRepoMock)
 }

func (crst *CourseRatingServiceTestSuite) SetupTest() {
	crst.transaction = crst.client.Db.Begin()
}

func (crst *CourseRatingServiceTestSuite) TearDownTest() {
	crst.transaction.Rollback()
}

func (crst *CourseRatingServiceTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		crst.client.Db.Unscoped().Where("1 = 1 ").Delete(model)
	}
	ctx := context.Background()
	crst.client.Close(ctx)
}

func (crst *CourseRatingServiceTestSuite) Test_CreateCourseRating_ShouldNotReturnError_WhenGivenValidCommand() {
	command := services.CreateCourseRatingCommand{ProfessorID: 1, ExperienceRating: 3, DifficultyRating: 4,Review: "GOOD ONE", CourseID: 3, UserID: 1, IsAnonymous: false}
	result := crst.courseratingserv.CreateCourseRating(command)
	crst.Nil(result)
}

func (crst *CourseRatingServiceTestSuite) Test_CreateCourseRating_ShouldReturnError_WhenGivenValidCommand() {
	command := services.CreateCourseRatingCommand{ProfessorID: 1, ExperienceRating: 3, DifficultyRating: 4,Review: "GOOD ONE", CourseID: 3, UserID: 1, IsAnonymous: false}
	internalErr := exceptions.NewInternalServerError()
	crst.coueseratingRepoMock.On("SaveCourseRating", mock.Anything).Return(&internalErr)
	result := crst.courseratingServMock.CreateCourseRating(command)
	crst.NotNil(result)
	crst.Equal(internalErr.Message, result.Message)
	crst.Equal(internalErr.TimeStamp, result.TimeStamp)
	crst.Equal(internalErr.StatusCode, result.StatusCode)
}

func TestCourseRatingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRatingServiceTestSuite))
}
