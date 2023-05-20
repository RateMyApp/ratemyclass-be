package utils

import "time"

type TimeUtil interface {
	GenerateTime(minutes uint) time.Time
}

type timeUtilImpl struct{}

func (self timeUtilImpl) GenerateTime(minutes uint) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func NewTimeUtil() timeUtilImpl {
	return timeUtilImpl{}
}
