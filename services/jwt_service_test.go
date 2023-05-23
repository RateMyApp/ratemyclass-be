package services_test

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/mocks"
	"github.com/ratemyapp/services"
	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JwtServiceTestSuite struct {
	suite.Suite
	config         config.AppConfig
	jwtUtil        utils.JwtUtil
	timeUtil       utils.TimeUtil
	timeUtilMock   mocks.TimeUtilMock
	jwtService     services.JwtService
	jwtUtilMock    mocks.JwtUtilMock
	jwtServiceMock services.JwtService
}

func (self *JwtServiceTestSuite) SetupTest() {
	os.Setenv("GO_ENV", "testing")
	self.config = config.InitAppConfig()

	// utils
	self.timeUtil = utils.NewTimeUtil()
	self.timeUtilMock = mocks.TimeUtilMock{}
	self.jwtUtil = utils.NewJwtUtil()
	self.jwtUtilMock = mocks.JwtUtilMock{}

	// services
	self.jwtService = services.NewJwtService(self.config, self.jwtUtil, self.timeUtil)
	self.jwtServiceMock = services.NewJwtService(self.config, &self.jwtUtilMock, self.timeUtil)
}

func (self *JwtServiceTestSuite) Test_GenerateAccessToken_ShouldReturnToken_WhenGivenCommand() {
	command := services.GenerateTokenCommand{Firstname: "TestF", Lastname: "TestL", Email: "test@gmail.com"}
	token, err := self.jwtService.GenerateAccessToken(command)

	self.Nil(err)
	self.NotEmpty(token)
}

func (self *JwtServiceTestSuite) Test_GenerateAccessToken_ShouldErr_WhenGivenCommand() {
	command := services.GenerateTokenCommand{Firstname: "TestF", Lastname: "TestL", Email: "test@gmail.com"}
	self.jwtUtilMock.On("GenerateJwtToken", mock.Anything, mock.Anything).Return("", errors.New("Failed Test"))
	token, err := self.jwtServiceMock.GenerateAccessToken(command)
	self.Empty(token)
	self.NotNil(err)
	self.Equal(err.StatusCode, http.StatusInternalServerError)
}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}
