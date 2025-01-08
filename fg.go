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

// Fg8Bit formats using the default formats for its operands and returns the
// resulting string with an 8-bit foreground color. Spaces are added between
// operands when neither is a string.
//
// color must be between 0 and 255 otherwise no color will be applied.
func Fg8Bit(color int, a ...any) string {
	fgColor := ansi.Foreground8Bit(color)

	if fgColor == "" {
		return fmt.Sprint(a...)
	}

	return fgColor + fmt.Sprint(a...) + ansi.Reset
}

// Fg8Bitf formats according to a format specifier and returns the resulting
// string with an 8-bit foreground color.
//
// color must be between 0 and 255 otherwise no color will be applied.
func Fg8Bitf(color int, format string, a ...any) string {
	fgColor := ansi.Foreground8Bit(color)

	if fgColor == "" {
		return fmt.Sprintf(format, a...)
	}

	return fgColor + fmt.Sprintf(format, a...) + ansi.Reset
}

// Fg24Bit formats using the default formats for its operands and returns the
// resulting string with an 24-bit foreground color. Spaces are added between
// operands when neither is a string.
//
// r, g, and b must be between 0 and 255 otherwise no color will be applied.
func Fg24Bit(r int, g int, b int, a ...any) string {
	fgColor := ansi.Foreground24Bit(r, g, b)

	if fgColor == "" {
		return fmt.Sprint(a...)
	}

	return fgColor + fmt.Sprint(a...) + ansi.Reset
}

// Fg24Bitf formats according to a format specifier and returns the resulting
// string with an 24-bit foreground color.
//
// r, g, and b must be between 0 and 255 otherwise no color will be applied.
func Fg24Bitf(r int, g int, b int, format string, a ...any) string {
	fgColor := ansi.Foreground24Bit(r, g, b)

	if fgColor == "" {
		return fmt.Sprintf(format, a...)
	}

	return fgColor + fmt.Sprintf(format, a...) + ansi.Reset
}
