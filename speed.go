// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

// Speed is a measurement of magnitude of velocity stored as an int64 nano
// Metre per Second.
//
// The highest representable value is 9.2Gm/s.
type Speed int64

// String returns the speed formatted as a string in m/s.
func (sp Speed) String() string {
	return nanoAsString(int64(sp)) + "m/s"
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
