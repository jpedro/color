package color

import (
    "os"
    "fmt"
    "strconv"
    "strings"
    "regexp"
)

var (
    defaultCode = "38;5;%s"
)

// Returns the shell output for a coloured text
func Paint(color string, text string) string {
    fallback := os.Getenv("COLOR_FALLBACK")
    if fallback == "" {
        fallback = "green"
    }

    // Check if we have a numeric color
    i, _ := strconv.Atoi(color)
    s := strconv.Itoa(i)
    if s == color {
        return FromNumber(color, text)
    }

    // Check if we have a term named color, like "@CornflowerBlue"
    if color[0] == '@' {
        termColor, ok := termNames[color[1:]]
        if ok {
            return FromHtml(termColor, text)
        }
        return FromName(fallback, text)
    }

    // Lower the case for the next rounds
    color = strings.ToLower(color)

    // Check if we have a basic color name
    if _, ok := basicNames[color]; ok {
        return FromName(color, text)
    }

    // Finally we do the regex things
    matches, _ := regexp.MatchString("^#[0-9a-f]{6}$", color)
    if matches {
        return FromHtml(color, text)
    }

    matches, _ = regexp.MatchString("^#[0-9a-f]{3}$", color);
    if matches {
        htmlColor := fmt.Sprintf(
            "#%s%s%s%s%s%s",
            string(color[1]),
            string(color[1]),
            string(color[2]),
            string(color[2]),
            string(color[3]),
            string(color[3]))
        return FromHtml(htmlColor, text)
    }

    return FromName(fallback, text)
}

// Returns the closest shell colour string from an html hex color (#rrggbb)
func FromHtml(color string, text string) string {
    r, g, b := Html2Rgb(color)
    return fmt.Sprintf("\x1b[38;2;%v;%v;%vm%s\x1b[0m", r, g, b, text)
}

// Returns the (r,g, b) tupple for an html hex color
func Html2Rgb(color string) (r uint8, g uint8, b uint8) {
    hexFormat := "#%02x%02x%02x"
    fmt.Sscanf(color, hexFormat, &r, &g, &b)
    return
}

// Returns the selected basic named color
func FromName(color string, text string) string {
    code := basicNames[color]

    // if code == "" {
    //  code = fmt.Sprintf(defaultCode, color)
    // }

    return fmt.Sprintf("\x1b[%sm%s\x1b[0m", code, text)
}

// Returns the selected numeric color
func FromNumber(color string, text string) string {
    code := fmt.Sprintf(defaultCode, color)
    return fmt.Sprintf("\x1b[%sm%s\x1b[0m", code, text)
}
