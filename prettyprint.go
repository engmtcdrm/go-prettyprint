package prettyprint

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
)

var (
	// IconComplete is the icon for a completed message.
	IconComplete = "[✓] "

	// IconAlert is the icon for an alerted message.
	IconAlert = "[!] "

	// IconFailed is the icon for a failed message.
	IconFailed = "[✗] "

	// IconInfo is the icon for an informational message.
	IconInfo = "[i] "
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

// Complete returns a string with a green checkmark icon and the given message.
func Complete(msg string) string {
	return Green(IconComplete) + msg
}

// Alert returns a string with a yellow exclamation icon and the given message.
func Alert(msg string) string {
	return Yellow(IconAlert) + msg
}

// RedAlert returns a string with a red exclamation icon and the given message.
func RedAlert(msg string) string {
	return Red(IconAlert) + msg
}

// Fail returns a string with a red X icon and the given message.
func Fail(msg string) string {
	return Red(IconFailed) + msg
}

// Info returns a string with a cyan info icon and the given message.
func Info(msg string) string {
	return Cyan(IconInfo) + msg
}

// Var returns a string with the given variable and value.
func Var(variable string, value string) string {
	return Info(Cyan(variable) + " is set to " + Green(value))
}

// VarQuote returns a string with the given variable and value, quoted.
func VarQuote(variable string, value string) string {
	return Info(fmt.Sprintf("\"%s\" is set to \"%s\"", Cyan(variable), Green(value)))
}
