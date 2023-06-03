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

type ProfessorRouteTestSuite struct {
	suite.Suite

	// db
	transaction *gorm.DB
	client      *dao.PostgresClient

	// repo
	professorRepo repositories.ProfessorRepository

	// service
	professorServ     services.ProfessorService
	professorServMock mocks.ProfessorServiceMock

	// router
	professorRouter     routers.Router
	professorRouterMock routers.Router
	ginRouter           *gin.Engine
	ginRouterMock       *gin.Engine
}

func (prts *ProfessorRouteTestSuite) SetupSuite() {
	// config
	os.Setenv("GO_ENV", "testing")

	// db
	_, postgresClient := dao.NewPostgresClient(config.InitAppConfig())
	prts.client = postgresClient
	prts.client.Init()

	// repo
	prts.professorRepo = repositories.NewProfessorRepository(prts.client)

	// service 
	prts.professorServ = services.NewProfessorService(prts.professorRepo)
	prts.professorServMock = mocks.ProfessorServiceMock{}

	// router
	prts.ginRouter = gin.Default()
	prts.ginRouterMock = gin.Default()
	prts.professorRouter = routers.NewProfessorRouter(prts.ginRouter, prts.professorServ)
	prts.professorRouterMock = routers.NewProfessorRouter(prts.ginRouterMock, &prts.professorServMock)

	// execute routers
	prts.professorRouter.ExecRoutes()
	prts.professorRouterMock.ExecRoutes()
}

func (prts *ProfessorRouteTestSuite) SetupTest() {
	prts.transaction = prts.client.Db.Begin()
}

func (prts *ProfessorRouteTestSuite) TearDownTest() {
	prts.transaction.Rollback()
}

func (prts *ProfessorRouteTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		prts.client.Db.Unscoped().Where(" 1 = 1").Delete(model)
	}

	ctx := context.Background()
	prts.client.Close(ctx)
}

func (prts *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenInvalidRequest() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "",
		"lastname":         "",
		"directoryListing": "",
		"email":            "",
		"department":       "",
	})
	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	prts.ginRouter.ServeHTTP(w, req)

	prts.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (prts *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus201_WhenGivenValidRequest() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "http://test.com",
		"email":            "test@gmail.com",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	prts.ginRouter.ServeHTTP(w, req)

	prts.Equal(http.StatusCreated, w.Result().StatusCode)
}

func (prts *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenAnInvalidEmail() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "http://test.com",
		"email":            "testgmail",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	prts.ginRouter.ServeHTTP(w, req)

	prts.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (prts *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenAnInvalidDirectoryListing() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "testcom",
		"email":            "test@gmail.com",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	prts.ginRouter.ServeHTTP(w, req)

	prts.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (prts *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturn500_WhenAnUnexpectedErrorOccursInServiceLayer() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "http://test.com",
		"email":            "test@gmail.com",
		"department":       "English",
	})


	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	internalErr:= exceptions.NewInternalServerError()
	prts.professorServMock.On("CreateProfessor", mock.Anything).Return(&internalErr)
	prts.ginRouterMock.ServeHTTP(w, req)

	prts.Equal(http.StatusInternalServerError, w.Result().StatusCode)
}
func TestProfessorRouteTestSuite(t *testing.T) {
	suite.Run(t, new(ProfessorRouteTestSuite))
}
