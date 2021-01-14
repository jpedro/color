package main

import (
    "fmt"

    "github.com/jpedro/color"
)

func main() {
    name := "green"
    text := "hello"
    fmt.Println(color.Paint(name, text))
}
