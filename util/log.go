package util

import (
	"log"
	"os"
)

var Logger Printer = newDefaultPrinter()

func Print(v ...interface{}) {
	Logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	Logger.Println(v...)
}

// ------------------------------------------------------------------------------

type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type printer struct {
	logger *log.Logger
}

func newDefaultPrinter() *printer {
	outputDestination := os.Stdout
	logMessagePrefix := ""
	loggerFlag := 0 // Do not prefix the date, time, etc. to the beginning of the log messages.
	return &printer{
		logger: log.New(outputDestination, logMessagePrefix, loggerFlag),
	}
}

func (p *printer) Print(v ...interface{}) {
	p.logger.Print(v...)
}

func (p *printer) Printf(format string, v ...interface{}) {
	p.logger.Printf(format, v...)
}

func (p *printer) Println(v ...interface{}) {
	p.logger.Println(v...)
}
