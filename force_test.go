// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestForce_String(t *testing.T) {
	if s := Newton.String(); s != "1N" {
		t.Fatalf("%#v", s)
	}
}

func TestForce_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Force
	}{
		{"1nN", 1 * NanoNewton},
		{"10nN", 10 * NanoNewton},
		{"100nN", 100 * NanoNewton},
		{"1uN", 1 * MicroNewton},
		{"10uN", 10 * MicroNewton},
		{"100uN", 100 * MicroNewton},
		{"1µN", 1 * MicroNewton},
		{"10µN", 10 * MicroNewton},
		{"100µN", 100 * MicroNewton},
		{"1mN", 1 * MilliNewton},
		{"10mN", 10 * MilliNewton},
		{"100mN", 100 * MilliNewton},
		{"1N", 1 * Newton},
		{"10N", 10 * Newton},
		{"100N", 100 * Newton},
		{"1kN", 1 * KiloNewton},
		{"10kN", 10 * KiloNewton},
		{"100kN", 100 * KiloNewton},
		{"1MN", 1 * MegaNewton},
		{"10MN", 10 * MegaNewton},
		{"100MN", 100 * MegaNewton},
		{"1GN", 1 * GigaNewton},
		{"12.345N", 12345 * MilliNewton},
		{"-12.345N", -12345 * MilliNewton},
		{"9.223372036854775807GN", 9223372036854775807 * NanoNewton},
		{"-9.223372036854775807GN", -9223372036854775807 * NanoNewton},
		{"1MN", 1 * MegaNewton},
		{"1nN", 1 * NanoNewton},
		{"1mlbf", 4448222 * NanoNewton},
		{"1lbf", 1 * PoundForce},
		{"1lbf", 4448221615 * NanoNewton},
		{"20lbf", 88964432305 * NanoNewton},
		{"1klbf", 4448221615261 * NanoNewton},
		{"1Mlbf", 4448221615261000 * NanoNewton},
		{"2Mlbf", 8896443230522000 * NanoNewton},
		{"2073496519lbf", 9223372034443058185 * NanoNewton},
		{"1.0000000000101lbf", 4448221615 * NanoNewton},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"2073496520lbf",
			"maximum value is 2.073496519Glbf",
		},
		{
			"-2073496520lbf",
			"minimum value is -2.073496519Glbf",
		},
		{
			"1234567.890123456789lbf",
			"converting to nano Newtons would overflow, consider using nN for maximum precision",
		},
		{
			"10TN",
			"maximum value is 9.223GN",
		},
		{
			"10EN",
			"unknown unit prefix; valid prefixes for \"N\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaN",
			"unknown unit prefix; valid prefixes for \"N\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eNewtonE",
			"unknown unit provided; need N or lbf",
		},
		{
			"10",
			"no unit provided; need N or lbf",
		},
		{
			"10n",
			"no unit provided; need N or lbf",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GN",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GN",
		},
		{
			"9.223372036854775808GN",
			"maximum value is 9.223GN",
		},
		{
			"-9.223372036854775808GN",
			"minimum value is -9.223GN",
		},
		{
			"9.223372036854775808GN",
			"maximum value is 9.223GN",
		},
		{
			"-9.223372036854775808GN",
			"minimum value is -9.223GN",
		},
		{
			"9.3GN",
			"maximum value is 9.223GN",
		},
		{
			"-9.3GN",
			"minimum value is -9.223GN",
		},
		{
			"1random",
			"unknown unit provided; need N or lbf",
		},
		{
			"N",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit N or lbf",
		},
		{
			"++1N",
			"contains multiple plus symbols",
		},
		{
			"--1N",
			"contains multiple minus symbols",
		},
		{
			"+-1N",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1N",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Force
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Force.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Force.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Force
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Force.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestForce_RoundTrip(t *testing.T) {
	x := 123 * Newton
	var y Force
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Force.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Force expected %s to equal %s", x, y)
	}
}
