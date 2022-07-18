package main

import (
	"fmt"
	"runtime"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println("Using color.Paint():")
	fmt.Println(color.Paint("orange", "Hello %s!", "Universe"))
	fmt.Println(color.Paint("33", "This is %s!", "blue"))
	fmt.Println(color.Parse("This is {green|green} but this {220|yellow}."))

	fmt.Println()
	fmt.Println("Using color.New():")
	fmt.Println(color.New().Background("27").Bold().Paint(" With a background "))
	fmt.Println(color.New().Paint("Vanilla color"))

	fmt.Println()
	fmt.Println(color.Gray("runtime.GOOS:    %s", runtime.GOOS))
	fmt.Println(color.Gray("runtime.GOARCH:  %s", runtime.GOARCH))
	fmt.Println(color.Gray("runtime.GOROOT:  %s", runtime.GOROOT()))
	fmt.Println(color.Gray("runtime.NumCPU:  %d", runtime.NumCPU()))
}
