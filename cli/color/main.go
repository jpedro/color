package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jpedro/color"
)

const (
	VERSION string  = "0.1.2"
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
		os.Exit(1)
		return

	case len(os.Args) == 2:
		name = fallback
		text = strings.Join(os.Args[1:], " ")

	default:
		name = os.Args[1]
		text = strings.Join(os.Args[2:], " ")
	}

	if name == "--version" {
		fmt.Printf("%s.\n", VERSION)
	}

	fmt.Printf("%s\n", color.Paint(name, text))
}
