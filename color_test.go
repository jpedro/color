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
	"os"
	"testing"
)

type testData struct {
	yes      bool
	text     string
	number   int
	decimals float64
}

// func (d data) String() string {
// 	return fmt.Sprintf("%s %t", d.text, d.yes)
// }

func TestPaintName(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Paint("green", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFloat(t *testing.T) {
	expected := escape + "32;1m123.46" + reset
	returned := Paint("green", 123.456)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintBool(t *testing.T) {
	expected := escape + "32;1mtrue" + reset
	returned := Paint("green", true)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRune(t *testing.T) {
	expected := escape + "32;1mr" + reset
	returned := Paint("green", 'r')
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintNumber(t *testing.T) {
	expected := escape + "32;1m123" + reset
	returned := Paint("green", 123)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintNumbers(t *testing.T) {
	expected := escape + "32;1m123" + reset
	returned := Paint("green", 123)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithArgs(t *testing.T) {
	expected := escape + "32;1mHello world!" + reset
	returned := Paint("green", "Hello %s!", "world")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithArgs2(t *testing.T) {
	expected := escape + "32;1mHello world -2!" + reset
	returned := Paint("green", "Hello %s %d!", "world", -2)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithStruct(t *testing.T) {
	d := testData{
		yes:      true,
		text:     "ok",
		number:   123,
		decimals: 123.4,
	}
	expected := escape + "32;1mHello {true ok 123 123.4}!" + reset
	returned := Paint("green", "Hello %v!", d)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintInteger(t *testing.T) {
	expected := escape + "38;5;27mHello" + reset
	returned := Paint("27", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTermName(t *testing.T) {
	expected := escape + "38;2;0;95;95mHello" + reset
	returned := Paint("@DeepSkyBlue1", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTermFail(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Paint("@xxx", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRgb1(t *testing.T) {
	expected := escape + "38;2;255;0;255mHello" + reset
	returned := Paint("#f0f", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRgb2(t *testing.T) {
	expected := escape + "38;2;255;0;255mHello" + reset
	returned := Paint("#ff00ff", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFail(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Paint("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFallback(t *testing.T) {
	os.Setenv("COLOR_FALLBACK", "yellow")
	expected := escape + "33;1mHello" + reset
	returned := Paint("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestGreen(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Green("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestYellow(t *testing.T) {
	expected := escape + "33;1mHello" + reset
	returned := Yellow("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestRed(t *testing.T) {
	expected := escape + "31;1mHello" + reset
	returned := Red("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}
