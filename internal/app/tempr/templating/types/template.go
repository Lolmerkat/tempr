package templating

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/goccy/go-yaml"
	"github.com/lolmerkat/tempr/internal/app/tempr/flags"
)

type Template struct {
	Name      string      `yaml:"name"`
	Author    string      `yaml:"author,omitempty"`
	Version   string      `yaml:"version,omitempty"`
	Languages []string    `yaml:"languages,omitempty"`
	Content   []FSElement `yaml:"content"`
}

type rawTemplate struct {
    Name      string      `yaml:"name"`
    Author    string      `yaml:"author,omitempty"`
    Version   string      `yaml:"version,omitempty"`
    Languages []string    `yaml:"languages,omitempty"`
    Content   []rawFSElement `yaml:"content"`
}

func (t Template) Expand(path string, logger *log.Logger) {
    templateDirPath := fmt.Sprintf("%s/%s", path, t.Name)
    err := os.Mkdir(templateDirPath, os.ModePerm)
    if err != nil && !os.IsExist(err) {
        logger.Errorf("Error creating project directory '%s': %v", templateDirPath, err)
    }

	if !*flags.DisableInfoFilePtr {
		t.CreateInfoFile(templateDirPath, logger)
	}

    for _, e := range t.Content {
        e.Expand(templateDirPath, logger)
    }
}

func FromYamlBytes(bytes []byte, logger *log.Logger) Template {
    var raw rawTemplate
    err := yaml.Unmarshal(bytes, &raw)
    if err != nil {
        logger.Fatalf("Error unmarshalling raw template: %v", err)
    }

    template := Template{}
    template.Name = raw.Name
    template.Author = raw.Author
    template.Version = raw.Version
    template.Languages = raw.Languages
    template.Content = make([]FSElement, 0, len(raw.Content))

    for _, c := range raw.Content {
        template.Content = append(template.Content, c.ToFSElement(logger))
    }
    return template
}

func LoadFromFile(path string, logger *log.Logger) Template {
    bytes, err := os.ReadFile(path)
    if err != nil {
        logger.Fatalf("Error reading file '%s': %v", path, err)
    }
    return FromYamlBytes(bytes, logger)
}

func (t Template) ToYamlBytes(logger *log.Logger) []byte {
    yamlBytes, err := yaml.MarshalWithOptions(t)
    if err != nil {
        logger.Fatalf("Error mashalling template '%s': %v", t.Name, err)
    }
    return yamlBytes
}

func (t Template) WriteToFile(path string, logger *log.Logger) {
    bytes := t.ToYamlBytes(logger)
    filePath := fmt.Sprintf("%s/%s.yml", path, t.Name)
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0664)
    if err != nil {
        logger.Fatalf("Error creating template file '%s': %v", filePath, err)
    }
    defer file.Close()

    n, err := file.Write(bytes)
    if err != nil {
        logger.Fatalf("Error writing template '%s' to '%s': %v", t.Name, filePath, err)
    }

    // Check if all bytes were written
    if len(bytes) != n {
        logger.Warnf(
            "Template '%s' not written completely to '%s' (only %d out of %d bytes)",
            t.Name, filePath, n, len(bytes))
    } else {
        logger.Infof("Template '%s' successfully written to '%s'", t.Name, filePath)
    }
}

func (t Template) GenerateInfoYaml(logger *log.Logger) ([]byte, error) {
	var infoFileData struct {
		Name      string      `yaml:"name"`
		Author    string      `yaml:"author,omitempty"`
		Version   string      `yaml:"version,omitempty"`
	}
	infoFileData.Name = t.Name
	infoFileData.Author = t.Author
	infoFileData.Version = t.Version

	commentMap := yaml.CommentMap {
		"$": []*yaml.Comment {
			{
				Texts: []string{
					" This project structure was created using github.com/lolmerkat/tempr",
					" This file contains information about the template used.",
					" ",
					" To support the project, keep this file in your codebase",
					" (and potentially commit it to your repository).",
					" If you don't want to, that's fine.",
					" Thank you for using my tool in any way.",
					" ",
				},
				Position: yaml.CommentHeadPosition,
			},
		},
	}

	bytes, err := yaml.MarshalWithOptions(infoFileData, yaml.WithComment(commentMap))
	if err != nil {
		logger.Warnf("Error creating info file yaml bytes: %v", err)
		return nil, err
	}
	return bytes, nil
}

func (t Template) CreateInfoFile(path string, logger *log.Logger) error {
	// get info file bytes
	infoFileBytes, err := t.GenerateInfoYaml(logger)
	if err != nil {
		return err
	}

	// create file
	filePath := fmt.Sprintf("%s/.tempr", path)
	file, err := os.Create(filePath)
	if err != nil {
		logger.Warnf("Error creating info file: %v", err)
		return err
	}
	defer file.Close()

	// write bytes
	n, err := file.Write(infoFileBytes)
	if err != nil {
		logger.Warnf("Error writing info file: %v", err)
		return err
	}

	// check if all bytes were written
	if len(infoFileBytes) != n {
		logger.Warnf(
			"Info file '%s' not written completely (only %d out of %d bytes)",
			filePath, n, len(infoFileBytes))
	} else {
		logger.Infof("Info file '%s' written successfully", filePath)
	}
	return nil
}
