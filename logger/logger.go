package logger

import (
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
)

func init() {
	file := os.Stdout
	log.SetFlags(log.Ltime)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Lmsgprefix|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Lmsgprefix|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Lmsgprefix|log.Lshortfile)
}
