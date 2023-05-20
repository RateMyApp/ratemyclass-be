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
	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
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
)

func beforeAll() {
	os.Setenv("GO_ENV", "testing")
	// config
	appConfig = config.InitAppConfig()

	// db client
	_, client := dao.NewPostgresClient(appConfig)
	postgresClient = client
	postgresClient.Init()
	testUser = models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	postgresClient.Db.Create(&testUser)

	// repositories
	userRepostory = repositories.NewUserRepository(postgresClient)
	userRepositoryMock = new(mocks.UserRepositoryMock)

	// utils
	timeUtil := utils.NewTimeUtil()
	timeUtilMock := new(mocks.TimeUtilMock)
	jwtUtil := utils.NewJwtUtil()

	// services
	jwtService := services.NewJwtService(appConfig, jwtUtil, timeUtil)
	jwtServiceMock := services.NewJwtService(appConfig, jwtUtil, timeUtilMock)
	authService = services.NewAuthServiceImpl(userRepostory, jwtService)
	authServiceMock = services.NewAuthServiceImpl(userRepositoryMock, jwtServiceMock)
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

func Test_RegisterUserCommand_ShouldReturnNoError_WhenUserDoesNotExist(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)
	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}

	err := authService.RegisterUser(command)

	assert.Nil(err)

	foundUser, err := userRepostory.FindUserByEmail(command.Email)

	assert.Nil(err)

	assert.Equal(command.Email, foundUser.Email)
	assert.Equal(command.Firstname, foundUser.Firstname)
	assert.Equal(command.Lastname, foundUser.Lastname)
	passwordEncryptErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(command.Password))
	assert.Nil(passwordEncryptErr)
}

func Test_RegisterUserCommand_ShouldReturnConflictAppErr_WhenUserAlreadyExists(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)

	command := services.RegisterCommand{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	err := authService.RegisterUser(command)

	assert.NotNil(err)
	assert.Equal(err.StatusCode, http.StatusConflict)
}

func Test_RegisterUserCommand_ShouldReturnInternalAppErr_WhenFindUserByEmailNotWorking(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)
	internalErr := exceptions.NewInternalServerError()
	mockCall := userRepositoryMock.On("FindUserByEmail", mock.Anything).Return(nil, &internalErr)

	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}
	err := authServiceMock.RegisterUser(command)
	// userRepositoryMock.AssertExpectations(t)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, err.StatusCode)
	mockCall.Unset()
}

func Test_RegisterUserCommand_ShouldReturnInternalAppErr_WhenSaveUserNotWorking(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)
	internalErr := exceptions.NewInternalServerError()
	userRepositoryMock.On("FindUserByEmail", mock.Anything).Return(nil, nil)
	userRepositoryMock.On("SaveUser", mock.Anything).Return(&internalErr)

	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}
	err := authServiceMock.RegisterUser(command)
	// userRepositoryMock.AssertExpectations(t)
	// userRepositoryMock.AssertExpectations(t)
	assert.NotNil(err)
	assert.Equal(http.StatusInternalServerError, err.StatusCode)
}
