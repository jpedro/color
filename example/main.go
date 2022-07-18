package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/jpedro/color"
)

func main() {
	fmt.Println("Using color.Paint()")
	fmt.Println(color.Paint("orange", "   Hello %s!", "Universe"))
	fmt.Println(color.Paint("33", "   This is %s!", "blue"))
	fmt.Println(color.Parse("   This is {green|green} but this {220|yellow}."))

	fmt.Println()
	fmt.Println("Using color.New()")
	fmt.Println(color.New().Background("27").Bold().Paint("   With a background "))
	fmt.Println(color.New().Paint("   Default color"))

	fmt.Println()
	fmt.Println("Runtime")
	report("GOOS", runtime.GOOS)
	report("GOARCH", runtime.GOARCH)
	report("GOROOT", runtime.GOROOT())
	report("NumCPU", strconv.Itoa(runtime.NumCPU()))
	report("NumCPU", runtime.NumCPU())
}

func report(header, value any) {
	fmt.Printf("    %-24s %s\n", color.Gray(header), color.Green(value))
}
