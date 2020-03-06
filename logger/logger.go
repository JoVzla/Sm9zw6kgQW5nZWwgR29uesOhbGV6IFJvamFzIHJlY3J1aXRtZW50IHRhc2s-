package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Create a new instance of the logger.
var Log = logrus.New()

//Initialize loggger
func InitLogger() {

	// Add this line for logging filename and line number
	Log.SetReportCaller(true)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	Log.SetLevel(logrus.InfoLevel)
}
