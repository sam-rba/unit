// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// Power is a measurement of power stored as a nano watts.
//
// The highest representable value is 9.2GW.
type Power int64

// String returns the power formatted as a string in watts.
func (p Power) String() string {
	return nanoAsString(int64(p)) + "W"
}

// Set sets the Power to the value represented by s. Units are to
// be provided in "W" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (p *Power) Set(s string) error {
	v, n, err := valueOfUnitString(s, nano)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s, "W", "w"); found != "" {
					return err
				}
				return notNumberUnitErr("W")
			case errOverflowsInt64:
				return maxValueErr(maxPower.String())
			case errOverflowsInt64Negative:
				return minValueErr(minPower.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "W", "w":
		*p = (Power)(v)
	case "":
		return noUnitErr("W")
	default:
		if found := hasSuffixes(s[n:], "W", "w"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("W")
	}

	return nil
}

const (
	// Watt is unit of power J/s, kg⋅m²⋅s⁻³
	NanoWatt  Power = 1
	MicroWatt Power = 1000 * NanoWatt
	MilliWatt Power = 1000 * MicroWatt
	Watt      Power = 1000 * MilliWatt
	KiloWatt  Power = 1000 * Watt
	MegaWatt  Power = 1000 * KiloWatt
	GigaWatt  Power = 1000 * MegaWatt

	maxPower = 9223372036854775807 * NanoWatt
	minPower = -9223372036854775807 * NanoWatt
)
