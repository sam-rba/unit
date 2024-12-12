package unit

import "testing"

func TestTemperature_String(t *testing.T) {
	if s := ZeroCelsius.String(); s != "0°C" {
		t.Fatalf("%#v", s)
	}
	if s := Temperature(0).String(); s != "-273.150°C" {
		t.Fatalf("%#v", s)
	}
}

func TestTemperature_K(t *testing.T) {
	if v := Temperature(0).K(); v != 0. {
		t.Fatal(v)
	}
	if v := Temperature(123 * Kelvin).K(); v != 123. {
		t.Fatal(v)
	}
}

func TestTemperature_C(t *testing.T) {
	if v := ZeroCelsius.C(); v != 0. {
		t.Fatal(v)
	}
	if v := Temperature(0).C(); v != -273.150 {
		t.Fatal(v)
	}
}

func TestTemperature_F(t *testing.T) {
	if v := ZeroFahrenheit.F(); v != 0. {
		t.Fatal(v)
	}
	if v := Temperature(0).F(); v != -459.67000045927 {
		t.Fatal(v)
	}
}
