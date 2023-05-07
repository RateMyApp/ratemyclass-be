package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
)

type ProfessorService interface {
	CreateProfessor(CreateProfessorCommand) *exceptions.AppError
}

type ProfessorServiceImpl struct {
	repo repositories.ProfessorRepository
}

func (ps *ProfessorServiceImpl) CreateProfessor(command CreateProfessorCommand) *exceptions.AppError {
	professor := models.Professor{Email: command.Email, DirectoryListing: command.DirectoryListing, Lastname: command.Lastname, Firstname: command.Firstname}
	return ps.repo.SaveProfessor(professor)
}

func NewProfessorService(repo repositories.ProfessorRepository) ProfessorService {
	return &ProfessorServiceImpl{repo: repo}
}
