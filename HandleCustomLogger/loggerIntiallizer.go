package customlogger

import (
	"log"
	"os"
)

// Calling custom logger function
var Log *CustomLogger = NewLogger(os.Stdout, "APP: ", log.LstdFlags)
