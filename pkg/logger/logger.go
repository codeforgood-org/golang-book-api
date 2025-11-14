package logger

import (
	"log"
	"os"
)

var (
	// Info logger for informational messages
	Info *log.Logger
	// Warning logger for warning messages
	Warning *log.Logger
	// Error logger for error messages
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
