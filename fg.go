package prettyprint

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
)

// Black formats using the default formats for its operands and returns the
// resulting string with a black foreground. Spaces are added between
// operands when neither is a string.
func Black(a ...any) string {
	return ansi.Black + fmt.Sprint(a...) + ansi.Reset
}

// Blackf formats according to a format specifier and returns the resulting
// string with a black foreground.
func Blackf(format string, a ...any) string {
	return ansi.Black + fmt.Sprintf(format, a...) + ansi.Reset
}

// Red formats using the default formats for its operands and returns the
// resulting string with a red foreground. Spaces are added between
// operands when neither is a string.
func Red(a ...any) string {
	return ansi.Red + fmt.Sprint(a...) + ansi.Reset
}

// Redf formats according to a format specifier and returns the resulting
// string with a red foreground.
func Redf(format string, a ...any) string {
	return ansi.Red + fmt.Sprintf(format, a...) + ansi.Reset
}

// Green formats using the default formats for its operands and returns the
// resulting string with a green foreground. Spaces are added between
// operands when neither is a string.
func Green(a ...any) string {
	return ansi.Green + fmt.Sprint(a...) + ansi.Reset
}

// Greenf formats according to a format specifier and returns the resulting
// string with a green foreground.
func Greenf(format string, a ...any) string {
	return ansi.Green + fmt.Sprintf(format, a...) + ansi.Reset
}

// Yellow formats using the default formats for its operands and returns the
// resulting string with a yellow foreground. Spaces are added between
// operands when neither is a string.
func Yellow(a ...any) string {
	return ansi.Yellow + fmt.Sprint(a...) + ansi.Reset
}

// Yellowf formats according to a format specifier and returns the resulting
// string with a yellow foreground.
func Yellowf(format string, a ...any) string {
	return ansi.Yellow + fmt.Sprintf(format, a...) + ansi.Reset
}

// Blue formats using the default formats for its operands and returns the
// resulting string with a blue foreground. Spaces are added between
// operands when neither is a string.
func Blue(a ...any) string {
	return ansi.Blue + fmt.Sprint(a...) + ansi.Reset
}

// Bluef formats according to a format specifier and returns the resulting
// string with a blue foreground.
func Bluef(format string, a ...any) string {
	return ansi.Blue + fmt.Sprintf(format, a...) + ansi.Reset
}

// Magenta formats using the default formats for its operands and returns the
// resulting string with a magenta foreground. Spaces are added between
// operands when neither is a string.
func Magenta(a ...any) string {
	return ansi.Magenta + fmt.Sprint(a...) + ansi.Reset
}

// Magentaf formats according to a format specifier and returns the resulting
// string with a magenta foreground.
func Magentaf(format string, a ...any) string {
	return ansi.Magenta + fmt.Sprintf(format, a...) + ansi.Reset
}

// Cyan formats using the default formats for its operands and returns the
// resulting string with a cyan foreground. Spaces are added between
// operands when neither is a string.
func Cyan(a ...any) string {
	return ansi.Cyan + fmt.Sprint(a...) + ansi.Reset
}

// Cyanf formats according to a format specifier and returns the resulting
// string with a cyan foreground.
func Cyanf(format string, a ...any) string {
	return ansi.Cyan + fmt.Sprintf(format, a...) + ansi.Reset
}

// White formats using the default formats for its operands and returns the
// resulting string with a white foreground. Spaces are added between
// operands when neither is a string.
func White(a ...any) string {
	return ansi.White + fmt.Sprint(a...) + ansi.Reset
}

// Whitef formats according to a format specifier and returns the resulting
// string with a white foreground.
func Whitef(format string, a ...any) string {
	return ansi.White + fmt.Sprintf(format, a...) + ansi.Reset
}
