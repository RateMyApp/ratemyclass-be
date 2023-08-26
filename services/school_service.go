package services

import (
	"github.com/ratemyapp/exceptions"
	"github.com/ratemyapp/models"
	"github.com/ratemyapp/repositories"
)

type SchoolService interface {
	CreateSchool(CreateSchoolCommand) *exceptions.AppError
	SearchSchoolsInfoByName(command string) ([]SearchSchoolInfoResult, *exceptions.AppError)
}

type schoolServiceImpl struct {
	schoolRepo repositories.SchoolRepository
}

func (ss *schoolServiceImpl) CreateSchool(command CreateSchoolCommand) *exceptions.AppError {
	newSchool := models.School{Name: command.Name, Address: command.Address, City: command.City, ProvinceOrState: command.ProvinceOrState, Country: command.Country}

	for _, val := range command.EmailDomains {
		newSchool.EmailDomains = append(newSchool.EmailDomains, models.EmailDomain{Domain: val})
	}

	err := ss.schoolRepo.SaveSchool(newSchool)
	if err != nil {
		return err
	}
	return nil
}

func (ss *schoolServiceImpl) SearchSchoolsInfoByName(command string) ([]SearchSchoolInfoResult, *exceptions.AppError) {
	var result []SearchSchoolInfoResult
	schools, err := ss.schoolRepo.SearchSchoolsByName(command)

	if err != nil {
		return result, err
	}

	result = []SearchSchoolInfoResult{}

	for _, v := range schools {
		location := ""

		if v.City != "" {
			location += v.City + ","
		}

		location += v.ProvinceOrState + "," + v.Country

		result = append(result, SearchSchoolInfoResult{Id: v.ID, Name: v.Name, Location: location})
	}
	return result, nil
}

func NewSchoolService(schoolRepo repositories.SchoolRepository) SchoolService {
	return &schoolServiceImpl{schoolRepo: schoolRepo}
}
