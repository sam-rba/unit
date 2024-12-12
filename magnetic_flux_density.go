// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// MagneticFluxDensity is a measurement of magnetic flux density, stored in Tesla.
//
// The highest representable value is 9.2GT.
type MagneticFluxDensity int64

// String returns the energy formatted as a string in Farad.
func (c MagneticFluxDensity) String() string {
	return nanoAsString(int64(c)) + "T"
}

// Set sets the MagneticFluxDensity to the value represented by s. Units are
// to be provided in "T" with an optional SI prefix: "p", "n", "u", "µ", "m",
// "k", "M", "G" or "T".
func (c *MagneticFluxDensity) Set(s string) error {
	v, n, err := valueOfUnitString(s, pico)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "T", "t"); found != "" {
					return err
				}
				return notNumberUnitErr("T")
			case errOverflowsInt64:
				return maxValueErr(maxMagneticFluxDensity.String())
			case errOverflowsInt64Negative:
				return minValueErr(minMagneticFluxDensity.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "T", "t":
		*c = (MagneticFluxDensity)(v)
	case "":
		return noUnitErr("T")
	default:
		if found := hasSuffixes(s[n:], "T", "t"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("T")
	}

	return nil
}

const (
	// Tesla is a unit of magnetic flux density.
	NanoTesla  MagneticFluxDensity = 1
	MicroTesla MagneticFluxDensity = 1000 * NanoTesla
	MilliTesla MagneticFluxDensity = 1000 * MicroTesla
	Tesla      MagneticFluxDensity = 1000 * MilliTesla
	KiloTesla  MagneticFluxDensity = 1000 * Tesla
	MegaTesla  MagneticFluxDensity = 1000 * KiloTesla
	GigaTesla  MagneticFluxDensity = 1000 * MegaTesla

	maxMagneticFluxDensity = 9223372036854775807 * NanoTesla
	minMagneticFluxDensity = -9223372036854775807 * NanoTesla
)
