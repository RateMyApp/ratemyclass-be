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

type UserRepositorySuite struct {
	suite.Suite
	appConfig      config.AppConfig
	postgresClient *dao.PostgresClient
	transaction    *gorm.DB
	testUser       models.User
	userRepostory  repositories.UserRepository
}

func (suite *UserRepositorySuite) SetupSuite() {
	os.Setenv("GO_ENV", "testing")
	suite.appConfig = config.InitAppConfig()
	_, client := dao.NewPostgresClient(suite.appConfig)
	suite.postgresClient = client
	suite.postgresClient.Init()
	suite.userRepostory = repositories.NewUserRepository(suite.postgresClient)
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.transaction = suite.postgresClient.Db.Begin()
}

func (suite *UserRepositorySuite) TearDownTest() {
	suite.transaction.Rollback()
}

func (suite *UserRepositorySuite) TearDownSuite() {
	for _, model := range *models.GetModels() {
		suite.postgresClient.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	suite.postgresClient.Close(ctx)
}

func (suite *UserRepositorySuite) Test_FindUserByEmail_ReturnUser_WhenGivenAValidEmail() {
	testUser := models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	suite.postgresClient.Db.Create(&testUser)
	user, err := suite.userRepostory.FindUserByEmail(testUser.Email)

	suite.Nil(err, "FindUserByEmail Error: Expected err to be nil")
	suite.NotNil(user, "FindUserByEmail Error: Expected user to not be nil")
	suite.NotEmpty(user.CreatedAt)
	suite.NotEmpty(user.UpdatedAt)
	suite.False(user.DeletedAt.Valid)
	suite.Equal(testUser.Password, user.Password)
	suite.Equal(testUser.Email, user.Email)
	suite.Equal(testUser.Firstname, user.Firstname)
	suite.Equal(testUser.Lastname, user.Lastname)
}

func (suite *UserRepositorySuite) Test_FindUserByEmail_ReturnNil_WhenGivenInvalidEmail() {
	user, err := suite.userRepostory.FindUserByEmail("notfound@gmail.com")
	suite.Nil(user, "FindUserByEmail Error: Expected user to be Nil")
	suite.Nil(err, "FindUserByEmail Error: Expected err to be Nil")
}

func (suite *UserRepositorySuite) Test_SaveUser_ReturnNil_WhenGivenAUserToSave() {
	newUser := models.User{Firstname: "TestFirstname", Lastname: "TestLastname", Email: "Test@email.com", Password: "Testpassword"}

	err := suite.userRepostory.SaveUser(newUser)
	suite.Nil(err, "SaveUser Error: Expected SaveUser to return no Error")

	foundUser, err := suite.userRepostory.FindUserByEmail(newUser.Email)
	suite.NotNil(foundUser)
	suite.Nil(err)
	suite.Equal(newUser.Lastname, foundUser.Lastname)
	suite.Equal(newUser.Firstname, foundUser.Firstname)
	suite.Equal(newUser.Email, foundUser.Email)
	suite.Equal(newUser.Password, foundUser.Password)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
