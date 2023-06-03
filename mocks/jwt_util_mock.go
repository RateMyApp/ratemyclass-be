package mocks

import (
	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/mock"
)

type JwtUtilMock struct {
	mock.Mock
}

func (self *JwtUtilMock) GenerateJwtToken(secret string, claim utils.JwtClaim) (string, error) {
	args := self.Called(secret, claim)

	return args.String(0), args.Error(1)
}

var _ utils.JwtUtil = (*JwtUtilMock)(nil)
