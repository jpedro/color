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
)

func TestColorizeName(t *testing.T) {
	expected := "\x1b[32;1mHello\x1b[0m"
	returned := Colorize("green", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeInteger(t *testing.T) {
	expected := "\x1b[38;5;27mHello\x1b[0m"
	returned := Colorize("27", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeRgb1(t *testing.T) {
	expected := "\x1b[38;2;255;0;255mHello\x1b[0m"
	returned := Colorize("#f0f", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeRgb2(t *testing.T) {
	expected := "\x1b[38;2;255;0;255mHello\x1b[0m"
	returned := Colorize("#ff00ff", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeFail(t *testing.T) {
	expected := "Hello"
	returned := Colorize("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}
