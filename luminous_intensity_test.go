// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestLuminousIntensity_String(t *testing.T) {
	if s := NanoCandela.String(); s != "1ncd" {
		t.Fatalf("%v", s)
	}
	if s := MicroCandela.String(); s != "1µcd" {
		t.Fatalf("%v", s)
	}
	if s := MilliCandela.String(); s != "1mcd" {
		t.Fatalf("%v", s)
	}
	if s := Candela.String(); s != "1cd" {
		t.Fatalf("%v", s)
	}
	if s := KiloCandela.String(); s != "1kcd" {
		t.Fatalf("%v", s)
	}
	if s := MegaCandela.String(); s != "1Mcd" {
		t.Fatalf("%v", s)
	}
	if s := GigaCandela.String(); s != "1Gcd" {
		t.Fatalf("%v", s)
	}
}

func TestLuminousIntensity_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected LuminousIntensity
	}{
		{"1ncd", 1 * NanoCandela},
		{"10ncd", 10 * NanoCandela},
		{"100ncd", 100 * NanoCandela},
		{"1ucd", 1 * MicroCandela},
		{"10ucd", 10 * MicroCandela},
		{"100ucd", 100 * MicroCandela},
		{"1µcd", 1 * MicroCandela},
		{"10µcd", 10 * MicroCandela},
		{"100µcd", 100 * MicroCandela},
		{"1mcd", 1 * MilliCandela},
		{"10mcd", 10 * MilliCandela},
		{"100mcd", 100 * MilliCandela},
		{"1cd", 1 * Candela},
		{"10cd", 10 * Candela},
		{"100cd", 100 * Candela},
		{"1kcd", 1 * KiloCandela},
		{"10kcd", 10 * KiloCandela},
		{"100kcd", 100 * KiloCandela},
		{"1Mcd", 1 * MegaCandela},
		{"10Mcd", 10 * MegaCandela},
		{"100Mcd", 100 * MegaCandela},
		{"1Gcd", 1 * GigaCandela},
		{"12.345cd", 12345 * MilliCandela},
		{"-12.345cd", -12345 * MilliCandela},
		{"9.223372036854775807Gcd", 9223372036854775807 * NanoCandela},
		{"-9.223372036854775807Gcd", -9223372036854775807 * NanoCandela},
		{"1Mcd", 1 * MegaCandela},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Tcd",
			"maximum value is 9.223Gcd",
		},
		{
			"10Ecd",
			"unknown unit prefix; valid prefixes for \"cd\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10Exacd",
			"unknown unit prefix; valid prefixes for \"cd\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ecdE",
			"unknown unit provided; need cd",
		},
		{
			"10",
			"no unit provided; need cd",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223Gcd",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223Gcd",
		},
		{
			"9.223372036854775808Gcd",
			"maximum value is 9.223Gcd",
		},
		{
			"-9.223372036854775808Gcd",
			"minimum value is -9.223Gcd",
		},
		{
			"9.223372036854775808Gcd",
			"maximum value is 9.223Gcd",
		},
		{
			"-9.223372036854775808Gcd",
			"minimum value is -9.223Gcd",
		},
		{
			"1random",
			"unknown unit provided; need cd",
		},
		{
			"cd",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit cd",
		},
		{
			"++1cd",
			"contains multiple plus symbols",
		},
		{
			"--1cd",
			"contains multiple minus symbols",
		},
		{
			"+-1cd",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1cd",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got LuminousIntensity
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: LuminousIntensity.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: LuminousIntensity.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got LuminousIntensity
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: LuminousIntensity.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestLuminousIntensity_RoundTrip(t *testing.T) {
	x := 123 * Candela
	var y LuminousIntensity
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("LuminousIntensity.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("LuminousIntensity expected %s to equal %s", x, y)
	}
}
