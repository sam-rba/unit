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
