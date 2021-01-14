# Color

Go cli to print colored text.

![Example](https://raw.githubusercontent.com/jpedro/color/master/.github/img/cast.gif "Example")


## Installation

    $ go get github.com/jpedro/color/cli/color

## Usage

    # Basic shell named color
    $ color green 'Hello green world!'

    # Use an extended shell int color code
    $ color 208 Welcome to an orange universe

    # Use the closest shell for an HTML color
    $ color "#0c0" "This is also green... 'ish"

    # One of the 256 extra named colors
    $ color "@SpringGreen5" "SpringGreen5 is one of the 256 named term colors"
