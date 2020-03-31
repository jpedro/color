package main

import (
    "os"
    "fmt"
    "strings"

    "github.com/jpedro/color"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: $0 <color> <text>")
        return
    }

    name := os.Args[1]
    text := strings.Join(os.Args[2:], " ")

    fmt.Printf("%s\n", color.Paint(name, text))
}
