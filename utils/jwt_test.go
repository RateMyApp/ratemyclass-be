package utils_test

import (
	"testing"
	"time"

	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/suite"
)

type JwtUtilTestSuite struct {
	suite.Suite
	jwtUtil utils.JwtUtil
}

func (self *JwtUtilTestSuite) SetupSuite() {
	self.jwtUtil = utils.NewJwtUtil()
}

func (self *JwtUtilTestSuite) TestGenerateJwtToken() {
	claim := utils.JwtClaim{
		Email:     "test@gmail.com",
		Firstname: "TestFirstname",
		Lastname:  "TestLastname",
		Exp:       time.Now().Add(3 * time.Minute),
	}
	token, error := self.jwtUtil.GenerateJwtToken("secret", claim)
	self.NotEmpty(token)
	self.Nil(error)
}

func TestJwtUtilTestSuite(t *testing.T) {
	suite.Run(t, new(JwtUtilTestSuite))
}
