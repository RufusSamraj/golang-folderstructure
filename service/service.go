package service

import (
	"main.go/dto"
	"main.go/errors"
	"main.go/logger"
	"main.go/mapper"
	"main.go/repository"
)

type SvcStc struct {
	repository repository.RepoInfc
}

type SvcInfc interface {
	GetStudent(req int) (dto.Student, *errors.AppError)
	CreateStudent(req dto.Student) (*errors.AppError, *errors.AppError)
}

func NewService(repo repository.RepoInfc) SvcInfc {
	return SvcStc{repository: repo}
}

func (s SvcStc) GetStudent(req int) (dto.Student, *errors.AppError) {

	repoResponse, repoErr := s.repository.GetStudent(req)
	if repoErr != nil {
		logger.InfoLogger.Println("Error in repository call: ", repoErr)
		return dto.Student{}, errors.NewUnexpectedError(repoErr.Message)
	}

	response := mapper.StudentEntityToDto(repoResponse)

	return response, nil
}

func (s SvcStc) CreateStudent(req dto.Student) (*errors.AppError, *errors.AppError) {

	repoRequest := mapper.StudentDtoToEntity(req)

	repoResponse, repoErr := s.repository.CreateStudent(repoRequest)
	if repoErr != nil {
		logger.InfoLogger.Println("Error in repository call: ", repoErr)
		return nil, errors.NewUnexpectedError(repoErr.Message)
	}

	return repoResponse, nil
}
