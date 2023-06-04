package services

import (
	"log"
	"strconv"

	"github.com/ratemyapp/config"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/utils"
)

type JwtService interface {
	GenerateAccessToken(command GenerateTokenCommand) (string, *exceptions.AppError)
	VerifyAccessToken(token string)
}

type jwtServiceImpl struct {
	accessTokenSecret   string
	accessTokenDuration int
	timeUtil            utils.TimeUtil
	jwtUtil             utils.JwtUtil
}

func (self *jwtServiceImpl) VerifyAccessToken(string) {
}

func (self *jwtServiceImpl) GenerateAccessToken(command GenerateTokenCommand) (string, *exceptions.AppError) {
	appTime := self.timeUtil.GenerateTime(uint(self.accessTokenDuration))

	claim := utils.JwtClaim{
		Email:     command.Email,
		Firstname: command.Firstname,
		Lastname:  command.Lastname,
		Exp:       appTime,
	}

	// sign the token
	tokenString, err := self.jwtUtil.GenerateJwtToken(self.accessTokenSecret, claim)
	if err != nil {
		log.Println(err)
		exc := exceptions.NewInternalServerError()
		return "", &exc
	}

	return tokenString, nil
}

func NewJwtService(config config.AppConfig, jwtUtil utils.JwtUtil, timeUtil utils.TimeUtil) JwtService {
	time, _ := strconv.Atoi(config.TIME)

	if time < 0 {
		time = -time
	}

	return &jwtServiceImpl{
		timeUtil:            timeUtil,
		jwtUtil:             jwtUtil,
		accessTokenSecret:   config.JWT_SECRET,
		accessTokenDuration: time,
	}
}
