package templating

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Directory struct {
	Name     string      `yaml:"string"`
	Children []FSElement `yaml:"children"`
}

func (d *Directory) Expand(path string, logger *log.Logger) {
	dirPath := fmt.Sprintf("%s/%s", path, d.Name)
	err := os.Mkdir(dirPath, os.ModePerm)
	if err != nil && !os.IsExist(err){
		logger.Fatalf("Error creating dir '%s': %v", dirPath, err)
	}

	for _, child := range d.Children {
		child.Expand(dirPath, logger)
	}
	logger.Infof("Created %s", dirPath)
}
