package routers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/ratemyapp/routers"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CourseRatingRouteTestSuite struct {
	suite.Suite
	courseratingrepo   repositories.CourseRatingRepository
	courseratingserv   services.CourseRatingService
	ginRouter          *gin.Engine
	transaction        *gorm.DB
	client             *dao.PostgresClient
	courseratingrouter routers.Router
}

func (self *CourseRatingRouteTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	self.ginRouter = gin.Default()
	_, postgresClient := dao.NewPostgresClient(config.InitAppConfig())
	self.client = postgresClient
	self.client.Init()
	self.courseratingrepo = repositories.NewCourseRatingRepository(self.client)
	self.courseratingserv = services.NewCourseRatingService(self.courseratingrepo)
	self.courseratingrouter = routers.NewCourseRatingRouter(self.ginRouter, self.courseratingserv)
	self.courseratingrouter.ExecRoutes()
}

func (self *CourseRatingRouteTestSuite) SetupTest() {
	self.transaction = self.client.Db.Begin()
}

func (self *CourseRatingRouteTestSuite) TearDownTest() {
	self.transaction.Rollback()
}

func (self *CourseRatingRouteTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		self.client.Db.Unscoped().Where(" 1 = 1").Delete(model)
	}

	ctx := context.Background()
	self.client.Close(ctx)
}

func (self *CourseRatingRouteTestSuite) Test_CreateCourseRatingRoute_ShouldReturnStatus400_WhenGivenInvalidRequest() {
	w := httptest.NewRecorder()
	// create json
	body, _ := json.Marshal(map[string]interface{}{
		"professorid":      "",
		"experiencerating": "",
		"difficultyrating": "",
		"review":           "",
		"courseid":         "",
		"userid":           "",
		"isanonymous":      "",
	})
	// create post request
	req, _ := http.NewRequest("POST", "/api/v1/courserating", bytes.NewReader(body))
	// call endpoint
	self.ginRouter.ServeHTTP(w, req)
	// check if bad request
	self.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (self *CourseRatingRouteTestSuite) Test_CreateCourseRatingRoute_ShouldReturnStatus201_WhenGivenValidRequest() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"professorid":      "2",
		"experiencerating": 5,
		"difficultyrating": 4,
		"review":           "GOOD PROFESSOR",
		"courseid":         3,
		"userid":           2,
		"isanonymous":      true,
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	self.ginRouter.ServeHTTP(w, req)

	self.Equal(http.StatusCreated, w.Result().StatusCode)
}

func TestCourseRatingRouteTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRatingRouteTestSuite))
}
