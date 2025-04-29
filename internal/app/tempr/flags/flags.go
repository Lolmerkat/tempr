package flags

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/lolmerkat/tempr/internal/app/tempr"
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
	flag.Usage = usage
	flag.Parse()
	//vvv  INFO: Flag handlers that should be executed immediately vvv
	handleLogLevel(logLevelPtr)
}

func usage() {
	out := tempr.Logger()
	usageHeaderStyle := lipgloss.NewStyle().Bold(true)
	fmt.Fprintf(os.Stderr, "\033[1;34mUsage:\033[0m\n") // Blue bold header

	out.Printf(usageHeaderStyle.Render("Usage:"))
	fmt.Fprintf(os.Stderr, "  %s [flags]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\033[1;33mFlags:\033[0m\n") // Yellow bold "Flags:"

	flag.VisitAll(func(f *flag.Flag) {
		// Example: green flag name, cyan shorthand, magenta default
		fmt.Fprintf(os.Stderr, "  \033[32m--%s\033[0m", f.Name)
		if f.Shorthand != "" {
			fmt.Fprintf(os.Stderr, ", \033[36m-%s\033[0m", f.Shorthand)
		}
		fmt.Fprintf(os.Stderr, " \033[35m(default: %s)\033[0m\n      %s\n", f.DefValue, f.Usage)
	})
}
