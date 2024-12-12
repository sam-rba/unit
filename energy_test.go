// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestEnergy_String(t *testing.T) {
	if s := NanoJoule.String(); s != "1nJ" {
		t.Fatalf("%v", s)
	}
	if s := MicroJoule.String(); s != "1µJ" {
		t.Fatalf("%v", s)
	}
	if s := MilliJoule.String(); s != "1mJ" {
		t.Fatalf("%v", s)
	}
	if s := Joule.String(); s != "1J" {
		t.Fatalf("%v", s)
	}
	if s := KiloJoule.String(); s != "1kJ" {
		t.Fatalf("%v", s)
	}
	if s := MegaJoule.String(); s != "1MJ" {
		t.Fatalf("%v", s)
	}
	if s := GigaJoule.String(); s != "1GJ" {
		t.Fatalf("%v", s)
	}
}

func TestEnergy_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Energy
	}{
		{"1nJ", 1 * NanoJoule},
		{"10nJ", 10 * NanoJoule},
		{"100nJ", 100 * NanoJoule},
		{"1uJ", 1 * MicroJoule},
		{"10uJ", 10 * MicroJoule},
		{"100uJ", 100 * MicroJoule},
		{"1µJ", 1 * MicroJoule},
		{"10µJ", 10 * MicroJoule},
		{"100µJ", 100 * MicroJoule},
		{"1mJ", 1 * MilliJoule},
		{"10mJ", 10 * MilliJoule},
		{"100mJ", 100 * MilliJoule},
		{"1J", 1 * Joule},
		{"1j", 1 * Joule},
		{"10J", 10 * Joule},
		{"100J", 100 * Joule},
		{"1kJ", 1 * KiloJoule},
		{"1kj", 1 * KiloJoule},
		{"10kJ", 10 * KiloJoule},
		{"100kJ", 100 * KiloJoule},
		{"1MJ", 1 * MegaJoule},
		{"10MJ", 10 * MegaJoule},
		{"100MJ", 100 * MegaJoule},
		{"1GJ", 1 * GigaJoule},
		{"12.345J", 12345 * MilliJoule},
		{"-12.345J", -12345 * MilliJoule},
		{"9.223372036854775807GJ", 9223372036854775807 * NanoJoule},
		{"-9.223372036854775807GJ", -9223372036854775807 * NanoJoule},
		{"1MJ", 1 * MegaJoule},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TJ",
			"maximum value is 9.223GJ",
		},
		{
			"10EJ",
			"unknown unit prefix; valid prefixes for \"J\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaJ",
			"unknown unit prefix; valid prefixes for \"J\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eJouleE",
			"unknown unit provided; need J",
		},
		{
			"10",
			"no unit provided; need J",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GJ",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GJ",
		},
		{
			"9.223372036854775808GJ",
			"maximum value is 9.223GJ",
		},
		{
			"-9.223372036854775808GJ",
			"minimum value is -9.223GJ",
		},
		{
			"9.223372036854775808GJ",
			"maximum value is 9.223GJ",
		},
		{
			"-9.223372036854775808GJ",
			"minimum value is -9.223GJ",
		},
		{
			"1random",
			"unknown unit provided; need J",
		},
		{
			"J",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit J",
		},
		{
			"++1J",
			"contains multiple plus symbols",
		},
		{
			"--1J",
			"contains multiple minus symbols",
		},
		{
			"+-1J",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1J",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got Energy
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Energy.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Energy.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Energy
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Energy.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestEnergy_RoundTrip(t *testing.T) {
	x := 123 * Joule
	var y Energy
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Energy.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Energy expected %s to equal %s", x, y)
	}
}
