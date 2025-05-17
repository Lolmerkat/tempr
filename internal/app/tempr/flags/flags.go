package flags

import (
	flag "github.com/spf13/pflag"
)

var LogLevelPtr *int
var DisableInfoFilePtr *bool

func Declare() {
	// log-level
	logLevelPtr = flag.IntP("log-level", "l", 3,
		`Defines the log level
	0: Fatal errors
	1: Non-Fatal errors
	2: Warnings
	3: Program information
	4: Debug information
`)

	disableInfoFilePtr = flag.Bool("no-info-file", false,
		"Disables the generation of the '.tempr' info file")
}

func Handle() {
	flag.Parse()
	// TODO: ADD STYLING
	//vvv  INFO: Flag handlers that should be executed immediately vvv
	handleLogLevel(logLevelPtr)
	handleInfoFileCreation(disableInfoFilePtr)
}
