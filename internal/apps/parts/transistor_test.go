package parts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmitterVoltage(t *testing.T) {
	transistor := new(Transistor)

	// エミッタ-コレクタ間とベース-エミッタ間の両方に電圧をかけた場合
	transistor.Collector = 5.0
	transistor.Base = 0.1
	actual := transistor.EmitterVoltage()
	expected := 5.0
	assert.Equal(t, expected, actual)

	// ベース-エミッタ間にだけ電圧をかけた場合
	transistor.Collector = 0.0
	actual = transistor.EmitterVoltage()
	expected = 0.1
	assert.Equal(t, expected, actual)

	// エミッタ-コレクタ間にだけ電圧をかけた場合
	transistor.Collector = 5.0
	transistor.Base = 0.0
	actual = transistor.EmitterVoltage()
	expected = 0.0
	assert.Equal(t, expected, actual)
}
