package repositories_test

import (
	"context"
	"os"
	"testing"
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/dao"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/stretchr/testify/suite"
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
	suite.testUser = models.User{Email: "test@gmail.com", Firstname: "Test1", Lastname: "TestTwo", Password: "hello123"}
	suite.postgresClient.Db.Create(&suite.testUser)
	suite.userRepostory = repositories.NewUserRepository(suite.postgresClient)
}

func (suite *UserRepositorySuite) SetupTest(){
	suite.transaction = suite.postgresClient.Db.Begin()
}

func(suite *UserRepositorySuite) TearDownTest(){
	suite.transaction.Rollback()
}

func (suite *UserRepositorySuite) TearDownSuite(){
	for _, model := range *models.GetModels() {
		suite.postgresClient.Db.Unscoped().Where("1 = 1").Delete(model)
	}
	ctx := context.Background()
	suite.postgresClient.Close(ctx)
}

func(suite *UserRepositorySuite) Test_FindUserByEmail_ReturnUser_WhenGivenAValidEmail(t *testing.T) {
	
	assert := assert.New(t)
	user, err := suite.userRepostory.FindUserByEmail(suite.testUser.Email)

	assert.Nil(err, "FindUserByEmail Error: Expected err to be nil")
	assert.NotNil(user, "FindUserByEmail Error: Expected user to not be nil")
	assert.NotEmpty(user.CreatedAt)
	assert.NotEmpty(user.UpdatedAt)
	assert.False(user.DeletedAt.Valid)
	assert.Equal(suite.testUser.Password, user.Password)
	assert.Equal(suite.testUser.Email, user.Email)
	assert.Equal(suite.testUser.Firstname, user.Firstname)
	assert.Equal(suite.testUser.Lastname, user.Lastname)
}

func (suite *UserRepositorySuite)Test_FindUserByEmail_ReturnNil_WhenGivenInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	user, err := suite.userRepostory.FindUserByEmail("notfound@gmail.com")
	assert.Nil(user, "FindUserByEmail Error: Expected user to be Nil")
	assert.Nil(err, "FindUserByEmail Error: Expected err to be Nil")
}

func(suite *UserRepositorySuite)Test_SaveUser_ReturnNil_WhenGivenAUserToSave(t *testing.T) {
	
	assert := assert.New(t)

	newUser := models.User{Firstname: "TestFirstname", Lastname: "TestLastname", Email: "Test@email.com", Password: "Testpassword"}

	err := suite.userRepostory.SaveUser(newUser)
	assert.Nil(err, "SaveUser Error: Expected SaveUser to return no Error")

	foundUser, err := suite.userRepostory.FindUserByEmail(newUser.Email)
	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(newUser.Lastname, foundUser.Lastname)
	assert.Equal(newUser.Firstname, foundUser.Firstname)
	assert.Equal(newUser.Email, foundUser.Email)
	assert.Equal(newUser.Password, foundUser.Password)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}

