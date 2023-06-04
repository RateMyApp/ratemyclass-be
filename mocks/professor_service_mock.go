package mocks

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/services"
	"github.com/stretchr/testify/mock"
)

type ProfessorServiceMock struct {
	mock.Mock
}

func (psm *ProfessorServiceMock) CreateProfessor(command services.CreateProfessorCommand) *exceptions.AppError {
	args := psm.Called(command)
	
	var return0 *exceptions.AppError = nil
	
	if val, ok := args.Get(0).(*exceptions.AppError) ; ok{
			return0 = val
	}
	return return0
}

var _ services.ProfessorService = (*ProfessorServiceMock)(nil)