// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestCapacitance_String(t *testing.T) {
	if s := PicoFarad.String(); s != "1pF" {
		t.Fatalf("%v", s)
	}
	if s := NanoFarad.String(); s != "1nF" {
		t.Fatalf("%v", s)
	}
	if s := MicroFarad.String(); s != "1µF" {
		t.Fatalf("%v", s)
	}
	if s := MilliFarad.String(); s != "1mF" {
		t.Fatalf("%v", s)
	}
	if s := Farad.String(); s != "1F" {
		t.Fatalf("%v", s)
	}
	if s := KiloFarad.String(); s != "1kF" {
		t.Fatalf("%v", s)
	}
	if s := MegaFarad.String(); s != "1MF" {
		t.Fatalf("%v", s)
	}
}

func TestElectricalCapacitance_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected ElectricalCapacitance
	}{
		{"1pF", 1 * PicoFarad},
		{"10pF", 10 * PicoFarad},
		{"100pF", 100 * PicoFarad},
		{"1nF", 1 * NanoFarad},
		{"10nF", 10 * NanoFarad},
		{"100nF", 100 * NanoFarad},
		{"1uF", 1 * MicroFarad},
		{"10uF", 10 * MicroFarad},
		{"100uF", 100 * MicroFarad},
		{"1µF", 1 * MicroFarad},
		{"10µF", 10 * MicroFarad},
		{"100µF", 100 * MicroFarad},
		{"1mF", 1 * MilliFarad},
		{"10mF", 10 * MilliFarad},
		{"100mF", 100 * MilliFarad},
		{"1F", 1 * Farad},
		{"1f", 1 * Farad},
		{"10F", 10 * Farad},
		{"100F", 100 * Farad},
		{"1kF", 1 * KiloFarad},
		{"1kf", 1 * KiloFarad},
		{"10kF", 10 * KiloFarad},
		{"100kF", 100 * KiloFarad},
		{"1MF", 1 * MegaFarad},
		{"12.345F", 12345 * MilliFarad},
		{"-12.345F", -12345 * MilliFarad},
		{"9.223372036854775807MF", 9223372036854775807 * PicoFarad},
		{"-9.223372036854775807MF", -9223372036854775807 * PicoFarad},
		{"1MF", 1 * MegaFarad},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TF",
			"maximum value is 9.223MF",
		},
		{
			"10EF",
			"unknown unit prefix; valid prefixes for \"F\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaF",
			"unknown unit prefix; valid prefixes for \"F\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eFaradE",
			"unknown unit provided; need F",
		},
		{
			"10",
			"no unit provided; need F",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223MF",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223MF",
		},
		{
			"9.223372036854775808MF",
			"maximum value is 9.223MF",
		},
		{
			"-9.223372036854775808MF",
			"minimum value is -9.223MF",
		},
		{
			"9.223372036854775808MF",
			"maximum value is 9.223MF",
		},
		{
			"-9.223372036854775808MF",
			"minimum value is -9.223MF",
		},
		{
			"1random",
			"unknown unit provided; need F",
		},
		{
			"F",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit F",
		},
		{
			"++1F",
			"contains multiple plus symbols",
		},
		{
			"--1F",
			"contains multiple minus symbols",
		},
		{
			"+-1F",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1F",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got ElectricalCapacitance
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: ElectricalCapacitance.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: ElectricalCapacitance.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got ElectricalCapacitance
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: ElectricalCapacitance.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestElectricalCapacitance_RoundTrip(t *testing.T) {
	x := 123 * Farad
	var y ElectricalCapacitance
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("ElectricalCapacitance.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("ElectricalCapacitance expected %s to equal %s", x, y)
	}
}
