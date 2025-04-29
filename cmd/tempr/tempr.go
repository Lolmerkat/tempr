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
	// TEST: Expanding a template
	// template.Expand(".", logger)
	// TEST: Writing a template
	// template.WriteToFile(".", logger)
	readTemplate := templating.LoadFromFile("./HTML5.yml", logger)
	readTemplate.Expand(".", logger)
}
