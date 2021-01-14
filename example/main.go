package main

import (
    "fmt"

    "github.com/jpedro/color"
)

func main() {
	name := "green"
	text := "hello"
    fmt.Printf("%s\n", color.Paint(name, text))
}
