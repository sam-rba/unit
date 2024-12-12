// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// ElectricalCapacitance is a measurement of capacitance stored as a pico farad.
//
// The highest representable value is 9.2MF.
type ElectricalCapacitance int64

// String returns the energy formatted as a string in Farad.
func (c ElectricalCapacitance) String() string {
	return picoAsString(int64(c)) + "F"
}

// Set sets the ElectricalCapacitance to the value represented by s. Units are
// to be provided in "F" with an optional SI prefix: "p", "n", "u", "µ", "m",
// "k", "M", "G" or "T".
func (c *ElectricalCapacitance) Set(s string) error {
	v, n, err := valueOfUnitString(s, pico)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "F", "f"); found != "" {
					return err
				}
				return notNumberUnitErr("F")
			case errOverflowsInt64:
				return maxValueErr(maxElectricalCapacitance.String())
			case errOverflowsInt64Negative:
				return minValueErr(minElectricalCapacitance.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "F", "f":
		*c = (ElectricalCapacitance)(v)
	case "":
		return noUnitErr("F")
	default:
		if found := hasSuffixes(s[n:], "F", "f"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("F")
	}

	return nil
}

const (
	// Farad is a unit of capacitance. kg⁻¹⋅m⁻²⋅s⁴A²
	PicoFarad  ElectricalCapacitance = 1
	NanoFarad  ElectricalCapacitance = 1000 * PicoFarad
	MicroFarad ElectricalCapacitance = 1000 * NanoFarad
	MilliFarad ElectricalCapacitance = 1000 * MicroFarad
	Farad      ElectricalCapacitance = 1000 * MilliFarad
	KiloFarad  ElectricalCapacitance = 1000 * Farad
	MegaFarad  ElectricalCapacitance = 1000 * KiloFarad

	maxElectricalCapacitance = 9223372036854775807 * PicoFarad
	minElectricalCapacitance = -9223372036854775807 * PicoFarad
)
