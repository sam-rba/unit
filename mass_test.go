// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"fmt"
	"testing"
)

func TestMass_String(t *testing.T) {
	if s := PoundMass.String(); s != "453.592g" {
		t.Fatalf("%#v", s)
	}
}

func TestMass_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Mass
	}{
		{"1ng", NanoGram},
		{"1ug", MicroGram},
		{"1µg", MicroGram},
		{"1mg", MilliGram},
		{"1g", Gram},
		{"1kg", KiloGram},
		{"1Mg", MegaGram},
		{"1Gg", GigaGram},
		{"1oz", OunceMass},
		{"1lb", PoundMass},
		// Maximum and minimum values that are allowed.
		{"9.223372036854775807Gg", 9223372036854775807},
		{"-9.223372036854775807Gg", -9223372036854775807},
		{"20334054lb", maxPoundMass * PoundMass},
		{"-20334054lb", minPoundMass * PoundMass},
		{"325344874oz", maxOunceMass * OunceMass},
		{"-325344874oz", minOunceMass * OunceMass},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Eg",
			"unknown unit prefix; valid prefixes for \"g\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10",
			"no unit provided; need g, lb or oz",
		},
		{
			fmt.Sprintf("%dlb", maxPoundMass+1),
			fmt.Sprintf("maximum value is %dlb", maxPoundMass),
		},
		{
			fmt.Sprintf("%dlb", minPoundMass-1),
			fmt.Sprintf("minimum value is %dlb", minPoundMass),
		},
		{
			fmt.Sprintf("%doz", maxOunceMass+1),
			fmt.Sprintf("maximum value is %doz", maxOunceMass),
		},
		{
			fmt.Sprintf("%doz", minOunceMass-1),
			fmt.Sprintf("minimum value is %doz", minOunceMass),
		},
		{
			fmt.Sprintf("%dlb", maxPoundMass+1),
			fmt.Sprintf("maximum value is %dlb", maxPoundMass),
		},
		{
			fmt.Sprintf("%dlb", minPoundMass-1),
			fmt.Sprintf("minimum value is %dlb", minPoundMass),
		},
		{
			"9.224Gg",
			"maximum value is 9.223Gg",
		},
		{
			"-9.224Gg",
			"minimum value is -9.223Gg",
		},
		{
			"9223372036854775808ng",
			"maximum value is 9.223Gg",
		},
		{
			"-9223372036854775808ng",
			"minimum value is -9.223Gg",
		},
		{
			"1random",
			"unknown unit provided; need g, lb or oz",
		},
		{
			"g",
			"not a number",
		},
		{
			"oz",
			"not a number",
		},
		{
			"lb",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit g, lb or oz",
		},
		{
			"++1g",
			"contains multiple plus symbols",
		},
		{
			"--1g",
			"contains multiple minus symbols",
		},
		{
			"+-1g",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1g",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
		{
			"20334055lb",
			"maximum value is 20334054lb",
		},
		{
			"325344875oz",
			"maximum value is 325344874oz",
		},
	}

	for i, tt := range succeeds {
		var got Mass
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Mass.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Mass.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Mass
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Mass.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestMass_RoundTrip(t *testing.T) {
	x := 123 * Gram
	var y Mass
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Mass.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Mass expected %s to equal %s", x, y)
	}
}
