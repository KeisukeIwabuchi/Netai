// Package combresistance calculate combined resistance
package combresistance

// SeriesCircuits is return combined resistance from series circuits case.
func SeriesCircuits(resistance ...float64) float64 {
	var sum float64

	for _, value := range resistance {
		sum += value
	}

	return sum
}

// ParallelCircuits is return combined resistance from parallel circuits.
func ParallelCircuits(resistance ...float64) float64 {
	var sum float64

	for _, value := range resistance {
		if value <= 0 {
			panic("Invalid resistance value.")
		}
		sum += (1 / value)
	}

	return 1 / sum
}
