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

func (jsts *JwtServiceTestSuite) SetupTest() {
	os.Setenv("GO_ENV", "testing")
	jsts.config = config.InitAppConfig()
			
	// utils
	jsts.timeUtil = utils.NewTimeUtil()
	jsts.timeUtilMock = mocks.TimeUtilMock{}
	jsts.jwtUtil = utils.NewJwtUtil()
	jsts.jwtUtilMock = mocks.JwtUtilMock{}

	// services
	jsts.jwtService = services.NewJwtService(jsts.config, jsts.jwtUtil, jsts.timeUtil)
	jsts.jwtServiceMock = services.NewJwtService(jsts.config, &jsts.jwtUtilMock, jsts.timeUtil)
}

func (jsts *JwtServiceTestSuite) Test_GenerateAccessToken_ShouldReturnToken_WhenGivenCommand() {
	command := services.GenerateTokenCommand{Firstname: "TestF", Lastname: "TestL", Email: "test@gmail.com"}
	token, err := jsts.jwtService.GenerateAccessToken(command)

	jsts.Nil(err)
	jsts.NotEmpty(token)
}

func (jsts *JwtServiceTestSuite) Test_GenerateAccessToken_ShouldErr_WhenGivenCommand() {
	command := services.GenerateTokenCommand{Firstname: "TestF", Lastname: "TestL", Email: "test@gmail.com"}
	jsts.jwtUtilMock.On("GenerateJwtToken", mock.Anything, mock.Anything).Return("", errors.New("Failed Test"))
	token, err := jsts.jwtServiceMock.GenerateAccessToken(command)
	jsts.Empty(token)
	jsts.NotNil(err)
	jsts.Equal(err.StatusCode, http.StatusInternalServerError)
}

func TestJwtServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JwtServiceTestSuite))
}
