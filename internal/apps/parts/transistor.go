package parts

// Transistor is electonic component
type Transistor struct {
	Collector float64
	Base      float64
}

// EmitterVoltage is return emitter voltage
func (t Transistor) EmitterVoltage() float64 {
	emitter := 0.0

	if t.Base > 0 {
		emitter = t.Base

		if t.Collector > 0 && t.Collector > t.Base {
			emitter = t.Collector
		}
	}

	return emitter
}
