// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"fmt"
	"testing"
)

func TestAngle_String(t *testing.T) {
	data := []struct {
		in       Angle
		expected string
	}{
		{0, "0°"},
		{Degree/10000 + Degree/2000, "0.001°"},
		{-Degree/10000 - Degree/2000, "-0.001°"},
		{Degree / 1000, "0.001°"},
		{-Degree / 1000, "-0.001°"},
		{Degree / 2, "0.500°"},
		{-Degree / 2, "-0.500°"},
		{Degree, "1.000°"},
		{-Degree, "-1.000°"},
		{10 * Degree, "10.00°"},
		{-10 * Degree, "-10.00°"},
		{100 * Degree, "100.0°"},
		{-100 * Degree, "-100.0°"},
		{1000 * Degree, "1000°"},
		{-1000 * Degree, "-1000°"},
		{100000000000 * Degree, "100000000000°"},
		{-100000000000 * Degree, "-100000000000°"},
		{maxAngle, "528460276055°"},
		{minAngle, "-528460276055°"},
		{Pi, "180.0°"},
		{Theta, "360.0°"},
		{Radian, "57.296°"},
	}
	for i, line := range data {
		if s := line.in.String(); s != line.expected {
			t.Fatalf("%d: Degree(%d).String() = %s != %s", i, int64(line.in), s, line.expected)
		}
	}
}

func TestAngle_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Angle
	}{
		{"1nrad", NanoRadian},
		{"10nrad", 10 * NanoRadian},
		{"100nrad", 100 * NanoRadian},
		{"1urad", 1 * MicroRadian},
		{"10urad", 10 * MicroRadian},
		{"100urad", 100 * MicroRadian},
		{"1µrad", 1 * MicroRadian},
		{"10µrad", 10 * MicroRadian},
		{"10µRad", 10 * MicroRadian},
		{"100µrad", 100 * MicroRadian},
		{"1mrad", 1 * MilliRadian},
		{"10mrad", 10 * MilliRadian},
		{"100mrad", 100 * MilliRadian},
		{"1rad", 1 * Radian},
		{"1Rad", 1 * Radian},
		{"10rad", 10 * Radian},
		{"100rad", 100 * Radian},
		{"1krad", 1000 * Radian},
		{"10krad", 10000 * Radian},
		{"100krad", 100000 * Radian},
		{"1Mrad", 1000000 * Radian},
		{"10Mrad", 10000000 * Radian},
		{"100Mrad", 100000000 * Radian},
		{"1Grad", 1000000000 * Radian},
		{"12.345rad", 12345 * MilliRadian},
		{"-12.345rad", -12345 * MilliRadian},
		{fmt.Sprintf("%dnrad", maxAngle), maxAngle},
		{"1deg", 1 * Degree},
		{"1Deg", 1 * Degree},
		{"1Mdeg", 1000000 * Degree},
		{"1MDeg", 1000000 * Degree},
		{"100Gdeg", 100000000000 * Degree},
		{"500Gdeg", 500000000000 * Degree},
		{maxAngle.String(), 528460276055 * Degree},
		{minAngle.String(), -528460276055 * Degree},
		{"1mdeg", Degree / 1000},
		{"1udeg", Degree / 1000000},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Erad",
			"unknown unit prefix; valid prefixes for \"rad\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10Exarad",
			"unknown unit prefix; valid prefixes for \"rad\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eRadianE",
			"unknown unit provided; need Rad, Deg or °",
		},
		{
			"10",
			"no unit provided; need Rad, Deg or °",
		},
		{
			fmt.Sprintf("%dnrad", uint64(maxAngle)+1),
			"maximum value is 528460276055°",
		},
		{
			fmt.Sprintf("-%dnrad", uint64(maxAngle)+1),
			"minimum value is -528460276055°",
		},
		{
			"528460276056deg",
			"maximum value is 528460276055°",
		},
		{
			"-528460276056deg",
			"minimum value is -528460276055°",
		},
		{
			"-9.223372036854775808Grad",
			// TODO(maruel): Investigate.
			"minimum value is -528460276055°",
		},
		{
			"9.223372036854775808Grad",
			"maximum value is 528460276055°",
		},
		{
			"9.224GRad",
			"maximum value is 9.223GRad",
		},
		{
			"-9.224GRad",
			"minimum value is -9.223GRad",
		},
		{
			"1random",
			"unknown unit provided; need Rad, Deg or °",
		},
		{
			"rad",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit Rad, Deg or °",
		},
		{
			"++1rad",
			"contains multiple plus symbols",
		},
		{
			"--1rad",
			"contains multiple minus symbols",
		},
		{
			"+-1rad",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1rad",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Angle
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Angle.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Angle.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Angle
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Angle.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestAngle_RoundTrip(t *testing.T) {
	x := 123 * Degree
	var y Angle
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Angle.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Angle expected %s to equal %s", x, y)
	}
}
