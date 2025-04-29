package ui

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/lipgloss"

	"github.com/lolmerkat/tempr"
	"github.com/lolmerkat/tempr/internal/pkg/util/color"
)

type Splash struct {
	Text	string
}

func GetSplash(logger *log.Logger) Splash {
	splash, err := tempr.Files.ReadFile("assets/ascii_title.txt")
	if err != nil {
		log.Errorf("Error when reading file 'assets/ascii_title.txt': %v", err)
	}

	return Splash{
		Text: string(splash),
	}
}

func (s Splash) AsStringArray() []string {
	return strings.Split(s.Text, "\n")
}

func (s Splash) Print(logger *log.Logger) {
	lines := s.AsStringArray()
	start := color.ByRGB(0, 255, 100) //#009664
	end := color.ByRGB(0, 200, 250) //#0096c9
	lineColors := start.RGB.ColorStepsTo(end.RGB, len(lines))

	for i, l := range lines {
		hexColor := lineColors[i].RGB.ToHexColor()
		color := lipgloss.Color(hexColor.HexString)
		style := lipgloss.NewStyle().Foreground(color)
		logger.Print(style.Render(l))
	}
}
