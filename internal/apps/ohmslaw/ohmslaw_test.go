package ohmslaw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatecVoltage(t *testing.T) {
	actual := CalculatecVoltage(0.001, 1000)
	expected := 1.0

	assert.Equal(t, expected, actual)
}

func TestCalculateCurrent(t *testing.T) {
	actual := CalculateCurrent(5, 100)
	expected := 0.05

	assert.Equal(t, expected, actual)
}

func TestCalculateResistance(t *testing.T) {
	actual := CalculateResistance(5, 0.1)
	expected := 50.0

	assert.Equal(t, expected, actual)
}
