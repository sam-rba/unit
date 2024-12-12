// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "testing"

func TestPressure_String(t *testing.T) {
	if s := NanoPascal.String(); s != "1nPa" {
		t.Fatalf("%v", s)
	}
	if s := MicroPascal.String(); s != "1µPa" {
		t.Fatalf("%v", s)
	}
	if s := MilliPascal.String(); s != "1mPa" {
		t.Fatalf("%v", s)
	}
	if s := Pascal.String(); s != "1Pa" {
		t.Fatalf("%v", s)
	}
	if s := KiloPascal.String(); s != "1kPa" {
		t.Fatalf("%v", s)
	}
	if s := MegaPascal.String(); s != "1MPa" {
		t.Fatalf("%v", s)
	}
	if s := GigaPascal.String(); s != "1GPa" {
		t.Fatalf("%v", s)
	}
}

func TestPressure_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Pressure
	}{
		{"1nPa", 1 * NanoPascal},
		{"10nPa", 10 * NanoPascal},
		{"100nPa", 100 * NanoPascal},
		{"1uPa", 1 * MicroPascal},
		{"10uPa", 10 * MicroPascal},
		{"100uPa", 100 * MicroPascal},
		{"1µPa", 1 * MicroPascal},
		{"10µPa", 10 * MicroPascal},
		{"100µPa", 100 * MicroPascal},
		{"1mPa", 1 * MilliPascal},
		{"10mPa", 10 * MilliPascal},
		{"100mPa", 100 * MilliPascal},
		{"1Pa", 1 * Pascal},
		{"10Pa", 10 * Pascal},
		{"100Pa", 100 * Pascal},
		{"1kPa", 1 * KiloPascal},
		{"10kPa", 10 * KiloPascal},
		{"100kPa", 100 * KiloPascal},
		{"1MPa", 1 * MegaPascal},
		{"10MPa", 10 * MegaPascal},
		{"100MPa", 100 * MegaPascal},
		{"1GPa", 1 * GigaPascal},
		{"12.345Pa", 12345 * MilliPascal},
		{"-12.345Pa", -12345 * MilliPascal},
		{"9.223372036854775807GPa", 9223372036854775807 * NanoPascal},
		{"-9.223372036854775807GPa", -9223372036854775807 * NanoPascal},
		{"1MPa", 1 * MegaPascal},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10TPa",
			"maximum value is 9.223GPa",
		},
		{
			"10EPa",
			"unknown unit prefix; valid prefixes for \"Pa\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaPa",
			"unknown unit prefix; valid prefixes for \"Pa\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ePascalE",
			"unknown unit provided; need Pa",
		},
		{
			"10",
			"no unit provided; need Pa",
		},
		{
			"9223372036854775808",
			"maximum value is 9.223GPa",
		},
		{
			"-9223372036854775808",
			"minimum value is -9.223GPa",
		},
		{
			"9.223372036854775808GPa",
			"maximum value is 9.223GPa",
		},
		{
			"-9.223372036854775808GPa",
			"minimum value is -9.223GPa",
		},
		{
			"9.223372036854775808GPa",
			"maximum value is 9.223GPa",
		},
		{
			"-9.223372036854775808GPa",
			"minimum value is -9.223GPa",
		},
		{
			"1random",
			"unknown unit provided; need Pa",
		},
		{
			"Pa",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit Pa",
		},
		{
			"++1Pa",
			"contains multiple plus symbols",
		},
		{
			"--1Pa",
			"contains multiple minus symbols",
		},
		{
			"+-1Pa",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1Pa",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got Pressure
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Pressure.Set(%s) got unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Pressure.Set(%s) expected: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Pressure
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Pressure.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestPressure_RoundTrip(t *testing.T) {
	x := 123 * Pascal
	var y Pressure
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Pressure.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Pressure expected %s to equal %s", x, y)
	}
}

func TestPressure_Pa(t *testing.T) {
	if v := Pressure(123 * Pascal).Pa(); v != 123. {
		t.Fatal(v)
	}
}

func TestPressure_KPa(t *testing.T) {
	if v := Pressure(123 * KiloPascal).KPa(); v != 123. {
		t.Fatal(v)
	}
}

func TestPressure_MBar(t *testing.T) {
	if v := Pressure(123 * MilliBar).MBar(); v != 123. {
		t.Fatal(v)
	}
}

func TestPressure_Bar(t *testing.T) {
	if v := Pressure(123 * Bar).Bar(); v != 123. {
		t.Fatal(v)
	}
}
