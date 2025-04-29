package tempr

import (
	"os"

	"github.com/charmbracelet/log"
)

var logger = log.NewWithOptions(os.Stdout, log.Options{
	ReportTimestamp: true,
	TimeFormat: "15:04:05.000",
})

func Logger() *log.Logger {
	return logger
}
