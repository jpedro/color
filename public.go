package color

import (
	"fmt"
	"strings"
)

type ColorFlags int

const (
	ColorBold = iota + 1
	ColorFaint
	ColorItalic
	ColorUnderline
	ColorBlink
	ColorUnknown
	ColorReverse
	ColorReseverd
	ColorStrike
)

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

func Paint(color string, message any, args ...any) string {
	text := getText(message, args...)
	code := getCode(color)

	return fmt.Sprintf(codeEscape+"3%sm%s"+codeReset, code, text)
}
