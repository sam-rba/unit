// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestElectricResistance_String(t *testing.T) {
	if s := Ohm.String(); s != "1Ω" {
		t.Fatalf("%#v", s)
	}
}

func TestElectricResistance_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected ElectricResistance
	}{
		{"1nOhm", 1 * NanoOhm},
		{"10nOhm", 10 * NanoOhm},
		{"100nOhm", 100 * NanoOhm},
		{"1uOhm", 1 * MicroOhm},
		{"10uOhm", 10 * MicroOhm},
		{"100uOhm", 100 * MicroOhm},
		{"1µOhm", 1 * MicroOhm},
		{"10µOhm", 10 * MicroOhm},
		{"100µOhm", 100 * MicroOhm},
		{"1mOhm", 1 * MilliOhm},
		{"10mOhm", 10 * MilliOhm},
		{"100mOhm", 100 * MilliOhm},
		{"1Ohm", 1 * Ohm},
		{"1ohm", 1 * Ohm},
		{"10Ohm", 10 * Ohm},
		{"100Ohm", 100 * Ohm},
		{"1kOhm", 1 * KiloOhm},
		{"10kOhm", 10 * KiloOhm},
		{"100kOhm", 100 * KiloOhm},
		{"1MOhm", 1 * MegaOhm},
		{"10MOhm", 10 * MegaOhm},
		{"100MOhm", 100 * MegaOhm},
		{"1GOhm", 1 * GigaOhm},
		{"12.345Ohm", 12345 * MilliOhm},
		{"-12.345Ohm", -12345 * MilliOhm},
		{"9.223372036854775807GOhm", 9223372036854775807 * NanoOhm},
		{"-9.223372036854775807GOhm", -9223372036854775807 * NanoOhm},
		{"1MΩ", 1 * MegaOhm},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TOhm",
			"maximum value is 9.223GΩ",
		},
		{
			"10EOhm",
			"unknown unit prefix; valid prefixes for \"Ohm\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaOhm",
			"unknown unit prefix; valid prefixes for \"Ohm\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eOhmE",
			"unknown unit provided; need Ohm or Ω",
		},
		{
			"10",
			"no unit provided; need Ohm or Ω",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GΩ",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GΩ",
		},
		{
			"9.223372036854775808GOhm",
			"maximum value is 9.223GΩ",
		},
		{
			"-9.223372036854775808GOhm",
			"minimum value is -9.223GΩ",
		},
		{
			"9.223372036854775808GOhm",
			"maximum value is 9.223GΩ",
		},
		{
			"-9.223372036854775808GOhm",
			"minimum value is -9.223GΩ",
		},
		{
			"1random",
			"unknown unit provided; need Ohm or Ω",
		},
		{
			"Ohm",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit Ohm or Ω",
		},
		{
			"++1Ohm",
			"contains multiple plus symbols",
		},
		{
			"--1Ohm",
			"contains multiple minus symbols",
		},
		{
			"+-1Ohm",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1Ohm",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got ElectricResistance
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: ElectricResistance.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: ElectricResistance.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got ElectricResistance
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: ElectricResistance.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestElectricResistance_RoundTrip(t *testing.T) {
	x := 123 * Ohm
	var y ElectricResistance
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("ElectricResistance.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("ElectricResistance expected %s to equal %s", x, y)
	}
}
