package flags

import (
	"github.com/lolmerkat/tempr/internal/app/tempr"
)

func handleInfoFileCreation(disableInfoFile *bool) {
	logger := tempr.Logger()
	logger.Debugf("Info file creation: %t", *disableInfoFile)
}
