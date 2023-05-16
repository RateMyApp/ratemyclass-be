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

type ProfessorRouteTestSuite struct {
	suite.Suite
	professorRepo   repositories.ProfessorRepository
	professorServ   services.ProfessorService
	ginRouter       *gin.Engine
	transaction     *gorm.DB
	client          *dao.PostgresClient
	professorRouter routers.Router
}

func (self *ProfessorRouteTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	self.ginRouter = gin.Default()
	_, postgresClient := dao.NewPostgresClient(config.InitAppConfig())
	self.client = postgresClient
	self.client.Init()
	self.professorRepo = repositories.NewProfessorRepository(self.client)
	self.professorServ = services.NewProfessorService(self.professorRepo)
	self.professorRouter = routers.NewProfessorRouter(self.ginRouter, self.professorServ)
	self.professorRouter.ExecRoutes()
}

func (self *ProfessorRouteTestSuite) SetupTest() {
	self.transaction = self.client.Db.Begin()
}

func (self *ProfessorRouteTestSuite) TearDownTest() {
	self.transaction.Rollback()
}

func (self *ProfessorRouteTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		self.client.Db.Unscoped().Where(" 1 = 1").Delete(model)
	}

	ctx := context.Background()
	self.client.Close(ctx)
}

func (self *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenInvalidRequest() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "",
		"lastname":         "",
		"directoryListing": "",
		"email":            "",
		"department":       "",
	})
	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	self.ginRouter.ServeHTTP(w, req)

	self.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (self *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus201_WhenGivenValidRequest() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "http://test.com",
		"email":            "test@gmail.com",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	self.ginRouter.ServeHTTP(w, req)

	self.Equal(http.StatusCreated, w.Result().StatusCode)
}

func (self *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenAnInvalidEmail() {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "http://test.com",
		"email":            "testgmail",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	self.ginRouter.ServeHTTP(w, req)

	self.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func (self *ProfessorRouteTestSuite) Test_CreateProfRoute_ShouldReturnStatus400_WhenGivenAnInvalidDirectoryListing() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"firstname":        "Firstname",
		"lastname":         "Lastname",
		"directoryListing": "testcom",
		"email":            "test@gmail.com",
		"department":       "English",
	})

	req, _ := http.NewRequest("POST", "/api/v1/prof", bytes.NewReader(body))

	self.ginRouter.ServeHTTP(w, req)

	self.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func TestProfessorRouteTestSuite(t *testing.T) {
	suite.Run(t, new(ProfessorRouteTestSuite))
}
