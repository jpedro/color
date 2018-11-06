package main

import (
	"fmt"
	"os"
	"strings"
)

var colors = map[string]string{
	"red":    "31;1",
	"green":  "32;1",
	"yellow": "33;1",
	"blue":   "38;5;27;1",
}
var colorCode = "38;5;%s;1"

func main() {
	if len(os.Args) < 3 {
		return
	}

	color := os.Args[1]

	fmt.Printf("%s\n", colorize(color, strings.Join(os.Args[2:], " ")))
}

func colorize(color string, text string) string {
	code := colors[color]

	if code == "" {
		code = fmt.Sprintf(colorCode, color)
	}

	return fmt.Sprintf("\x1b[%sm%s\x1b[0m", code, text)
}
