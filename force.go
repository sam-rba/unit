// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"errors"
	"unicode/utf8"
)

// Force is a measurement of interaction that will change the motion of an
// object stored as an int64 nano Newton.
//
// A measurement of Force is a vector and has a direction but this unit only
// represents the magnitude. The orientation needs to be stored as a Quaternion
// independently.
//
// The highest representable value is 9.2TN.
type Force int64

// String returns the force formatted as a string in Newton.
func (f Force) String() string {
	return nanoAsString(int64(f)) + "N"
}

// Set sets the Force to the value represented by s. Units are to
// be provided in "N", or "lbf" (Pound force) with an optional SI prefix: "p",
// "n", "u", "µ", "m", "k", "M", "G" or "T".
func (f *Force) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "N", "lbf"); found != "" {
					return err
				}
				return notNumberUnitErr("N or lbf")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxForce.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minForce.String())
			}
		}
		return err
	}

	var si prefix
	if n != len(s) {
		r, rsize := utf8.DecodeRuneInString(s[n:])
		if r <= 1 || rsize == 0 {
			return errors.New("unexpected end of string")
		}
		var siSize int
		si, siSize = parseSIPrefix(r)

		n += siSize
	}

	switch s[n:] {
	case "lbf":
		poundForce := decimal{
			base: 4448221615261,
			exp:  -3,
			neg:  false,
		}
		lbf, loss := decimalMul(d, poundForce)
		if loss > 9 {
			return errors.New("converting to nano Newtons would overflow, consider using nN for maximum precision")
		}
		v, overflow := dtoi(lbf, int(si))
		if overflow {
			if lbf.neg {
				return minValueErr("-2.073496519Glbf")
			}
			return maxValueErr("2.073496519Glbf")
		}
		*f = (Force)(v)
	case "N":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr(minForce.String())
			}
			return maxValueErr(maxForce.String())
		}
		*f = (Force)(v)
	case "":
		return noUnitErr("N or lbf")
	default:
		if found := hasSuffixes(s[n:], "N", "lbf"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("N or lbf")
	}
	return nil
}

const (
	// Newton is kg⋅m/s².
	NanoNewton  Force = 1
	MicroNewton Force = 1000 * NanoNewton
	MilliNewton Force = 1000 * MicroNewton
	Newton      Force = 1000 * MilliNewton
	KiloNewton  Force = 1000 * Newton
	MegaNewton  Force = 1000 * KiloNewton
	GigaNewton  Force = 1000 * MegaNewton

	EarthGravity Force = 9806650 * MicroNewton

	// Conversion between Newton and imperial units.
	// Pound is both a unit of mass and weight (force). The suffix Force is added
	// to disambiguate the measurement it represents.
	PoundForce Force = 4448221615 * NanoNewton

	maxForce Force = (1 << 63) - 1
	minForce Force = -((1 << 63) - 1)
)
