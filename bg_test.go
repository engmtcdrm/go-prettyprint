package prettyprint_test

import (
	"testing"

	"github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestBlackBg(t *testing.T) {
	expected := "\x1b[40mtest\x1b[0m"
	result := prettyprint.BlackBg("test")
	assert.Equal(t, expected, result)
}

func TestBlackBgf(t *testing.T) {
	expected := "\x1b[40mtest 123\x1b[0m"
	result := prettyprint.BlackBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestRedBg(t *testing.T) {
	expected := "\x1b[41mtest\x1b[0m"
	result := prettyprint.RedBg("test")
	assert.Equal(t, expected, result)
}

func TestRedBgf(t *testing.T) {
	expected := "\x1b[41mtest 123\x1b[0m"
	result := prettyprint.RedBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestGreenBg(t *testing.T) {
	expected := "\x1b[42mtest\x1b[0m"
	result := prettyprint.GreenBg("test")
	assert.Equal(t, expected, result)
}

func TestGreenBgf(t *testing.T) {
	expected := "\x1b[42mtest 123\x1b[0m"
	result := prettyprint.GreenBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestYellowBg(t *testing.T) {
	expected := "\x1b[43mtest\x1b[0m"
	result := prettyprint.YellowBg("test")
	assert.Equal(t, expected, result)
}

func TestYellowBgf(t *testing.T) {
	expected := "\x1b[43mtest 123\x1b[0m"
	result := prettyprint.YellowBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBlueBg(t *testing.T) {
	expected := "\x1b[44mtest\x1b[0m"
	result := prettyprint.BlueBg("test")
	assert.Equal(t, expected, result)
}

func TestBlueBgf(t *testing.T) {
	expected := "\x1b[44mtest 123\x1b[0m"
	result := prettyprint.BlueBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestMagentaBg(t *testing.T) {
	expected := "\x1b[45mtest\x1b[0m"
	result := prettyprint.MagentaBg("test")
	assert.Equal(t, expected, result)
}

func TestMagentaBgf(t *testing.T) {
	expected := "\x1b[45mtest 123\x1b[0m"
	result := prettyprint.MagentaBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestCyanBg(t *testing.T) {
	expected := "\x1b[46mtest\x1b[0m"
	result := prettyprint.CyanBg("test")
	assert.Equal(t, expected, result)
}

func TestCyanBgf(t *testing.T) {
	expected := "\x1b[46mtest 123\x1b[0m"
	result := prettyprint.CyanBgf("test %d", 123)
	assert.Equal(t, expected, result)
}
func TestWhiteBg(t *testing.T) {
	expected := "\x1b[47mtest\x1b[0m"
	result := prettyprint.WhiteBg("test")
	assert.Equal(t, expected, result)
}

func TestWhiteBgf(t *testing.T) {
	expected := "\x1b[47mtest 123\x1b[0m"
	result := prettyprint.WhiteBgf("test %d", 123)
	assert.Equal(t, expected, result)
}
