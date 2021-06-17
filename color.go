package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	defaultCode = "38;5;%s"
	escape      = "\x1b["
	reset       = "\x1b[0m"
)

var (
	rgb6Regex  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	rgb3Regex  = regexp.MustCompile(`^#[0-9a-f]{3}$`)
	parseRegex = regexp.MustCompile(`{([^\}]+)}`)
	groupRegex = regexp.MustCompile(`([^\|]+)\|(.+)`)
)

type Color struct {
	foreground string
	background string
	bold       bool
	underline  bool
	format     string
}

func NewColor() *Color {
	color := &Color{}
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
func (c *Color) Paint(text interface{}, args ...interface{}) string {
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
func Parse(text string, args ...interface{}) string {
	replace := text
	matches := parseRegex.FindAllStringSubmatch(text, -1)
	for _, found := range matches {
		groups := groupRegex.FindAllStringSubmatch(found[1], -1)
		color := groups[0][1]
		group := groups[0][2]
		colored := Paint(color, group)
		replace = strings.Replace(replace, "{"+color+"|"+group+"}", colored, 1)
	}

	message := getText(replace, args...)
	return message
}

// Returns shell coloured output for text and args
func Paint(color string, text interface{}, args ...interface{}) string {
	message := getText(text, args...)

	fallback := os.Getenv("COLOR_FALLBACK")
	if fallback == "" {
		fallback = "green"
	}

	// Check if we have a numeric color
	i, _ := strconv.Atoi(color)
	s := strconv.Itoa(i)
	if s == color {
		return fromNumber(color, message)
	}

	// Check if we have a term named color, like "@CornflowerBlue"
	if color[0] == '@' {
		termColor, ok := termNames[color[1:]]
		if ok {
			return fromHtml(termColor, message)
		}
		return fromName(fallback, message)
	}

	// Lower the case for the next rounds
	color = strings.ToLower(color)

	// Check if we have a basic color name
	if _, ok := basicNames[color]; ok {
		return fromName(color, message)
	}

	// Finally we do the regex things
	// matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	if rgb6Regex.MatchString(color) {
		return fromHtml(color, message)
	}

	if rgb3Regex.MatchString(color) {
		htmlColor := fmt.Sprintf(
			"#%s%s%s%s%s%s",
			string(color[1]),
			string(color[1]),
			string(color[2]),
			string(color[2]),
			string(color[3]),
			string(color[3]))
		return fromHtml(htmlColor, message)
	}

	return fromName(fallback, message)
}

// Shortcut for color.Paint("green", text)
func Green(text interface{}, args ...interface{}) string {
	return Paint("green", text, args...)
}

// Shortcut for color.Paint("yellow", text)
func Yellow(text interface{}, args ...interface{}) string {
	return Paint("yellow", text, args...)
}

// Shortcut for color.Paint("red", text)
func Red(text interface{}, args ...interface{}) string {
	return Paint("red", text, args...)
}

// Shortcut for color.Paint("cyan", text)
func Cyan(text interface{}, args ...interface{}) string {
	return Paint("cyan", text, args...)
}

// Shortcut for color.Paint("blue", text)
func Blue(text interface{}, args ...interface{}) string {
	return Paint("blue", text, args...)
}

// Shortcut for color.Paint("magenta", text)
func Magenta(text interface{}, args ...interface{}) string {
	return Paint("magenta", text, args...)
}

// Shortcut for color.Paint("gray", text)
func Gray(text interface{}, args ...interface{}) string {
	return Paint("gray", text, args...)
}

// Shortcut for color.Paint("pale", text)
func Pale(text interface{}, args ...interface{}) string {
	return Paint("pale", text, args...)
}

// Concatenates text and args
func getText(text interface{}, args ...interface{}) string {
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
func fromHtml(color string, text string) string {
	r, g, b := html2Rgb(color)
	return fmt.Sprintf(escape+"38;2;%v;%v;%vm%s"+reset, r, g, b, text)
}

// Returns the (r,g, b) tupple for an html hex color
func html2Rgb(color string) (r uint8, g uint8, b uint8) {
	hexFormat := "#%02x%02x%02x"
	fmt.Sscanf(color, hexFormat, &r, &g, &b)
	return
}

// Returns the selected basic named color
func fromName(color string, text string) string {
	code := basicNames[color]

	// if code == "" {
	//  code = fmt.Sprintf(defaultCode, color)
	// }

	return fmt.Sprintf(escape+"%sm%s"+reset, code, text)
}

// Returns the selected numeric color
func fromNumber(color string, text string) string {
	code := fmt.Sprintf(defaultCode, color)
	return fmt.Sprintf(escape+"%sm%s"+reset, code, text)
}
