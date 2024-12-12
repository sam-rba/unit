// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

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
