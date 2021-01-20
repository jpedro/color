# Color

[![Travis Status](https://travis-ci.org/jpedro/color.svg?branch=master)](https://travis-ci.org/jpedro/color)
[![Github Status](https://github.com/jpedro/color/workflows/main/badge.svg)](https://github.com/jpedro/color/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/color?status.svg)](https://godoc.org/github.com/jpedro/color)

Go library to add shell color to text.


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
    fmt.Println(color.Paint(name, text))
}
```

Check the [cli/color](cli/color) for the cli tool usage.


## Todos

- [x] Create fallback color
- [ ] Name all the 256 term colors
