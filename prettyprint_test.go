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

func TestComplete(t *testing.T) {
	expected := "\x1b[32m[✓] \x1b[0mtest"
	result := prettyprint.Complete("test")
	assert.Equal(t, expected, result)
}

func TestAlert(t *testing.T) {
	expected := "\x1b[33m[!] \x1b[0mtest"
	result := prettyprint.Alert("test")
	assert.Equal(t, expected, result)
}

func TestRedAlert(t *testing.T) {
	expected := "\x1b[31m[!] \x1b[0mtest"
	result := prettyprint.RedAlert("test")
	assert.Equal(t, expected, result)
}

func TestFail(t *testing.T) {
	expected := "\x1b[31m[✗] \x1b[0mtest"
	result := prettyprint.Fail("test")
	assert.Equal(t, expected, result)
}

func TestInfo(t *testing.T) {
	expected := "\x1b[36m[i] \x1b[0mtest"
	result := prettyprint.Info("test")
	assert.Equal(t, expected, result)
}

func TestVar(t *testing.T) {
	expected := "\x1b[36m[i] \x1b[0m\x1b[36mvariable\x1b[0m is set to \x1b[32mvalue\x1b[0m"
	result := prettyprint.Var("variable", "value")
	assert.Equal(t, expected, result)
}

func TestVarQuote(t *testing.T) {
	expected := "\x1b[36m[i] \x1b[0m\"\x1b[36mvariable\x1b[0m\" is set to \"\x1b[32mvalue\x1b[0m\""
	result := prettyprint.VarQuote("variable", "value")
	assert.Equal(t, expected, result)
}
