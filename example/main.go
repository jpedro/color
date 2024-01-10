package main

import (
	"fmt"
	"runtime"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Paint("27", "This is %s!", "blue"))
	fmt.Println(color.Format("This is {green|green} but this {220|yellow}."))

	fmt.Println(color.New().Background("27").Bold().Paint("OK"))
	fmt.Println(color.New().Paint("This should be green"))
	fmt.Println(color.New("yellow", "blue").Paint("This should be yellow on blue"))

	fmt.Println(color.Green("Runtime OS: %s", runtime.GOOS))
}
