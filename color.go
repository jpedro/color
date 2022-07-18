package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	// "encoding/hex"
)

const (
	escape 	 string = "\x1b["
	reset  	 string = "\x1b[0m"
	extended string = "38;5"
	hexed    string = "38;2"
)

var (
	// regexRgb6  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	// regexRgb3  = regexp.MustCompile(`^#[0-9a-f]{3}$`)
	// regexRgb   = regexp.MustCompile(`^#[0-9a-f]{3,6}$`)
	regexParse = regexp.MustCompile(`{([^\}]+)}`)
	regexGroup = regexp.MustCompile(`([^\|]+)\|(.+)`)
	fallback   = os.Getenv("COLOR_FALLBACK")
)

type Color struct {
	foreground string
	background string
	bold       bool
	underline  bool
	format     string
}

type Rgb struct {
	R uint8
	G uint8
	B uint8
}

func New() *Color {
	if fallback == "" {
		fallback = "green"
	}

	color := &Color{
		// foreground: fallback,
	}
	// fmt.Printf("Color %v.\n", color)
	return color
}

// Select a color for your foreground
func (c *Color) Foreground(text string) *Color {
	c.foreground = text
	return c
}

// Select a color for your background
func (c *Color) Background(text string) *Color {
	c.background = text
	return c
}

// Turns the text bold
func (c *Color) Bold() *Color {
	c.bold = true
	return c
}

// Turns the text underlined
func (c *Color) Underline() *Color {
	c.underline = true
	return c
}

// Paints text and args according to its settings
func (c *Color) Paint(text any, args ...any) string {
	message := getText(text, args...)

	if c.format == "" {
		format := ""

		if c.foreground != "" {
			format = fmt.Sprintf("%s;38;5;%s", format, c.foreground)
		}

		if c.background != "" {
			format = fmt.Sprintf("%s;48;5;%s", format, c.background)
		}

		if c.bold {
			format = fmt.Sprintf("%s;1", format)
		}

		if c.underline {
			format = fmt.Sprintf("%s;4", format)
		}

		if format == "" {
			return message
		}

		c.format = format[1:]
	}

	return fmt.Sprintf("%s%sm%s%s", escape, c.format, message, reset)
}

// Returns a painted string with groups like `{green|this should be green}`
// replaced with `this should be green` in green color
func Parse(text string, args ...any) string {
	replace := text
	matches := regexParse.FindAllStringSubmatch(text, -1)
	for _, found := range matches {
		groups := regexGroup.FindAllStringSubmatch(found[1], -1)
		color := groups[0][1]
		group := groups[0][2]
		colored := Paint(color, group)
		replace = strings.Replace(replace, "{"+color+"|"+group+"}", colored, 1)
	}

	message := getText(replace, args...)
	return message
}


func Code(color string) string {
	fallback := os.Getenv("COLOR_FALLBACK")
	if fallback == "" {
		fallback = "green"
	}

	// Check if we have a numeric color
	number, _ := strconv.Atoi(color)
	text := strconv.Itoa(number)
	if text == color {
		return CodeFromNumber(color)
	}

	// Check if we have a term named color, like "@CornflowerBlue"
	if color[0] == '@' {
		termColor, ok := termNames[color[1:]]
		if ok {
			return CodeFromHex(termColor)
		}
		return CodeFromName(fallback)
	}

	// Lower the case for the next rounds
	color = strings.ToLower(color)

	if color[0] == '#' {
		return CodeFromHex(color)
	}

	// Check if we have a basic color name
	if _, ok := basicNames[color]; ok {
		return CodeFromName(color)
	}

	return CodeFromName(fallback)
}

// Returns shell coloured output for text and args
func Paint(color string, text any, args ...any) string {
	message := getText(text, args...)
	code := Code(color)

	return fmt.Sprintf("%s%sm%s%s", escape, code, message, reset)
}

// Shortcut for color.Paint("green", text)
func Green(text any, args ...any) string {
	return Paint("green", text, args...)
}

// Shortcut for color.Paint("yellow", text)
func Yellow(text any, args ...any) string {
	return Paint("yellow", text, args...)
}

// Shortcut for color.Paint("red", text)
func Red(text any, args ...any) string {
	return Paint("red", text, args...)
}

// Shortcut for color.Paint("cyan", text)
func Cyan(text any, args ...any) string {
	return Paint("cyan", text, args...)
}

// Shortcut for color.Paint("blue", text)
func Blue(text any, args ...any) string {
	return Paint("blue", text, args...)
}

// Shortcut for color.Paint("magenta", text)
func Magenta(text any, args ...any) string {
	return Paint("magenta", text, args...)
}

// Shortcut for color.Paint("gray", text)
func Gray(text any, args ...any) string {
	return Paint("gray", text, args...)
}

// Shortcut for color.Paint("pale", text)
func Pale(text any, args ...any) string {
	return Paint("pale", text, args...)
}

// Concatenates text and args
func getText(text any, args ...any) string {
	message := ""

	switch value := text.(type) {
	case string:
		message = value
	case uint, int:
	case int8, uint8:
	case int16, uint16:
	case int32, uint32:
	case int64, uint64:
		message = fmt.Sprintf("%d", value)
	case float32, float64:
		message = fmt.Sprintf("%0.2f", value)
	case bool:
		message = fmt.Sprintf("%t", value)
	default:
		message = fmt.Sprintf("%v", value)
	}

	if len(args) > 0 {
		// fmt.Printf("ARGS %v\n", args)
		message = fmt.Sprintf(message, args...)
	}

	return message
}

// Returns an (r, g, b) tupple for an html hex color
func Hex2Rgb(hex string) *Rgb {
	len := len(hex)

	if len != 4 && len != 7 {
		return nil
	}

	if len == 4 {
		hex = fmt.Sprintf(
			"#%s%s%s%s%s%s",
			string(hex[1]),
			string(hex[1]),
			string(hex[2]),
			string(hex[2]),
			string(hex[3]),
			string(hex[3]),
		)
	}

	rgb := Rgb{}
	format := "#%02x%02x%02x"
	fmt.Sscanf(hex, format, &rgb.R, &rgb.G, &rgb.B)
	return &rgb

	// val, err := hex.DecodeString(color[1:])
    // if err != nil {
    //     fmt.Printf("Error: %v", err)
	// 	return nil
    // }

    // return &Rgb{val[0], val[1], val[2]}
}

// Returns the closest shell colour string from an html hex color (#rrggbb)
// func FromHex(color string, text string) string {
func CodeFromRgb(rgb Rgb) string {
	return fmt.Sprintf("%s;%d;%d;%d", hexed, rgb.R, rgb.G, rgb.B)
}

func CodeFromHex(color string) string {
	rgb := Hex2Rgb(color)
	return fmt.Sprintf("%s;%d;%d;%d", hexed, rgb.R, rgb.G, rgb.B)
}

// Returns the selected basic named color
// func CodeFromName(color string, text string) string {
func CodeFromName(color string) string {
	code := basicNames[color]
	return code
	// if code == "" {
	//  code = fmt.Sprintf(defaultCode, color)
	// }
	// return fmt.Sprintf(escape+"%sm", code)
}

// Returns the selected numeric color
// func FromNumber(color string, text string) string {
func CodeFromNumber(color string) string {
	return fmt.Sprintf("%s;%s", extended, color)
}
