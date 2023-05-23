package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaim struct {
	Email     string
	Firstname string
	Lastname  string
	Exp       time.Time
}
type jwtClaim struct {
	Firstname string
	Lastname  string
	jwt.RegisteredClaims
}

type JwtUtil interface {
	GenerateJwtToken(secret string, claim JwtClaim) (string, error)
}

type jwtGoImpl struct {
	algo   jwt.SigningMethod
	issuer string
}

func (self jwtGoImpl) GenerateJwtToken(secret string, claim JwtClaim) (string, error) {
	// create new jwtToken
	token := jwt.NewWithClaims(self.algo, jwtClaim{
		claim.Firstname,
		claim.Lastname,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(claim.Exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    self.issuer,
			Subject:   claim.Email,
		},
	},
	)

	// sign the token
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func NewJwtUtil() JwtUtil {
	return jwtGoImpl{algo: jwt.SigningMethodHS256}
}
