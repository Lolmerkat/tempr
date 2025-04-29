package main

import (
	"github.com/lolmerkat/tempr/internal/app/tempr"
	"github.com/lolmerkat/tempr/internal/app/tempr/flags"
	"github.com/lolmerkat/tempr/internal/app/tempr/ui"
)

func main() {
	flags.Declare()
	flags.Handle()

	logger := tempr.Logger()
	ui.GetSplash(logger).Print(logger)

	// TEST:
	template :=	getTestTemplate()
	template.Expand(".", logger)
}
