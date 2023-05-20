package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaim struct {
	Email     string
	Firstname string
	Lastname  string
	Exp       time.Time
}

type JwtUtil interface {
	GenerateJwtToken(secret string, claim JwtClaim) (string, error)
}

type jwtGoImpl struct {
	algo jwt.SigningMethod
}

func (self jwtGoImpl) GenerateJwtToken(secret string, claim JwtClaim) (string, error) {
	// create new jwtToken
	token := jwt.New(self.algo)
	claims := token.Claims.(jwt.MapClaims)

	// edit claims
	claims["email"] = claim.Email
	claims["firstname"] = claim.Firstname
	claims["lastname"] = claim.Lastname
	claims["exp"] = claim.Exp

	// sign the token
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}

func NewJwtUtil() JwtUtil {
	return jwtGoImpl{algo: jwt.SigningMethodEdDSA}
}
