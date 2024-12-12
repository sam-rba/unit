// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"testing"
	"time"
)

func TestFrequency_String(t *testing.T) {
	data := []struct {
		in       Frequency
		expected string
	}{
		{minFrequency, "-9.223THz"},
		{-Hertz, "-1Hz"},
		{0, "0Hz"},
		{Hertz, "1Hz"},
		{1666500 * MicroHertz, "1.666Hz"},
		{1666501 * MicroHertz, "1.667Hz"},
		{MegaHertz, "1MHz"},
		{GigaHertz, "1GHz"},
		{999999500 * KiloHertz, "999.999GHz"},
		{999999501 * KiloHertz, "1THz"},
		{1000500 * MegaHertz, "1THz"},
		{1000501 * MegaHertz, "1.001THz"},
		{1001 * GigaHertz, "1.001THz"},
		{1000 * GigaHertz, "1THz"},
		{maxFrequency, "9.223THz"},
	}
	for i, line := range data {
		if v := line.in.String(); v != line.expected {
			t.Fatalf("%d: Frequency(%d).String() = %s != %s", i, line.in, v, line.expected)
		}
	}
}

func TestFrequency_Period(t *testing.T) {
	data := []struct {
		in       Frequency
		expected time.Duration
	}{
		{0, 0},
		{MicroHertz, 277*time.Hour + 46*time.Minute + 40*time.Second},
		{MilliHertz, 16*time.Minute + 40*time.Second},
		{999999 * MicroHertz, 1000001 * time.Microsecond},
		{Hertz, time.Second},
		{1000001 * MicroHertz, 999999 * time.Microsecond},
		{MegaHertz, time.Microsecond},
		{23 * MegaHertz, 43 * time.Nanosecond},
		{100 * MegaHertz, 10 * time.Nanosecond},
		{150 * MegaHertz, 7 * time.Nanosecond},
		{GigaHertz, time.Nanosecond},
		{2 * GigaHertz, time.Nanosecond},
		{20000000 * KiloHertz, 0},
		{TeraHertz, 0},
		{maxFrequency, 0},
	}
	for i, line := range data {
		if v := line.in.Period(); v != line.expected {
			t.Fatalf("%d: Frequency(%d).Period() = %s != %s", i, line.in, v, line.expected)
		}
		if v := (-line.in).Period(); v != -line.expected {
			t.Fatalf("%d: Frequency(%d).Period() = %s != %s", i, -line.in, v, -line.expected)
		}
	}
}

func TestFrequency_Duration(t *testing.T) {
	// TODO(maruel): To be removed in v4.0.0.
	if MicroHertz.Duration() != MicroHertz.Period() {
		t.Fatal("should have the same result")
	}
}

func TestFrequency_PeriodToFrequency(t *testing.T) {
	data := []struct {
		in       time.Duration
		expected Frequency
	}{
		{0, 0},
		{time.Nanosecond, GigaHertz},
		{time.Microsecond, MegaHertz},
		{time.Millisecond, KiloHertz},
		{999990000 * time.Nanosecond, 1000010 * MicroHertz},
		{999999500 * time.Nanosecond, 1000001 * MicroHertz},
		{999999501 * time.Nanosecond, 1000000 * MicroHertz},
		{time.Second, Hertz},
		{1000000000 * time.Nanosecond, Hertz},
		{1000000500 * time.Nanosecond, Hertz},
		{1000000501 * time.Nanosecond, 999999 * MicroHertz},
		{time.Minute, 16667 * MicroHertz},
		{time.Hour, 278 * MicroHertz},
	}
	for i, line := range data {
		if v := PeriodToFrequency(line.in); v != line.expected {
			t.Fatalf("%d: PeriodToFrequency(%s) = %d != %d", i, line.in, v, line.expected)
		}
		if v := PeriodToFrequency(-line.in); v != -line.expected {
			t.Fatalf("%d: PeriodToFrequency(%s) = %d != %d", i, -line.in, v, -line.expected)
		}
	}
}

func TestFrequency_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Frequency
	}{
		{"1uHz", 1 * MicroHertz},
		{"10uHz", 10 * MicroHertz},
		{"100uHz", 100 * MicroHertz},
		{"1µHz", 1 * MicroHertz},
		{"10µHz", 10 * MicroHertz},
		{"100µHz", 100 * MicroHertz},
		{"1mHz", 1 * MilliHertz},
		{"10mHz", 10 * MilliHertz},
		{"100mHz", 100 * MilliHertz},
		{"1hz", 1 * Hertz},
		{"1Hz", 1 * Hertz},
		{"10", 10 * Hertz},
		{"10Hz", 10 * Hertz},
		{"100Hz", 100 * Hertz},
		{"1kHz", 1 * KiloHertz},
		{"1khz", 1 * KiloHertz},
		{"1k", 1 * KiloHertz},
		{"10kHz", 10 * KiloHertz},
		{"100kHz", 100 * KiloHertz},
		{"1MHz", 1 * MegaHertz},
		{"10MHz", 10 * MegaHertz},
		{"100MHz", 100 * MegaHertz},
		{"1GHz", 1 * GigaHertz},
		{"10GHz", 10 * GigaHertz},
		{"100GHz", 100 * GigaHertz},
		{"1THz", 1 * TeraHertz},
		{"12.345Hz", 12345 * MilliHertz},
		{"-12.345Hz", -12345 * MilliHertz},
		{"9.223372036854775807THz", 9223372036854775807 * MicroHertz},
		{"-9.223372036854775807THz", -9223372036854775807 * MicroHertz},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10THz",
			"maximum value is 9.223THz",
		},
		{
			"10EHz",
			"unknown unit prefix; valid prefixes for \"Hz\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10ExaHz",
			"unknown unit prefix; valid prefixes for \"Hz\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10eHzE",
			"unknown unit provided; need Hz",
		},
		{
			"922337203685477580",
			"maximum value is 9.223THz",
		},
		{
			"-922337203685477580",
			"minimum value is -9.223THz",
		},
		{
			"9.223372036854775808THz",
			"maximum value is 9.223THz",
		},
		{
			"-9.223372036854775808THz",
			"minimum value is -9.223THz",
		},
		{
			"9.223372036854775808THertz",
			"maximum value is 9.223THz",
		},
		{
			"-9.223372036854775808THertz",
			"minimum value is -9.223THz",
		},
		{
			"1random",
			"unknown unit provided; need Hz",
		},
		{
			"Hz",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit Hz",
		},
		{
			"++1Hz",
			"contains multiple plus symbols",
		},
		{
			"--1Hz",
			"contains multiple minus symbols",
		},
		{
			"+-1Hz",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1Hz",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got Frequency
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Frequency.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Frequency.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Frequency
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Frequency.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestFrequency_RoundTrip(t *testing.T) {
	x := 123 * Hertz
	var y Frequency
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Frequency.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Frequency expected %s to equal %s", x, y)
	}
}
