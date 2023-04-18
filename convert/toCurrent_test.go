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
