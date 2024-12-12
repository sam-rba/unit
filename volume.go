package unit

import (
	"errors"
	"unicode/utf8"
)

// Volume is a measurement of volume stored as an int64 nano litre.
//
// The highest representable value is 9.2GL.
type Volume int64

// String returns the volume formatted as a string in litres.
func (v Volume) String() string {
	return nanoAsString(int64(v)) + "L"
}

// Set sets the Volume to the value represented by s. The unit to be provided is "L"
// with an optional SI prefix: "n", "u", "µ", "m", "k", "M", or "G".
func (v *Volume) Set(s string) error {
	d, n, err := atod(s)
	if err != nil {
		if e, ok := err.(*parseError); ok {
			switch e.error {
			case errNotANumber:
				if found := hasSuffixes(s[n:], "L"); found != "" {
					return err
				}
				return notNumberUnitErr("L")
			case errOverflowsInt64:
				return maxValueErr(maxVolume.String())
			case errOverflowsInt64Negative:
				return minValueErr(minVolume.String())
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
	case "L":
		x, overflow := dtoi(d, int(si-nano))
		if overflow {
			if d.neg {
				return minValueErr(minVolume.String())
			}
			return maxValueErr(maxVolume.String())
		}
		*v = Volume(x)
	case "":
		return noUnitErr("L")
	default:
		if found := hasSuffixes(s[n:], "L"); found != "" {
			return unknownUnitPrefixErr(found, "n,u,µ,m,k,M, or G")
		}
		return incorrectUnitErr("L")
	}
	return nil
}

const (
	NanoLitre  Volume = 1
	MicroLitre Volume = 1000 * NanoLitre
	MilliLitre Volume = 1000 * MicroLitre
	Litre      Volume = 1000 * MilliLitre
	KiloLitre  Volume = 1000 * Litre
	MegaLitre  Volume = 1000 * KiloLitre
	GigaLitre  Volume = 1000 * MegaLitre

	CubicCentimetre Volume = MilliLitre

	maxVolume Volume = (1 << 63) - 1
	minVolume Volume = -((1 << 63) - 1)
)
