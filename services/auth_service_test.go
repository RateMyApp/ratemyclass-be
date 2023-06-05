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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthServiceTestSuite struct {
	suite.Suite
	appconfig          config.AppConfig
	postgresClient     *dao.PostgresClient
	transaction        *gorm.DB
	testUser           models.User
	userRepostory      repositories.UserRepository
	authService        services.AuthService
	userRepositoryMock *mocks.UserRepositoryMock
	authServiceMock1   services.AuthService
	authServiceMock2   services.AuthService
	jwtService         services.JwtService
	jwtServiceMock     *mocks.JwtServiceMock
}

func (asts *AuthServiceTestSuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	// config
	asts.appconfig = config.InitAppConfig()

	// db client
	_, client := dao.NewPostgresClient(asts.appconfig)
	asts.postgresClient = client
	asts.postgresClient.Init()
	asts.testUser = models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	asts.postgresClient.Db.Create(&asts.testUser)

	// repositories
	asts.userRepostory = repositories.NewUserRepository(asts.postgresClient)
	asts.userRepositoryMock = new(mocks.UserRepositoryMock)

	// utils
	timeUtil := utils.NewTimeUtil()
	jwtUtil := utils.NewJwtUtil()

	// services
	asts.jwtService = services.NewJwtService(asts.appconfig, jwtUtil, timeUtil)
	asts.jwtServiceMock = &mocks.JwtServiceMock{}
	asts.authService = services.NewAuthServiceImpl(asts.userRepostory, asts.jwtService)
	asts.authServiceMock1 = services.NewAuthServiceImpl(asts.userRepositoryMock, asts.jwtServiceMock)
	asts.authServiceMock2 = services.NewAuthServiceImpl(asts.userRepostory, asts.jwtServiceMock)

}
func (asts *AuthServiceTestSuite) SetupTest() {

	asts.transaction = asts.postgresClient.Db.Begin()
}
func (asts *AuthServiceTestSuite) TearDownTest() {
	asts.transaction.Rollback()
}
func (asts *AuthServiceTestSuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		asts.postgresClient.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	asts.postgresClient.Close(ctx)
}

func (asts *AuthServiceTestSuite) Test_RegisterUserCommand_ShouldReturnNoError_WhenUserDoesNotExist() {
	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}

	err := asts.authService.RegisterUser(command)

	asts.Nil(err)

	foundUser, err := asts.userRepostory.FindUserByEmail(command.Email)

	asts.Nil(err)

	asts.Equal(command.Email, foundUser.Email)
	asts.Equal(command.Firstname, foundUser.Firstname)
	asts.Equal(command.Lastname, foundUser.Lastname)
	passwordEncryptErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(command.Password))
	asts.Nil(passwordEncryptErr)
}

func (asts *AuthServiceTestSuite) Test_RegisterUserCommand_ShouldReturnConflictAppErr_WhenUserAlreadyExists() {

	command := services.RegisterCommand{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	err := asts.authService.RegisterUser(command)

	asts.NotNil(err)
	asts.Equal(err.StatusCode, http.StatusConflict)
}

func (asts *AuthServiceTestSuite) Test_RegisterUserCommand_ShouldReturnInternalAppErr_WhenFindUserByEmailNotWorking() {
	internalErr := exceptions.NewInternalServerError()
	mockCall := asts.userRepositoryMock.On("FindUserByEmail", mock.Anything).Return(nil, &internalErr)

	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}
	err := asts.authServiceMock1.RegisterUser(command)
	// userRepositoryMock.AssertExpectations(t)
	asts.NotNil(err)
	asts.Equal(http.StatusInternalServerError, err.StatusCode)
	mockCall.Unset()
}

func (asts *AuthServiceTestSuite) Test_RegisterUserCommand_ShouldReturnInternalAppErr_WhenSaveUserNotWorking() {
	internalErr := exceptions.NewInternalServerError()
	asts.userRepositoryMock.On("FindUserByEmail", mock.Anything).Return(nil, nil)
	asts.userRepositoryMock.On("SaveUser", mock.Anything).Return(&internalErr)

	command := services.RegisterCommand{Email: "test2@gmail.com", Password: "hello123", Firstname: "TestFirstname", Lastname: "TestLastname"}
	err := asts.authServiceMock1.RegisterUser(command)
	// userRepositoryMock.AssertExpectations(t)
	// userRepositoryMock.AssertExpectations(t)
	asts.NotNil(err)
	asts.Equal(http.StatusInternalServerError, err.StatusCode)
}

func (asts *AuthServiceTestSuite) Test_LoginUser_ShouldReturnUnAuthorizedException_WhenUserEmailIsInValid() {
	registerCommand := services.RegisterCommand{Email: "test1@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "password"}
	asts.authService.RegisterUser(registerCommand)

	loginCommand := services.LoginCommand{Email: "test2@gmail.com", Password: "password"}
	err, userInfo := asts.authService.LoginUser(loginCommand)
	asts.Nil(userInfo)
	asts.NotNil(err)
	asts.Equal(http.StatusUnauthorized, err.StatusCode)
}

func (asts *AuthServiceTestSuite) Test_LoginUser_ShouldReturnUnAuthorizedException_WhenPasswordisInvalid() {
	registerCommand := services.RegisterCommand{Email: "test3@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "password"}
	asts.authService.RegisterUser(registerCommand)

	loginCommand := services.LoginCommand{Email: "test3@gmail.com", Password: "incorrectpassword"}
	err, userInfo := asts.authService.LoginUser(loginCommand)
	asts.Nil(userInfo)
	asts.NotNil(err)
	asts.Equal(http.StatusUnauthorized, err.StatusCode)
}

func (asts *AuthServiceTestSuite) Test_LoginUser_ShouldReturnInternalServerException_WhenGeneratingATokenFails() {
	registerCommand := services.RegisterCommand{Email: "test1@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "password"}
	asts.authService.RegisterUser(registerCommand)
	internalError := exceptions.NewInternalServerError()
	asts.jwtServiceMock.On("GenerateAccessToken", mock.Anything).Return("", &internalError)
	loginCommand := services.LoginCommand{Email: "test1@gmail.com", Password: "password"}
	err, userInfo := asts.authServiceMock2.LoginUser(loginCommand)
	asts.Nil(userInfo)
	if asts.NotNil(err) {
		asts.Equal(http.StatusInternalServerError, err.StatusCode)
	}
}

func (asts *AuthServiceTestSuite) Test_LoginUser_ShouldReturnUserInfo_WhenEmailAndPasswordMatch() {
	registerCommand := services.RegisterCommand{Email: "test1@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "password"}
	asts.authService.RegisterUser(registerCommand)

	loginCommand := services.LoginCommand{Email: "test1@gmail.com", Password: "password"}
	err, userInfo := asts.authService.LoginUser(loginCommand)
	asts.Nil(err)
	asts.NotNil(userInfo)
	asts.Equal(registerCommand.Email, userInfo.Email)
	asts.Equal(registerCommand.Firstname, userInfo.Firstname)
	asts.Equal(registerCommand.Lastname, userInfo.Lastname)
	asts.NotEmpty(userInfo.AccessToken)
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
