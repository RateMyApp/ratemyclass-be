package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ratemyapp/config"
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// Service to Register Users
	RegisterUser(command RegisterCommand) *exceptions.AppError
	CheckLogin(command LoginCommand) (*exceptions.AppError, UserDetails)
	GenerateJWTtoken(command UserDetails) (string, error)
	// VerifyJWTtoken(command UserDetails)
}

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}


func (as *AuthServiceImpl) RegisterUser(command RegisterCommand) *exceptions.AppError {
	// check for existing user
	existingUser, err := as.userRepo.FindUserByEmail(command.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		ce := exceptions.NewConflictError("User Already Exists")
		return &ce
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(command.Password), 8)

	command.Password = string(hashPassword)

	var user models.User = models.User{Firstname: command.Firstname, Lastname: command.Lastname, Email: command.Email, Password: command.Password}

	serr := as.userRepo.SaveUser(user)
	if serr != nil {
		return serr
	}
	return nil
}
func (as *AuthServiceImpl) CheckLogin(command LoginCommand) (*exceptions.AppError, UserDetails){
	// check if email exists 
	existingUser, err := as.userRepo.FindUserByEmail(command.Email)
	var userdetails UserDetails

	if err != nil{
		return err,userdetails
	}
	// user does not exist, throw error
	if existingUser == nil{
		ce:=exceptions.NewConflictError("User does not exist")
		return &ce,userdetails
	}
	// if password is the same then return user details
	if existingUser.Password == command.Password{
		userdetails.Email = existingUser.Email
		userdetails.Firstname = existingUser.Firstname
		userdetails.Lastname = existingUser.Lastname
		return nil,userdetails
	}else{
		return nil,userdetails
	}
}
 
func (as *AuthServiceImpl) GenerateJWTtoken(command UserDetails) (string,error) {

	secretkey:= config.InitAppConfig()
	timeStr:= secretkey.TIME
	timeInt, _ := strconv.Atoi(timeStr)

	//create new jwtToken
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)

	//edit claims
	claims["email"] = command.Email
	claims["firstname"] = command.Firstname
	claims["lastname"] = command.Lastname
	claims["exp"] = time.Now().Add(time.Duration(timeInt) * time.Minute)
	
	//sign the token

	tokenString, err := token.SignedString(secretkey.JWT_SECRET)
	return tokenString,err


}
func NewAuthServiceImpl(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}
