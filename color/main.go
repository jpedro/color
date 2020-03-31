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

    if len(os.Args) < 2 {
        fmt.Println("Usage: $0 [color] <text>")
        return
    }

    if len(os.Args) == 2 {
        fallback := os.Getenv("COLOR_FALLBACK")
        if fallback == "" {
            fallback = "green"
        }
        name = fallback
        text = strings.Join(os.Args[1:], " ")

    } else {
        name = os.Args[1]
        text = strings.Join(os.Args[2:], " ")
    }

    fmt.Printf("%s\n", color.Paint(name, text))
}
