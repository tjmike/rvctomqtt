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
