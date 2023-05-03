package convert

import "testing"

func TestToVolts(t *testing.T) {
	var in uint16 = 0
	var out = ToVolts(in)
	if out != 0 {
		t.Errorf("Expected ZERO volts but got %f", out)
	}

	in = 1
	out = ToVolts(in)
	if out != 0.05 {
		t.Errorf("Expected 0.05 volts but got %f", out)
	}

}
func TestToCurrent(t *testing.T) {
	var in = ZERO_AMPS
	var out = ToCurrent(in)
	if out != 0 {
		t.Errorf("Expected ZERO amps but got %f", out)
	}

	out = ToCurrent(in + 1)
	if out != 0.001 {
		t.Errorf("Expected 0.001 amps but got %f", out)
	}

	out = ToCurrent(in - 1)
	if out != -0.001 {
		t.Errorf("Expected 0.001 amps but got %f", out)
	}

}

func TestToPercent(t *testing.T) {
	var in uint8 = 0x01
	var out = ToPercent(in)
	if out != 0.5 {
		t.Errorf("Expected 0.5 volts but got %f", out)
	}
}

func TestFromPercent(t *testing.T) {
	var v = 0.5
	var out = FromPercent(v)
	if out != 0x01 {
		t.Errorf("Expected 1  but got %x", out)
	}

	v = 126.0
	out = FromPercent(v)
	if out != 250 {
		t.Errorf("Expected 250  but got %x", out)
	}

	v = -1.0
	out = FromPercent(v)
	if out != 0 {
		t.Errorf("Expected 0  but got %x", out)
	}
}

func TestToDegreesC(t *testing.T) {
	{
		var in uint8 = 0x0A
		var out = ToDegreesC(in)
		if out != 1 {
			t.Errorf("Expected 1 degrees but got %f", out)
		}
	}

	{
		// uint16 -273 1735 0.03125 °C -
		var in uint16 = 0x1
		var out = ToDegreesC16(in)
		if out != (-273 + 0.03125) {
			t.Errorf("Expected 0.3125 degrees but got %f", out)
		}
	}

	{
		// uint16 -273 1735 0.03125 °C -
		var in float64 = 1
		var out = CelsiusToFahrenheit(in)
		if out != (33.8) {
			t.Errorf("Expected 33.8 degrees F but got %f", out)
		}
	}
}
