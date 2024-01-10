// Swine leberkas venison
//
// Burgdoggen sirloin biltong chuck drumstick shank capicola porchetta. Turkey pork loin
// chuck fatback jowl. T-bone short ribs turducken cupim, brisket cow pork belly leberkas.
// Landjaeger ham hock fatback pig corned beef bresaola beef ribs. Pork pork chop boudin
// strip steak landjaeger, pork belly kevin pork loin capicola ham. Pastrami spare ribs
// porchetta, drumstick leberkas t-bone short loin doner filet mignon hamburger corned
// beef. Venison short loin flank, cupim fatback spare ribs pork loin buffalo turducken
// tail.package main
// This package name is just to be different
package color_test

import (
	"os"
	"testing"

	"github.com/jpedro/color"
)

const (
	codeEscape = "\x1b["
	codeReset  = "\x1b[0m"
)

func TestPaintName(t *testing.T) {
	expected := codeEscape + "32;1mHello" + codeReset
	returned := color.Paint("green", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFloat(t *testing.T) {
	expected := codeEscape + "32;1m123.46" + codeReset
	returned := color.Paint("green", 123.456)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintBool(t *testing.T) {
	expected := codeEscape + "32;1mtrue" + codeReset
	returned := color.Paint("green", true)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRune(t *testing.T) {
	expected := codeEscape + "32;1mr" + codeReset
	returned := color.Paint("green", 'r')
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintNumber(t *testing.T) {
	expected := codeEscape + "32;1m123" + codeReset
	returned := color.Paint("green", 123)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintNumbers(t *testing.T) {
	expected := codeEscape + "32;1m123" + codeReset
	returned := color.Paint("green", 123)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithArgs(t *testing.T) {
	expected := codeEscape + "32;1mHello world!" + codeReset
	returned := color.Paint("green", "Hello %s!", "world")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithArgs2(t *testing.T) {
	expected := codeEscape + "32;1mHello world -2!" + codeReset
	returned := color.Paint("green", "Hello %s %d!", "world", -2)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTextWithStruct(t *testing.T) {
	s := struct {
		yes      bool
		text     string
		number   int
		decimals float64
	}{
		yes:      true,
		text:     "ok",
		number:   123,
		decimals: 123.4,
	}
	expected := codeEscape + "32;1mHello {true ok 123 123.4}!" + codeReset
	returned := color.Paint("green", "Hello %v!", s)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintInteger(t *testing.T) {
	expected := codeEscape + "38;5;27mHello" + codeReset
	returned := color.Paint("27", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTermName(t *testing.T) {
	expected := codeEscape + "38;2;0;95;95mHello" + codeReset
	returned := color.Paint("@DeepSkyBlue1", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintTermFail(t *testing.T) {
	expected := codeEscape + "32;1mHello" + codeReset
	returned := color.Paint("@xxx", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRgb1(t *testing.T) {
	expected := codeEscape + "38;2;255;0;255mHello" + codeReset
	returned := color.Paint("#f0f", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintRgb2(t *testing.T) {
	expected := codeEscape + "38;2;255;0;255mHello" + codeReset
	returned := color.Paint("#ff00ff", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFail(t *testing.T) {
	expected := codeEscape + "32;1mHello" + codeReset
	returned := color.Paint("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintFallback(t *testing.T) {
	os.Setenv("COLOR_FALLBACK", "yellow")
	expected := codeEscape + "33;1mHello" + codeReset
	returned := color.Paint("fail", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestGreen(t *testing.T) {
	expected := codeEscape + "32;1mHello" + codeReset
	returned := color.Green("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestYellow(t *testing.T) {
	expected := codeEscape + "33;1mHello" + codeReset
	returned := color.Yellow("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestRed(t *testing.T) {
	expected := codeEscape + "31;1mHello" + codeReset
	returned := color.Red("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestParse(t *testing.T) {
	expected := "Hello " + codeEscape + "31;1mFAIL" + codeReset + " and " + codeEscape + "32;1mPASS" + codeReset + "!"
	returned := color.Format("Hello {red|FAIL} and {green|PASS}!")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestParseArgs(t *testing.T) {
	expected := "Hello " + codeEscape + "31;1mFAIL" + codeReset + " and " + codeEscape + "38;5;27mPASS" + codeReset + "!"
	returned := color.Format("Hello {red|%s} and {27|%s}!", "FAIL", "PASS")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestParseArgsNumber(t *testing.T) {
	expected := "Hello " + codeEscape + "31;1mFAIL" + codeReset + " and " + codeEscape + "38;5;27mPASS" + codeReset + "!"
	returned := color.Format("Hello {red|%s} and {27|%s}!", "FAIL", "PASS")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestNew(t *testing.T) {
	color := color.New().
		Foreground("230").
		Background("208").
		Bold().
		Blink().
		Underline()

	expected := codeEscape + "38;5;230;48;5;208;1;4;5mTesting blue." + codeReset
	returned := color.Paint("Testing %s.", "blue")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}
