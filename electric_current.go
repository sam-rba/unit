// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// ElectricCurrent is a measurement of a flow of electric charge stored as an
// int64 nano Ampere.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2GA.
type ElectricCurrent int64

// String returns the current formatted as a string in Ampere.
func (c ElectricCurrent) String() string {
	return nanoAsString(int64(c)) + "A"
}

// Set sets the ElectricCurrent to the value represented by s. Units are to
// be provided in "A" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (c *ElectricCurrent) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "A", "a"); found != "" {
					return err
				}
				return notNumberUnitErr("A")
			case errOverflowsInt64:
				return maxValueErr(maxElectricCurrent.String())
			case errOverflowsInt64Negative:
				return minValueErr(minElectricCurrent.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "A", "a":
		*c = (ElectricCurrent)(v)
	case "":
		return noUnitErr("A")
	default:
		if found := hasSuffixes(s[n:], "A"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("A")
	}

	return nil
}

const (
	NanoAmpere  ElectricCurrent = 1
	MicroAmpere ElectricCurrent = 1000 * NanoAmpere
	MilliAmpere ElectricCurrent = 1000 * MicroAmpere
	Ampere      ElectricCurrent = 1000 * MilliAmpere
	KiloAmpere  ElectricCurrent = 1000 * Ampere
	MegaAmpere  ElectricCurrent = 1000 * KiloAmpere
	GigaAmpere  ElectricCurrent = 1000 * MegaAmpere

	maxElectricCurrent = 9223372036854775807 * NanoAmpere
	minElectricCurrent = -9223372036854775807 * NanoAmpere
)
