package prettyprint

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
)

var (
	// IconComplete is the icon for a completed task.
	IconComplete = "[✓] "

	// IconAlert is the icon for an alert.
	IconAlert = "[!] "

	// IconFailed is the icon for a failed task.
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
