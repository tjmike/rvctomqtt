package convert

const ZERO_AMPS uint32 = 0x77359400
const ZERO_AMPSF = float64(ZERO_AMPS)

// uint32 –2,000,000.000A 2,221,081.200A 0.001 A 0A = 0x77359400
func ToCurrent(in uint32) float64 {
	var out = (float64(in) - ZERO_AMPSF) * 0.001
	return out

}

// uint16 , 0,3212.5 , 0.050 V
func ToVolts(in uint16) float64 {
	var out = float64(in) * 0.05
	return out
}
func ToPercent(in uint8) float64 {
	var out = float64(in) / 2
	return out
}
func FromPercent(in float64) uint8 {
	// The max range for percent is 0-125 so we clip to that range before conversion
	var clipped = in
	if clipped > 125 {
		clipped = 125
	} else if clipped < 0 {
		clipped = 0
	}
	return uint8(clipped * 2)
}

func ToDegreesC(in uint8) float64 {
	var out = float64(in) / 10.0
	return out
}

// -273 -> 1735 (0.03125 °C ) (1/32)
func ToDegreesC16(in uint16) float64 {
	var out = (float64(in) / 32) - 273
	return out
}

func CelsiusToFahrenheit(in float64) float64 {
	return (in * 1.8) + 32.0
}
