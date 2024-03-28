package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	infoLogFile, err := os.OpenFile("logger/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open info log file: %v", err)
	}
	errorLogFile, err := os.OpenFile("logger/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	flag := log.Ldate | log.Ltime | log.Lshortfile

	InfoLogger = log.New(infoLogFile, "INFO: ", flag)
	ErrorLogger = log.New(errorLogFile, "ERROR: ", flag)
}