# Color

[![Build Status](https://travis-ci.org/jpedro/color.svg?branch=master)](https://travis-ci.org/jpedro/color)
[![GoDoc](https://godoc.org/github.com/jpedro/color?status.svg)](https://godoc.org/github.com/jpedro/color)

Go library for adding color to text.


## Usage

```go
package main

import (
    "fmt"

    "github.com/jpedro/color"
)

func main() {
    name := "green"
    text := "hello"
    fmt.Printf("%s\n", color.Paint(name, text))
}
```

Check the [cli/color](cli/color) for the cli tool usage.


## Todos

- [x] Create fallback color
- [ ] Complete the names of the 256 term colors
