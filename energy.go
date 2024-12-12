// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// Energy is a measurement of work stored as a nano joules.
//
// The highest representable value is 9.2GJ.
type Energy int64

// String returns the energy formatted as a string in Joules.
func (e Energy) String() string {
	return nanoAsString(int64(e)) + "J"
}

const (
	// Joule is a unit of work. kg⋅m²⋅s⁻²
	NanoJoule  Energy = 1
	MicroJoule Energy = 1000 * NanoJoule
	MilliJoule Energy = 1000 * MicroJoule
	Joule      Energy = 1000 * MilliJoule
	KiloJoule  Energy = 1000 * Joule
	MegaJoule  Energy = 1000 * KiloJoule
	GigaJoule  Energy = 1000 * MegaJoule

	// BTU (British thermal unit) is the heat required to raise the temperature
	// of one pound of water by one degree Fahrenheit. This is the ISO value.
	BTU Energy = 1055060 * MilliJoule

	WattSecond   Energy = Joule
	WattHour     Energy = 3600 * Joule
	KiloWattHour Energy = 3600 * KiloJoule

	maxEnergy = 9223372036854775807 * NanoJoule
	minEnergy = -9223372036854775807 * NanoJoule
)
