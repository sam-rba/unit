// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// LuminousIntensity is a measurement of the quantity of visible light energy
// emitted per unit solid angle with wavelength power weighted by a luminosity
// function which represents the human eye's response to different wavelengths.
// The CIE 1931 luminosity function is the SI standard for candela.
//
// LuminousIntensity is stored as nano candela.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gcd.
type LuminousIntensity int64

// String returns the energy formatted as a string in Candela.
func (i LuminousIntensity) String() string {
	return nanoAsString(int64(i)) + "cd"
}

// Set sets the LuminousIntensity to the value represented by s. Units are to
// be provided in "cd" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (i *LuminousIntensity) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "cd"); found != "" {
					return err
				}
				return notNumberUnitErr("cd")
			case errOverflowsInt64:
				return maxValueErr(maxLuminousIntensity.String())
			case errOverflowsInt64Negative:
				return minValueErr(minLuminousIntensity.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "cd":
		*i = (LuminousIntensity)(v)
	case "":
		return noUnitErr("cd")
	default:
		if found := hasSuffixes(s[n:], "cd"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("cd")
	}

	return nil
}

const (
	// Candela is a unit of luminous intensity. cd
	NanoCandela  LuminousIntensity = 1
	MicroCandela LuminousIntensity = 1000 * NanoCandela
	MilliCandela LuminousIntensity = 1000 * MicroCandela
	Candela      LuminousIntensity = 1000 * MilliCandela
	KiloCandela  LuminousIntensity = 1000 * Candela
	MegaCandela  LuminousIntensity = 1000 * KiloCandela
	GigaCandela  LuminousIntensity = 1000 * MegaCandela

	maxLuminousIntensity = 9223372036854775807 * NanoCandela
	minLuminousIntensity = -9223372036854775807 * NanoCandela
)
