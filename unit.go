// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

func prefixZeros(digits, v int) string {
	// digits is expected to be around 2~3.
	s := strconv.Itoa(v)
	for len(s) < digits {
		// O(n²) but since digits is expected to run 2~3 times at most, it doesn't
		// matter.
		s = "0" + s
	}
	return s
}

// nanoAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func nanoAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "G"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "m"
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "µ"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "n"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// microAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func microAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "T"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "G"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "m"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "µ"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// picoAsString converts a value in S.I. unit in a string with the predefined
// prefix.
func picoAsString(v int64) string {
	sign := ""
	if v < 0 {
		if v == -9223372036854775808 {
			v++
		}
		sign = "-"
		v = -v
	}
	var frac int
	var base int
	var precision int64
	unit := ""
	switch {
	case v >= 999999500000000001:
		precision = v % 1000000000000000
		base = int(v / 1000000000000000)
		if precision > 500000000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "M"
	case v >= 999999500000001:
		precision = v % 1000000000000
		base = int(v / 1000000000000)
		if precision > 500000000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "k"
	case v >= 999999500001:
		precision = v % 1000000000
		base = int(v / 1000000000)
		if precision > 500000000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = ""
	case v >= 999999501:
		precision = v % 1000000
		base = int(v / 1000000)
		if precision > 500000 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "m"
	case v >= 1000000:
		precision = v % 1000
		base = int(v / 1000)
		if precision > 500 {
			base++
		}
		frac = (base % 1000)
		base = base / 1000
		unit = "µ"
	case v >= 1000:
		frac = int(v) % 1000
		base = int(v) / 1000
		unit = "n"
	default:
		if v == 0 {
			return "0"
		}
		base = int(v)
		unit = "p"
	}
	if frac == 0 {
		return sign + strconv.Itoa(base) + unit
	}
	return sign + strconv.Itoa(base) + "." + prefixZeros(3, frac) + unit
}

// Decimal is the representation of decimal number.
type decimal struct {
	// base hold the significant digits.
	base uint64
	// exponent is the left or right decimal shift. (powers of ten).
	exp int
	// neg it true if the number is negative.
	neg bool
}

// Positive powers of 10 in the form such that powerOF10[index] = 10^index.
var powerOf10 = [...]uint64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
}

// Maximum value for a int64.
const maxInt64 = (1 << 63) - 1

var maxInt64Str = "9223372036854775807"

var (
	errOverflowsInt64         = errors.New("exceeds maximum")
	errOverflowsInt64Negative = errors.New("exceeds minimum")
	errNotANumber             = errors.New("not a number")
)

// Converts from decimal to int64.
//
// Scale is combined with the decimal exponent to maximise the resolution and is
// in powers of ten.
//
// Returns true if the value overflowed.
func dtoi(d decimal, scale int) (int64, bool) {
	// Get the total magnitude of the number.
	// a^x * b^y = a*b^(x+y) since scale is of the order unity this becomes
	// 1^x * b^y = b^(x+y).
	// mag must be positive to use as index in to powerOf10 array.
	u := d.base
	mag := d.exp + scale
	if mag < 0 {
		mag = -mag
	}
	var n int64
	if mag > 18 {
		return 0, true
	}
	// Divide is = 10^(-mag)
	switch {
	case d.exp+scale < 0:
		u = (u + powerOf10[mag]/2) / powerOf10[mag]
	case mag == 0:
		if u > maxInt64 {
			return 0, true
		}
	default:
		check := u * powerOf10[mag]
		if check/powerOf10[mag] != u || check > maxInt64 {
			return 0, true
		}
		u *= powerOf10[mag]
	}

	n = int64(u)
	if d.neg {
		n = -n
	}
	return n, false
}

