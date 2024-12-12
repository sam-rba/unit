// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestDistance_String(t *testing.T) {
	if s := Mile.String(); s != "1.609km" {
		t.Fatalf("%#v", s)
	}
}

func TestDistance_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Distance
	}{
		{"1nm", 1 * NanoMetre},
		{"10nm", 10 * NanoMetre},
		{"100nm", 100 * NanoMetre},
		{"1um", 1 * MicroMetre},
		{"10um", 10 * MicroMetre},
		{"100um", 100 * MicroMetre},
		{"1µm", 1 * MicroMetre},
		{"10µm", 10 * MicroMetre},
		{"100µm", 100 * MicroMetre},
		{"1mm", 1 * MilliMetre},
		{"1mm", 1 * MilliMetre},
		{"10mm", 10 * MilliMetre},
		{"100mm", 100 * MilliMetre},
		{"1m", 1 * Metre},
		{"10m", 10 * Metre},
		{"100m", 100 * Metre},
		{"1km", 1 * KiloMetre},
		{"10km", 10 * KiloMetre},
		{"100km", 100 * KiloMetre},
		{"1Mm", 1 * MegaMetre},
		{"1Mm", 1 * MegaMetre},
		{"10Mm", 10 * MegaMetre},
		{"100Mm", 100 * MegaMetre},
		{"1Gm", 1 * GigaMetre},
		{"12.345m", 12345 * MilliMetre},
		{"-12.345m", -12345 * MilliMetre},
		{"9.223372036854775807Gm", 9223372036854775807 * NanoMetre},
		{"-9.223372036854775807Gm", -9223372036854775807 * NanoMetre},
		{"1Mm", 1 * MegaMetre},
		{"5Mile", 8046720000000 * NanoMetre},
		{"5mile", 8046720000000 * NanoMetre},
		{"3ft", 914400000 * NanoMetre},
		{"10Yard", 9144000000 * NanoMetre},
		{"5731.137678988Mile", 9223372036853264 * NanoMetre},
		{"-5731.137678988Mile", -9223372036853264 * NanoMetre},
		{"1.008680231502051MYard", 922337203685475 * NanoMetre},
		{"1Yard", 914400 * MicroMetre},
		{"1yard", 914400 * MicroMetre},
		{"-1008680.231502051Yard", -922337203685475 * NanoMetre},
		{"3026040.694506158ft", 922337203685477 * NanoMetre},
		{"-3.026040694506158Mft", -922337203685477 * NanoMetre},
		{"36.312488334073900Min", 922337203685477 * NanoMetre},
		{"-36312488.334073900in", -922337203685477 * NanoMetre},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Tm",
			"maximum value is 9.223Gm",
		},
		{
			"10Em",
			"unknown unit prefix; valid prefixes for \"m\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10Exam",
			"unknown unit prefix; valid prefixes for \"m\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eMetreE",
			"unknown unit provided; need m, Mile, in, ft or Yard",
		},
		{
			"10",
			"no unit provided; need m, Mile, in, ft or Yard",
		},
		{
			"9.3Gm",
			"maximum value is 9.223Gm",
		},
		{
			"-9.3Gm",
			"minimum value is -9.223Gm",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223Gm",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223Gm",
		},
		{
			"9.223372036854775808Gm",
			"maximum value is 9.223Gm",
		},
		{
			"-9.223372036854775808Gm",
			"minimum value is -9.223Gm",
		},
		{
			"9.223372036854775808Gm",
			"maximum value is 9.223Gm",
		},
		{
			"-9.223372036854775808Gm",
			"minimum value is -9.223Gm",
		},
		{
			"5731.137678989Mile",
			"maximum value is 5731Mile",
		},
		{
			"-5731.1376789889Mile",
			"minimum value is -5731Mile",
		},
		{
			"1.008680231502053MYard",
			"maximum value is 1 Million Yard",
		},
		{
			"-1008680.231502053Yard",
			"minimum value is -1 Million Yard",
		},
		{
			"3026040.694506159ft",
			"maximum value is 3 Million ft",
		},
		{
			"-3.026040694506159Mft",
			"minimum value is -3 Million ft",
		},
		{
			"36.312488334073901Min",
			"maximum value is 36 Million inch",
		},
		{
			"-36312488.334073901in",
			"minimum value is -36 Million inch",
		},
		{
			"1random",
			"unknown unit prefix; valid prefixes for \"m\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"m",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit m, Mile, in, ft or Yard",
		},
		{
			"cd",
			"does not contain number or unit m, Mile, in, ft or Yard",
		},
		{
			"++1m",
			"contains multiple plus symbols",
		},
		{
			"--1m",
			"contains multiple minus symbols",
		},
		{
			"+-1m",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1m",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x31, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Distance
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Distance.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Distance.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Distance
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Distance.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestDistance_RoundTrip(t *testing.T) {
	x := 123 * Metre
	var y Distance
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Distance.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Distance expected %s to equal %s", x, y)
	}
}
