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

// IntenseBlack formats using the default formats for its operands and returns the
// resulting string with a intense black (gray) foreground. Spaces are added between
// operands when neither is a string.
func IntenseBlack(a ...any) string {
	return ansi.IntenseBlack + fmt.Sprint(a...) + ansi.Reset
}

// IntenseBlackf formats according to a format specifier and returns the resulting
// string with a intense black (gray) foreground.
func IntenseBlackf(format string, a ...any) string {
	return ansi.IntenseBlack + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseRed formats using the default formats for its operands and returns the
// resulting string with a intense red foreground. Spaces are added between
// operands when neither is a string.
func IntenseRed(a ...any) string {
	return ansi.IntenseRed + fmt.Sprint(a...) + ansi.Reset
}

// IntenseRedf formats according to a format specifier and returns the resulting
// string with a intense red foreground.
func IntenseRedf(format string, a ...any) string {
	return ansi.IntenseRed + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseGreen formats using the default formats for its operands and returns the
// resulting string with a intense green foreground. Spaces are added between
// operands when neither is a string.
func IntenseGreen(a ...any) string {
	return ansi.IntenseGreen + fmt.Sprint(a...) + ansi.Reset
}

// IntenseGreenf formats according to a format specifier and returns the resulting
// string with a intense green foreground.
func IntenseGreenf(format string, a ...any) string {
	return ansi.IntenseGreen + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseYellow formats using the default formats for its operands and returns the
// resulting string with a intense yellow foreground. Spaces are added between
// operands when neither is a string.
func IntenseYellow(a ...any) string {
	return ansi.IntenseYellow + fmt.Sprint(a...) + ansi.Reset
}

// IntenseYellowf formats according to a format specifier and returns the resulting
// string with a intense yellow foreground.
func IntenseYellowf(format string, a ...any) string {
	return ansi.IntenseYellow + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseBlue formats using the default formats for its operands and returns the
// resulting string with a intense blue foreground. Spaces are added between
// operands when neither is a string.
func IntenseBlue(a ...any) string {
	return ansi.IntenseBlue + fmt.Sprint(a...) + ansi.Reset
}

// IntenseBluef formats according to a format specifier and returns the resulting
// string with a intense blue foreground.
func IntenseBluef(format string, a ...any) string {
	return ansi.IntenseBlue + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseMagenta formats using the default formats for its operands and returns the
// resulting string with a intense magenta foreground. Spaces are added between
// operands when neither is a string.
func IntenseMagenta(a ...any) string {
	return ansi.IntenseMagenta + fmt.Sprint(a...) + ansi.Reset
}

// IntenseMagentaf formats according to a format specifier and returns the resulting
// string with a intense magenta foreground.
func IntenseMagentaf(format string, a ...any) string {
	return ansi.IntenseMagenta + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseCyan formats using the default formats for its operands and returns the
// resulting string with a intense cyan foreground. Spaces are added between
// operands when neither is a string.
func IntenseCyan(a ...any) string {
	return ansi.IntenseCyan + fmt.Sprint(a...) + ansi.Reset
}

// IntenseCyanf formats according to a format specifier and returns the resulting
// string with a intense cyan foreground.
func IntenseCyanf(format string, a ...any) string {
	return ansi.IntenseCyan + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseWhite formats using the default formats for its operands and returns the
// resulting string with a intense white foreground. Spaces are added between
// operands when neither is a string.
func IntenseWhite(a ...any) string {
	return ansi.IntenseWhite + fmt.Sprint(a...) + ansi.Reset
}

// IntenseWhitef formats according to a format specifier and returns the resulting
// string with a intense white foreground.
func IntenseWhitef(format string, a ...any) string {
	return ansi.IntenseWhite + fmt.Sprintf(format, a...) + ansi.Reset
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
