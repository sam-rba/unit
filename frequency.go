// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import "time"

// Frequency is a measurement of cycle per second, stored as an int64 micro
// Hertz.
//
// The highest representable value is 9.2THz.
type Frequency int64

// String returns the frequency formatted as a string in Hertz.
func (f Frequency) String() string {
	return microAsString(int64(f)) + "Hz"
}

// Period returns the duration of one cycle at this frequency.
//
// Frequency above GigaHertz cannot be represented as Duration.
//
// A 0Hz frequency returns a 0s period.
func (f Frequency) Period() time.Duration {
	if f == 0 {
		return 0
	}
	if f < 0 {
		return (time.Second*time.Duration(Hertz) - time.Duration(f/2)) / time.Duration(f)
	}
	return (time.Second*time.Duration(Hertz) + time.Duration(f/2)) / time.Duration(f)
}

// Duration returns the duration of one cycle at this frequency.
//
// Deprecated: This method is removed in v4.0.0. Use Period() instead.
func (f Frequency) Duration() time.Duration {
	return f.Period()
}

// PeriodToFrequency returns the frequency for a period of this interval.
//
// A 0s period returns a 0Hz frequency.
func PeriodToFrequency(p time.Duration) Frequency {
	if p == 0 {
		return 0
	}
	if p < 0 {
		return (Frequency(time.Second)*Hertz - Frequency(p/2)) / Frequency(p)
	}
	return (Frequency(time.Second)*Hertz + Frequency(p/2)) / Frequency(p)
}

const (
	// Hertz is 1/s.
	MicroHertz Frequency = 1
	MilliHertz Frequency = 1000 * MicroHertz
	Hertz      Frequency = 1000 * MilliHertz
	KiloHertz  Frequency = 1000 * Hertz
	MegaHertz  Frequency = 1000 * KiloHertz
	GigaHertz  Frequency = 1000 * MegaHertz
	TeraHertz  Frequency = 1000 * GigaHertz

	// RPM is revolutions per minute. It is used to quantify angular frequency.
	RPM Frequency = 16667 * MicroHertz

	maxFrequency = 9223372036854775807 * MicroHertz
	minFrequency = -9223372036854775807 * MicroHertz
)
