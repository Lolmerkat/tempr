package main

import (
	"github.com/lolmerkat/tempr/internal/app/tempr"
	"github.com/lolmerkat/tempr/internal/app/tempr/ui"
	"github.com/lolmerkat/tempr/internal/app/tempr/flags"
)

func main() {
	flags.Declare()
	flags.Handle()

	logger := tempr.Logger()
	ui.GetSplash(logger).Print(logger)
}
