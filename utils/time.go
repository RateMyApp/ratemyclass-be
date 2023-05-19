package utils

import "time"

type TimeUtil interface {
	GenerateTime(minutes uint) time.Time
}

type timeUtil struct{}

func (self timeUtil) GenerateTime(minutes uint) time.Time {
	return time.Now().Add(time.Duration(minutes) * time.Minute)
}

func NewTimeUtil() timeUtil {
	return timeUtil{}
}
