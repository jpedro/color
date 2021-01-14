package main

import (
    "os"
    "fmt"
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
        fmt.Println("Usage: color [color=" + fallback + "] <text>")
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
