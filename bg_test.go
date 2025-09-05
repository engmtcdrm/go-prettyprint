package prettyprint_test

import (
	"testing"

	pp "github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestBlackBg(t *testing.T) {
	expected := "\x1b[40mtest\x1b[0m"
	result := pp.BlackBg("test")
	assert.Equal(t, expected, result)
}

func TestBlackBgf(t *testing.T) {
	expected := "\x1b[40mtest 123\x1b[0m"
	result := pp.BlackBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestRedBg(t *testing.T) {
	expected := "\x1b[41mtest\x1b[0m"
	result := pp.RedBg("test")
	assert.Equal(t, expected, result)
}

func TestRedBgf(t *testing.T) {
	expected := "\x1b[41mtest 123\x1b[0m"
	result := pp.RedBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestGreenBg(t *testing.T) {
	expected := "\x1b[42mtest\x1b[0m"
	result := pp.GreenBg("test")
	assert.Equal(t, expected, result)
}

func TestGreenBgf(t *testing.T) {
	expected := "\x1b[42mtest 123\x1b[0m"
	result := pp.GreenBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestYellowBg(t *testing.T) {
	expected := "\x1b[43mtest\x1b[0m"
	result := pp.YellowBg("test")
	assert.Equal(t, expected, result)
}

func TestYellowBgf(t *testing.T) {
	expected := "\x1b[43mtest 123\x1b[0m"
	result := pp.YellowBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBlueBg(t *testing.T) {
	expected := "\x1b[44mtest\x1b[0m"
	result := pp.BlueBg("test")
	assert.Equal(t, expected, result)
}

func TestBlueBgf(t *testing.T) {
	expected := "\x1b[44mtest 123\x1b[0m"
	result := pp.BlueBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestMagentaBg(t *testing.T) {
	expected := "\x1b[45mtest\x1b[0m"
	result := pp.MagentaBg("test")
	assert.Equal(t, expected, result)
}

func TestMagentaBgf(t *testing.T) {
	expected := "\x1b[45mtest 123\x1b[0m"
	result := pp.MagentaBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestCyanBg(t *testing.T) {
	expected := "\x1b[46mtest\x1b[0m"
	result := pp.CyanBg("test")
	assert.Equal(t, expected, result)
}

func TestCyanBgf(t *testing.T) {
	expected := "\x1b[46mtest 123\x1b[0m"
	result := pp.CyanBgf("test %d", 123)
	assert.Equal(t, expected, result)
}
func TestWhiteBg(t *testing.T) {
	expected := "\x1b[47mtest\x1b[0m"
	result := pp.WhiteBg("test")
	assert.Equal(t, expected, result)
}

func TestWhiteBgf(t *testing.T) {
	expected := "\x1b[47mtest 123\x1b[0m"
	result := pp.WhiteBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlackBg(t *testing.T) {
	expected := "\x1b[100mtest\x1b[0m"
	result := pp.IntenseBlackBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBlackBgf(t *testing.T) {
	expected := "\x1b[100mtest 123\x1b[0m"
	result := pp.IntenseBlackBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseRedBg(t *testing.T) {
	expected := "\x1b[101mtest\x1b[0m"
	result := pp.IntenseRedBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseRedBgf(t *testing.T) {
	expected := "\x1b[101mtest 123\x1b[0m"
	result := pp.IntenseRedBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseGreenBg(t *testing.T) {
	expected := "\x1b[102mtest\x1b[0m"
	result := pp.IntenseGreenBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseGreenBgf(t *testing.T) {
	expected := "\x1b[102mtest 123\x1b[0m"
	result := pp.IntenseGreenBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseYellowBg(t *testing.T) {
	expected := "\x1b[103mtest\x1b[0m"
	result := pp.IntenseYellowBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseYellowBgf(t *testing.T) {
	expected := "\x1b[103mtest 123\x1b[0m"
	result := pp.IntenseYellowBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseBlueBg(t *testing.T) {
	expected := "\x1b[104mtest\x1b[0m"
	result := pp.IntenseBlueBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseBlueBgf(t *testing.T) {
	expected := "\x1b[104mtest 123\x1b[0m"
	result := pp.IntenseBlueBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaBg(t *testing.T) {
	expected := "\x1b[105mtest\x1b[0m"
	result := pp.IntenseMagentaBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseMagentaBgf(t *testing.T) {
	expected := "\x1b[105mtest 123\x1b[0m"
	result := pp.IntenseMagentaBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseCyanBg(t *testing.T) {
	expected := "\x1b[106mtest\x1b[0m"
	result := pp.IntenseCyanBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseCyanBgf(t *testing.T) {
	expected := "\x1b[106mtest 123\x1b[0m"
	result := pp.IntenseCyanBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestIntenseWhiteBg(t *testing.T) {
	expected := "\x1b[107mtest\x1b[0m"
	result := pp.IntenseWhiteBg("test")
	assert.Equal(t, expected, result)
}

func TestIntenseWhiteBgf(t *testing.T) {
	expected := "\x1b[107mtest 123\x1b[0m"
	result := pp.IntenseWhiteBgf("test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBg8Bit(t *testing.T) {
	expected := "\x1b[48;5;42mtest\x1b[0m"
	result := pp.Bg8Bit(42, "test")
	assert.Equal(t, expected, result)
}

func TestBg8Bitf(t *testing.T) {
	expected := "\x1b[48;5;42mtest 123\x1b[0m"
	result := pp.Bg8Bitf(42, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBg24Bit(t *testing.T) {
	expected := "\x1b[48;2;255;128;64mtest\x1b[0m"
	result := pp.Bg24Bit(255, 128, 64, "test")
	assert.Equal(t, expected, result)
}

func TestBg24Bitf(t *testing.T) {
	expected := "\x1b[48;2;255;128;64mtest 123\x1b[0m"
	result := pp.Bg24Bitf(255, 128, 64, "test %d", 123)
	assert.Equal(t, expected, result)
}

func TestBg8BitInvalidColor(t *testing.T) {
	// Test with color value > 255 (should return plain text)
	result := pp.Bg8Bit(256, "test")
	assert.Equal(t, "test", result)

	// Test with negative color value (should return plain text)
	result = pp.Bg8Bit(-1, "test")
	assert.Equal(t, "test", result)
}

func TestBg8BitfInvalidColor(t *testing.T) {
	// Test with color value > 255 (should return plain text)
	result := pp.Bg8Bitf(256, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative color value (should return plain text)
	result = pp.Bg8Bitf(-1, "test %d", 123)
	assert.Equal(t, "test 123", result)
}

func TestBg24BitInvalidColor(t *testing.T) {
	// Test with R value > 255 (should return plain text)
	result := pp.Bg24Bit(256, 128, 64, "test")
	assert.Equal(t, "test", result)

	// Test with G value > 255 (should return plain text)
	result = pp.Bg24Bit(255, 256, 64, "test")
	assert.Equal(t, "test", result)

	// Test with B value > 255 (should return plain text)
	result = pp.Bg24Bit(255, 128, 256, "test")
	assert.Equal(t, "test", result)

	// Test with negative R value (should return plain text)
	result = pp.Bg24Bit(-1, 128, 64, "test")
	assert.Equal(t, "test", result)

	// Test with negative G value (should return plain text)
	result = pp.Bg24Bit(255, -1, 64, "test")
	assert.Equal(t, "test", result)

	// Test with negative B value (should return plain text)
	result = pp.Bg24Bit(255, 128, -1, "test")
	assert.Equal(t, "test", result)
}

func TestBg24BitfInvalidColor(t *testing.T) {
	// Test with R value > 255 (should return plain text)
	result := pp.Bg24Bitf(256, 128, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with G value > 255 (should return plain text)
	result = pp.Bg24Bitf(255, 256, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with B value > 255 (should return plain text)
	result = pp.Bg24Bitf(255, 128, 256, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative R value (should return plain text)
	result = pp.Bg24Bitf(-1, 128, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative G value (should return plain text)
	result = pp.Bg24Bitf(255, -1, 64, "test %d", 123)
	assert.Equal(t, "test 123", result)

	// Test with negative B value (should return plain text)
	result = pp.Bg24Bitf(255, 128, -1, "test %d", 123)
	assert.Equal(t, "test 123", result)
}

func TestBg8BitBoundaryValues(t *testing.T) {
	// Test with color value 0 (should work)
	result := pp.Bg8Bit(0, "test")
	expected := "\x1b[48;5;0mtest\x1b[0m"
	assert.Equal(t, expected, result)

	// Test with color value 255 (should work)
	result = pp.Bg8Bit(255, "test")
	expected = "\x1b[48;5;255mtest\x1b[0m"
	assert.Equal(t, expected, result)
}

func TestBg24BitBoundaryValues(t *testing.T) {
	// Test with all values 0 (should work)
	result := pp.Bg24Bit(0, 0, 0, "test")
	expected := "\x1b[48;2;0;0;0mtest\x1b[0m"
	assert.Equal(t, expected, result)

	// Test with all values 255 (should work)
	result = pp.Bg24Bit(255, 255, 255, "test")
	expected = "\x1b[48;2;255;255;255mtest\x1b[0m"
	assert.Equal(t, expected, result)
}
