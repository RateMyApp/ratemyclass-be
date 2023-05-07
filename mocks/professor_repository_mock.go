package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/stretchr/testify/mock"
)

type ProfessorRepository struct {
	mock.Mock
}

func (pr *ProfessorRepository) SaveProfessor(professor models.Professor) *exceptions.AppError {
	args := pr.Called(professor)

	var arg0 *exceptions.AppError

	if val, ok := args.Get(0).(*exceptions.AppError); ok {
		arg0 = val
	} else {
		arg0 = nil
	}

	return arg0
}
