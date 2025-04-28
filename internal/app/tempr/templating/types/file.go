package templating

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type File struct {
	Name		string		`yaml:"name"`
	Content		[]string	`yaml:"content"`
}

func (f *File) Expand(path string, logger *log.Logger) {
	filePath := fmt.Sprintf("%s/%s", path, f.Name)
	file, err := os.Create(filePath)
	if err != nil {
		logger.Fatalf("Error creating file '%s': %v", filePath, err)
	}
	defer file.Close()

	for _, line := range f.Content {
		file.WriteString(line + "\n")
	}
	logger.Infof("Created %s", filePath)
}
