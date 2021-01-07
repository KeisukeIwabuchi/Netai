package parts

// TransistorType is Enum for selecting the type of transistor
type TransistorType int

// NPN or PNP
const (
	NPN TransistorType = 1
	PNP TransistorType = -1
)

// Transistor is electonic component
type Transistor struct {
	Type      TransistorType
	Emitter   float64
	Collector float64
	Base      float64
}

// EmitterVoltage is return emitter voltage
func (t Transistor) EmitterVoltage() float64 {
	if t.Type == PNP {
		return t.Emitter
	}

	voltage := 0.0

	if t.Base > 0 {
		voltage = t.Base

		if t.Collector > 0 && t.Collector > t.Base {
			voltage = t.Collector
		}
	}

	return voltage
}
