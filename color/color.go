package main

import (
	"fmt"
	"strconv"
)

// Foreground colors.
const (
	ColorBlack Color = iota + 30
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

// Color represents a text color.
type Color uint8

// Add adds the coloring to the given string.
func (c Color) Add(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(c), s)
}


func main() {
	colors := []Color{
		ColorBlack,
		ColorRed,
		ColorBlue,
		ColorCyan,
		ColorGreen,
		ColorMagenta,
		ColorWhite,
		ColorYellow,
	}

	for _, color := range colors {
		fmt.Println(Color(color).Add("Hello Color" + strconv.Itoa(int(color))))

	}
}