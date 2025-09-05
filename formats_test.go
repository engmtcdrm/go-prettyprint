package prettyprint_test

import (
	"testing"

	"github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestBold(t *testing.T) {
	expected := "\x1b[1mtest\x1b[0m"
	result := prettyprint.Bold("test")
	assert.Equal(t, expected, result)
}

func TestBoldf(t *testing.T) {
	expected := "\x1b[1mtest 123\x1b[0m"
	result := prettyprint.Boldf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestDim(t *testing.T) {
	expected := "\x1b[2mtest\x1b[0m"
	result := prettyprint.Dim("test")
	assert.Equal(t, expected, result)
}

func TestDimf(t *testing.T) {
	expected := "\x1b[2mtest 123\x1b[0m"
	result := prettyprint.Dimf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestItalic(t *testing.T) {
	expected := "\x1b[3mtest\x1b[0m"
	result := prettyprint.Italic("test")
	assert.Equal(t, expected, result)
}

func TestItalicf(t *testing.T) {
	expected := "\x1b[3mtest 123\x1b[0m"
	result := prettyprint.Italicf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestUnderline(t *testing.T) {
	expected := "\x1b[4mtest\x1b[0m"
	result := prettyprint.Underline("test")
	assert.Equal(t, expected, result)
}

func TestUnderlinef(t *testing.T) {
	expected := "\x1b[4mtest 123\x1b[0m"
	result := prettyprint.Underlinef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestReverse(t *testing.T) {
	expected := "\x1b[7mtest\x1b[0m"
	result := prettyprint.Reverse("test")
	assert.Equal(t, expected, result)
}

func TestReversef(t *testing.T) {
	expected := "\x1b[7mtest 123\x1b[0m"
	result := prettyprint.Reversef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestStrike(t *testing.T) {
	expected := "\x1b[9mtest\x1b[0m"
	result := prettyprint.Strike("test")
	assert.Equal(t, expected, result)
}

func TestStrikef(t *testing.T) {
	expected := "\x1b[9mtest 123\x1b[0m"
	result := prettyprint.Strikef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestDoubleUnderline(t *testing.T) {
	expected := "\x1b[21mtest\x1b[0m"
	result := prettyprint.DoubleUnderline("test")
	assert.Equal(t, expected, result)
}

func TestDoubleUnderlinef(t *testing.T) {
	expected := "\x1b[21mtest 123\x1b[0m"
	result := prettyprint.DoubleUnderlinef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestFormat(t *testing.T) {
	// Test with no formatting functions
	result := prettyprint.Format("test")
	assert.Equal(t, "test", result)

	// Test with single formatting function
	result = prettyprint.Format("test", prettyprint.Bold)
	// Current implementation: "test" + Bold("test") = "test" + "\x1b[1mtest\x1b[0m"
	expected := "test\x1b[1mtest\x1b[0m"
	assert.Equal(t, expected, result)

	// Test with multiple formatting functions
	result = prettyprint.Format("test", prettyprint.Bold, prettyprint.Red)
	// Current implementation appends each formatted result
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "test")
	assert.Contains(t, result, "\x1b[1m")  // Bold
	assert.Contains(t, result, "\x1b[31m") // Red
}
