// Bacon ipsum dolor amet corned
//
// This text shall ne hereby ignored by https://godoc.org/
//
// Beef short loin sausage ground round venison pig. Sirloin bresaola ham
// meatloaf leberkas landjaeger. Rump jowl cow turkey, shoulder andouille filet
// mignon chicken tail porchetta. Tail pork chop strip steak, andouille
// tenderloin short ribs alcatra. Turkey frankfurter ham hock boudin. Pork belly
// capicola hamburger ham hock burgdoggen fatback pancetta swine picanha
// turducken landjaeger pastrami shank shankle shoulder.

/*
Binary for colouring output.


Installation

To install the ``color`` binary run:

    go get github.com/jpedro/color/color


Usage

The first argument is the color, the rest is the text. If the color is an integer
it will be passed to as the __38;5;x__ code.

    color <color> <text...>

You can also pass a named color. The list is:

    - red
    - green
    - yellow
    - blue
    - magenta
    - cyan

Examples

A header needs a text after. Otherwise the code below renders it a normal paragraph.
But here are the examples:

    color 208 WELCOME TO MY ORANGE UNIVERSE
    color green 'Hello green world!'

*/
package main
