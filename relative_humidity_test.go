// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"fmt"
	"testing"
)

func TestRelativeHumidity_String(t *testing.T) {
	data := []struct {
		in       RelativeHumidity
		expected string
	}{
		{TenthMicroRH, "0%rH"},
		{MicroRH, "0%rH"},
		{10 * MicroRH, "0%rH"},
		{100 * MicroRH, "0%rH"},
		{1000 * MicroRH, "0.1%rH"},
		{506000 * MicroRH, "50.6%rH"},
		{90 * PercentRH, "90%rH"},
		{100 * PercentRH, "100%rH"},
		// That's a lot of humidity. This is to test the value doesn't overflow
		// int32 too quickly.
		{1000 * PercentRH, "1000%rH"},
		// That's really dry.
		{-501000 * MicroRH, "-50.1%rH"},
	}
	for i, line := range data {
		if s := line.in.String(); s != line.expected {
			t.Fatalf("%d: RelativeHumidity(%d).String() = %s != %s", i, int64(line.in), s, line.expected)
		}
	}
}

func TestRelativeHumidity_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected RelativeHumidity
	}{
		{"10u%rH", PercentRH / 100000},
		{"1m%rH", PercentRH / 1000},
		{"1%rH", PercentRH},
		{"10%rH", 10 * PercentRH},
		{"100%rH", 100 * PercentRH},
		{"10u%", PercentRH / 100000},
		{"1m%", PercentRH / 1000},
		{"1%", PercentRH},
		{"10%", 10 * PercentRH},
		{"100%", 100 * PercentRH},
		{fmt.Sprintf("%du%%rH", int64(maxRelativeHumidity)*10), maxRelativeHumidity},
		{fmt.Sprintf("%du%%rH", int64(minRelativeHumidity)*10), minRelativeHumidity},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10E%rH",
			"unknown unit prefix; valid prefixes for \"%rH\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10",
			"no unit provided; need %rH or %",
		},
		{
			"21474836.48m%rH",
			"maximum value is 100%rH",
		},
		{
			"-21474836.48m%rH",
			"minimum value is 0%rH",
		},
		{
			"90224T%rH",
			"maximum value is 100%rH",
		},
		{
			"-90224T%rH",
			"minimum value is 0%rH",
		},
		{
			"1random",
			"unknown unit provided; need %rH or %",
		},
		{
			"%rH",
			"not a number",
		},
		{
			"%",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit %rH or %",
		},
		{
			"++1%rH",
			"contains multiple plus symbols",
		},
		{
			"--1%rH",
			"contains multiple minus symbols",
		},
		{
			"+-1%rH",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1%rH",
			"contains multiple decimal points",
		},
	}

	for i, tt := range succeeds {
		var got RelativeHumidity
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: RelativeHumidity.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: RelativeHumidity.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got RelativeHumidity
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: RelativeHumidity.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestRelativeHumidity_RoundTrip(t *testing.T) {
	x := 23 * PercentRH
	var y RelativeHumidity
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("RelativeHumidity.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("RelativeHumidity expected %s to equal %s", x, y)
	}
}
