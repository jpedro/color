// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
package color

import (
    "testing"
    "os"
)

func TestPaintName(t *testing.T) {
    expected := "\x1b[32;1mHello\x1b[0m"
    returned := Paint("green", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintInteger(t *testing.T) {
    expected := "\x1b[38;5;27mHello\x1b[0m"
    returned := Paint("27", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintTermName(t *testing.T) {
    expected := "\x1b[38;2;0;95;95mHello\x1b[0m"
    returned := Paint("@DeepSkyBlue1", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintRgb1(t *testing.T) {
    expected := "\x1b[38;2;255;0;255mHello\x1b[0m"
    returned := Paint("#f0f", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintRgb2(t *testing.T) {
    expected := "\x1b[38;2;255;0;255mHello\x1b[0m"
    returned := Paint("#ff00ff", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintFail(t *testing.T) {
    expected := "\x1b[32;1mHello\x1b[0m"
    returned := Paint("fail", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}

func TestPaintFallback(t *testing.T) {
    os.Setenv("COLOR_FALLBACK", "yellow")
    expected := "\x1b[33;1mHello\x1b[0m"
    returned := Paint("fail", "Hello")
    if expected != returned {
        t.Error("Expected", expected, "got", returned)
    }
}
