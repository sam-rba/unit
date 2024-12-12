// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"errors"
	"strconv"
	"unicode/utf8"
)

// Angle is the measurement of the difference in orientation between two vectors
// stored as an int64 nano radian.
//
// A negative angle is valid.
//
// The highest representable value is a bit over 9.223GRad or 500,000,000,000°.
type Angle int64

// String returns the angle formatted as a string in degree.
func (a Angle) String() string {
	// Angle is not a S.I. unit, so it must not be prefixed by S.I. prefixes.
	if a == 0 {
		return "0°"
	}
	// Round.
	prefix := ""
	if a < 0 {
		a = -a
		prefix = "-"
	}
	switch {
	case a < Degree:
		v := ((a * 1000) + Degree/2) / Degree
		return prefix + "0." + prefixZeros(3, int(v)) + "°"
	case a < 10*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(3, int(v)) + "°"
	case a < 100*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(2, int(v)) + "°"
	case a < 1000*Degree:
		v := ((a * 1000) + Degree/2) / Degree
		i := v / 1000
		v = v - i*1000
		return prefix + strconv.FormatInt(int64(i), 10) + "." + prefixZeros(1, int(v)) + "°"
	case a > maxAngle-Degree:
		u := (uint64(a) + uint64(Degree)/2) / uint64(Degree)
		v := int64(u)
		return prefix + strconv.FormatInt(int64(v), 10) + "°"
	default:
		v := (a + Degree/2) / Degree
		return prefix + strconv.FormatInt(int64(v), 10) + "°"
	}
}

// Set sets the Angle to the value represented by s. Units are to be provided in
// "rad", "deg" or "°" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (a *Angle) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "Rad", "rad", "Deg", "deg", "°"); found != "" {
					return err
				}
				return notNumberUnitErr("Rad, Deg or °")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxAngle.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minAngle.String())
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
	case "Deg", "deg", "°":
		degreePerRadian := decimal{
			base: 17453293,
			exp:  0,
			neg:  false,
		}
		deg, _ := decimalMul(d, degreePerRadian)
		// Impossible for precision loss to exceed 9 since the number of
		// significant figures in degrees per radian is only 8.
		v, overflow := dtoi(deg, int(si))
		if overflow {
			if deg.neg {
				return minValueErr(minAngle.String())
			}
			return maxValueErr(maxAngle.String())
		}
		*a = (Angle)(v)
	case "Rad", "rad":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr("-9.223G" + s[n:])
			}
			return maxValueErr("9.223G" + s[n:])
		}
		*a = (Angle)(v)
	case "":
		return noUnitErr("Rad, Deg or °")
	default:
		if found := hasSuffixes(s[n:], "Rad", "rad", "Deg", "deg", "°"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("Rad, Deg or °")
	}
	return nil
}

const (
	NanoRadian  Angle = 1
	MicroRadian Angle = 1000 * NanoRadian
	MilliRadian Angle = 1000 * MicroRadian
	Radian      Angle = 1000 * MilliRadian

	// Theta is 2π. This is equivalent to 360°.
	Theta  Angle = 6283185307 * NanoRadian
	Pi     Angle = 3141592653 * NanoRadian
	Degree Angle = 17453293 * NanoRadian

	maxAngle Angle = 9223372036854775807
	minAngle Angle = -9223372036854775807
)
