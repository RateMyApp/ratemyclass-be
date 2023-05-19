package mocks

import (
	"time"

	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/mock"
)

type TimeUtilMock struct {
	mock.Mock
}

func (self *TimeUtilMock) GenerateTime(minutes uint) time.Time {
	args := self.Called(minutes)

	return args.Get(0).(time.Time)
}

var _ utils.TimeUtil = (*TimeUtilMock)(nil)
