package flags

import (
	flag "github.com/spf13/pflag"
)

var helpPtr *bool
var logLevelPtr *int

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
}

func Handle() {
	flag.Parse()
	//vvv  INFO: Flag handlers that should be executed immediately vvv
	handleLogLevel(logLevelPtr)
}
