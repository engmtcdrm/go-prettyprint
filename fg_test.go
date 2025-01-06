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
