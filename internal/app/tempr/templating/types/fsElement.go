package templating

import (
	"github.com/charmbracelet/log"
)

type FSElement interface {
	Expand(path string, logger *log.Logger)
}

type rawFSElement struct {
	Name		string			`yaml:"name"`
	Content		[]string		`yaml:"content,omitempty"`
	Children	[]rawFSElement	`yaml:"children,omitempty"`
}

func (raw rawFSElement) ToFSElement(logger *log.Logger) FSElement {
	logger.Debugf("ToFSElement: Processing '%s'", raw.Name)
	if raw.Content != nil {
		logger.Debugf("ToFSElement: Processing '%s' - File found", raw.Name)
		return &File {
			Name: raw.Name,
			Content: raw.Content,
		}
	}

	logger.Debugf("ToFSElement: Processing '%s' - Directory found", raw.Name)
	dir := &Directory {
		Name: raw.Name,
		Children: make([]FSElement, 0, len(raw.Children)),
	}

	for _, child := range raw.Children {
		dir.Children = append(dir.Children, child.ToFSElement(logger))
	}

	return dir
}
