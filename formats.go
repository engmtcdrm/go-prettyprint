package prettyprint

import (
	"fmt"

	"github.com/engmtcdrm/go-ansi"
)

// Bold formats using the default formats for its operands and returns the
// resulting string in bold. Spaces are added between operands when neither is a string.
func Bold(a ...any) string {
	return ansi.Bold + fmt.Sprint(a...) + ansi.Reset
}

// Boldf formats according to a format specifier and returns the resulting
// string in bold.
func Boldf(format string, a ...any) string {
	return ansi.Bold + fmt.Sprintf(format, a...) + ansi.Reset
}

// Dim formats using the default formats for its operands and returns the
// resulting string in dim. Spaces are added between operands when neither is a string.
func Dim(a ...any) string {
	return ansi.Dim + fmt.Sprint(a...) + ansi.Reset
}

// Dimf formats according to a format specifier and returns the resulting
// string in dim.
func Dimf(format string, a ...any) string {
	return ansi.Dim + fmt.Sprintf(format, a...) + ansi.Reset
}

// Italic formats using the default formats for its operands and returns the
// resulting string in italic. Spaces are added between operands when neither is a string.
func Italic(a ...any) string {
	return ansi.Italic + fmt.Sprint(a...) + ansi.Reset
}

// Italicf formats according to a format specifier and returns the resulting
// string in italic.
func Italicf(format string, a ...any) string {
	return ansi.Italic + fmt.Sprintf(format, a...) + ansi.Reset
}

// Underline formats using the default formats for its operands and returns the
// resulting string underlined. Spaces are added between operands when neither is a string.
func Underline(a ...any) string {
	return ansi.Underline + fmt.Sprint(a...) + ansi.Reset
}

// Underlinef formats according to a format specifier and returns the resulting
// string underlined.
func Underlinef(format string, a ...any) string {
	return ansi.Underline + fmt.Sprintf(format, a...) + ansi.Reset
}

// Reverse formats using the default formats for its operands and returns the
// resulting string with reversed foreground and background. Spaces are added between
// operands when neither is a string.
func Reverse(a ...any) string {
	return ansi.Reverse + fmt.Sprint(a...) + ansi.Reset
}

// Reversef formats according to a format specifier and returns the resulting
// string with reversed foreground and background.
func Reversef(format string, a ...any) string {
	return ansi.Reverse + fmt.Sprintf(format, a...) + ansi.Reset
}

// Strike formats using the default formats for its operands and returns the
// resulting string with a strikethrough. Spaces are added between
// operands when neither is a string.
func Strike(a ...any) string {
	return ansi.Strike + fmt.Sprint(a...) + ansi.Reset
}

// Strikef formats according to a format specifier and returns the resulting
// string with a strikethrough.
func Strikef(format string, a ...any) string {
	return ansi.Strike + fmt.Sprintf(format, a...) + ansi.Reset
}

// DoubleUnderline formats using the default formats for its operands and returns the
// resulting string with a double underline. Spaces are added between
// operands when neither is a string.
func DoubleUnderline(a ...any) string {
	return ansi.DoubleUnderline + fmt.Sprint(a...) + ansi.Reset
}

// DoubleUnderlinef formats according to a format specifier and returns the resulting
// string with a double underline.
func DoubleUnderlinef(format string, a ...any) string {
	return ansi.DoubleUnderline + fmt.Sprintf(format, a...) + ansi.Reset
}

// Format applies the given formatting functions to the message in order and returns the resulting string.
// If no formatting functions are provided, the original message is returned.
func Format(msg string, a ...func(...any) string) string {
	if len(a) == 0 {
		return msg
	}

	result := msg
	for _, f := range a {
		result = f(result)
	}

	return result
}
