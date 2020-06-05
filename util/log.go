package util

import (
	"log"
	"os"
)

var Logger *log.Logger = getNewDefaultLogger()

func Print(v ...interface{}) {
	Logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	Logger.Println(v...)
}

func getNewDefaultLogger() *log.Logger {
	outputDestination := os.Stdout
	logMessagePrefix := ""
	loggerFlag := 0 // Do not prefix the date, time, etc. to the beginning of the log messages.
	return log.New(outputDestination, logMessagePrefix, loggerFlag)
}
