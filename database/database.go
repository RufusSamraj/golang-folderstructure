package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"log/slog"
	"main.go/constant"
	"os"
	"time"
)

func Connect() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", constant.HOST, constant.USER, constant.PASS, constant.DBNAME, constant.PORT)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		fmt.Println("Cannot connect to database", err)
		return nil, err
	}

	return db, nil
}

func Close(db *gorm.DB) {
	pgDb, err := db.DB()
	if err != nil {
		slog.Error("Cannot get underlying sql sb to close")
	}

	pgDb.Close()
}
