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
	"fmt"
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

func TestCodeFromName(t *testing.T) {
	expected := "32;1"
	returned := CodeFromName("green")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestCodeFromNumber(t *testing.T) {
	expected := "38;5;76"
	returned := CodeFromNumber("76")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestHex2RgbNil(t *testing.T) {
	returned := Hex2Rgb("#wrong")
	if returned != nil {
		t.Error("Expected nil", "got", returned)
	}
}

func TestHex2Rgb6(t *testing.T) {
	expected := Rgb{255, 0, 255}
	returned := Hex2Rgb("#ff00ff")
	if expected.R != returned.R || expected.G != returned.G || expected.B != returned.B {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestHex2Rgb3(t *testing.T) {
	expected := Rgb{255, 0, 255}
	returned := Hex2Rgb("#f0f")
	if expected.R != returned.R || expected.G != returned.G || expected.B != returned.B {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestCodeFromHex(t *testing.T) {
	expected := fmt.Sprintf("%s;%d;%d;%d", hexed, 255, 0, 255)
	returned := CodeFromHex("#ff00ff")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestCodeFromRgb(t *testing.T) {
	rgb := Rgb{255, 0, 255}
	expected := fmt.Sprintf("%s;%d;%d;%d", hexed, 255, 0, 255)
	returned := CodeFromRgb(rgb)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestCodeName(t *testing.T) {
	expected := "32;1"
	returned := Code("green")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithName(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Paint("green", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithNumber(t *testing.T) {
	expected := escape + "38;5;76mHello" + reset
	returned := Paint("76", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithHex6(t *testing.T) {
	expected := escape + "38;2;255;0;0mHello" + reset
	returned := Paint("#ff0000", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithHex3(t *testing.T) {
	expected := escape + "38;2;255;0;0mHello" + reset
	returned := Paint("#f00", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithExtra(t *testing.T) {
	expected := escape + "38;2;255;0;255mHello" + reset
	returned := Paint("@Pink", "Hello")
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

func TestPaintWithArgs(t *testing.T) {
	expected := escape + "32;1mHello world!" + reset
	returned := Paint("green", "Hello %s!", "world")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintWithArgs2(t *testing.T) {
	expected := escape + "32;1mHello world -2!" + reset
	returned := Paint("green", "Hello %s %d!", "world", -2)
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPaintStruct(t *testing.T) {
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

// func TestPaintInteger(t *testing.T) {
// 	expected := escape + "38;5;27m123" + reset
// 	returned := Paint("27", 123)
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// }

// func TestPaintWithExtra(t *testing.T) {
// 	expected := escape + "38;2;0;95;95mHello" + reset
// 	returned := Paint("@DeepSkyBlue1", "Hello")
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// }

func TestPaintTermFail(t *testing.T) {
	expected := escape + "32;1mHello" + reset
	returned := Paint("@xxx", "Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

// func TestPaintRgb1(t *testing.T) {
// 	expected := escape + "38;2;255;0;255mHello" + reset
// 	returned := Paint("#f0f", "Hello")
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// }

// func TestPaintRgb2(t *testing.T) {
// 	expected := escape + "38;2;255;0;255mHello" + reset
// 	returned := Paint("#ff00ff", "Hello")
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// }

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

func TestCyan(t *testing.T) {
	expected := escape + "36;1mHello" + reset
	returned := Cyan("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestBlue(t *testing.T) {
	expected := escape + "34;1mHello" + reset
	returned := Blue("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestMagenta(t *testing.T) {
	expected := escape + "35;1mHello" + reset
	returned := Magenta("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestGray(t *testing.T) {
	expected := escape + "38;5;242mHello" + reset
	returned := Gray("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestPale(t *testing.T) {
	expected := escape + "38;5;246mHello" + reset
	returned := Pale("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestParse(t *testing.T) {
	expected := "Hello " + escape + "31;1mFAIL" + reset + " and " + escape + "32;1mPASS" + reset + "!"
	returned := Parse("Hello {red|FAIL} and {green|PASS}!")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
	// fmt.Println(returned)
}

// func TestParseArgs(t *testing.T) {
// 	// expected := fmt.Sprintf("Hello %s31;1mFAIL%s!", escape, reset)
// 	// returned := Parse("Hello {red|%s}!", "FAIL")
// 	expected := fmt.Sprintf("Hello %s31;1mFAIL%s and %s32;1mPASS%s!", escape, reset, escape, reset)
// 	returned := Parse("Hello {red|%s} and {green|%s}!", "FAIL", "PASS")
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// 	// fmt.Println(returned)
// }

// func TestParseArgsNumber(t *testing.T) {
// 	// expected := "Hello " + escape + "31;1mFAIL" + reset + " and " + escape + "38;5;76mPASS" + reset + "!"
// 	expected := fmt.Sprintf("Hello %s31;1mFAIL%s and %s38;5;76mPASS%s!", escape, reset, escape, reset)
// 	returned := Parse("Hello {red|%s} and {76|%s}!", "FAIL", "PASS")
// 	if expected != returned {
// 		t.Error("Expected", expected, "got", returned)
// 	}
// }

func TestNewEmpty(t *testing.T) {
	c := New()
	expected := "Hello"
	returned := c.Paint("Hello")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}

func TestNewCustom(t *testing.T) {
	c := New().Foreground("230").Background("208").Bold().Underline()
	expected := escape + "38;5;230;48;5;208;1;4mTesting blue." + reset
	returned := c.Paint("Testing %s.", "blue")
	if expected != returned {
		t.Error("Expected", expected, "got", returned)
	}
}
