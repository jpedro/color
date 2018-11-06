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
