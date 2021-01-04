package ohmslaw

// CalculatecVoltage is return Voltage
func CalculatecVoltage(current, resistance float64) float64 {
	return current * resistance
}

// CalculateCurrent is return Current
func CalculateCurrent(voltage, resistance float64) float64 {
	if resistance <= 0 {
		panic("Invalid resistance.")
	}
	return voltage / resistance
}

// CalculateResistance is return Resistance
func CalculateResistance(voltage, current float64) float64 {
	if current <= 0 {
		panic("Invalid current.")
	}
	return voltage / current
}
