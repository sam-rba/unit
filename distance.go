// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"errors"
	"unicode/utf8"
)

// Distance is a measurement of length stored as an int64 nano metre.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gm.
type Distance int64

// String returns the distance formatted as a string in metre.
func (d Distance) String() string {
	return nanoAsString(int64(d)) + "m"
}

// Set sets the Distance to the value represented by s. Units are to
// be provided in "m", "Mile", "Yard", "in", or "ft" with an optional SI
// prefix: "p", "n", "u", "µ", "m", "k", "M", "G" or "T".
func (d *Distance) Set(s string) error {
	dc, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "in", "ft", "Yard", "yard", "Mile", "mile", "m"); found != "" {
					return err
				}
				return notNumberUnitErr("m, Mile, in, ft or Yard")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxDistance.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minDistance.String())
			}
		}
		return err
	}
	si := prefix(unit)
	if n != len(s) {
		r, rsize := utf8.DecodeRuneInString(s[n:])
		if r <= 1 || rsize == 0 {
			return errors.New("unexpected end of string")
		}
		var siSize int
		si, siSize = parseSIPrefix(r)
		if si == milli || si == mega {
			switch s[n:] {
			case "m", "Mile", "mile":
				si = unit
			}
		}
		if si != unit {
			n += siSize
		}
	}
	v, overflow := dtoi(dc, int(si-nano))
	if overflow {
		if dc.neg {
			return minValueErr(minDistance.String())
		}
		return maxValueErr(maxDistance.String())
	}
	switch s[n:] {
	case "m":
		*d = (Distance)(v)
	case "Mile", "mile":
		switch {
		case v > maxMiles:
			return maxValueErr("5731Mile")
		case v < minMiles:
			return minValueErr("-5731Mile")
		case v >= 0:
			*d = (Distance)((v*1609344 + 500) / 1000)
		default:
			*d = (Distance)((v*1609344 - 500) / 1000)
		}
	case "Yard", "yard":
		switch {
		case v > maxYards:
			return maxValueErr("1 Million Yard")
		case v < minYards:
			return minValueErr("-1 Million Yard")
		case v >= 0:
			*d = (Distance)((v*9144 + 5000) / 10000)
		default:
			*d = (Distance)((v*9144 - 5000) / 10000)
		}
	case "ft":
		switch {
		case v > maxFeet:
			return maxValueErr("3 Million ft")
		case v < minFeet:
			return minValueErr("-3 Million ft")
		case v >= 0:
			*d = (Distance)((v*3048 + 5000) / 10000)
		default:
			*d = (Distance)((v*3048 - 5000) / 10000)
		}
	case "in":
		switch {
		case v > maxInches:
			return maxValueErr("36 Million inch")
		case v < minInches:
			return minValueErr("-36 Million inch")
		case v >= 0:
			*d = (Distance)((v*254 + 5000) / 10000)
		default:
			*d = (Distance)((v*254 - 5000) / 10000)
		}
	case "":
		return noUnitErr("m, Mile, in, ft or Yard")
	default:
		if found := hasSuffixes(s[n:], "in", "ft", "Yard", "Mile", "m"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("m, Mile, in, ft or Yard")
	}
	return nil
}

const (
	NanoMetre  Distance = 1
	MicroMetre Distance = 1000 * NanoMetre
	MilliMetre Distance = 1000 * MicroMetre
	Metre      Distance = 1000 * MilliMetre
	KiloMetre  Distance = 1000 * Metre
	MegaMetre  Distance = 1000 * KiloMetre
	GigaMetre  Distance = 1000 * MegaMetre

	// Conversion between Metre and imperial units.
	Thou Distance = 25400 * NanoMetre
	Inch Distance = 1000 * Thou
	Foot Distance = 12 * Inch
	Yard Distance = 3 * Foot
	Mile Distance = 1760 * Yard

	maxDistance       = 9223372036854775807 * NanoMetre
	minDistance       = -9223372036854775807 * NanoMetre
	maxMiles    int64 = (int64(maxDistance) - 500) / int64((Mile)/1000000) // ~Max/1609344
	minMiles    int64 = (int64(minDistance) + 500) / int64((Mile)/1000000) // ~Min/1609344
	maxYards    int64 = (int64(maxDistance) - 5000) / int64((Yard)/100000) // ~Max/9144
	minYards    int64 = (int64(minDistance) + 5000) / int64((Yard)/100000) // ~Min/9144
	maxFeet     int64 = (int64(maxDistance) - 5000) / int64((Foot)/100000) // ~Max/3048
	minFeet     int64 = (int64(minDistance) + 5000) / int64((Foot)/100000) // ~Min/3048
	maxInches   int64 = (int64(maxDistance) - 5000) / int64((Inch)/100000) // ~Max/254
	minInches   int64 = (int64(minDistance) + 5000) / int64((Inch)/100000) // ~Min/254
)
