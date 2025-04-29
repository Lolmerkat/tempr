package templating

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Template struct {
	Name      string      `yaml:"name"`
	Author    string      `yaml:"author,omitempty"`
	Version   string      `yaml:"version,omitempty"`
	Languages []string    `yaml:"languages,omitempty"`
	Content   []FSElement `yaml:"content"`
}

func (t Template) Expand(path string, logger *log.Logger) {
	templateDirPath := fmt.Sprintf("%s/%s", path, t.Name)
	err := os.Mkdir(templateDirPath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		logger.Fatalf("Error creating project directory '%s': %v", templateDirPath, err)
	}

	for _, e := range t.Content {
		e.Expand(templateDirPath, logger)
	}
}
