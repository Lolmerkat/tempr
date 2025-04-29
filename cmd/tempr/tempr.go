package main

import (
	"fmt"
	"os"

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
	templateDirPath := fmt.Sprintf("./%s", template.Name)
	err := os.Mkdir(templateDirPath, os.ModeAppend)
	if err != nil && !os.IsExist(err) {
		logger.Fatalf("Error creating '%s': %v", templateDirPath, err)
	}
	for _, e := range template.Content {
		e.Expand(templateDirPath, logger)
	}
}
