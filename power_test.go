// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestPower_String(t *testing.T) {
	if s := NanoWatt.String(); s != "1nW" {
		t.Fatalf("%v", s)
	}
	if s := MicroWatt.String(); s != "1µW" {
		t.Fatalf("%v", s)
	}
	if s := MilliWatt.String(); s != "1mW" {
		t.Fatalf("%v", s)
	}
	if s := Watt.String(); s != "1W" {
		t.Fatalf("%v", s)
	}
	if s := KiloWatt.String(); s != "1kW" {
		t.Fatalf("%v", s)
	}
	if s := MegaWatt.String(); s != "1MW" {
		t.Fatalf("%v", s)
	}
	if s := GigaWatt.String(); s != "1GW" {
		t.Fatalf("%v", s)
	}
}

func TestPower_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Power
	}{
		{"1nW", 1 * NanoWatt},
		{"10nW", 10 * NanoWatt},
		{"100nW", 100 * NanoWatt},
		{"1uW", 1 * MicroWatt},
		{"10uW", 10 * MicroWatt},
		{"100uW", 100 * MicroWatt},
		{"1µW", 1 * MicroWatt},
		{"10µW", 10 * MicroWatt},
		{"100µW", 100 * MicroWatt},
		{"1mW", 1 * MilliWatt},
		{"10mW", 10 * MilliWatt},
		{"100mW", 100 * MilliWatt},
		{"1w", 1 * Watt},
		{"1W", 1 * Watt},
		{"10W", 10 * Watt},
		{"100W", 100 * Watt},
		{"1kW", 1 * KiloWatt},
		{"1kw", 1 * KiloWatt},
		{"10kW", 10 * KiloWatt},
		{"100kW", 100 * KiloWatt},
		{"1MW", 1 * MegaWatt},
		{"10MW", 10 * MegaWatt},
		{"100MW", 100 * MegaWatt},
		{"1GW", 1 * GigaWatt},
		{"12.345W", 12345 * MilliWatt},
		{"-12.345W", -12345 * MilliWatt},
		{"9.223372036854775807GW", 9223372036854775807 * NanoWatt},
		{"-9.223372036854775807GW", -9223372036854775807 * NanoWatt},
		{"1MW", 1 * MegaWatt},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TW",
			"maximum value is 9.223GW",
		},
		{
			"10EW",
			"unknown unit prefix; valid prefixes for \"W\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaW",
			"unknown unit prefix; valid prefixes for \"W\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eWattE",
			"unknown unit provided; need W",
		},
		{
			"10",
			"no unit provided; need W",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GW",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GW",
		},
		{
			"9.223372036854775808GW",
			"maximum value is 9.223GW",
		},
		{
			"-9.223372036854775808GW",
			"minimum value is -9.223GW",
		},
		{
			"9.223372036854775808GW",
			"maximum value is 9.223GW",
		},
		{
			"-9.223372036854775808GW",
			"minimum value is -9.223GW",
		},
		{
			"1random",
			"unknown unit provided; need W",
		},
		{
			"W",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit W",
		},
		{
			"++1W",
			"contains multiple plus symbols",
		},
		{
			"--1W",
			"contains multiple minus symbols",
		},
		{
			"+-1W",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1W",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got Power
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Power.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Power.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Power
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Power.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestPower_RoundTrip(t *testing.T) {
	x := 123 * Watt
	var y Power
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Power.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Power expected %s to equal %s", x, y)
	}
}
