// package main

// import (
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     rgb := os.Args[1]
//     hex := Hex(rgb)
//     fmt.Printf("%s %s \033[38;2;%sm%s\033[0m", rgb, hex, hex, strings.Join(os.Args[2:], " "))
// }

// func Hex(rgb string) string {
//     var r, g, b uint8
//     hexFormat := "#%02x%02x%02x"
//     fmt.Sscanf(rgb, hexFormat, &r, &g, &b)
//     return fmt.Sprintf("%v;%v;%v", r, g, b)
// }
