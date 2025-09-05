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

// IntenseBlackBg formats using the default formats for its operands and returns the
// resulting string with a intense black background. Spaces are added between
// operands when neither is a string.
func IntenseBlackBg(a ...any) string {
	return ansi.IntenseBlackBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseBlackBgf formats according to a format specifier and returns the resulting
// string with a intense black background.
func IntenseBlackBgf(format string, a ...any) string {
	return ansi.IntenseBlackBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseRedBg formats using the default formats for its operands and returns the
// resulting string with a intense red background. Spaces are added between
// operands when neither is a string.
func IntenseRedBg(a ...any) string {
	return ansi.IntenseRedBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseRedBgf formats according to a format specifier and returns the resulting
// string with a intense red background.
func IntenseRedBgf(format string, a ...any) string {
	return ansi.IntenseRedBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseGreenBg formats using the default formats for its operands and returns the
// resulting string with a intense green background. Spaces are added between
// operands when neither is a string.
func IntenseGreenBg(a ...any) string {
	return ansi.IntenseGreenBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseGreenBgf formats according to a format specifier and returns the resulting
// string with a intense green background.
func IntenseGreenBgf(format string, a ...any) string {
	return ansi.IntenseGreenBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseYellowBg formats using the default formats for its operands and returns the
// resulting string with a intense yellow background. Spaces are added between
// operands when neither is a string.
func IntenseYellowBg(a ...any) string {
	return ansi.IntenseYellowBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseYellowBgf formats according to a format specifier and returns the resulting
// string with a intense yellow background.
func IntenseYellowBgf(format string, a ...any) string {
	return ansi.IntenseYellowBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseBlueBg formats using the default formats for its operands and returns the
// resulting string with a intense blue background. Spaces are added between
// operands when neither is a string.
func IntenseBlueBg(a ...any) string {
	return ansi.IntenseBlueBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseBlueBgf formats according to a format specifier and returns the resulting
// string with a intense blue background.
func IntenseBlueBgf(format string, a ...any) string {
	return ansi.IntenseBlueBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseMagentaBg formats using the default formats for its operands and returns the
// resulting string with a intense magenta background. Spaces are added between
// operands when neither is a string.
func IntenseMagentaBg(a ...any) string {
	return ansi.IntenseMagentaBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseMagentaBgf formats according to a format specifier and returns the resulting
// string with a intense magenta background.
func IntenseMagentaBgf(format string, a ...any) string {
	return ansi.IntenseMagentaBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseCyanBg formats using the default formats for its operands and returns the
// resulting string with a intense cyan background. Spaces are added between
// operands when neither is a string.
func IntenseCyanBg(a ...any) string {
	return ansi.IntenseCyanBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseCyanBgf formats according to a format specifier and returns the resulting
// string with a intense cyan background.
func IntenseCyanBgf(format string, a ...any) string {
	return ansi.IntenseCyanBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// IntenseWhiteBg formats using the default formats for its operands and returns the
// resulting string with a intense white background. Spaces are added between
// operands when neither is a string.
func IntenseWhiteBg(a ...any) string {
	return ansi.IntenseWhiteBg + fmt.Sprint(a...) + ansi.Reset
}

// IntenseWhiteBgf formats according to a format specifier and returns the resulting
// string with a intense white background.
func IntenseWhiteBgf(format string, a ...any) string {
	return ansi.IntenseWhiteBg + fmt.Sprintf(format, a...) + ansi.Reset
}

// Background8Bit returns a string with the given 8-bit background color and the given message.
func Background8Bit(color int, msg string) string {
	return ansi.Background8Bit(color) + msg + ansi.Reset
}

// Background8Bitf returns a string with the given 8-bit background color and the given formatted message.
func Background8Bitf(color int, msg string, a ...any) string {
	return ansi.Background8Bit(color) + fmt.Sprintf(msg, a...) + ansi.Reset
}

// Background24Bit returns a string with the given 24-bit background color and the given message.
func Background24Bit(r, g, b int, msg string) string {
	return ansi.Background24Bit(r, g, b) + msg + ansi.Reset
}

// Background24Bitf returns a string with the given 24-bit background color and the given formatted message.
func Background24Bitf(r, g, b int, msg string, a ...any) string {
	return ansi.Background24Bit(r, g, b) + fmt.Sprintf(msg, a...) + ansi.Reset
}
