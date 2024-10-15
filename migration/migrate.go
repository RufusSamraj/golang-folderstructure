package migration

import (
	"main.go/database"
	"main.go/entity"
	"main.go/errors"
)

func AutoMigrate() *errors.AppError {
	db, dbErr := database.Connect()
	if dbErr != nil {
		return errors.NewUnexpectedError(dbErr.Error())
	}

	defer database.Close(db)

	err := db.AutoMigrate(&entity.StudentEntity{})
	if err != nil {
		return errors.NewUnexpectedError(err.Error())
	}

	return nil
}
