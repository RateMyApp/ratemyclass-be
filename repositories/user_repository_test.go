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
)

var (
	appConfig      config.AppConfig
	postgresClient *dao.PostgresClient
	transaction    *gorm.DB
	testUser       models.User
	userRepostory  repositories.UserRepository
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

func Test_FindUserByEmail_ReturnUser_WhenGivenAValidEmail(t *testing.T) {
	beforeEach()
	defer afterEach()
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

func Test_FindUserByEmail_ReturnNil_WhenGivenInvalidEmail(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)

	user, err := userRepostory.FindUserByEmail("notfound@gmail.com")
	assert.Nil(user, "FindUserByEmail Error: Expected user to be Nil")
	assert.Nil(err, "FindUserByEmail Error: Expected err to be Nil")
}

func Test_SaveUser_ReturnNil_WhenGivenAUserToSave(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)

	newUser := models.User{Firstname: "TestFirstname", Lastname: "TestLastname", Email: "Test@email.com", Password: "Testpassword"}

	err := userRepostory.SaveUser(newUser)
	assert.Nil(err, "SaveUser Error: Expected SaveUser to return no Error")

	foundUser, err := userRepostory.FindUserByEmail(newUser.Email)
	assert.NotNil(foundUser)
	assert.Nil(err)
	assert.Equal(newUser.Lastname, foundUser.Lastname)
	assert.Equal(newUser.Firstname, foundUser.Firstname)
	assert.Equal(newUser.Email, foundUser.Email)
	assert.Equal(newUser.Password, foundUser.Password)
}

func TestMain(m *testing.M) {
	beforeAll()
	defer afterAll()
	m.Run()
}
