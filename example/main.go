package main

import (
	"fmt"
	"runtime"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Paint("27", "This is %s!", "blue"))
	fmt.Println(color.Parse("This is {green|green} but this {220|yellow}."))

	fmt.Println(color.NewColor().Background("27").Bold().Paint("OK"))
	fmt.Println(color.NewColor().Paint("OK"))

	fmt.Println(color.Green("Runtime OS: %s", runtime.GOOS))
}
