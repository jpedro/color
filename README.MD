# Color

[![Build Status](https://travis-ci.org/jpedro/color.svg?branch=master)](https://travis-ci.org/jpedro/color)
[![GoDoc](https://godoc.org/github.com/jpedro/color?status.svg)](https://godoc.org/github.com/jpedro/color)

Library and example binary for colouring shell output.

[![asciicast](https://asciinema.org/a/I41lfgMMXr2z9QmAjyFmWQYCd.svg)](https://asciinema.org/a/I41lfgMMXr2z9QmAjyFmWQYCd)


## Installation

    go get github.com/jpedro/color/color

## Usage

    color green 'Hello green world!'
    color 208 Welcome to an orange universe
    color "#0c0" "This is also green"
    color "@SpringGreen5" "SpringGreen5 is one of the 256 named term colors"


## Todos

- [x] Create fallback color
- [ ] Complete the names of the 256 term colors
