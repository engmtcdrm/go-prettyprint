package prettyprint

import "fmt"

var (
	// IconComplete is the icon for a completed message.
	IconComplete = "[✓]"

	// IconAlert is the icon for an alerted message.
	IconAlert = "[!]"

	// IconFailed is the icon for a failed message.
	IconFailed = "[✗]"

	// IconInfo is the icon for an informational message.
	IconInfo = "[i]"
)

// Complete returns a string with a green checkmark icon and the given message.
func Complete(msg string) string {
	return Green(IconComplete) + " " + msg
}

// Alert returns a string with a yellow exclamation icon and the given message.
func Alert(msg string) string {
	return Yellow(IconAlert) + " " + msg
}

// RedAlert returns a string with a red exclamation icon and the given message.
func RedAlert(msg string) string {
	return Red(IconAlert) + " " + msg
}

// Fail returns a string with a red X icon and the given message.
func Fail(msg string) string {
	return Red(IconFailed) + " " + msg
}

// Info returns a string with a cyan info icon and the given message.
func Info(msg string) string {
	return Cyan(IconInfo) + " " + msg
}

// Var returns a string with the given variable and value.
func Var(variable string, value string) string {
	return Info(Cyan(variable) + " is set to " + Green(value))
}

// VarQuote returns a string with the given variable and value, quoted.
func VarQuote(variable string, value string) string {
	return Info(fmt.Sprintf("\"%s\" is set to \"%s\"", Cyan(variable), Green(value)))
}
