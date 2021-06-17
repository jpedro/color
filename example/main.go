package main

import (
	"fmt"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Green("This is green"))
	fmt.Println(color.Parse("This is {green:green} but this {yellow:yellow}"))
}
