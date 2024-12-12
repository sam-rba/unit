// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// ElectricResistance is a measurement of the difficulty to pass an electric
// current through a conductor stored as an int64 nano Ohm.
//
// The highest representable value is 9.2GΩ.
type ElectricResistance int64

// String returns the resistance formatted as a string in Ohm.
func (r ElectricResistance) String() string {
	return nanoAsString(int64(r)) + "Ω"
}

// Set sets the ElectricResistance to the value represented by s. Units are to
// be provided in "Ohm", or "Ω" with an optional SI prefix: "p", "n", "u", "µ",
// "m", "k", "M", "G" or "T".
func (r *ElectricResistance) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "Ohm", "ohm", "Ω"); found != "" {
					return err
				}
				return notNumberUnitErr("Ohm or Ω")
			case errOverflowsInt64:
				return maxValueErr(maxElectricResistance.String())
			case errOverflowsInt64Negative:
				return minValueErr(minElectricResistance.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "Ohm", "ohm", "Ω":
		*r = (ElectricResistance)(v)
	case "":
		return noUnitErr("Ohm or Ω")
	default:
		if found := hasSuffixes(s[n:], "Ohm", "ohm", "Ω"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("Ohm or Ω")
	}
	return nil
}

const (
	// Ohm is V/A, kg⋅m²/s³/A².
	NanoOhm  ElectricResistance = 1
	MicroOhm ElectricResistance = 1000 * NanoOhm
	MilliOhm ElectricResistance = 1000 * MicroOhm
	Ohm      ElectricResistance = 1000 * MilliOhm
	KiloOhm  ElectricResistance = 1000 * Ohm
	MegaOhm  ElectricResistance = 1000 * KiloOhm
	GigaOhm  ElectricResistance = 1000 * MegaOhm

	maxElectricResistance = 9223372036854775807 * NanoOhm
	minElectricResistance = -9223372036854775807 * NanoOhm
)
