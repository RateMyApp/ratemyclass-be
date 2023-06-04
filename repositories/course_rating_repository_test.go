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

type CourseRatingRepositoryTestSuite struct {
	suite.Suite
	appConfig     config.AppConfig
	client        *dao.PostgresClient
	transaction   *gorm.DB
	courseratingrepo repositories.CourseRatingRepository
}

func (prts *CourseRatingRepositoryTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	prts.appConfig = config.InitAppConfig()
	_, postgresClient := dao.NewPostgresClient(prts.appConfig)
	prts.client = postgresClient
	prts.client.Init()
	prts.courseratingrepo = repositories.NewCourseRatingRepository(postgresClient)
}

func (prts *CourseRatingRepositoryTestSuite) SetupTest() {
	prts.transaction = prts.client.Db.Begin()
}

func (prts *CourseRatingRepositoryTestSuite) TearDownTest() {
	prts.transaction.Rollback()
}

func (prts *CourseRatingRepositoryTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		prts.client.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	prts.client.Close(ctx)
}

func (prts *CourseRatingRepositoryTestSuite) Test_SaveCourse_ShouldReturnNoError_WhenGivenAValidCourse() {
	newCourse:=models.CourseRating{ProfessorID: 5, ExperienceRating: 4, DifficultyRating: 3,Review: "good course , very technical", CourseID: 1,UserID: 99, IsAnonymous: true}
	result := prts.courseratingrepo.SaveCourseRating(newCourse)
	prts.Nil(result)
}

func TestCourseRatingRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRatingRepositoryTestSuite))
}
