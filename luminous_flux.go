// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// LuminousFlux is a measurement of total quantity of visible light energy
// emitted with wavelength power weighted by a luminosity function which
// represents a model of the human eye's response to different wavelengths.
// The CIE 1931 luminosity function is the standard for lumens.
//
// LuminousFlux is stored as nano lumens.
//
// The highest representable value is 9.2Glm.
type LuminousFlux int64

// String returns the energy formatted as a string in Lumens.
func (f LuminousFlux) String() string {
	return nanoAsString(int64(f)) + "lm"
}

// Set sets the LuminousFlux to the value represented by s. Units are to
// be provided in "lm" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (f *LuminousFlux) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "lm"); found != "" {
					return err
				}
				return notNumberUnitErr("lm")
			case errOverflowsInt64:
				return maxValueErr(maxLuminousFlux.String())
			case errOverflowsInt64Negative:
				return minValueErr(minLuminousFlux.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "lm":
		*f = (LuminousFlux)(v)
	case "":
		return noUnitErr("lm")
	default:
		if found := hasSuffixes(s[n:], "lm"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("lm")
	}

	return nil
}

const (
	// Lumen is a unit of luminous flux. cd⋅sr
	NanoLumen  LuminousFlux = 1
	MicroLumen LuminousFlux = 1000 * NanoLumen
	MilliLumen LuminousFlux = 1000 * MicroLumen
	Lumen      LuminousFlux = 1000 * MilliLumen
	KiloLumen  LuminousFlux = 1000 * Lumen
	MegaLumen  LuminousFlux = 1000 * KiloLumen
	GigaLumen  LuminousFlux = 1000 * MegaLumen

	maxLuminousFlux = 9223372036854775807 * NanoLumen
	minLuminousFlux = -9223372036854775807 * NanoLumen
)
