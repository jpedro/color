package main

import (
	"fmt"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Green("This is green"))
}
