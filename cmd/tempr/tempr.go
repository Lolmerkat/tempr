package main

import (
	_ "embed"
	"os"

	"github.com/charmbracelet/log"
	"github.com/lolmerkat/tempr/internal/app/tempr/ui"
)

func main() {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
		TimeFormat: "15:04:05.000",
	})

	ui.GetSplash(logger).Print(logger)
}
