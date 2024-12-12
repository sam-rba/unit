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

// Speed is a measurement of magnitude of velocity stored as an int64 nano
// Metre per Second.
//
// The highest representable value is 9.2Gm/s.
type Speed int64

// String returns the speed formatted as a string in m/s.
func (sp Speed) String() string {
	return nanoAsString(int64(sp)) + "m/s"
}

// Set sets the Speed to the value represented by s. Units are to be provided in
// "mps"(meters per second), "m/s", "kph", "fps", or "mph" with an optional SI
// prefix: "p", "n", "u", "µ", "m", "k", "M", "G" or "T".
func (sp *Speed) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "m/s", "mps", "kph", "fps", "mph"); found != "" {
					return err
				}
				return notNumberUnitErr("m/s, mps, kph, fps or mph")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxSpeed.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minSpeed.String())
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
		if si == milli {
			switch s[n:] {
			case "m/s", "mps", "mph":
				si = unit
				siSize = 0
			}
		}
		if si == kilo {
			switch s[n:] {
			case "kph":
				si = unit
				siSize = 0
			}
		}
		n += siSize
	}
	switch s[n:] {
	case "m/s", "mps":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr(minSpeed.String())
			}
			return maxValueErr(maxSpeed.String())
		}
		*sp = (Speed)(v)
	case "kph":
		mpsPerkph := decimal{
			base: uint64(KilometrePerHour),
			exp:  0,
			neg:  false,
		}
		kph, _ := decimalMul(d, mpsPerkph)
		v, overflow := dtoi(kph, int(si))
		if overflow {
			if kph.neg {
				return minValueErr(strconv.FormatInt(int64(minKilometrePerHour), 10) + "kph")
			}
			return maxValueErr(strconv.FormatInt(int64(maxKilometrePerHour), 10) + "kph")
		}
		*sp = (Speed)(v)
	case "fps":
		mpsPerfps := decimal{
			base: uint64(FootPerSecond / 1000),
			exp:  3,
			neg:  false,
		}
		oz, _ := decimalMul(d, mpsPerfps)
		v, overflow := dtoi(oz, int(si))
		if overflow {
			if oz.neg {
				return minValueErr(strconv.FormatInt(int64(minFootPerSecond), 10) + "fps")
			}
			return maxValueErr(strconv.FormatInt(int64(maxFootPerSecond), 10) + "fps")
		}
		*sp = (Speed)(v)
	case "mph":
		mpsPermph := decimal{
			base: uint64(MilePerHour / 1000),
			exp:  3,
			neg:  false,
		}
		oz, _ := decimalMul(d, mpsPermph)
		v, overflow := dtoi(oz, int(si))
		if overflow {
			if oz.neg {
				return minValueErr(strconv.FormatInt(int64(minMilePerHour), 10) + "mph")
			}
			return maxValueErr(strconv.FormatInt(int64(maxMilePerHour), 10) + "mph")
		}
		*sp = (Speed)(v)
	case "":
		return noUnitErr("m/s, mps, kph, fps or mph")
	default:
		if found := hasSuffixes(s[n:], "m/s", "mps", "kph", "fps", "mph"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("m/s, mps, kph, fps or mph")
	}
	return nil
}

const (
	// MetrePerSecond is m/s.
	NanoMetrePerSecond  Speed = 1
	MicroMetrePerSecond Speed = 1000 * NanoMetrePerSecond
	MilliMetrePerSecond Speed = 1000 * MicroMetrePerSecond
	MetrePerSecond      Speed = 1000 * MilliMetrePerSecond
	KiloMetrePerSecond  Speed = 1000 * MetrePerSecond
	MegaMetrePerSecond  Speed = 1000 * KiloMetrePerSecond
	GigaMetrePerSecond  Speed = 1000 * MegaMetrePerSecond

	LightSpeed Speed = 299792458 * MetrePerSecond

	KilometrePerHour Speed = 277777778 * NanoMetrePerSecond
	MilePerHour      Speed = 447040 * MicroMetrePerSecond
	FootPerSecond    Speed = 304800 * MicroMetrePerSecond

	maxSpeed Speed = (1 << 63) - 1
	minSpeed Speed = -((1 << 63) - 1)

	// Min Max KilometrePerHour are in kph.
	minKilometrePerHour Speed = -33204139306
	maxKilometrePerHour Speed = 33204139306
	// Min Max MilePerHour are in mph.
	minMilePerHour Speed = -20632095644
	maxMilePerHour Speed = 20632095644
	// Min Max FootPerSecond are in fps.
	minFootPerSecond Speed = -30260406945
	maxFootPerSecond Speed = 30260406945
)
