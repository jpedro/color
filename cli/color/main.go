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
	var cmd string
	var arg string

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
		cmd = os.Args[1]
		if cmd == "--version" {
			fmt.Printf("%s\n", VERSION)
			os.Exit(0)
		}

		arg = cmd
		cmd = fallback
		fmt.Printf("CMD %s.\n", cmd)

	default:
		cmd = os.Args[1]
		arg = strings.Join(os.Args[2:], " ")
	}

	fmt.Printf("%s\n", color.Paint(cmd, arg))
}