// Converts a string to a decimal form. The return int is how many bytes of the
// string are considered numeric. The string may contain +-0 prefixes and
// arbitrary suffixes as trailing non number characters are ignored.
// Significant digits are stored without leading or trailing zeros, rather a
// base and exponent is used. Significant digits are stored as uint64, max size
// of significant digits is int64
func atod(s string) (decimal, int, error) {
	var d decimal
	start := 0
	dp := 0
	end := len(s)
	seenDigit := false
	seenZero := false
	isPoint := false
	seenPlus := false

	// Strip leading zeros, +/- and mark DP.
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '-':
			if seenDigit {
				end = i
				break
			}
			if seenPlus {
				return decimal{}, 0, &parseError{
					errors.New("contains both plus and minus symbols"),
				}
			}
			if d.neg {
				return decimal{}, 0, &parseError{
					errors.New("contains multiple minus symbols"),
				}
			}
			d.neg = true
			start++
		case s[i] == '+':
			if seenDigit {
				end = i
				break
			}
			if d.neg {
				return decimal{}, 0, &parseError{
					errors.New("contains both plus and minus symbols"),
				}
			}
			if seenPlus {
				return decimal{}, 0, &parseError{
					errors.New("contains multiple plus symbols"),
				}
			}
			seenPlus = true
			start++
		case s[i] == '.':
			if isPoint {
				return decimal{}, 0, &parseError{
					errors.New("contains multiple decimal points"),
				}
			}
			isPoint = true
			dp = i
			if !seenDigit {
				start++
			}
		case s[i] == '0':
			if !seenDigit {
				start++
			}
			seenZero = true
		case s[i] >= '1' && s[i] <= '9':
			seenDigit = true
		default:
			if !seenDigit && !seenZero {
				return decimal{}, 0, &parseError{errNotANumber}
			}
			end = i
		}
	}

	last := end
	seenDigit = false
	exp := 0
	// Strip non significant zeros to find base exponent.
	for i := end - 1; i > start-1; i-- {
		switch {
		case s[i] >= '1' && s[i] <= '9':
			seenDigit = true
		case s[i] == '.':
			if !seenDigit {
				end--
			}
		case s[i] == '0':
			if !seenDigit {
				if i > dp {
					end--
				}
				if i <= dp || dp == 0 {
					exp++
				}
			}
		default:
			last--
			end--
		}
	}

	for i := start; i < end; i++ {
		c := s[i]
		// Check that is is a digit.
		if c >= '0' && c <= '9' {
			// *10 is decimal shift left.
			d.base *= 10
			// Convert ascii digit into number.
			check := d.base + uint64(c-'0')
			// Check should always be larger than u unless we have overflowed.
			// Similarly if check > max it will overflow when converted to int64.
			if check < d.base || check > maxInt64 {
				if d.neg {
					return decimal{}, 0, &parseError{errOverflowsInt64Negative}
				}
				return decimal{}, 0, &parseError{errOverflowsInt64}
			}
			d.base = check
		} else if c != '.' {
			return decimal{}, 0, &parseError{errNotANumber}
		}
	}
	if !isPoint {
		d.exp = exp
	} else {
		if dp > start && dp < end {
			// Decimal Point is in the middle of a number.
			end--
		}
		// Find the exponent based on decimal point distance from left and the
		// length of the number.
		d.exp = (dp - start) - (end - start)
		if dp <= start {
			// Account for numbers of the form 1 > n < -1 eg 0.0001.
			d.exp++
		}
	}
	return d, last, nil
}

// valueOfUnitString is a helper for converting a string and a prefix in to a
// physic unit. It can be used when characters of the units do not conflict with
// any of the SI prefixes.
func valueOfUnitString(s string, base prefix) (int64, int, error) {
	d, n, err := atod(s)
	if err != nil {
		return 0, n, err
	}
	si := prefix(unit)
	if n != len(s) {
		r, rsize := utf8.DecodeRuneInString(s[n:])
		if r <= 1 || rsize == 0 {
			return 0, 0, &parseError{
				errors.New("unexpected end of string"),
			}
		}
		var siSize int
		si, siSize = parseSIPrefix(r)
		n += siSize
	}
	v, overflow := dtoi(d, int(si-base))
	if overflow {
		if d.neg {
			return -maxInt64, 0, &parseError{errOverflowsInt64Negative}
		}
		return maxInt64, 0, &parseError{errOverflowsInt64}
	}
	return v, n, nil
}

