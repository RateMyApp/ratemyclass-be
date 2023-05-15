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
	appConfig = config.InitAppConfig()
	_, client := dao.NewPostgresClient(appConfig)
	postgresClient = client
	postgresClient.Init()
	userRepostory = repositories.NewUserRepository(postgresClient)
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

func Test_FindUserByEmail_ReturnUser_WhenGivenAValidEmail(t *testing.T) {
	beforeEach()
	defer afterEach()
	testUser = models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	postgresClient.Db.Create(&testUser)
	assert := assert.New(t)
	user, err := userRepostory.FindUserByEmail(testUser.Email)

	assert.Nil(err, "FindUserByEmail Error: Expected err to be nil")
	assert.NotNil(user, "FindUserByEmail Error: Expected user to not be nil")
	assert.NotEmpty(user.CreatedAt)
	assert.NotEmpty(user.UpdatedAt)
	assert.False(user.DeletedAt.Valid)
	assert.Equal(testUser.Password, user.Password)
	assert.Equal(testUser.Email, user.Email)
	assert.Equal(testUser.Firstname, user.Firstname)
	assert.Equal(testUser.Lastname, user.Lastname)

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
