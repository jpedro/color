package color

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	codeEscape   = "\033["
	codeReset    = codeEscape + "0m"
	codeExtended = "8;5;%s"
	// codeForeground = "38;5;%s"
	// codeBackground = "48;5;%s"
)

var (
	fallbackColor string

	// rgb6Regex  = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	// rgb3Regex  = regexp.MustCompile(`^#[0-9a-f]{3}$`)
	parseRegex = regexp.MustCompile(`{([^\}]+)}`)
	groupRegex = regexp.MustCompile(`([^\|]+)\|(.+)`)
)

func init() {
	loadEnv()
}

func loadEnv() {
	fallbackColor = os.Getenv("COLOR_FALLBACK")
	if fallbackColor == "" {
		fallbackColor = "green"
	}
}

func getCode(color string) string {
	color = strings.ToLower(color)

	// Check if we have a term named color, like "@CornflowerBlue"
	if color[0] == '@' {
		name, ok := namesHtml[color[1:]]
		if ok {
			return fromHtml(name[1:])
		}
		return fromCommon(fallbackColor)
	}

	if color[0] == '#' && len(color) == 7 {
		return fromHtml(color[1:])
	}

	if color[0] == '#' && len(color) == 4 {
		html := fmt.Sprintf(
			"%s%s%s%s%s%s",
			string(color[1]),
			string(color[1]),
			string(color[2]),
			string(color[2]),
			string(color[3]),
			string(color[3]),
		)
		return fromHtml(html)
	}

	// Check if we have a common color name
	value, found := namesCommon[color]
	if found {
		return value
	}

	// Check if we have a numeric color
	val, err := strconv.Atoi(color)
	if err == nil {
		str := strconv.Itoa(val)
		if str == color {
			return fromExtended(color)
		}
		return fallbackColor
	}

	// Finally we do the regex things

	// matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
	// if rgb6Regex.MatchString(color) {
	// 	return getCodeFromHtml(color)
	// }

	// if rgb3Regex.MatchString(color) {
	// 	html := fmt.Sprintf(
	// 		"#%s%s%s%s%s%s",
	// 		string(color[1]),
	// 		string(color[1]),
	// 		string(color[2]),
	// 		string(color[2]),
	// 		string(color[3]),
	// 		string(color[3]))
	// 	return getCodeFromHtml(html)
	// }

	return fromCommon(fallbackColor)
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

// // func getCodeFromHtml(color string, foreground bool) string {
// func getCodeFromHtml(color string) string {
// 	r, g, b := getRgbFromHtml(color)
// 	// if foreground {
// 	return fmt.Sprintf("38;2;%v;%v;%v", r, g, b)
// 	// }
// 	// return fmt.Sprintf("48;2;%v;%v;%v", r, g, b)
// }

// func getRgbFromHtml(color string) (uint8, uint8, uint8) {
// 	var r, g, b uint8
// 	fmt.Sscanf(color, "#%02x%02x%02x", &r, &g, &b)
// 	return r, g, b
// }

// func getCodeFromName(color string) string {
// 	code := namesBasic[color]
// 	return code
// }

// func getCodeFromNumber(color string) string {
// 	return fmt.Sprintf(codeForeground, color)
// }

// BETTER WAY
func fromHtml(color string) string {
	r, g, b := getRgb(color)
	return fmt.Sprintf("8;2;%v;%v;%v", r, g, b)
}

func getRgb(color string) (uint8, uint8, uint8) {
	var r, g, b uint8
	fmt.Sscanf(color, "%02x%02x%02x", &r, &g, &b)
	return r, g, b
}

func fromCommon(color string) string {
	code := namesCommon[color]
	return code
}

func fromExtended(color string) string {
	return fmt.Sprintf(codeExtended, color)
}
