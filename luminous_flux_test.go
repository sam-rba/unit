// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestFlux_String(t *testing.T) {
	if s := NanoLumen.String(); s != "1nlm" {
		t.Fatalf("%v", s)
	}
	if s := MicroLumen.String(); s != "1µlm" {
		t.Fatalf("%v", s)
	}
	if s := MilliLumen.String(); s != "1mlm" {
		t.Fatalf("%v", s)
	}
	if s := Lumen.String(); s != "1lm" {
		t.Fatalf("%v", s)
	}
	if s := KiloLumen.String(); s != "1klm" {
		t.Fatalf("%v", s)
	}
	if s := MegaLumen.String(); s != "1Mlm" {
		t.Fatalf("%v", s)
	}
	if s := GigaLumen.String(); s != "1Glm" {
		t.Fatalf("%v", s)
	}
}

func TestLuminousFlux_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected LuminousFlux
	}{
		{"1nlm", 1 * NanoLumen},
		{"10nlm", 10 * NanoLumen},
		{"100nlm", 100 * NanoLumen},
		{"1ulm", 1 * MicroLumen},
		{"10ulm", 10 * MicroLumen},
		{"100ulm", 100 * MicroLumen},
		{"1µlm", 1 * MicroLumen},
		{"10µlm", 10 * MicroLumen},
		{"100µlm", 100 * MicroLumen},
		{"1mlm", 1 * MilliLumen},
		{"10mlm", 10 * MilliLumen},
		{"100mlm", 100 * MilliLumen},
		{"1lm", 1 * Lumen},
		{"10lm", 10 * Lumen},
		{"100lm", 100 * Lumen},
		{"1klm", 1 * KiloLumen},
		{"10klm", 10 * KiloLumen},
		{"100klm", 100 * KiloLumen},
		{"1Mlm", 1 * MegaLumen},
		{"10Mlm", 10 * MegaLumen},
		{"100Mlm", 100 * MegaLumen},
		{"1Glm", 1 * GigaLumen},
		{"12.345lm", 12345 * MilliLumen},
		{"-12.345lm", -12345 * MilliLumen},
		{"9.223372036854775807Glm", 9223372036854775807 * NanoLumen},
		{"-9.223372036854775807Glm", -9223372036854775807 * NanoLumen},
		{"1Mlm", 1 * MegaLumen},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Tlm",
			"maximum value is 9.223Glm",
		},
		{
			"10Elm",
			"unknown unit prefix; valid prefixes for \"lm\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10Exalm",
			"unknown unit prefix; valid prefixes for \"lm\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10elmE",
			"unknown unit provided; need lm",
		},
		{
			"10",
			"no unit provided; need lm",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223Glm",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223Glm",
		},
		{
			"9.223372036854775808Glm",
			"maximum value is 9.223Glm",
		},
		{
			"-9.223372036854775808Glm",
			"minimum value is -9.223Glm",
		},
		{
			"9.223372036854775808Glm",
			"maximum value is 9.223Glm",
		},
		{
			"-9.223372036854775808Glm",
			"minimum value is -9.223Glm",
		},
		{
			"1random",
			"unknown unit provided; need lm",
		},
		{
			"lm",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit lm",
		},
		{
			"++1lm",
			"contains multiple plus symbols",
		},
		{
			"--1lm",
			"contains multiple minus symbols",
		},
		{
			"+-1lm",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1lm",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got LuminousFlux
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: LuminousFlux.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: LuminousFlux.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got LuminousFlux
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: LuminousFlux.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestLuminousFlux_RoundTrip(t *testing.T) {
	x := 123 * Lumen
	var y LuminousFlux
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("LuminousFlux.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("LuminousFlux expected %s to equal %s", x, y)
	}
}
