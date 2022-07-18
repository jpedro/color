package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	escape   = "\x1b["
	reset    = "\x1b[0m"
	extended = "38;5;"
	hexed    = "38;2;"
)

var (
	regexRgb6  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	regexRgb3  = regexp.MustCompile(`^#[0-9a-f]{3}$`)
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

func New() *Color {
	if fallback == "" {
		fallback = "green"
	}

	color := &Color{
		foreground: fallback,
	}
	fmt.Printf("Color %v.\n", color)
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
	i, _ := strconv.Atoi(color)
	s := strconv.Itoa(i)
	if s == color {
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

	// Check if we have a basic color name
	if _, ok := basicNames[color]; ok {
		return CodeFromName(color)
	}

	// Finally we do the regex things
	// matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	if regexRgb6.MatchString(color) {
		return CodeFromHex(color)
	}

	if regexRgb3.MatchString(color) {
		htmlColor := fmt.Sprintf(
			"#%s%s%s%s%s%s",
			string(color[1]),
			string(color[1]),
			string(color[2]),
			string(color[2]),
			string(color[3]),
			string(color[3]))
		return CodeFromHex(htmlColor)
	}

	return CodeFromName(fallback)
}

// Returns shell coloured output for text and args
func Paint(color string, text any, args ...any) string {
	message := getText(text, args...)
	code := Code(color)

	return fmt.Sprintf("%s%sm%s%s", escape, code, message, reset)

	// // Check if we have a numeric color
	// i, _ := strconv.Atoi(color)
	// s := strconv.Itoa(i)
	// if s == color {
	// 	return FromNumber(color, message)
	// }

	// // Check if we have a term named color, like "@CornflowerBlue"
	// if color[0] == '@' {
	// 	termColor, ok := termNames[color[1:]]
	// 	if ok {
	// 		return FromHex(termColor, message)
	// 	}
	// 	return FromName(fallback, message)
	// }

	// // Lower the case for the next rounds
	// color = strings.ToLower(color)

	// // Check if we have a basic color name
	// if _, ok := basicNames[color]; ok {
	// 	return FromName(color, message)
	// }

	// // Finally we do the regex things
	// // matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	// if regexRgb6.MatchString(color) {
	// 	return FromHex(color, message)
	// }

	// if regexRgb3.MatchString(color) {
	// 	htmlColor := fmt.Sprintf(
	// 		"#%s%s%s%s%s%s",
	// 		string(color[1]),
	// 		string(color[1]),
	// 		string(color[2]),
	// 		string(color[2]),
	// 		string(color[3]),
	// 		string(color[3]))
	// 	return FromHex(htmlColor, message)
	// }

	// return FromName(fallback, message)
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
	case uint:
	case int:
		message = fmt.Sprintf("%d", value)
	case float32:
	case float64:
		message = fmt.Sprintf("%0.2f", value)
	case bool:
		message = fmt.Sprintf("%t", value)
	case rune:
		// fmt.Printf("RUNE %s\n", text)
		message = string(value)
	case string:
		// fmt.Printf("STRING %s\n", text)
		message = value
	default:
		message = fmt.Sprintf("%v", value)
	}

	if len(args) > 0 {
		// fmt.Printf("ARGS %v\n", args)
		message = fmt.Sprintf(message, args...)
	}

	return message
}

// Returns the closest shell colour string from an html hex color (#rrggbb)
// func FromHex(color string, text string) string {
func CodeFromHex(color string) string {
	r, g, b := Hex2Rgb(color)
	return fmt.Sprintf("%s;%v;%v;%v", hexed, r, g, b)
}

// Returns an (r, g, b) tupple for an html hex color
func Hex2Rgb(color string) (r, g, b uint8) {
	hexFormat := "#%02x%02x%02x"
	fmt.Sscanf(color, hexFormat, &r, &g, &b)
	return
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
