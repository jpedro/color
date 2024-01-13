package color

import (
	"fmt"
)

type Color struct {
	foreground string
	background string
	format     string
	flags      int
	bold       bool // 1
	faint      bool // 2
	italic     bool // 3
	underline  bool // 4
	blink      bool // 5
	reverse    bool // 7
	strike     bool // 9
}

func New(args ...string) *Color {
	color := &Color{}
	if len(args) > 0 {
		color.foreground = args[0]
	}
	if len(args) > 1 {
		color.background = args[1]
	}

	return color
}

func (c *Color) Foreground(text string) *Color {
	c.foreground = text
	return c
}

func (c *Color) Background(text string) *Color {
	c.background = text
	return c
}

func (c *Color) Bold() *Color {
	c.flags = c.flags | (1 << (ColorBold - 1))
	c.bold = true
	return c
}

func (c *Color) Faint() *Color {
	c.flags = c.flags | (1 << (ColorFaint - 1))
	c.faint = true
	return c
}

func (c *Color) Italic() *Color {
	c.flags = c.flags | (1 << (ColorItalic - 1))
	c.italic = true
	return c
}

func (c *Color) Underline() *Color {
	c.flags = c.flags | (1 << (ColorUnderline - 1))
	c.underline = true
	return c
}

func (c *Color) Blink() *Color {
	c.flags = c.flags | (1 << (ColorBlink - 1))
	c.blink = true
	return c
}

func (c *Color) Reverse() *Color {
	c.flags = c.flags | (1 << (ColorReseverd - 1))
	c.reverse = true
	return c
}

func (c *Color) Strike() *Color {
	c.flags = c.flags | (1 << (ColorStrike - 1))
	c.strike = true
	return c
}

// Paints text and args according to its settings
func (c *Color) Paint(message any, args ...any) string {
	text := getText(message, args...)

	if c.format != "" {
		return fmt.Sprintf("%s%sm%s%s", codeEscape, c.format, text, codeReset)
	}

	format := ""

	if c.foreground != "" {
		// format = fmt.Sprintf("%s;38;5;%s", format, c.foreground)
		// format = fmt.Sprintf("%s;38;5;%s", format, getCode(c.foreground))
		// format = "3" + getCode(c.foreground)
		code := getCode(c.foreground)
		format = fmt.Sprintf("%s;3%s", format, code)
	}

	if c.background != "" {
		// format = fmt.Sprintf("%s;48;5;%s", format, c.background)
		code := getCode(c.background)
		format = fmt.Sprintf("%s;4%s", format, code)
	}

	// if c.flags & (1 << (ColorBold - 1)) == ColorBold{
	if c.bold {
		format = fmt.Sprintf("%s;1", format)
	}

	if c.faint {
		format = fmt.Sprintf("%s;2", format)
	}

	if c.italic {
		format = fmt.Sprintf("%s;3", format)
	}

	if c.underline {
		format = fmt.Sprintf("%s;4", format)
	}

	if c.blink {
		format = fmt.Sprintf("%s;5", format)
	}

	if c.reverse {
		format = fmt.Sprintf("%s;7", format)
	}

	if c.strike {
		format = fmt.Sprintf("%s;9", format)
	}

	if format == "" {
		return text
	}

	c.format = format[1:]

	return fmt.Sprintf("%s%sm%s%s", codeEscape, c.format, text, codeReset)
}
