// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"fmt"
	"testing"
)

func TestSpeed_String(t *testing.T) {
	if s := MilePerHour.String(); s != "447.040mm/s" {
		t.Fatalf("%#v", s)
	}
}

func TestSpeed_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Speed
	}{
		{"1nmps", NanoMetrePerSecond},
		{"1umps", MicroMetrePerSecond},
		{"1µmps", MicroMetrePerSecond},
		{"1mmps", MilliMetrePerSecond},
		{"1mps", MetrePerSecond},
		{"1kmps", KiloMetrePerSecond},
		{"1Mmps", MegaMetrePerSecond},
		{"1Gmps", GigaMetrePerSecond},
		{"1nm/s", NanoMetrePerSecond},
		{"1um/s", MicroMetrePerSecond},
		{"1µm/s", MicroMetrePerSecond},
		{"1mm/s", MilliMetrePerSecond},
		{"1m/s", MetrePerSecond},
		{"1km/s", KiloMetrePerSecond},
		{"1Mm/s", MegaMetrePerSecond},
		{"1Gm/s", GigaMetrePerSecond},
		{"1mph", MilePerHour},
		{"1fps", FootPerSecond},
		{"1kph", KilometrePerHour},
		// Maximum and minimum values that are allowed.
		{fmt.Sprintf("%dnmps", minSpeed), minSpeed},
		{fmt.Sprintf("%dnmps", maxSpeed), maxSpeed},
		{fmt.Sprintf("%dkph", minKilometrePerHour), minKilometrePerHour * KilometrePerHour},
		{fmt.Sprintf("%dkph", maxKilometrePerHour), maxKilometrePerHour * KilometrePerHour},
		{fmt.Sprintf("%dmph", minMilePerHour), minMilePerHour * MilePerHour},
		{fmt.Sprintf("%dmph", maxMilePerHour), maxMilePerHour * MilePerHour},
		{fmt.Sprintf("%dfps", minFootPerSecond), minFootPerSecond * FootPerSecond},
		{fmt.Sprintf("%dfps", maxFootPerSecond), maxFootPerSecond * FootPerSecond},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10Gm/s",
			"maximum value is 9.223Gm/s",
		},
		{
			"10Em/s",
			"unknown unit prefix; valid prefixes for \"m/s\" are p,n,u,µ,m,k,M,G or T",
		},
		{
			"10",
			"no unit provided; need m/s, mps, kph, fps or mph",
		},
		{
			fmt.Sprintf("%dkph", maxKilometrePerHour+1),
			fmt.Sprintf("maximum value is %dkph", maxKilometrePerHour),
		},
		{
			fmt.Sprintf("%dkph", minKilometrePerHour-1),
			fmt.Sprintf("minimum value is %dkph", minKilometrePerHour),
		},
		{
			fmt.Sprintf("%dmph", maxMilePerHour+1),
			fmt.Sprintf("maximum value is %dmph", maxMilePerHour),
		},
		{
			fmt.Sprintf("%dmph", minMilePerHour-1),
			fmt.Sprintf("minimum value is %dmph", minMilePerHour),
		},
		{
			fmt.Sprintf("%dfps", maxFootPerSecond+1),
			fmt.Sprintf("maximum value is %dfps", maxFootPerSecond),
		},
		{
			fmt.Sprintf("%dfps", minFootPerSecond-1),
			fmt.Sprintf("minimum value is %dfps", minFootPerSecond),
		},
		{
			"9.224Gm/s",
			"maximum value is 9.223Gm/s",
		},
		{
			"-9.224Gm/s",
			"minimum value is -9.223Gm/s",
		},
		{
			"9223372036854775808nm/s",
			"maximum value is 9.223Gm/s",
		},
		{
			"-9223372036854775808nm/s",
			"minimum value is -9.223Gm/s",
		},
		{
			"1random",
			"unknown unit provided; need m/s, mps, kph, fps or mph",
		},
		{
			"m/s",
			"not a number",
		},
		{
			"fps",
			"not a number",
		},
		{
			"mph",
			"not a number",
		},
		{
			"kph",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit m/s, mps, kph, fps or mph",
		},
		{
			"++1m/s",
			"contains multiple plus symbols",
		},
		{
			"--1m/s",
			"contains multiple minus symbols",
		},
		{
			"+-1m/s",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1m/s",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Speed
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Speed.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Speed.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Speed
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Speed.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}

func TestSpeed_RoundTrip(t *testing.T) {
	x := 123 * MetrePerSecond
	var y Speed
	if err := y.Set(x.String()); err != nil {
		t.Fatalf("Speed.Set(stringer) failed: %v", err)
	}
	if x != y {
		t.Fatalf("Speed expected %s to equal %s", x, y)
	}
}
