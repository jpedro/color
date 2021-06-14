package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpedro/color"
)

func main() {
	var name string
	var text string

	fallback := os.Getenv("COLOR_FALLBACK")
	if fallback == "" {
		fallback = "green"
	}

	switch {
	case len(os.Args) < 2:
		fmt.Printf("Usage: color [color=%s] <text>\n", fallback)
		return

	case len(os.Args) == 2:
		name = fallback
		text = strings.Join(os.Args[1:], " ")

	default:
		name = os.Args[1]
		text = strings.Join(os.Args[2:], " ")
	}

	fmt.Printf("%s\n", color.Paint(name, text))
}
