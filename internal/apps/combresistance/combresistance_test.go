package combresistance

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeriesCircuits(t *testing.T) {
	actual := SeriesCircuits(100.0, 200.0)
	expected := 300.0

	assert.Equal(t, expected, actual)
}

func TestParallelCircuits(t *testing.T) {
	// float64の精度の関係でコケるので小数点以下1桁に丸めてる
	actual := math.Round(ParallelCircuits(40, 40, 30) * 10 / 10)
	expected := 12.0

	assert.Equal(t, expected, actual)
}
