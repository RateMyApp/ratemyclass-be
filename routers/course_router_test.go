package routers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/mocks"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/ratemyapp/routers"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CourseRouterTestSuite struct {
	suite.Suite
	// config
	config config.AppConfig

	// dao
	transaction *gorm.DB
	dbclient    *dao.PostgresClient

	// repository
	courseRepo repositories.CourseRepository

	// services
	courseService services.CourseService
	courseServiceMock mocks.CourseServiceMock

	// router
	courseRouter routers.Router
	courseRouterMock routers.Router
	ginRouter    *gin.Engine
	ginRouterMock *gin.Engine
}

func (crts *CourseRouterTestSuite) SetupSuite() {
	crts.config = config.InitAppConfig()

	// db 
	_, crts.dbclient = dao.NewPostgresClient(crts.config)
	crts.dbclient.Init()

	// repo
	crts.courseRepo = repositories.NewCoursRepository(crts.dbclient)

	// services
	crts.courseService = services.NewCourseService(crts.courseRepo)
	crts.courseServiceMock = mocks.CourseServiceMock{}

	// routers 
	crts.ginRouter = gin.Default()
	crts.ginRouterMock = gin.Default()
	crts.courseRouterMock = routers.NewCourseRouter(crts.ginRouterMock, &crts.courseServiceMock )
	crts.courseRouterMock.ExecRoutes()
	crts.courseRouter = routers.NewCourseRouter(crts.ginRouter, crts.courseService)
	crts.courseRouter.ExecRoutes()
}

func (crts *CourseRouterTestSuite) SetupTest() {
	crts.transaction = crts.dbclient.Db.Begin()
}

func (crts *CourseRouterTestSuite) TearDownTest() {
	crts.transaction.Rollback()
}

func (crts *CourseRouterTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		crts.dbclient.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	context := context.Background()
	crts.dbclient.Close(context)
}

func (crts *CourseRouterTestSuite) Test_CreateCourseRoute_ShouldReturn200_WhenGivenAValidReq() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"code":  "ECE657",
		"name":  "Artificial Intelligence",
		"units": 0.42,
	})
	req, _ := http.NewRequest("POST", "/api/v1/course", bytes.NewReader(body))

	crts.ginRouter.ServeHTTP(w, req)
	testBody, _ := io.ReadAll(w.Body)
	log.Println(string( testBody ))
	crts.Equal(http.StatusCreated, w.Result().StatusCode)
}


func (crts *CourseRouterTestSuite) Test_CreateCourseRoute_ShouldReturn400_WhenGivenAInvalidReq(){
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"code":  "",
		"name":  "",
		"units": 0.42,
	})
	req, _ := http.NewRequest("POST", "/api/v1/course", bytes.NewReader(body))

	crts.ginRouter.ServeHTTP(w, req)
	crts.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (crts *CourseRouterTestSuite) Test_CreateCourseRoute_ShouldReturn400_WhenGivenUnitGreaterThan2DecimalPlaces(){
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"code":  "ECE657",
		"name":  "Artificial Intelligence",
		"units": 0.423243423,
	})
	req, _ := http.NewRequest("POST", "/api/v1/course", bytes.NewReader(body))

	crts.ginRouter.ServeHTTP(w, req)
	crts.Equal(http.StatusBadRequest, w.Result().StatusCode)
}
func (crts *CourseRouterTestSuite) Test_CreateCourseRoute_ShouldReturn500_WhenAnUnexpectedErrorOccursInServiceLayer(){
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"code":  "ECE657",
		"name":  "Artificial Intelligence",
		"units": 0.45,
	})
	req, _ := http.NewRequest("POST", "/api/v1/course", bytes.NewReader(body))

	internalErr:= exceptions.NewInternalServerError()
	crts.courseServiceMock.On("CreateCourse", mock.Anything).Return(&internalErr)
	crts.ginRouterMock.ServeHTTP(w, req)
	crts.Equal(http.StatusInternalServerError, w.Result().StatusCode)
}

func TestCourseRouterTestSuite(t *testing.T) {
	suite.Run(t, new(CourseRouterTestSuite))
}
