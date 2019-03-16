package color

import (
  "fmt"
  "regexp"
  "strings"
)

var (
  colorNames = map[string]string{
    "red"     : "31",
    "green"   : "32",
    "yellow"  : "33",
    "blue"    : "34",
    "magenta" : "35",
    "cyan"    : "36",
  }
  defaultCode = "38;5;%s"
)

// Returns the wrapped shell outout for a text coloured
func Colorize(color string, text string) string {
  regexRgb1 := "^#[0-9a-f]{3}$"
  regexRgb2 := "^#[0-9a-f]{6}$"
  regexCode := "^[0-9]{2,3}$"
  // regexName := "^[a-z]*$"

  color = strings.ToLower(color)

  if m, _ := regexp.MatchString(regexRgb1, color); m {
    color = fmt.Sprintf(
      "#%s%s%s%s%s%s",
      string(color[1]),
      string(color[1]),
      string(color[2]),
      string(color[2]),
      string(color[3]),
      string(color[3]))
    return FromRgb(color, text)
  }

  if m, _ := regexp.MatchString(regexRgb2, color); m {
    return FromRgb(color, text)
  }

  if m, _ := regexp.MatchString(regexCode, color); m {
    return FromCode(color, text)
  }

  if _, ok := colorNames[color]; ok {
      return FromName(color, text)
  }

  return text
}

// Returns the closest shell colour string from an (r,g, b) tupple html color
func FromRgb(color string, text string) string {
    r, g, b := Html2Rgb(color)
    return fmt.Sprintf("\x1b[38;2;%v;%v;%vm%s\x1b[0m", r, g, b, text)
}

// Returns the (r,g, b) tupple html color from an html hex color string
func Html2Rgb(color string) (r uint8, g uint8, b uint8) {
    hexFormat := "#%02x%02x%02x"
    fmt.Sscanf(color, hexFormat, &r, &g, &b)
    return
}

// Returns the shell color code from a named color
func FromName(color string, text string) string {
  code := colorNames[color]

  // if code == "" {
  //  code = fmt.Sprintf(defaultCode, color)
  // }

  return fmt.Sprintf("\x1b[%sm%s\x1b[0m", code, text)
}

// Wraps text with the shell escape code for the selected color
func FromCode(color string, text string) string {
  code := fmt.Sprintf(defaultCode, color)
  return fmt.Sprintf("\x1b[%sm%s\x1b[0m", code, text)
}
