package prettyprint

import "fmt"

var (
	IconComplete = "[✓]" // IconComplete is the icon for a completed message.
	IconAlert    = "[!]" // IconAlert is the icon for an alerted message.
	IconFailed   = "[✗]" // IconFailed is the icon for a failed message.
	IconInfo     = "[i]" // IconInfo is the icon for an informational message.
)

// Complete returns a string with a green checkmark icon and the given message.
func Complete(msg string) string {
	return Green(IconComplete) + " " + msg
}

// Completef returns a string with a green checkmark icon and the given formatted message.
func Completef(msg string, a ...any) string {
	return Green(IconComplete) + " " + fmt.Sprintf(msg, a...)
}

// Alert returns a string with a yellow exclamation icon and the given message.
func Alert(msg string) string {
	return Yellow(IconAlert) + " " + msg
}

// Alertf returns a string with a yellow exclamation icon and the given formatted message.
func Alertf(msg string, a ...any) string {
	return Yellow(IconAlert) + " " + fmt.Sprintf(msg, a...)
}

// RedAlert returns a string with a red exclamation icon and the given message.
func RedAlert(msg string) string {
	return Red(IconAlert) + " " + msg
}

// RedAlertf returns a string with a red exclamation icon and the given formatted message.
func RedAlertf(msg string, a ...any) string {
	return Red(IconAlert) + " " + fmt.Sprintf(msg, a...)
}

// Fail returns a string with a red X icon and the given message.
func Fail(msg string) string {
	return Red(IconFailed) + " " + msg
}

// Failf returns a string with a red X icon and the given formatted message.
func Failf(msg string, a ...any) string {
	return Red(IconFailed) + " " + fmt.Sprintf(msg, a...)
}

// Info returns a string with a cyan info icon and the given message.
func Info(msg string) string {
	return Cyan(IconInfo) + " " + msg
}

// Infof returns a string with a cyan info icon and the given formatted message.
func Infof(msg string, a ...any) string {
	return Cyan(IconInfo) + " " + fmt.Sprintf(msg, a...)
}

// Var returns a string with the given variable and value.
func Var(variable string, value string) string {
	return Info(Cyan(variable) + " is set to " + Green(value))
}

// VarSingleQuote returns a string with the given variable and value, single quoted.
func VarSingleQuote(variable string, value string) string {
	return Info(fmt.Sprintf("'%s' is set to '%s'", Cyan(variable), Green(value)))
}

// VarDoubleQuote returns a string with the given variable and value, double quoted.
func VarDoubleQuote(variable string, value string) string {
	return Info(fmt.Sprintf("\"%s\" is set to \"%s\"", Cyan(variable), Green(value)))
}

// VarQuote returns a string with the given variable and value, double quoted.
// Deprecated: use VarDoubleQuote instead.
func VarQuote(variable string, value string) string {
	return Info(fmt.Sprintf("\"%s\" is set to \"%s\"", Cyan(variable), Green(value)))
}
