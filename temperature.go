// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

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

// Celsius returns the temperature as a floating number of °Celsius.
func (t Temperature) Celsius() float64 {
	return float64(t-ZeroCelsius) / float64(Celsius)
}

// Fahrenheit returns the temperature as a floating number of °Fahrenheit.
func (t Temperature) Fahrenheit() float64 {
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
