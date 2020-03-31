# Color

[![Build Status](https://travis-ci.org/jpedro/color.svg?branch=master)](https://travis-ci.org/jpedro/color)
[![GoDoc](https://godoc.org/github.com/jpedro/color?status.svg)](https://godoc.org/github.com/jpedro/color)

Binary for colouring shell output.


## Installation

    go get github.com/jpedro/color/color

## Usage

    color green 'Hello green world!'
    color 208 WELCOME TO MY ORANGE UNIVERSE
    color "#0c0" "This is also green"
    color "@SpringGreen5" "SpringGreen5 is apparently one of the 256 named Term colors"


## Todos

- [x] Create fallback color
- [ ] Complete the names of the 256 term colors
