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

func TestIntenseBlackBg(t *testing.T) {
	expected := "\x1b[100mtest\x1b[0m"
	result := prettyprint.IntenseBlackBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBlackBgf(t *testing.T) {
	expected := "\x1b[100mtest 123\x1b[0m"
	result := prettyprint.IntenseBlackBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseRedBg(t *testing.T) {
	expected := "\x1b[101mtest\x1b[0m"
	result := prettyprint.IntenseRedBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseRedBgf(t *testing.T) {
	expected := "\x1b[101mtest 123\x1b[0m"
	result := prettyprint.IntenseRedBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseGreenBg(t *testing.T) {
	expected := "\x1b[102mtest\x1b[0m"
	result := prettyprint.IntenseGreenBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseGreenBgf(t *testing.T) {
	expected := "\x1b[102mtest 123\x1b[0m"
	result := prettyprint.IntenseGreenBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseYellowBg(t *testing.T) {
	expected := "\x1b[103mtest\x1b[0m"
	result := prettyprint.IntenseYellowBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseYellowBgf(t *testing.T) {
	expected := "\x1b[103mtest 123\x1b[0m"
	result := prettyprint.IntenseYellowBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlueBg(t *testing.T) {
	expected := "\x1b[104mtest\x1b[0m"
	result := prettyprint.IntenseBlueBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBlueBgf(t *testing.T) {
	expected := "\x1b[104mtest 123\x1b[0m"
	result := prettyprint.IntenseBlueBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaBg(t *testing.T) {
	expected := "\x1b[105mtest\x1b[0m"
	result := prettyprint.IntenseMagentaBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaBgf(t *testing.T) {
	expected := "\x1b[105mtest 123\x1b[0m"
	result := prettyprint.IntenseMagentaBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseCyanBg(t *testing.T) {
	expected := "\x1b[106mtest\x1b[0m"
	result := prettyprint.IntenseCyanBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseCyanBgf(t *testing.T) {
	expected := "\x1b[106mtest 123\x1b[0m"
	result := prettyprint.IntenseCyanBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseWhiteBg(t *testing.T) {
	expected := "\x1b[107mtest\x1b[0m"
	result := prettyprint.IntenseWhiteBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseWhiteBgf(t *testing.T) {
	expected := "\x1b[107mtest 123\x1b[0m"
	result := prettyprint.IntenseWhiteBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBackground8Bit(t *testing.T) {
	expected := "\x1b[48;5;42mtest\x1b[0m"
	result := prettyprint.Background8Bit(42, "test")
	assert.Equal(t, expected, result)
}

func TestBackground8Bitf(t *testing.T) {
	expected := "\x1b[48;5;42mtest 123\x1b[0m"
	result := prettyprint.Background8Bitf(42, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBackground24Bit(t *testing.T) {
	expected := "\x1b[48;2;255;128;64mtest\x1b[0m"
	result := prettyprint.Background24Bit(255, 128, 64, "test")
	assert.Equal(t, expected, result)
}

func TestBackground24Bitf(t *testing.T) {
	expected := "\x1b[48;2;255;128;64mtest 123\x1b[0m"
	result := prettyprint.Background24Bitf(255, 128, 64, "test %d", 123)
	assert.Equal(t, expected, result)
}
