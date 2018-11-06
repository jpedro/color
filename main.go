// Binary for colouring output
//
// USAGE
//    c <color> <args>
//
// EXAMPLES
//    c green "Hello green world!"
//    c 210 WELCOPME TO MY PINK UNIVERSE
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
var colorCode string = "38;5;%s;1"

func main() {
	color := os.Args[1]
	code := colors[color]

	if code == "" {
		code = fmt.Sprintf(colorCode, color)
	}

	fmt.Printf("\x1b[%sm%s\x1b[0m\n", code, strings.Join(os.Args[2:], " "))
}
