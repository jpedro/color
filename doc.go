/*
Binary for colouring output.


Installation

To install ``c`` just run:
   go get github.com/jpedro/c


Usage

The first argument is the color, the rest is the text. If the color is an integer
it will be passed to as the __38;5;x__ code.
   c <color> <text...>


Examples

A header needs a text after. Otherwise the code below renders it a normal paragraph.
   c green "Hello green world!"
   c 201 WELCOME TO MY PINK UNIVERSE

*/
package main
