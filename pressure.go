// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// Pressure is a measurement of force applied to a surface per unit
// area (stress) stored as an int64 nano Pascal.
//
// The highest representable value is 9.2GPa.
type Pressure int64

// String returns the pressure formatted as a string in Pascal.
func (p Pressure) String() string {
	return nanoAsString(int64(p)) + "Pa"
}

// Set sets the Pressure to the value represented by s. Units are to
// be provided in "Pa" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (p *Pressure) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "Pa"); found != "" {
					return err
				}
				return notNumberUnitErr("Pa")
			case errOverflowsInt64:
				return maxValueErr(maxPressure.String())
			case errOverflowsInt64Negative:
				return minValueErr(minPressure.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "Pa":
		*p = (Pressure)(v)
	case "":
		return noUnitErr("Pa")
	default:
		if found := hasSuffixes(s[n:], "Pa"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("Pa")
	}

	return nil
}

// Pa returns the pressure as a floating number of Pascals.
func (p Pressure) Pa() float64 {
	return float64(p) / float64(Pascal)
}

// KPa returns the pressure as a floating number of KiloPascals.
func (p Pressure) KPa() float64 {
	return float64(p) / float64(KiloPascal)
}

// MBar returns the pressure as a floating number of MilliBar.
func (p Pressure) MBar() float64 {
	return float64(p) / float64(MilliBar)
}

// Bar returns the pressure as a floating number of Bar.
func (p Pressure) Bar() float64 {
	return float64(p) / float64(Bar)
}

const (
	// Pascal is N/m², kg/m/s².
	NanoPascal  Pressure = 1
	MicroPascal Pressure = 1000 * NanoPascal
	MilliPascal Pressure = 1000 * MicroPascal
	Pascal      Pressure = 1000 * MilliPascal
	KiloPascal  Pressure = 1000 * Pascal
	MegaPascal  Pressure = 1000 * KiloPascal
	GigaPascal  Pressure = 1000 * MegaPascal

	MilliBar Pressure = 100 * Pascal
	Bar      Pressure = 1000 * MilliBar

	maxPressure = 9223372036854775807 * NanoPascal
	minPressure = -9223372036854775807 * NanoPascal
)
