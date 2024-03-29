# Color

<!--
[![Travis Status](https://travis-ci.org/jpedro/color.svg?branch=master)](https://travis-ci.org/jpedro/color)
-->
[![Test](https://github.com/jpedro/color/actions/workflows/test.yaml/badge.svg)](https://github.com/jpedro/color/actions/workflows/test.yaml)
[![GoDoc](https://godoc.org/github.com/jpedro/color?status.svg)](https://godoc.org/github.com/jpedro/color)

Color is a library for colouring things.


## Usage

```go
package main

import (
    "fmt"
    "runtime"

    "github.com/jpedro/color"
)

func main() {
    name := "green"
    text := "Runtime OS: %s"
    fmt.Println(color.Paint(name, text, runtime.GOOS))
}
```

Check the [cli/color](cli/color) for your terminal needs.


## Todos

- [x] Create fallback color
- [x] Add color groups. Now via `color.Parse()`
- [x] Add custom color painters. Now via `color.NewColor()`
- [ ] Name all the 256 term colors
