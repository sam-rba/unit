package unit

import (
	"fmt"
	"testing"
)

func TestTemperature_String(t *testing.T) {
	if s := ZeroCelsius.String(); s != "0°C" {
		t.Fatalf("%#v", s)
	}
	if s := Temperature(0).String(); s != "-273.150°C" {
		t.Fatalf("%#v", s)
	}
}

func TestTemperature_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Temperature
	}{
		{"0C", ZeroCelsius},
		{"0K", 0},
		{"0F", ZeroFahrenheit},
		{"1K", Kelvin},
		{"100C", ZeroCelsius + 100*Celsius},
		{"-40F", ZeroCelsius - 40*Celsius},
		{fmt.Sprintf("%dnC", int64(maxCelsius)), ZeroCelsius + maxCelsius},
		{"-273.15C", 0},
		{fmt.Sprintf("%dnK", int64(maxTemperature)), maxTemperature},
		{fmt.Sprintf("%dnK", int64(minTemperature)), 0},
		{fmt.Sprintf("%dF", int64(maxFahrenheit)), 9223372033887869742},
		{"-459.67F", 0},
		{"1GK", GigaKelvin},
		{"1kC", ZeroCelsius + 1000*Celsius},
		{"16kF", 9144261111118},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"-1nK",
			"minimum value is 0K",
		},
		{
			fmt.Sprintf("%dnC", int64(maxCelsius+1)),
			"maximum value is 9223371763°C",
		},
		{
			fmt.Sprintf("%dnC", int64(-ZeroCelsius-1)),
			"minimum value is -273.15°C",
		},
		{
			"9223372036854775808nK",
			"maximum value is 9.223GK",
		},
		{
			"-9223372036854775808nK",
			"minimum value is -273.150°C",
		},
		{
			fmt.Sprintf("%dF", int64(maxFahrenheit+1)),
			"maximum value is 16602069204F",
		},
		{
			"-459.671F",
			"minimum value is -459.67F",
		},
		{
			fmt.Sprintf("%dF", int64(maxCelsius)),
			"maximum value is 16602069204F",
		},
		{
			"-273.151C",
			"minimum value is -273.15°C",
		},
		{
			"9.224GK",
			"maximum value is 9223372036K",
		},
		{
			"-9.224GK",
			"minimum value is 0K",
		},
		{
			"9.224GC",
			"maximum value is 9223371763°C",
		},
		{
			"-9.224GC",
			"minimum value is -273.15°C",
		},
		{
			"-9.224TF",
			"minimum value is -459.67F",
		},
		{
			"10E°C",
			"unknown unit prefix; valid prefixes for \"°C\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10",
			"no unit provided; need K, °C, C, °F or F",
		},
		{
			"1random",
			"unknown unit provided; need K, °C, C, °F or F",
		},
		{
			"C",
			"not a number",
		},
		{
			"°C",
			"not a number",
		},
		{
			"K",
			"not a number",
		},
		{
			"F",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit K, °C, C, °F or F",
		},
		{
			"++1°C",
			"contains multiple plus symbols",
		},
		{
			"--1°C",
			"contains multiple minus symbols",
		},
		{
			"+-1°C",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1°C",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Temperature
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Temperature.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Temperature.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Temperature
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Temperature.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestTemperature_RoundTrip(t *testing.T) {
	x := 123 * Celsius
	var y Temperature
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Temperature.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Temperature expected %s to equal %s", x, y)
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
