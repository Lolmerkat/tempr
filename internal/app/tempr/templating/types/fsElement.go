package templating

import (
	"github.com/charmbracelet/log"
)

type FSElement interface {
	Expand(path string, logger *log.Logger)
}

type RawFSElement struct {
	Name		string			`yaml:"name"`
	Content		[]string		`yaml:"content"`
	Children	[]RawFSElement	`yaml:"children,omitempty"`
}

func (raw RawFSElement) ToFSElement() FSElement {
	if raw.Content != nil {
		return &File {
			Name: raw.Name,
			Content: raw.Content,
		}
	}

	dir := &Directory {
		Name: raw.Name,
		Children: make([]FSElement, 0, len(raw.Children)),
	}

	for _, child := range raw.Children {
		dir.Children = append(dir.Children, child.ToFSElement())
	}

	return dir
}
