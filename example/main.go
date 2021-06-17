package main

import (
	"fmt"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Paint("218", "Also %s!", "orange"))
	fmt.Println(color.Parse("This is {green|green} but this {27|yellow}."))
}