// decimalMul calcululates the product of two decimals; a and b, keeping the
// base less than maxInt64. Returns the number of times a figure was trimmed
// from either base coefficients. This function is to aid in the multiplication
// of numbers whose product have more than 18 significant figures. The minimum
// accuracy of the end product that has been truncated is 9 significant figures.
func decimalMul(a, b decimal) (decimal, uint) {
	switch {
	case a.base == 0 || b.base == 0:
		// Anything multiplied by zero is zero. Special case to set exponent to
		// zero.
		return decimal{}, 0
	case a.base > (1<<64)-6 || b.base > (1<<64)-6:
		// In normal usage base will never be greater than 1<<63. However since
		// base could be large as (1<<64 -1) this is to prevent an infinite loop
		// when ((1<<64)-6)+5 overflows in the truncate least significant digit
		// loop during rounding without adding addition bounds checking at that
		// point.
		break
	default:
		exp := a.exp + b.exp
		neg := a.neg != b.neg
		ab := a.base
		bb := b.base
		for i := uint(0); i < 21; i++ {
			if ab <= 1 || bb <= 1 {
				// This will always fit inside uint64.
				return decimal{ab * bb, exp, neg}, i
			}
			if base := ab * bb; (base/ab == bb) && base < maxInt64 {
				// Return if product did not overflow or exceed int64.
				return decimal{base, exp, neg}, i
			}
			// Truncate least significant digit in product.
			if bb > ab {
				bb = (bb + 5) / 10
				// Compact trailing zeros if any.
				for bb > 0 && bb%10 == 0 {
					bb /= 10
					exp++
				}
			} else {
				ab = (ab + 5) / 10
				// Compact trailing zeros if any.
				for ab > 0 && ab%10 == 0 {
					ab /= 10
					exp++
				}
			}
			exp++
		}
	}
	return decimal{}, 21
}

// hasSuffixes returns the first suffix found and the prefix content.
func hasSuffixes(s string, suffixes ...string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return suffix
		}
	}
	return ""
}

type parseError struct {
	error
}

func noUnitErr(valid string) error {
	return errors.New("no unit provided; need " + valid)
}

func incorrectUnitErr(valid string) error {
	return errors.New("unknown unit provided; need " + valid)
}

func unknownUnitPrefixErr(unit, valid string) error {
	return errors.New("unknown unit prefix; valid prefixes for \"" + unit + "\" are " + valid)
}

func maxValueErr(valid string) error {
	return errors.New("maximum value is " + valid)
}

func minValueErr(valid string) error {
	return errors.New("minimum value is " + valid)
}

func notNumberUnitErr(unit string) error {
	return errors.New("does not contain number or unit " + unit)
}

type prefix int

const (
	pico  prefix = -12
	nano  prefix = -9
	micro prefix = -6
	milli prefix = -3
	unit  prefix = 0
	deca  prefix = 1
	hecto prefix = 2
	kilo  prefix = 3
	mega  prefix = 6
	giga  prefix = 9
	tera  prefix = 12
)

func parseSIPrefix(r rune) (prefix, int) {
	switch r {
	case 'p':
		return pico, len("p")
	case 'n':
		return nano, len("n")
	case 'u':
		return micro, len("u")
	case 'µ':
		return micro, len("µ")
	case 'm':
		return milli, len("m")
	case 'k':
		return kilo, len("k")
	case 'M':
		return mega, len("M")
	case 'G':
		return giga, len("G")
	case 'T':
		return tera, len("T")
	default:
		return unit, 0
	}
}
