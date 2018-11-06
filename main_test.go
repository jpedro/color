// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
import (
	"testing"
)

func TestColorizeNamed(t *testing.T) {
	expected := "\x1b[32;1mHello\x1b[0m"
	returned := colorize("green", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeInteger(t *testing.T) {
	expected := "\x1b[38;5;27;1mHello\x1b[0m"
	returned := colorize("27", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestColorizeFailk(t *testing.T) {
	expected := "\x1b[38;5;fail;1mHello\x1b[0m"
	returned := colorize("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}
