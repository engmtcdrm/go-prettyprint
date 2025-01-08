package prettyprint

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
)

// BlackBg formats using the default formats for its operands and returns the
// resulting string with a black background. Spaces are added between
// operands when neither is a string.
func BlackBg(a ...any) string {
	return ansi.BlackBg + fmt.Sprint(a...) + ansi.Reset
}

// BlackBgf formats according to a format specifier and returns the resulting
// string with a black background.
func BlackBgf(format string, a ...any) string {
	return ansi.BlackBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// RedBg formats using the default formats for its operands and returns the
// resulting string with a red background. Spaces are added between
// operands when neither is a string.
func RedBg(a ...any) string {
	return ansi.RedBg + fmt.Sprint(a...) + ansi.Reset
}

// RedBgf formats according to a format specifier and returns the resulting
// string with a red background.
func RedBgf(format string, a ...any) string {
	return ansi.RedBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// GreenBg formats using the default formats for its operands and returns the
// resulting string with a green background. Spaces are added between
// operands when neither is a string.
func GreenBg(a ...any) string {
	return ansi.GreenBg + fmt.Sprint(a...) + ansi.Reset
}

// GreenBgf formats according to a format specifier and returns the resulting
// string with a green background.
func GreenBgf(format string, a ...any) string {
	return ansi.GreenBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// YellowBg formats using the default formats for its operands and returns the
// resulting string with a yellow background. Spaces are added between
// operands when neither is a string.
func YellowBg(a ...any) string {
	return ansi.YellowBg + fmt.Sprint(a...) + ansi.Reset
}

// YellowBgf formats according to a format specifier and returns the resulting
// string with a yellow background.
func YellowBgf(format string, a ...any) string {
	return ansi.YellowBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// BlueBg formats using the default formats for its operands and returns the
// resulting string with a blue background. Spaces are added between
// operands when neither is a string.
func BlueBg(a ...any) string {
	return ansi.BlueBg + fmt.Sprint(a...) + ansi.Reset
}

// BlueBgf formats according to a format specifier and returns the resulting
// string with a blue background.
func BlueBgf(format string, a ...any) string {
	return ansi.BlueBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// MagentaBg formats using the default formats for its operands and returns the
// resulting string with a magenta background. Spaces are added between
// operands when neither is a string.
func MagentaBg(a ...any) string {
	return ansi.MagentaBg + fmt.Sprint(a...) + ansi.Reset
}

// MagentaBgf formats according to a format specifier and returns the resulting
// string with a magenta background.
func MagentaBgf(format string, a ...any) string {
	return ansi.MagentaBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// CyanBg formats using the default formats for its operands and returns the
// resulting string with a cyan background. Spaces are added between
// operands when neither is a string.
func CyanBg(a ...any) string {
	return ansi.CyanBg + fmt.Sprint(a...) + ansi.Reset
}

// CyanBgf formats according to a format specifier and returns the resulting
// string with a cyan background.
func CyanBgf(format string, a ...any) string {
	return ansi.CyanBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// WhiteBg formats using the default formats for its operands and returns the
// resulting string with a white background. Spaces are added between
// operands when neither is a string.
func WhiteBg(a ...any) string {
	return ansi.WhiteBg + fmt.Sprint(a...) + ansi.Reset
}

// WhiteBgf formats according to a format specifier and returns the resulting
// string with a white background.
func WhiteBgf(format string, a ...any) string {
	return ansi.WhiteBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// Bg8Bit formats using the default formats for its operands and returns the
// resulting string with an 8-bit background color. Spaces are added between
// operands when neither is a string.
//
// color must be between 0 and 255 otherwise no color will be applied.
func Bg8Bit(color int, a ...any) string {
	bgColor := ansi.Background8Bit(color)

	if bgColor == "" {
		return fmt.Sprint(a...)
	}

	return bgColor + fmt.Sprint(a...) + ansi.Reset
}

// Bg8Bitf formats according to a format specifier and returns the resulting
// string with an 8-bit background color.
//
// color must be between 0 and 255 otherwise no color will be applied.
func Bg8Bitf(color int, format string, a ...any) string {
	bgColor := ansi.Background8Bit(color)

	if bgColor == "" {
		return fmt.Sprintf(format, a...)
	}

	return bgColor + fmt.Sprintf(format, a...) + ansi.Reset
}

// Bg24Bit formats using the default formats for its operands and returns the
// resulting string with an 24-bit background color. Spaces are added between
// operands when neither is a string.
//
// r, g, and b must be between 0 and 255 otherwise no color will be applied.
func Bg24Bit(r int, g int, b int, a ...any) string {
	bgColor := ansi.Background24Bit(r, g, b)

	if bgColor == "" {
		return fmt.Sprint(a...)
	}

	return bgColor + fmt.Sprint(a...) + ansi.Reset
}

// Bg24Bitf formats according to a format specifier and returns the resulting
// string with an 24-bit background color.
//
// r, g, and b must be between 0 and 255 otherwise no color will be applied.
func Bg24Bitf(r int, g int, b int, format string, a ...any) string {
	bgColor := ansi.Background24Bit(r, g, b)

	if bgColor == "" {
		return fmt.Sprintf(format, a...)
	}

	return bgColor + fmt.Sprintf(format, a...) + ansi.Reset
}