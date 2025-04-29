package main

import (
	"github.com/lolmerkat/tempr/internal/app/tempr"
	"github.com/lolmerkat/tempr/internal/app/tempr/ui"
)

func main() {

	logger := tempr.Logger()
	ui.GetSplash(logger).Print(logger)
}
