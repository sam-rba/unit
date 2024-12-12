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

// Mass is a measurement of mass stored as an int64 nano gram.
//
// This is one of the base unit in the International System of Units.
//
// The highest representable value is 9.2Gg.
type Mass int64

// String returns the mass formatted as a string in gram.
func (m Mass) String() string {
	return nanoAsString(int64(m)) + "g"
}

// Set sets the Mass to the value represented by s. Units are to be provided in
// "g", "lb" or "oz" with an optional SI prefix: "p", "n", "u", "µ", "m", "k",
// "M", "G" or "T".
func (m *Mass) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "g", "lb", "oz"); found != "" {
					return err
				}
				return notNumberUnitErr("g, lb or oz")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxMass.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minMass.String())
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
	case "g":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr(minMass.String())
			}
			return maxValueErr(maxMass.String())
		}
		*m = (Mass)(v)
	case "lb":
		gramsPerlb := decimal{
			base: uint64(PoundMass),
			exp:  0,
			neg:  false,
		}
		lbs, _ := decimalMul(d, gramsPerlb)
		v, overflow := dtoi(lbs, int(si))
		if overflow {
			if lbs.neg {
				return minValueErr(strconv.FormatInt(int64(minPoundMass), 10) + "lb")
			}
			return maxValueErr(strconv.FormatInt(int64(maxPoundMass), 10) + "lb")
		}
		*m = (Mass)(v)
	case "oz":
		gramsPerOz := decimal{
			base: uint64(OunceMass),
			exp:  0,
			neg:  false,
		}
		oz, _ := decimalMul(d, gramsPerOz)
		v, overflow := dtoi(oz, int(si))
		if overflow {
			if oz.neg {
				return minValueErr(strconv.FormatInt(int64(minOunceMass), 10) + "oz")
			}
			return maxValueErr(strconv.FormatInt(int64(maxOunceMass), 10) + "oz")
		}
		*m = (Mass)(v)
	case "":
		return noUnitErr("g, lb or oz")
	default:
		if found := hasSuffixes(s[n:], "g", "lb", "oz"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("g, lb or oz")
	}
	return nil
}

const (
	NanoGram  Mass = 1
	MicroGram Mass = 1000 * NanoGram
	MilliGram Mass = 1000 * MicroGram
	Gram      Mass = 1000 * MilliGram
	KiloGram  Mass = 1000 * Gram
	MegaGram  Mass = 1000 * KiloGram
	GigaGram  Mass = 1000 * MegaGram
	Tonne     Mass = MegaGram

	// Conversion between Gram and imperial units.
	// Ounce is both a unit of mass, weight (force) or volume depending on
	// context. The suffix Mass is added to disambiguate the measurement it
	// represents.
	OunceMass Mass = 28349523125 * NanoGram
	// Pound is both a unit of mass and weight (force). The suffix Mass is added
	// to disambiguate the measurement it represents.
	PoundMass Mass = 16 * OunceMass

	Slug Mass = 14593903 * MilliGram

	maxMass Mass = (1 << 63) - 1
	minMass Mass = -((1 << 63) - 1)

	// min and max Pound mass are in lb.
	minPoundMass Mass = -20334054
	maxPoundMass Mass = 20334054
	// min and max Ounce mass are in oz.
	minOunceMass Mass = -325344874
	maxOunceMass Mass = 325344874
)
