package prettyprint_test

import (
	"testing"

	"github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestBlack(t *testing.T) {
	expected := "\x1b[30mtest\x1b[0m"
	result := prettyprint.Black("test")
	assert.Equal(t, expected, result)
}

func TestBlackf(t *testing.T) {
	expected := "\x1b[30mtest 123\x1b[0m"
	result := prettyprint.Blackf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestRed(t *testing.T) {
	expected := "\x1b[31mtest\x1b[0m"
	result := prettyprint.Red("test")
	assert.Equal(t, expected, result)
}

func TestRedf(t *testing.T) {
	expected := "\x1b[31mtest 123\x1b[0m"
	result := prettyprint.Redf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestGreen(t *testing.T) {
	expected := "\x1b[32mtest\x1b[0m"
	result := prettyprint.Green("test")
	assert.Equal(t, expected, result)
}

func TestGreenf(t *testing.T) {
	expected := "\x1b[32mtest 123\x1b[0m"
	result := prettyprint.Greenf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestYellow(t *testing.T) {
	expected := "\x1b[33mtest\x1b[0m"
	result := prettyprint.Yellow("test")
	assert.Equal(t, expected, result)
}

func TestYellowf(t *testing.T) {
	expected := "\x1b[33mtest 123\x1b[0m"
	result := prettyprint.Yellowf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBlue(t *testing.T) {
	expected := "\x1b[34mtest\x1b[0m"
	result := prettyprint.Blue("test")
	assert.Equal(t, expected, result)
}

func TestBluef(t *testing.T) {
	expected := "\x1b[34mtest 123\x1b[0m"
	result := prettyprint.Bluef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestMagenta(t *testing.T) {
	expected := "\x1b[35mtest\x1b[0m"
	result := prettyprint.Magenta("test")
	assert.Equal(t, expected, result)
}

func TestMagentaf(t *testing.T) {
	expected := "\x1b[35mtest 123\x1b[0m"
	result := prettyprint.Magentaf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestCyan(t *testing.T) {
	expected := "\x1b[36mtest\x1b[0m"
	result := prettyprint.Cyan("test")
	assert.Equal(t, expected, result)
}

func TestCyanf(t *testing.T) {
	expected := "\x1b[36mtest 123\x1b[0m"
	result := prettyprint.Cyanf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestWhite(t *testing.T) {
	expected := "\x1b[37mtest\x1b[0m"
	result := prettyprint.White("test")
	assert.Equal(t, expected, result)
}

func TestWhitef(t *testing.T) {
	expected := "\x1b[37mtest 123\x1b[0m"
	result := prettyprint.Whitef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlack(t *testing.T) {
	expected := "\x1b[90mtest\x1b[0m"
	result := prettyprint.IntenseBlack("test")
	assert.Equal(t, expected, result)
}
func TestIntenseBlackf(t *testing.T) {
	expected := "\x1b[90mtest 123\x1b[0m"
	result := prettyprint.IntenseBlackf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseRed(t *testing.T) {
	expected := "\x1b[91mtest\x1b[0m"
	result := prettyprint.IntenseRed("test")
	assert.Equal(t, expected, result)
}

func TestIntenseRedf(t *testing.T) {
	expected := "\x1b[91mtest 123\x1b[0m"
	result := prettyprint.IntenseRedf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseGreen(t *testing.T) {
	expected := "\x1b[92mtest\x1b[0m"
	result := prettyprint.IntenseGreen("test")
	assert.Equal(t, expected, result)
}

func TestIntenseGreenf(t *testing.T) {
	expected := "\x1b[92mtest 123\x1b[0m"
	result := prettyprint.IntenseGreenf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseYellow(t *testing.T) {
	expected := "\x1b[93mtest\x1b[0m"
	result := prettyprint.IntenseYellow("test")
	assert.Equal(t, expected, result)
}

func TestIntenseYellowf(t *testing.T) {
	expected := "\x1b[93mtest 123\x1b[0m"
	result := prettyprint.IntenseYellowf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlue(t *testing.T) {
	expected := "\x1b[94mtest\x1b[0m"
	result := prettyprint.IntenseBlue("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBluef(t *testing.T) {
	expected := "\x1b[94mtest 123\x1b[0m"
	result := prettyprint.IntenseBluef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseMagenta(t *testing.T) {
	expected := "\x1b[95mtest\x1b[0m"
	result := prettyprint.IntenseMagenta("test")
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaf(t *testing.T) {
	expected := "\x1b[95mtest 123\x1b[0m"
	result := prettyprint.IntenseMagentaf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseCyan(t *testing.T) {
	expected := "\x1b[96mtest\x1b[0m"
	result := prettyprint.IntenseCyan("test")
	assert.Equal(t, expected, result)
}

func TestIntenseCyanf(t *testing.T) {
	expected := "\x1b[96mtest 123\x1b[0m"
	result := prettyprint.IntenseCyanf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseWhite(t *testing.T) {
	expected := "\x1b[97mtest\x1b[0m"
	result := prettyprint.IntenseWhite("test")
	assert.Equal(t, expected, result)
}

func TestIntenseWhitef(t *testing.T) {
	expected := "\x1b[97mtest 123\x1b[0m"
	result := prettyprint.IntenseWhitef("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestForeground8Bit(t *testing.T) {
	expected := "\x1b[38;5;42mtest\x1b[0m"
	result := prettyprint.Foreground8Bit(42, "test")
	assert.Equal(t, expected, result)
}

func TestForeground8Bitf(t *testing.T) {
	expected := "\x1b[38;5;42mtest 123\x1b[0m"
	result := prettyprint.Foreground8Bitf(42, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestForeground24Bit(t *testing.T) {
	expected := "\x1b[38;2;255;128;64mtest\x1b[0m"
	result := prettyprint.Foreground24Bit(255, 128, 64, "test")
	assert.Equal(t, expected, result)
}

func TestForeground24Bitf(t *testing.T) {
	expected := "\x1b[38;2;255;128;64mtest 123\x1b[0m"
	result := prettyprint.Foreground24Bitf(255, 128, 64, "test %d", 123)
	assert.Equal(t, expected, result)
}
