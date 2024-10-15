package repository

import (
	"main.go/database"
	"main.go/entity"
	"main.go/errors"
	"main.go/logger"
)

type RepoStc struct {
}

type RepoInfc interface {
	GetStudent(req int) (entity.StudentEntity, *errors.AppError)
	CreateStudent(req entity.StudentEntity) (*errors.AppError, *errors.AppError)
}

func NewRepository() RepoInfc {
	return RepoStc{}
}

func (r RepoStc) GetStudent(req int) (entity.StudentEntity, *errors.AppError) {
	dbClient, dbErr := database.Connect()
	if dbErr != nil {
		logger.InfoLogger.Println("Error in db connection: ", dbErr)
		return entity.StudentEntity{}, errors.NewUnexpectedError(dbErr.Error())
	}

	defer database.Close(dbClient)
	var dbResponse entity.StudentEntity

	queryErr := dbClient.Where("roll_no = ?", req).Find(&dbResponse).Error
	if queryErr != nil {
		logger.InfoLogger.Println("Error in query: ", queryErr)
		return entity.StudentEntity{}, errors.NewUnexpectedError(queryErr.Error())
	}

	return dbResponse, nil
}

func (r RepoStc) CreateStudent(req entity.StudentEntity) (*errors.AppError, *errors.AppError) {
	dbClient, dbErr := database.Connect()
	if dbErr != nil {
		logger.InfoLogger.Println("Error in db connection: ", dbErr)
		return nil, errors.NewUnexpectedError(dbErr.Error())
	}

	defer database.Close(dbClient)

	queryErr := dbClient.Create(&req).Error
	if queryErr != nil {
		logger.InfoLogger.Println("Error in query: ", queryErr)
		return nil, errors.NewUnexpectedError(queryErr.Error())
	}

	return &errors.AppError{
		Code:    200,
		Message: "Inserted successfully",
	}, nil
}
