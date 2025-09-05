package prettyprint_test

import (
	"testing"

	pp "github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestBlack(t *testing.T) {
	expected := "\x1b[30mtest\x1b[0m"
	result := pp.Black("test")
	assert.Equal(t, expected, result)
}

func TestBlackf(t *testing.T) {
	expected := "\x1b[30mtest 123\x1b[0m"
	result := pp.Blackf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestRed(t *testing.T) {
	expected := "\x1b[31mtest\x1b[0m"
	result := pp.Red("test")
	assert.Equal(t, expected, result)
}

func TestRedf(t *testing.T) {
	expected := "\x1b[31mtest 123\x1b[0m"
	result := pp.Redf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestGreen(t *testing.T) {
	expected := "\x1b[32mtest\x1b[0m"
	result := pp.Green("test")
	assert.Equal(t, expected, result)
}

func TestGreenf(t *testing.T) {
	expected := "\x1b[32mtest 123\x1b[0m"
	result := pp.Greenf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestYellow(t *testing.T) {
	expected := "\x1b[33mtest\x1b[0m"
	result := pp.Yellow("test")
	assert.Equal(t, expected, result)
}

func TestYellowf(t *testing.T) {
	expected := "\x1b[33mtest 123\x1b[0m"
	result := pp.Yellowf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBlue(t *testing.T) {
	expected := "\x1b[34mtest\x1b[0m"
	result := pp.Blue("test")
	assert.Equal(t, expected, result)
}

func TestBluef(t *testing.T) {
	expected := "\x1b[34mtest 123\x1b[0m"
	result := pp.Bluef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestMagenta(t *testing.T) {
	expected := "\x1b[35mtest\x1b[0m"
	result := pp.Magenta("test")
	assert.Equal(t, expected, result)
}

func TestMagentaf(t *testing.T) {
	expected := "\x1b[35mtest 123\x1b[0m"
	result := pp.Magentaf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestCyan(t *testing.T) {
	expected := "\x1b[36mtest\x1b[0m"
	result := pp.Cyan("test")
	assert.Equal(t, expected, result)
}

func TestCyanf(t *testing.T) {
	expected := "\x1b[36mtest 123\x1b[0m"
	result := pp.Cyanf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestWhite(t *testing.T) {
	expected := "\x1b[37mtest\x1b[0m"
	result := pp.White("test")
	assert.Equal(t, expected, result)
}

func TestWhitef(t *testing.T) {
	expected := "\x1b[37mtest 123\x1b[0m"
	result := pp.Whitef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlack(t *testing.T) {
	expected := "\x1b[90mtest\x1b[0m"
	result := pp.IntenseBlack("test")
	assert.Equal(t, expected, result)
}
func TestIntenseBlackf(t *testing.T) {
	expected := "\x1b[90mtest 123\x1b[0m"
	result := pp.IntenseBlackf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseRed(t *testing.T) {
	expected := "\x1b[91mtest\x1b[0m"
	result := pp.IntenseRed("test")
	assert.Equal(t, expected, result)
}

func TestIntenseRedf(t *testing.T) {
	expected := "\x1b[91mtest 123\x1b[0m"
	result := pp.IntenseRedf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseGreen(t *testing.T) {
	expected := "\x1b[92mtest\x1b[0m"
	result := pp.IntenseGreen("test")
	assert.Equal(t, expected, result)
}

func TestIntenseGreenf(t *testing.T) {
	expected := "\x1b[92mtest 123\x1b[0m"
	result := pp.IntenseGreenf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseYellow(t *testing.T) {
	expected := "\x1b[93mtest\x1b[0m"
	result := pp.IntenseYellow("test")
	assert.Equal(t, expected, result)
}

func TestIntenseYellowf(t *testing.T) {
	expected := "\x1b[93mtest 123\x1b[0m"
	result := pp.IntenseYellowf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlue(t *testing.T) {
	expected := "\x1b[94mtest\x1b[0m"
	result := pp.IntenseBlue("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBluef(t *testing.T) {
	expected := "\x1b[94mtest 123\x1b[0m"
	result := pp.IntenseBluef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseMagenta(t *testing.T) {
	expected := "\x1b[95mtest\x1b[0m"
	result := pp.IntenseMagenta("test")
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaf(t *testing.T) {
	expected := "\x1b[95mtest 123\x1b[0m"
	result := pp.IntenseMagentaf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseCyan(t *testing.T) {
	expected := "\x1b[96mtest\x1b[0m"
	result := pp.IntenseCyan("test")
	assert.Equal(t, expected, result)
}

func TestIntenseCyanf(t *testing.T) {
	expected := "\x1b[96mtest 123\x1b[0m"
	result := pp.IntenseCyanf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseWhite(t *testing.T) {
	expected := "\x1b[97mtest\x1b[0m"
	result := pp.IntenseWhite("test")
	assert.Equal(t, expected, result)
}

func TestIntenseWhitef(t *testing.T) {
	expected := "\x1b[97mtest 123\x1b[0m"
	result := pp.IntenseWhitef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestFg8Bit(t *testing.T) {
	expected := "\x1b[38;5;42mtest\x1b[0m"
	result := pp.Fg8Bit(42, "test")
	assert.Equal(t, expected, result)
}

func TestFg8Bitf(t *testing.T) {
	expected := "\x1b[38;5;42mtest 123\x1b[0m"
	result := pp.Fg8Bitf(42, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestFg24Bit(t *testing.T) {
	expected := "\x1b[38;2;255;128;64mtest\x1b[0m"
	result := pp.Fg24Bit(255, 128, 64, "test")
	assert.Equal(t, expected, result)
}

func TestFg24Bitf(t *testing.T) {
	expected := "\x1b[38;2;255;128;64mtest 123\x1b[0m"
	result := pp.Fg24Bitf(255, 128, 64, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestFg8BitInvalidColor(t *testing.T) {
	// Test with color value > 255 (should return plain text)
	result := pp.Fg8Bit(256, "test")
	assert.Equal(t, "test", result)

	// Test with negative color value (should return plain text)
	result = pp.Fg8Bit(-1, "test")
	assert.Equal(t, "test", result)
}

func TestFg8BitfInvalidColor(t *testing.T) {
	// Test with color value > 255 (should return plain text)
	result := pp.Fg8Bitf(256, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative color value (should return plain text)
	result = pp.Fg8Bitf(-1, "test %d", 123)
	assert.Equal(t, "test 123", result)
}

func TestFg24BitInvalidColor(t *testing.T) {
	// Test with R value > 255 (should return plain text)
	result := pp.Fg24Bit(256, 128, 64, "test")
	assert.Equal(t, "test", result)

	// Test with G value > 255 (should return plain text)
	result = pp.Fg24Bit(255, 256, 64, "test")
	assert.Equal(t, "test", result)

	// Test with B value > 255 (should return plain text)
	result = pp.Fg24Bit(255, 128, 256, "test")
	assert.Equal(t, "test", result)

	// Test with negative R value (should return plain text)
	result = pp.Fg24Bit(-1, 128, 64, "test")
	assert.Equal(t, "test", result)

	// Test with negative G value (should return plain text)
	result = pp.Fg24Bit(255, -1, 64, "test")
	assert.Equal(t, "test", result)

	// Test with negative B value (should return plain text)
	result = pp.Fg24Bit(255, 128, -1, "test")
	assert.Equal(t, "test", result)
}

func TestFg24BitfInvalidColor(t *testing.T) {
	// Test with R value > 255 (should return plain text)
	result := pp.Fg24Bitf(256, 128, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with G value > 255 (should return plain text)
	result = pp.Fg24Bitf(255, 256, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with B value > 255 (should return plain text)
	result = pp.Fg24Bitf(255, 128, 256, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative R value (should return plain text)
	result = pp.Fg24Bitf(-1, 128, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative G value (should return plain text)
	result = pp.Fg24Bitf(255, -1, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative B value (should return plain text)
	result = pp.Fg24Bitf(255, 128, -1, "test %d", 123)
	assert.Equal(t, "test 123", result)
}

func TestFg8BitBoundaryValues(t *testing.T) {
	// Test with color value 0 (should work)
	result := pp.Fg8Bit(0, "test")
	expected := "\x1b[38;5;0mtest\x1b[0m"
	assert.Equal(t, expected, result)

	// Test with color value 255 (should work)
	result = pp.Fg8Bit(255, "test")
	expected = "\x1b[38;5;255mtest\x1b[0m"
	assert.Equal(t, expected, result)
}

func TestFg24BitBoundaryValues(t *testing.T) {
	// Test with all values 0 (should work)
	result := pp.Fg24Bit(0, 0, 0, "test")
	expected := "\x1b[38;2;0;0;0mtest\x1b[0m"
	assert.Equal(t, expected, result)

	// Test with all values 255 (should work)
	result = pp.Fg24Bit(255, 255, 255, "test")
	expected = "\x1b[38;2;255;255;255mtest\x1b[0m"
	assert.Equal(t, expected, result)
}
