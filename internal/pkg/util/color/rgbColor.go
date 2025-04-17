package color

import "fmt"

type RGBColor struct {
	Red   int
	Green int
	Blue  int
}

func ByRGB(red int, green int, blue int) Color {
	return Color{
		RGB: RGBColor{
			Red:   red,
			Green: green,
			Blue:  blue,
		},
	}
}

func (c RGBColor) IntTuple() (red int, green int, blue int) {
	return c.Red, c.Green, c.Blue
}

func (start RGBColor) ColorStepsTo(end RGBColor, steps int) []Color {
	redStepSize := (end.Red - start.Red) / (steps + 1)
	greenStepSize := (end.Green - start.Green) / (steps + 1)
	blueStepSize := (end.Blue - start.Blue) / (steps + 1)

	colors := make([]Color, steps)
	for i := range steps {
		red := start.Red + (i * redStepSize)
		green := start.Green + (i * greenStepSize)
		blue := start.Blue + (i * blueStepSize)
		colors[i] = ByRGB(red, green, blue)
	}

	return colors
}

func (c RGBColor) ToHexColor() HexColor {
	r, g, b := c.IntTuple()
	hexString := fmt.Sprintf("#%02X%02X%02X", r, g, b)
	return HexColor {
		HexString: hexString,
	}
}
