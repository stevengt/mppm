package util

import (
	"io"
	"log"
	"os"
)

var Logger WritePrinter = newDefaultWritePrinter()

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

// Contains common printing methods.
//
// The io.Writer interface is extended, as well, so that it
// can be used to redirect output from cobra commands.
type WritePrinter interface {
	io.Writer
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type writePrinter struct {
	logger *log.Logger
}

func newDefaultWritePrinter() *writePrinter {
	outputDestination := os.Stdout
	logMessagePrefix := ""
	loggerFlag := 0 // Do not prefix the date, time, etc. to the beginning of the log messages.
	return &writePrinter{
		logger: log.New(outputDestination, logMessagePrefix, loggerFlag),
	}
}

func (p *writePrinter) Print(v ...interface{}) {
	p.logger.Print(v...)
}

func (p *writePrinter) Printf(format string, v ...interface{}) {
	p.logger.Printf(format, v...)
}

func (p *writePrinter) Println(v ...interface{}) {
	p.logger.Println(v...)
}

func (printer *writePrinter) Write(p []byte) (n int, err error) {
	printer.Print(string(p))
	return len(p), nil
}
