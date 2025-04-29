package main

import (
	"github.com/lolmerkat/tempr/internal/app/tempr"
	"github.com/lolmerkat/tempr/internal/app/tempr/flags"
	templating "github.com/lolmerkat/tempr/internal/app/tempr/templating/types"
	"github.com/lolmerkat/tempr/internal/app/tempr/ui"
)

func main() {
	flags.Declare()
	flags.Handle()

	logger := tempr.Logger()
	ui.GetSplash(logger).Print(logger)

	// template :=	getTestTemplate()
	// TEST: EXPANDING A TEMPLATE
	// template.Expand(".", logger)

	// TEST: WRITE A TEMPLATE TO A FILE ON DISK
	// template.WriteToFile(".", logger)

	// TEST: LOADING A TEMPLATE FROM FILE AND EXPADING IT
	readTemplate := templating.LoadFromFile("./HTML5.yml", logger)
	readTemplate.Expand(".", logger)
}
