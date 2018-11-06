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

// Bacon ipsum dolor amet corned
// beef short loin sausage ground round venison pig. Sirloin bresaola ham meatloaf
// leberkas landjaeger. Rump jowl cow turkey, shoulder andouille filet mignon
// chicken tail porchetta. Tail pork chop strip steak, andouille tenderloin short
// ribs alcatra. Turkey frankfurter ham hock boudin. Pork belly capicola hamburger
// ham hock burgdoggen fatback pancetta swine picanha turducken landjaeger pastrami
// shank shankle shoulder.
func main() {
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
