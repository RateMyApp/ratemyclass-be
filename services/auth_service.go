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

func NewAuthServiceImpl(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{userRepo: userRepo}
}
