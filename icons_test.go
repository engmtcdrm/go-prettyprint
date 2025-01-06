package prettyprint_test

import (
	"testing"

	"github.com/engmtcdrm/go-prettyprint"
	"github.com/stretchr/testify/assert"
)

func TestComplete(t *testing.T) {
	expected := "\x1b[32m[✓]\x1b[0m test"
	result := prettyprint.Complete("test")
	assert.Equal(t, expected, result)
}

func TestAlert(t *testing.T) {
	expected := "\x1b[33m[!]\x1b[0m test"
	result := prettyprint.Alert("test")
	assert.Equal(t, expected, result)
}

func TestRedAlert(t *testing.T) {
	expected := "\x1b[31m[!]\x1b[0m test"
	result := prettyprint.RedAlert("test")
	assert.Equal(t, expected, result)
}

func TestFail(t *testing.T) {
	expected := "\x1b[31m[✗]\x1b[0m test"
	result := prettyprint.Fail("test")
	assert.Equal(t, expected, result)
}

func TestInfo(t *testing.T) {
	expected := "\x1b[36m[i]\x1b[0m test"
	result := prettyprint.Info("test")
	assert.Equal(t, expected, result)
}

func TestVar(t *testing.T) {
	expected := "\x1b[36m[i]\x1b[0m \x1b[36mvariable\x1b[0m is set to \x1b[32mvalue\x1b[0m"
	result := prettyprint.Var("variable", "value")
	assert.Equal(t, expected, result)
}

func TestVarQuote(t *testing.T) {
	expected := "\x1b[36m[i]\x1b[0m \"\x1b[36mvariable\x1b[0m\" is set to \"\x1b[32mvalue\x1b[0m\""
	result := prettyprint.VarQuote("variable", "value")
	assert.Equal(t, expected, result)
}
