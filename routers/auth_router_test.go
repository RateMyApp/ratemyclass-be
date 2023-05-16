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
	"github.com/ratemyapp/mocks"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/ratemyapp/routers"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	appConfig          config.AppConfig
	postgresClient     *dao.PostgresClient
	transaction        *gorm.DB
	testUser           models.User
	userRepostory      repositories.UserRepository
	authService        services.AuthService
	userRepositoryMock *mocks.UserRepositoryMock
	authServiceMock    services.AuthService
	ginRouter          *gin.Engine
	authRouter         routers.Router
)

func beforeAll() {
	os.Setenv("GO_ENV", "testing")
	appConfig = config.InitAppConfig()
	_, client := dao.NewPostgresClient(appConfig)
	postgresClient = client
	postgresClient.Init()
	testUser = models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	postgresClient.Db.Create(&testUser)
	userRepostory = repositories.NewUserRepository(postgresClient)
	authService = services.NewAuthServiceImpl(userRepostory)
	gin.SetMode(gin.ReleaseMode)
	ginRouter = gin.Default()
	ginRouter.Use(gin.Recovery())
	authRouter = routers.NewAuthRouter(ginRouter, authService)
	authRouter.ExecRoutes()
}

func beforeEach() {
	transaction = postgresClient.Db.Begin()
}

func afterEach() {
	transaction.Rollback()
}

func afterAll() {
	for _, model := range *models.GetModels() {
		postgresClient.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	postgresClient.Close(ctx)
}

func TestMain(m *testing.M) {
	beforeAll()
	defer afterAll()
	m.Run()
}

func Test_RegisterRoute_ShouldReturnValidationErr_WhenGivenInvalidRequest(t *testing.T) {
	beforeEach()
	defer afterEach()

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"email":     "",
		"password":  "",
		"firstname": "",
		"lastname":  "",
	})
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(body))

	ginRouter.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, w.Result().StatusCode)
}

func Test_RegisterRoute_ShouldReturnConflictError_WhenGivenExistingEmail(t *testing.T) {
	beforeEach()
	defer afterEach()

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"email":     "test@gmail.com",
		"password":  "hello123",
		"firstname": "TestFirstname",
		"lastname":  "TestLastname",
	})
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(body))

	ginRouter.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(http.StatusConflict, w.Result().StatusCode)
}

func Test_RegisterRoute_ShouldReturnCreatedSuccess_WhenGivenValidRequest(t *testing.T) {
	beforeEach()
	defer afterEach()

	w := httptest.NewRecorder()
	body, _ := json.Marshal(map[string]string{
		"email":     "newtest@gmail.com",
		"password":  "hello123",
		"firstname": "TestFirstname",
		"lastname":  "TestLastname",
	})
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewReader(body))

	ginRouter.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(http.StatusCreated, w.Result().StatusCode)
}
