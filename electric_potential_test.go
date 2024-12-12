// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestElectricPotential_String(t *testing.T) {
	if s := Volt.String(); s != "1V" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricPotential_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected ElectricPotential
	}{
		{"1nV", 1 * NanoVolt},
		{"10nV", 10 * NanoVolt},
		{"100nV", 100 * NanoVolt},
		{"1uV", 1 * MicroVolt},
		{"10uV", 10 * MicroVolt},
		{"100uV", 100 * MicroVolt},
		{"1µV", 1 * MicroVolt},
		{"10µV", 10 * MicroVolt},
		{"100µV", 100 * MicroVolt},
		{"1mV", 1 * MilliVolt},
		{"10mV", 10 * MilliVolt},
		{"100mV", 100 * MilliVolt},
		{"1V", 1 * Volt},
		{"1v", 1 * Volt},
		{"10V", 10 * Volt},
		{"100V", 100 * Volt},
		{"1kV", 1 * KiloVolt},
		{"1kv", 1 * KiloVolt},
		{"10kV", 10 * KiloVolt},
		{"100kV", 100 * KiloVolt},
		{"1MV", 1 * MegaVolt},
		{"10MV", 10 * MegaVolt},
		{"100MV", 100 * MegaVolt},
		{"1GV", 1 * GigaVolt},
		{"12.345V", 12345 * MilliVolt},
		{"-12.345V", -12345 * MilliVolt},
		{"9.223372036854775807GV", 9223372036854775807 * NanoVolt},
		{"-9.223372036854775807GV", -9223372036854775807 * NanoVolt},
		{"1MV", 1 * MegaVolt},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TV",
			"maximum value is 9.223GV",
		},
		{
			"10EV",
			"unknown unit prefix; valid prefixes for \"V\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eVoltE",
			"unknown unit provided; need V",
		},
		{
			"10",
			"no unit provided; need V",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GV",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GV",
		},
		{
			"9.223372036854775808TV",
			"maximum value is 9.223GV",
		},
		{
			"-9.223372036854775808GV",
			"minimum value is -9.223GV",
		},
		{
			"9.223372036854775808GV",
			"maximum value is 9.223GV",
		},
		{
			"-9.223372036854775808GOhm",
			"minimum value is -9.223GV",
		},
		{
			"1random",
			"unknown unit provided; need V",
		},
		{
			"V",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit V",
		},
		{
			"++1V",
			"contains multiple plus symbols",
		},
		{
			"--1V",
			"contains multiple minus symbols",
		},
		{
			"+-1V",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1V",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got ElectricPotential
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: ElectricPotential.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: ElectricPotential.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got ElectricPotential
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: ElectricPotential.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestElectricPotential_RoundTrip(t *testing.T) {
	x := 123 * Volt
	var y ElectricPotential
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("ElectricPotential.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("ElectricPotential expected %s to equal %s", x, y)
	}
}
