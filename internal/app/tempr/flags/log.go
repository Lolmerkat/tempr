package flags

import (
	"github.com/charmbracelet/log"
	"github.com/lolmerkat/tempr/internal/app/tempr"
)

func handleLogLevel(level *int) {
	logger := tempr.Logger()

	levelMap := map[int]log.Level {
		0: log.FatalLevel,
		1: log.ErrorLevel,
		2: log.WarnLevel,
		3: log.InfoLevel,
		4: log.DebugLevel,
	}

	logger.SetLevel(levelMap[*level])
}
