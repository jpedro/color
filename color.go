package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	codeEscape     = "\033["
	codeReset      = codeEscape + "0m"
	codeForeground = "38;5;%s"
	codeBackground = "48;5;%s"
)

var (
	fallbackColor string

	rgb6Regex  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	rgb3Regex  = regexp.MustCompile(`^#[0-9a-f]{3}$`)
	parseRegex = regexp.MustCompile(`{([^\}]+)}`)
	groupRegex = regexp.MustCompile(`([^\|]+)\|(.+)`)
)

type Color struct {
	foreground string
	background string
	bold       bool // 1
	faint      bool // 2
	italic     bool // 3
	underline  bool // 4
	blink      bool // 5
	reverse    bool // 7
	strike     bool // 9
	format     string
}

func init() {
	fallbackColor = os.Getenv("COLOR_FALLBACK")
	if fallbackColor == "" {
		fallbackColor = "green"
	}
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
	c.bold = true
	return c
}

func (c *Color) Underline() *Color {
	c.underline = true
	return c
}

func (c *Color) Blink() *Color {
	c.blink = true
	return c
}

func (c *Color) Strike() *Color {
	c.strike = true
	return c
}

// Paints text and args according to its settings
func (c *Color) Paint(message any, args ...any) string {
	text := getText(message, args...)

	if c.format == "" {
		format := ""

		if c.foreground != "" {
			// format = fmt.Sprintf("%s;38;5;%s", format, c.foreground)
			// format = fmt.Sprintf("%s;38;5;%s", format, getCode(c.foreground))
			format = getCode(c.foreground)
		}

		if c.background != "" {
			format = fmt.Sprintf("%s;48;5;%s", format, c.background)
		}

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
	}

	return fmt.Sprintf("%s%sm%s%s", codeEscape, c.format, text, codeReset)
}

// Returns a painted string with groups like `{green|this should be green}`
// replaced with `this should be green` in green
func Format(text string, args ...any) string {
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
func Paint(color string, message any, args ...any) string {
	text := getText(message, args...)
	code := getCode(color)

	return fmt.Sprintf(codeEscape+"%sm%s"+codeReset, code, text)

	// // Check if we have a numeric color
	// i, _ := strconv.Atoi(color)
	// s := strconv.Itoa(i)
	// if s == color {
	// 	return fromNumber(color, message)
	// }

	// // Check if we have a term named color, like "@CornflowerBlue"
	// if color[0] == '@' {
	// 	termColor, ok := namesTerm[color[1:]]
	// 	if ok {
	// 		return fromHtml(termColor, message)
	// 	}
	// 	return fromName(fallback, message)
	// }

	// // Lower the case for the next rounds
	// color = strings.ToLower(color)

	// // Check if we have a basic color name
	// if _, ok := namesBasic[color]; ok {
	// 	return fromName(color, message)
	// }

	// // Finally we do the regex things
	// // matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	// if rgb6Regex.MatchString(color) {
	// 	return fromHtml(color, message)
	// }

	// if rgb3Regex.MatchString(color) {
	// 	htmlColor := fmt.Sprintf(
	// 		"#%s%s%s%s%s%s",
	// 		string(color[1]),
	// 		string(color[1]),
	// 		string(color[2]),
	// 		string(color[2]),
	// 		string(color[3]),
	// 		string(color[3]))
	// 	return fromHtml(htmlColor, message)
	// }

	// return fromName(fallback, message)
}

func Green(text any, args ...any) string {
	return Paint("green", text, args...)
}

func Yellow(text any, args ...any) string {
	return Paint("yellow", text, args...)
}

func Red(text any, args ...any) string {
	return Paint("red", text, args...)
}

func Cyan(text any, args ...any) string {
	return Paint("cyan", text, args...)
}

func Blue(text any, args ...any) string {
	return Paint("blue", text, args...)
}

func Magenta(text any, args ...any) string {
	return Paint("magenta", text, args...)
}

func Gray(text any, args ...any) string {
	return Paint("gray", text, args...)
}

func Pale(text any, args ...any) string {
	return Paint("pale", text, args...)
}

func getCode(color string) string {
	color = strings.ToLower(color)

	// Check if we have a term named color, like "@CornflowerBlue"
	if color[0] == '@' {
		name, ok := namesTerm[color[1:]]
		if ok {
			return getCodeFromHtml(name)
		}
		return getCodeFromName(fallbackColor)
	}

	// Check if we have a basic color name
	value, found := namesBasic[color]
	if found {
		return value
	}

	// Check if we have a numeric color
	val, err := strconv.Atoi(color)
	if err != nil {
		str := strconv.Itoa(val)
		if str == color {
			return getCodeFromNumber(color)
		}
		return fallbackColor
	}

	// Finally we do the regex things

	// matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	if rgb6Regex.MatchString(color) {
		return getCodeFromHtml(color)
	}

	if rgb3Regex.MatchString(color) {
		rgb := fmt.Sprintf(
			"#%s%s%s%s%s%s",
			string(color[1]),
			string(color[1]),
			string(color[2]),
			string(color[2]),
			string(color[3]),
			string(color[3]))
		return getCodeFromHtml(rgb)
	}

	return getCodeFromName(fallbackColor)
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
func getCodeFromHtml(color string) string {
	r, g, b := getCodeFromRgb(color)
	// return fmt.Sprintf(codeEscape+"38;2;%v;%v;%vm%s"+codeReset, r, g, b, text)
	return fmt.Sprintf("38;2;%v;%v;%v", r, g, b)
}

// Returns an (r, g, b) tupple for an html hex color
func getCodeFromRgb(color string) (uint8, uint8, uint8) {
	var r, g, b uint8
	fmt.Sscanf(color, "#%02x%02x%02x", &r, &g, &b)
	return r, g, b
}

// Returns the selected basic named color
func getCodeFromName(color string) string {
	code := namesBasic[color]

	// if code == "" {
	//  code = fmt.Sprintf(defaultCode, color)
	// }

	// return fmt.Sprintf(codeEscape+"%sm%s"+codeReset, code, text)
	return code
}

func getCodeFromNumber(color string) string {
	return fmt.Sprintf(codeForeground, color)
}
