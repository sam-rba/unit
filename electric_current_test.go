// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestElectricCurrent_String(t *testing.T) {
	if s := Ampere.String(); s != "1A" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricCurrent_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected ElectricCurrent
	}{
		{"1nA", 1 * NanoAmpere},
		{"10nA", 10 * NanoAmpere},
		{"100nA", 100 * NanoAmpere},
		{"1uA", 1 * MicroAmpere},
		{"10uA", 10 * MicroAmpere},
		{"100uA", 100 * MicroAmpere},
		{"1µA", 1 * MicroAmpere},
		{"10µA", 10 * MicroAmpere},
		{"100µA", 100 * MicroAmpere},
		{"1mA", 1 * MilliAmpere},
		{"10mA", 10 * MilliAmpere},
		{"100mA", 100 * MilliAmpere},
		{"1A", 1 * Ampere},
		{"1a", 1 * Ampere},
		{"10A", 10 * Ampere},
		{"100A", 100 * Ampere},
		{"1kA", 1 * KiloAmpere},
		{"1ka", 1 * KiloAmpere},
		{"10kA", 10 * KiloAmpere},
		{"100kA", 100 * KiloAmpere},
		{"1MA", 1 * MegaAmpere},
		{"10MA", 10 * MegaAmpere},
		{"100MA", 100 * MegaAmpere},
		{"1GA", 1 * GigaAmpere},
		{"12.345A", 12345 * MilliAmpere},
		{"-12.345A", -12345 * MilliAmpere},
		{"9.223372036854775807GA", 9223372036854775807 * NanoAmpere},
		{"-9.223372036854775807GA", -9223372036854775807 * NanoAmpere},
		{"1A", 1 * Ampere},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TA",
			"maximum value is 9.223GA",
		},
		{
			"10EA",
			"unknown unit prefix; valid prefixes for \"A\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eAmpE",
			"unknown unit provided; need A",
		},
		{
			"10",
			"no unit provided; need A",
		},
		{
			"922337203685477580",
			"maximum value is 9.223GA",
		},
		{
			"-922337203685477580",
			"minimum value is -9.223GA",
		},
		{
			"9.223372036854775808GA",
			"maximum value is 9.223GA",
		},
		{
			"-9.223372036854775808GA",
			"minimum value is -9.223GA",
		},
		{
			"9.223372036854775808GA",
			"maximum value is 9.223GA",
		},
		{
			"-9.223372036854775808GA",
			"minimum value is -9.223GA",
		},
		{
			"1random",
			"unknown unit provided; need A",
		},
		{
			"A",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit A",
		},
		{
			"++1A",
			"contains multiple plus symbols",
		},
		{
			"--1A",
			"contains multiple minus symbols",
		},
		{
			"+-1A",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1A",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got ElectricCurrent
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: ElectricCurrent.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: ElectricCurrent.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got ElectricCurrent
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: ElectricCurrent.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestElectricCurrent_RoundTrip(t *testing.T) {
	x := 123 * Ampere
	var y ElectricCurrent
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("ElectricCurrent.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("ElectricCurrent expected %s to equal %s", x, y)
	}
}
