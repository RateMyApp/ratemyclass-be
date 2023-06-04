package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// Service to Register Users
	RegisterUser(command RegisterCommand) *exceptions.AppError
	LoginUser(command LoginCommand) (*exceptions.AppError, *UserDetails)
}

type AuthServiceImpl struct {
	userRepo   repositories.UserRepository
	jwtService JwtService
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

func (as *AuthServiceImpl) LoginUser(command LoginCommand) (*exceptions.AppError, *UserDetails) {
	// check if email exists
	existingUser, err := as.userRepo.FindUserByEmail(command.Email)
	if err != nil {
		return err, nil
	}

	// user does not exist, throw error
	if existingUser == nil {
		ce := exceptions.NewConflictError("Invalid Email or Password")
		return &ce, nil
	}

	// password is correct
	if passwordErr := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(command.Password)); passwordErr != nil {
		ce := exceptions.NewConflictError("Invalid Email or Password")
		return &ce, nil
	}

	accessToken, err := as.jwtService.GenerateAccessToken(
		GenerateTokenCommand{Firstname: existingUser.Firstname, Lastname: existingUser.Lastname, Email: existingUser.Email},
	)
	// error generating accessToken
	if err != nil {
		return err, nil
	}

	userdetails := UserDetails{
		Email:       existingUser.Email,
		Firstname:   existingUser.Firstname,
		Lastname:    existingUser.Lastname,
		AccessToken: accessToken,
	}

	return nil, &userdetails
}

func NewAuthServiceImpl(userRepo repositories.UserRepository, jwtService JwtService) AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}
