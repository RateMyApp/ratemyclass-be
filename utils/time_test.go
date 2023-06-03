package utils_test

import (
	"testing"
	"time"

	"github.com/ratemyapp/utils"
	"github.com/stretchr/testify/suite"
)

type TestUtilTestSuite struct {
	suite.Suite
	timeUtil utils.TimeUtil
}

func (self *TestUtilTestSuite) SetupSuite() {
	self.timeUtil = utils.NewTimeUtil()
}

func (self *TestUtilTestSuite) Test_GenerateTime_WithoutAddedTime() {
	prevTime := time.Now()
	currTime := self.timeUtil.GenerateTime(0)
	nextTime := time.Now()

	self.NotEmpty(currTime)
	self.GreaterOrEqual(currTime.Unix(), prevTime.Unix())
	self.LessOrEqual(currTime.Unix(), nextTime.Unix())
}

func (self *TestUtilTestSuite) Test_GenerateTime_WithAddedTime() {
	prevTime := time.Now().Add(3 * time.Minute)
	currTime := self.timeUtil.GenerateTime(4)
	nextTime := time.Now().Add(5 * time.Minute)

	self.NotEmpty(currTime)
	self.GreaterOrEqual(currTime.Unix(), prevTime.Unix())
	self.LessOrEqual(currTime.Unix(), nextTime.Unix())
}

func TestTimeUtilTestSuite(t *testing.T) {
	suite.Run(t, new(TestUtilTestSuite))
}
