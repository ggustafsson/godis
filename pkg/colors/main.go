// Terminal colors package.
//
// Contains functions that generate data structures with preset terminal color
// and attribute string values to allow for easy use with standard print
// functions. ANSI 16 colors and basic style attributes only. By default all
// values are set to empty string if 'NO_COLOR' environment variable is set or
// if program is not running inside of interactive TTY, i.e. colors are
// automatically disabled during redirection or piping.
//
// Function [InitAuto] is automatically executed for user convenience.
// Functions [InitOn] and [InitOff] can be used to override with specific
// behaviour, e.g. to support implementation of '--color=on/off' argument.
//
// Structure:
//
//	colors
//	|-- Attr
//	|   |-- Blink
//	|   |-- Bold
//	|   |-- Italic
//	|   |-- Reset
//	|   |-- Reverse
//	|   `-- Underline
//	|-- Bg
//	|   |-- Black
//	|   |-- Blue
//	|   |-- Cyan
//	|   |-- Green
//	|   |-- Magenta
//	|   |-- Red
//	|   |-- White
//	|   |-- Yellow
//	|   |-- BrightBlack
//	|   |-- BrightBlue
//	|   |-- BrightCyan
//	|   |-- BrightGreen
//	|   |-- BrightMagenta
//	|   |-- BrightRed
//	|   |-- BrightWhite
//	|   `-- BrightYellow
//	`-- Fg
//	    |-- Black
//	    |-- Blue
//	    |-- Cyan
//	    |-- Green
//	    |-- Magenta
//	    |-- Red
//	    |-- White
//	    |-- Yellow
//	    |-- BrightBlack
//	    |-- BrightBlue
//	    |-- BrightCyan
//	    |-- BrightGreen
//	    |-- BrightMagenta
//	    |-- BrightRed
//	    |-- BrightWhite
//	    `-- BrightYellow
//
// Usage:
//
//	import "github.com/ggustafsson/godis/pkg/colors"
//
//	fmt.Printf("%sHello, 世界%s\n", colors.Fg.BrightRed, colors.Attr.Reset)
//
// Author: Göran Gustafsson <gustafsson.g@gmail.com>
//
// License: BSD 3-Clause
package colors

import "os"

// Terminal style attributes.
type Attributes struct {
	Blink     string
	Bold      string
	Italic    string
	Reset     string
	Reverse   string
	Underline string
}

// Terminal background & foreground colors.
type Colors struct {
	Black   string
	Blue    string
	Cyan    string
	Green   string
	Magenta string
	Red     string
	White   string
	Yellow  string

	BrightBlack   string
	BrightBlue    string
	BrightCyan    string
	BrightGreen   string
	BrightMagenta string
	BrightRed     string
	BrightWhite   string
	BrightYellow  string
}

// Using global variables to keep state without hassle for users.
var (
	Attr Attributes
	Bg   Colors
	Fg   Colors
)

// Check if running inside of TTY. Equivalent to libc isatty().
func isTTY() bool {
	stat, _ := os.Stdout.Stat()
	return (stat.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

// Check if 'NO_COLOR' environment variable is set.
func noColorEnv() bool {
	_, env := os.LookupEnv("NO_COLOR")
	return env
}

// Run either [InitOn] or [InitOff].
//
// If program is running inside of interactive TTY and 'NO_COLOR' environment
// variable is not set use function [InitOn], otherwise use [InitOff].
func InitAuto() {
	if isTTY() && !noColorEnv() {
		InitOn()
	} else {
		InitOff()
	}
}

// Set data structures with preset attribute and color values.
func InitOn() {
	Attr = Attributes{
		Reset:     "\033[0m",
		Bold:      "\033[1m",
		Italic:    "\033[3m",
		Underline: "\033[4m",
		Blink:     "\033[5m",
		Reverse:   "\033[7m",
	}
	Bg = Colors{
		Black:   "\033[40m",
		Red:     "\033[41m",
		Green:   "\033[42m",
		Yellow:  "\033[43m",
		Blue:    "\033[44m",
		Magenta: "\033[45m",
		Cyan:    "\033[46m",
		White:   "\033[47m",

		BrightBlack:   "\033[100m",
		BrightRed:     "\033[101m",
		BrightGreen:   "\033[102m",
		BrightYellow:  "\033[103m",
		BrightBlue:    "\033[104m",
		BrightMagenta: "\033[105m",
		BrightCyan:    "\033[106m",
		BrightWhite:   "\033[107m",
	}
	Fg = Colors{
		Black:   "\033[30m",
		Red:     "\033[31m",
		Green:   "\033[32m",
		Yellow:  "\033[33m",
		Blue:    "\033[34m",
		Magenta: "\033[35m",
		Cyan:    "\033[36m",
		White:   "\033[37m",

		BrightBlack:   "\033[90m",
		BrightRed:     "\033[91m",
		BrightGreen:   "\033[92m",
		BrightYellow:  "\033[93m",
		BrightBlue:    "\033[94m",
		BrightMagenta: "\033[95m",
		BrightCyan:    "\033[96m",
		BrightWhite:   "\033[97m",
	}
}

// Set data structures with empty attribute and color values.
func InitOff() {
	// Use default type values, i.e. empty strings.
	Attr, Bg, Fg = Attributes{}, Colors{}, Colors{}
}

func init() {
	// Start automatic mode by default for user convenience.
	InitAuto()
}
