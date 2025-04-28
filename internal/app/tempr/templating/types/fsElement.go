package templating

import (
	"github.com/charmbracelet/log"
)

type FSElement interface {
	Expand(path string)
	Expand(path string, logger *log.Logger)
}

type RawFSElement struct {
	Name		string			`yaml:"name"`
	Content		[]string		`yaml:"content"`
	Children	[]RawFSElement	`yaml:"children,omitempty"`
}
