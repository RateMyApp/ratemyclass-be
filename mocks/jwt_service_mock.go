package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
)

type JwtServiceMock struct {
	mock.Mock
}

func (jsm *JwtServiceMock) GenerateAccessToken(command services.GenerateTokenCommand) (string, *exceptions.AppError) {
	args := jsm.Called(command)

	var (
		return0 string               = ""
		return1 *exceptions.AppError = nil
	)

	if val, ok := args.Get(0).(string); ok {
		return0 = val
	}
	if val, ok := args.Get(1).(*exceptions.AppError); ok {
		return1 = val
	}

	return return0, return1
}

var _ services.JwtService = (*JwtServiceMock)(nil)
