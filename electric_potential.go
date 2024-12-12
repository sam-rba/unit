// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// ElectricPotential is a measurement of electric potential stored as an int64
// nano Volt.
//
// The highest representable value is 9.2GV.
type ElectricPotential int64

// String returns the tension formatted as a string in Volt.
func (p ElectricPotential) String() string {
	return nanoAsString(int64(p)) + "V"
}

// Set sets the ElectricPotential to the value represented by s. Units are to
// be provided in "V" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (p *ElectricPotential) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "V", "v"); found != "" {
					return err
				}
				return notNumberUnitErr("V")
			case errOverflowsInt64:
				return maxValueErr(maxElectricPotential.String())
			case errOverflowsInt64Negative:
				return minValueErr(minElectricPotential.String())
			}
		}
		return err
	}
	switch s[n:] {
	case "V", "v":
		*p = (ElectricPotential)(v)
	case "":
		return noUnitErr("V")
	default:
		if found := hasSuffixes(s[n:], "V", "v"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("V")
	}
	return nil
}

const (
	// Volt is W/A, kg⋅m²/s³/A.
	NanoVolt  ElectricPotential = 1
	MicroVolt ElectricPotential = 1000 * NanoVolt
	MilliVolt ElectricPotential = 1000 * MicroVolt
	Volt      ElectricPotential = 1000 * MilliVolt
	KiloVolt  ElectricPotential = 1000 * Volt
	MegaVolt  ElectricPotential = 1000 * KiloVolt
	GigaVolt  ElectricPotential = 1000 * MegaVolt

	maxElectricPotential = 9223372036854775807 * NanoVolt
	minElectricPotential = -9223372036854775807 * NanoVolt
)
