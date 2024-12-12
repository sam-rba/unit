// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// MagneticFluxDensity is a measurement of magnetic flux density, stored in Tesla.
//
// The highest representable value is 9.2GT.
type MagneticFluxDensity int64

// String returns the energy formatted as a string in Farad.
func (c MagneticFluxDensity) String() string {
	return nanoAsString(int64(c)) + "T"
}

const (
	// Tesla is a unit of magnetic flux density.
	NanoTesla  MagneticFluxDensity = 1
	MicroTesla MagneticFluxDensity = 1000 * NanoTesla
	MilliTesla MagneticFluxDensity = 1000 * MicroTesla
	Tesla      MagneticFluxDensity = 1000 * MilliTesla
	KiloTesla  MagneticFluxDensity = 1000 * Tesla
	MegaTesla  MagneticFluxDensity = 1000 * KiloTesla
	GigaTesla  MagneticFluxDensity = 1000 * MegaTesla

	maxMagneticFluxDensity = 9223372036854775807 * NanoTesla
	minMagneticFluxDensity = -9223372036854775807 * NanoTesla
)
