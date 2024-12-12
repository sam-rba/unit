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

// Temperature is a measurement of hotness stored as a nano kelvin.
//
// Negative values are invalid.
//
// The highest representable value is 9.2GK.
type Temperature int64

// String returns the temperature formatted as a string in °Celsius.
func (t Temperature) String() string {
	if t < -ZeroCelsius || t > maxCelsius {
		return nanoAsString(int64(t)) + "K"
	}
	return nanoAsString(int64(t-ZeroCelsius)) + "°C"
}

// Set sets the Temperature to the value represented by s. Units are to be
// provided in "C", "°C", "F", "°F" or "K" with an optional SI prefix: "p", "n",
// "u", "µ", "m", "k", "M", "G" or "T".
func (t *Temperature) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "°C", "C", "°F", "F", "K"); found != "" {
					return err
				}
				return notNumberUnitErr("K, °C, C, °F or F")
			case errOverflowsInt64:
				// TODO(maruel): Look for suffix, and reuse it.
				return maxValueErr(maxTemperature.String())
			case errOverflowsInt64Negative:
				// TODO(maruel): Look for suffix, and reuse it.
				return minValueErr(minTemperature.String())
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
	case "F", "°F":
		// F to nK  nK = 555555555.556*F + 255372222222
		fPerK := decimal{
			base: 555555555556,
			exp:  -3,
			neg:  false,
		}
		f, _ := decimalMul(d, fPerK)
		v, overflow := dtoi(f, int(si))
		if overflow {
			if f.neg {
				return minValueErr("-459.67F")
			}
			return maxValueErr(strconv.FormatInt(int64(maxFahrenheit), 10) + "F")
		}
		// We need an extra check here to make sure that will not overflow with
		// the addition of ZeroFahrenheit.
		switch {
		case v > int64(maxTemperature-ZeroFahrenheit):
			return maxValueErr(strconv.FormatInt(int64(maxFahrenheit), 10) + "F")
		case v < int64(-ZeroFahrenheit):
			return minValueErr("-459.67F")
		}
		v += int64(ZeroFahrenheit)
		*t = (Temperature)(v)
	case "K":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr("0K")
			}
			return maxValueErr(strconv.FormatInt(int64(maxTemperature/1000000000), 10) + "K")
		}
		if v < 0 {
			return minValueErr("0K")
		}
		*t = (Temperature)(v)
	case "C", "°C":
		v, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr("-273.15°C")
			}
			return maxValueErr(strconv.FormatInt(int64(maxCelsius/1000000000), 10) + "°C")
		}
		// We need an extra check here to make sure that will not overflow with
		// the addition of ZeroCelsius.
		switch {
		case v > int64(maxCelsius):
			return maxValueErr(strconv.FormatInt(int64(maxCelsius/1000000000), 10) + "°C")
		case v < int64(-ZeroCelsius):
			return minValueErr("-273.15°C")
		}
		v += int64(ZeroCelsius)
		*t = (Temperature)(v)
	case "":
		return noUnitErr("K, °C, C, °F or F")
	default:
		if found := hasSuffixes(s[n:], "°C", "C", "°F", "F", "K"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("K, °C, C, °F or F")
	}
	return nil
}

// K returns the temperature as a floating number of °Kelvin.
func (t Temperature) K() float64 {
	return float64(t) / float64(Kelvin)
}

// C returns the temperature as a floating number of °Celsius.
func (t Temperature) C() float64 {
	return float64(t-ZeroCelsius) / float64(Celsius)
}

// F returns the temperature as a floating number of °Fahrenheit.
func (t Temperature) F() float64 {
	return float64(t-ZeroFahrenheit) / float64(Fahrenheit)
}

const (
	NanoKelvin  Temperature = 1
	MicroKelvin Temperature = 1000 * NanoKelvin
	MilliKelvin Temperature = 1000 * MicroKelvin
	Kelvin      Temperature = 1000 * MilliKelvin
	KiloKelvin  Temperature = 1000 * Kelvin
	MegaKelvin  Temperature = 1000 * KiloKelvin
	GigaKelvin  Temperature = 1000 * MegaKelvin

	// Conversion between Kelvin and Celsius.
	ZeroCelsius  Temperature = 273150 * MilliKelvin
	MilliCelsius Temperature = MilliKelvin
	Celsius      Temperature = Kelvin

	// Conversion between Kelvin and Fahrenheit.
	ZeroFahrenheit  Temperature = 255372222222 * NanoKelvin
	MilliFahrenheit Temperature = 555555 * NanoKelvin
	Fahrenheit      Temperature = 555555555 * NanoKelvin

	maxTemperature Temperature = (1 << 63) - 1
	minTemperature Temperature = 0

	// Maximum Celsius is 9223371763704775807°nC.
	maxCelsius Temperature = maxTemperature - ZeroCelsius

	// Maximum Fahrenheit is 16602069204F
	maxFahrenheit Temperature = 16602069204
)
