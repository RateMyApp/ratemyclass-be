package repositories_test

import (
	"context"
	"log"
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
	user, err := userRepostory.FindUserByEmail(testUser.Email)
	if err != nil {
		log.Println("hello again")
		t.Error(err)
	}
	if user == nil {
		t.Errorf("Expected User to not be nil")
	}
	if user.Email != testUser.Email {
		t.Errorf("Expect %v but got %v", testUser.Email, user.Email)
	}
}

func Test_FindUserByEmail_ReturnNil_WhenGivenInvalidEmail(t *testing.T) {
	beforeEach()
	defer afterEach()
	assert := assert.New(t)

	user, err := userRepostory.FindUserByEmail("notfound@gmail.com")
	assert.Nil(user, "FindUserByEmail Error: Expected user to be Nil")
	assert.Nil(err, "FindUserByEmail Error: Expected err to be Nil")
}

func TestMain(m *testing.M) {
	beforeAll()
	defer afterAll()
	m.Run()
}
