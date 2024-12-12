// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "strconv"

// RelativeHumidity is a humidity level measurement stored as an int32 fixed
// point integer at a precision of 0.00001%rH.
//
// Valid values are between 0% and 100%.
type RelativeHumidity int32

// String returns the humidity formatted as a string.
func (r RelativeHumidity) String() string {
	r /= MilliRH
	frac := int(r % 10)
	if frac == 0 {
		return strconv.Itoa(int(r)/10) + "%rH"
	}
	if frac < 0 {
		frac = -frac
	}
	return strconv.Itoa(int(r)/10) + "." + strconv.Itoa(frac) + "%rH"
}

// Set sets the RelativeHumidity to the value represented by s. Units are to
// be provided in "%rH" or "%" with an optional SI prefix: "p", "n", "u", "µ",
// "m", "k", "M", "G" or "T".
func (r *RelativeHumidity) Set(s string) error {
	// PercentRH is micro + deca.
	v, n, err := valueOfUnitString(s, micro+deca)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "%rH", "%"); found != "" {
					return err
				}
				return notNumberUnitErr("%rH or %")
			case errOverflowsInt64:
				return maxValueErr(maxRelativeHumidity.String())
			case errOverflowsInt64Negative:
				return minValueErr(minRelativeHumidity.String())
			}
		}
		return err
	}

	switch s[n:] {
	case "%rH", "%":
		// We need an extra check here to make sure that v will fit inside a
		// int32.
		switch {
		case v > int64(maxRelativeHumidity):
			return maxValueErr(maxRelativeHumidity.String())
		case v < int64(minRelativeHumidity):
			return minValueErr(minRelativeHumidity.String())
		}
		*r = (RelativeHumidity)(v)
	case "":
		return noUnitErr("%rH or %")
	default:
		if found := hasSuffixes(s[n:], "%rH", "%"); found != "" {
			return unknownUnitPrefixErr(found, "p,n,u,µ,m,k,M,G or T")
		}
		return incorrectUnitErr("%rH or %")
	}

	return nil
}

const (
	TenthMicroRH RelativeHumidity = 1                 // 0.00001%rH
	MicroRH      RelativeHumidity = 10 * TenthMicroRH // 0.0001%rH
	MilliRH      RelativeHumidity = 1000 * MicroRH    // 0.1%rH
	PercentRH    RelativeHumidity = 10 * MilliRH      // 1%rH

	maxRelativeHumidity RelativeHumidity = 100 * PercentRH
	minRelativeHumidity RelativeHumidity = 0
)
