// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit_test

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sam-rba/unit"
)

func ExampleAngle() {
	fmt.Println(unit.Degree)
	fmt.Println(unit.Pi)
	fmt.Println(unit.Theta)
	// Output:
	// 1.000°
	// 180.0°
	// 360.0°
}

func ExampleAngle_Set() {
	var a unit.Angle

	if err := a.Set("2°"); err != nil {
		log.Fatal(a)
	}
	fmt.Println(a)

	if err := a.Set("90deg"); err != nil {
		log.Fatal(a)
	}
	fmt.Println(a)

	if err := a.Set("1rad"); err != nil {
		log.Fatal(a)
	}
	fmt.Println(a)
	// Output:
	// 2.000°
	// 90.00°
	// 57.296°
}

func ExampleAngle_flag() {
	var a unit.Angle

	flag.Var(&a, "angle", "angle to set the servo to")
	flag.Parse()
}

func ExampleAngle_float64() {
	// A 45° angle. The +2 here is to help integer based rounding.
	v := (unit.Pi + 2) / 4

	// Convert to float64 as degree.
	fd := float64(v) / float64(unit.Degree)

	// Convert to float64 as radian.
	fr := float64(v) / float64(unit.Radian)

	fmt.Println(v)
	fmt.Printf("%.1fdeg\n", fd)
	fmt.Printf("%frad\n", fr)
	// Output:
	// 45.00°
	// 45.0deg
	// 0.785398rad
}

func ExampleDistance() {
	fmt.Println(unit.Inch)
	fmt.Println(unit.Foot)
	fmt.Println(unit.Mile)
	// Output:
	// 25.400mm
	// 304.800mm
	// 1.609km
}

func ExampleDistance_Set() {
	var d unit.Distance

	if err := d.Set("1ft"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	if err := d.Set("1m"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	if err := d.Set("9Mile"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	// Output:
	// 304.800mm
	// 1m
	// 14.484km

}

func ExampleDistance_flag() {
	var d unit.Distance

	flag.Var(&d, "distance", "x axis travel length")
	flag.Parse()
}

func ExampleDistance_float64() {
	// Distance between the Earth and the Moon.
	v := 384400 * unit.KiloMetre

	// Convert to float64 as meter.
	fm := float64(v) / float64(unit.Metre)

	// Convert to float64 as inches.
	fi := float64(v) / float64(unit.Inch)

	fmt.Println(v)
	fmt.Printf("%.0fm\n", fm)
	fmt.Printf("%.0fin\n", fi)
	// Output:
	// 384.400Mm
	// 384400000m
	// 15133858268in
}

func ExampleElectricalCapacitance() {
	fmt.Println(1 * unit.Farad)
	fmt.Println(22 * unit.PicoFarad)
	// Output:
	// 1F
	// 22pF
}

func ExampleElectricalCapacitance_Set() {
	var c unit.ElectricalCapacitance

	if err := c.Set("1F"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	if err := c.Set("22pF"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	// Output:
	// 1F
	// 22pF
}

func ExampleElectricalCapacitance_flag() {
	var c unit.ElectricalCapacitance

	flag.Var(&c, "mintouch", "minimum touch sensitivity")
	flag.Parse()
}

func ExampleElectricalCapacitance_float64() {
	// A typical condensator.
	v := 4700 * unit.NanoFarad

	// Convert to float64 as microfarad.
	f := float64(v) / float64(unit.MicroFarad)

	fmt.Println(v)
	fmt.Printf("%.1fµF\n", f)
	// Output:
	// 4.700µF
	// 4.7µF
}

func ExampleElectricCurrent() {
	fmt.Println(10010 * unit.MilliAmpere)
	fmt.Println(10 * unit.Ampere)
	fmt.Println(-10 * unit.MilliAmpere)
	// Output:
	// 10.010A
	// 10A
	// -10mA
}

func ExampleElectricCurrent_Set() {
	var e unit.ElectricCurrent

	if err := e.Set("12.5mA"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	if err := e.Set("2.4kA"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	if err := e.Set("2A"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)
	// Output:
	// 12.500mA
	// 2.400kA
	// 2A
}

func ExampleElectricCurrent_flag() {
	var m unit.ElectricCurrent

	flag.Var(&m, "motor", "rated motor current")
	flag.Parse()
}

func ExampleElectricCurrent_float64() {
	// The maximum current that can be drawn from all Raspberry Pi GPIOs combined.
	v := 51 * unit.MilliAmpere

	// Convert to float64 as ampere.
	f := float64(v) / float64(unit.Ampere)

	fmt.Println(v)
	fmt.Printf("%.3fA\n", f)
	// Output:
	// 51mA
	// 0.051A
}

func ExampleElectricPotential() {
	fmt.Println(10010 * unit.MilliVolt)
	fmt.Println(10 * unit.Volt)
	fmt.Println(-10 * unit.MilliVolt)
	// Output:
	// 10.010V
	// 10V
	// -10mV
}

func ExampleElectricPotential_Set() {
	var v unit.ElectricPotential
	if err := v.Set("250uV"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	if err := v.Set("100kV"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	if err := v.Set("12V"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
	// Output:
	// 250µV
	// 100kV
	// 12V
}

func ExampleElectricPotential_flag() {
	var v unit.ElectricPotential
	flag.Var(&v, "cutout", "battery full charge voltage")
	flag.Parse()
}

func ExampleElectricPotential_float64() {
	// The level of Raspberry Pi GPIO when high.
	v := 3300 * unit.MilliVolt

	// Convert to float64 as volt.
	f := float64(v) / float64(unit.Volt)

	fmt.Println(v)
	fmt.Printf("%.1fV\n", f)
	// Output:
	// 3.300V
	// 3.3V
}

func ExampleElectricResistance() {
	fmt.Println(10010 * unit.MilliOhm)
	fmt.Println(10 * unit.Ohm)
	fmt.Println(24 * unit.MegaOhm)
	// Output:
	// 10.010Ω
	// 10Ω
	// 24MΩ
}

func ExampleElectricResistance_Set() {
	var r unit.ElectricResistance

	if err := r.Set("33.3kOhm"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	if err := r.Set("1Ohm"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	if err := r.Set("5MOhm"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// Output:
	// 33.300kΩ
	// 1Ω
	// 5MΩ
}

func ExampleElectricResistance_flag() {
	var r unit.ElectricResistance

	flag.Var(&r, "shunt", "shunt resistor value")
	flag.Parse()
}

func ExampleElectricResistance_float64() {
	// A common resistor value.
	v := 4700 * unit.Ohm

	// Convert to float64 as ohm.
	f := float64(v) / float64(unit.Ohm)

	fmt.Println(v)
	fmt.Printf("%.0fOhm\n", f)
	// Output:
	// 4.700kΩ
	// 4700Ohm
}

func ExampleEnergy() {
	fmt.Println(1 * unit.Joule)
	fmt.Println(1 * unit.WattSecond)
	fmt.Println(1 * unit.KiloWattHour)
	// Output:
	// 1J
	// 1J
	// 3.600MJ
}

func ExampleEnergy_Set() {
	var e unit.Energy

	if err := e.Set("2.6kJ"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	if err := e.Set("45mJ"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	// Output:
	// 2.600kJ
	// 45mJ
}

func ExampleEnergy_flag() {
	var e unit.Energy

	flag.Var(&e, "capacity", "capacity of battery")
	flag.Parse()
}

func ExampleEnergy_float64() {
	// BTU is used in barbecue rating. It is a measure of the thermal energy
	// content of the of fuel consumed by the barbecue at its maximum rate. A
	// 35000 BTU barbecue is an average power.
	v := 35000 * unit.BTU

	// Convert to float64 as joule.
	f := float64(v) / float64(unit.Joule)

	fmt.Println(v)
	fmt.Printf("%.0f\n", f)
	// Output:
	// 36.927MJ
	// 36927100
}

func ExampleForce() {
	fmt.Println(10 * unit.MilliNewton)
	fmt.Println(unit.EarthGravity)
	fmt.Println(unit.PoundForce)
	// Output:
	// 10mN
	// 9.807N
	// 4.448N
}

func ExampleForce_Set() {
	var f unit.Force

	if err := f.Set("9.8N"); err != nil {
		log.Fatal(f)
	}
	fmt.Println(f)

	if err := f.Set("20lbf"); err != nil {
		log.Fatal(f)
	}
	fmt.Println(f)

	// Output:
	// 9.800N
	// 88.964N
}

func ExampleForce_flag() {
	var f unit.Force

	flag.Var(&f, "force", "load cell wakeup force")
	flag.Parse()
}

func ExampleForce_float64() {
	// The gravity of Earth.
	v := unit.EarthGravity

	// Convert to float64 as newton.
	fn := float64(v) / float64(unit.Newton)

	// Convert to float64 as lbf.
	fl := float64(v) / float64(unit.PoundForce)

	fmt.Println(v)
	fmt.Printf("%fN\n", fn)
	fmt.Printf("%flbf\n", fl)
	// Output:
	// 9.807N
	// 9.806650N
	// 2.204623lbf
}

func ExampleFrequency() {
	fmt.Println(10 * unit.MilliHertz)
	fmt.Println(101010 * unit.MilliHertz)
	fmt.Println(10 * unit.MegaHertz)
	fmt.Println(60 * unit.RPM)
	// Output:
	// 10mHz
	// 101.010Hz
	// 10MHz
	// 1Hz
}

func ExampleFrequency_Period() {
	fmt.Println(unit.MilliHertz.Period())
	fmt.Println(unit.MegaHertz.Period())
	// Output:
	// 16m40s
	// 1µs
}

func ExampleFrequency_Set() {
	var f unit.Frequency

	if err := f.Set("10MHz"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)

	if err := f.Set("10mHz"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)

	if err := f.Set("1kHz"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	// Output:
	// 10MHz
	// 10mHz
	// 1kHz
}

func ExampleFrequency_flag() {
	var pwm unit.Frequency

	flag.Var(&pwm, "pwm", "pwm frequency")
	flag.Parse()
}

func ExampleFrequency_float64() {
	// NTSC color subcarrier.
	v := (315*unit.MegaHertz + 44) / 88

	// Convert to float64 as hertz.
	f := float64(v) / float64(unit.Hertz)

	fmt.Println(v)
	fmt.Printf("%fHz\n", f)
	// Output:
	// 3.580MHz
	// 3579545.454545Hz
}

func ExamplePeriodToFrequency() {
	fmt.Println(unit.PeriodToFrequency(time.Microsecond))
	fmt.Println(unit.PeriodToFrequency(time.Minute))
	// Output:
	// 1MHz
	// 16.667mHz
}

func ExampleLuminousFlux() {
	fmt.Println(18282 * unit.Lumen)
	// Output:
	// 18.282klm
}

func ExampleLuminousFlux_Set() {
	var l unit.LuminousFlux

	if err := l.Set("25mlm"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)

	if err := l.Set("2.5Mlm"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)

	// Output:
	// 25mlm
	// 2.500Mlm
}

func ExampleLuminousFlux_flag() {
	var l unit.LuminousFlux

	flag.Var(&l, "low", "mood light level")
	flag.Parse()
}

func ExampleLuminousFlux_float64() {
	// Typical output of a 7W LED.
	v := 450 * unit.Lumen

	// Convert to float64 as lumen.
	f := float64(v) / float64(unit.Lumen)

	fmt.Println(v)
	fmt.Printf("%.0f\n", f)
	// Output:
	// 450lm
	// 450
}

func ExampleLuminousIntensity() {
	fmt.Println(12 * unit.Candela)
	// Output:
	// 12cd
}

func ExampleLuminousIntensity_Set() {
	var l unit.LuminousIntensity

	if err := l.Set("16cd"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)

	// Output:
	// 16cd
}

func ExampleLuminousIntensity_flag() {
	var l unit.LuminousIntensity

	flag.Var(&l, "dusk", "light level to turn on light")
	flag.Parse()
}

func ExampleLuminousIntensity_float64() {
	// A 7W LED generating 450lm in all directions (4π steradian) will have an
	// intensity of 450lm/4π = ~35.8cd.
	v := 35800 * unit.MilliCandela

	// Convert to float64 as candela.
	f := float64(v) / float64(unit.Candela)

	fmt.Println(v)
	fmt.Printf("%.1f\n", f)
	// Output:
	// 35.800cd
	// 35.8
}

func ExampleMass() {
	fmt.Println(10 * unit.MilliGram)
	fmt.Println(unit.OunceMass)
	fmt.Println(unit.PoundMass)
	fmt.Println(unit.Slug)
	// Output:
	// 10mg
	// 28.350g
	// 453.592g
	// 14.594kg
}

func ExampleMass_Set() {
	var m unit.Mass

	if err := m.Set("10mg"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	if err := m.Set("16.5kg"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	if err := m.Set("2.2oz"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	if err := m.Set("16lb"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)
	// Output:
	// 10mg
	// 16.500kg
	// 62.369g
	// 7.257kg
}

func ExampleMass_flag() {
	var m unit.Mass

	flag.Var(&m, "weight", "amount of cat food to dispense")
	flag.Parse()
}

func ExampleMass_float64() {
	// Weight of a Loonie (Canadian metal piece).
	v := 6270 * unit.MilliGram

	// Convert to float64 as gram.
	f := float64(v) / float64(unit.Gram)

	fmt.Println(v)
	fmt.Printf("%.2fg\n", f)
	// Output:
	// 6.270g
	// 6.27g
}

func ExamplePower() {
	fmt.Println(1 * unit.Watt)
	fmt.Println(16 * unit.MilliWatt)
	fmt.Println(1210 * unit.MegaWatt)
	// Output:
	// 1W
	// 16mW
	// 1.210GW
}

func ExamplePower_Set() {
	var p unit.Power

	if err := p.Set("25mW"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	if err := p.Set("1W"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	if err := p.Set("1.21GW"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	// Output:
	// 25mW
	// 1W
	// 1.210GW
}

func ExamplePower_flag() {
	var p unit.Power

	flag.Var(&p, "power", "heater maximum power")
	flag.Parse()
}

func ExamplePower_float64() {
	// Maximum emitted power by Bluetooth class 2 device.
	v := 2500 * unit.MicroWatt

	// Convert to float64 as watt.
	f := float64(v) / float64(unit.Watt)

	fmt.Println(v)
	fmt.Printf("%.4f\n", f)
	// Output:
	// 2.500mW
	// 0.0025
}

func ExamplePressure() {
	fmt.Println(101010 * unit.Pascal)
	fmt.Println(101 * unit.KiloPascal)
	// Output:
	// 101.010kPa
	// 101kPa
}

func ExamplePressure_Set() {
	var p unit.Pressure

	if err := p.Set("300kPa"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	if err := p.Set("16MPa"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	// Output:
	// 300kPa
	// 16MPa
}

func ExamplePressure_flag() {
	var p unit.Pressure

	flag.Var(&p, "setpoint", "pressure for pump to maintain")
	flag.Parse()
}

func ExamplePressure_float64() {
	// A typical tire pressure in North America (33psi).
	v := 227526990 * unit.MicroPascal

	// Convert to float64 as pascal.
	f := float64(v) / float64(unit.Pascal)

	fmt.Println(v)
	fmt.Printf("%f\n", f)
	// Output:
	// 227.527Pa
	// 227.526990
}

func ExampleRelativeHumidity() {
	fmt.Println(506 * unit.MilliRH)
	fmt.Println(20 * unit.PercentRH)
	// Output:
	// 50.6%rH
	// 20%rH
}

func ExampleRelativeHumidity_Set() {
	var r unit.RelativeHumidity

	if err := r.Set("50.6%rH"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)

	if err := r.Set("20%"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// Output:
	// 50.6%rH
	// 20%rH
}

func ExampleRelativeHumidity_flag() {
	var h unit.RelativeHumidity

	flag.Var(&h, "humidity", "green house humidity high alarm level")
	flag.Parse()
}

func ExampleRelativeHumidity_float64() {
	// Yearly humidity average of Nadi, Fiji.
	v := 800 * unit.MilliRH

	// Convert to float64 as %.
	f := float64(v) / float64(unit.PercentRH)

	fmt.Println(v)
	fmt.Printf("%.0f\n", f)
	// Output:
	// 80%rH
	// 80
}

func ExampleSpeed() {
	fmt.Println(10 * unit.MilliMetrePerSecond)
	fmt.Println(unit.LightSpeed)
	fmt.Println(unit.KilometrePerHour)
	fmt.Println(unit.MilePerHour)
	fmt.Println(unit.FootPerSecond)
	// Output:
	// 10mm/s
	// 299.792Mm/s
	// 277.778mm/s
	// 447.040mm/s
	// 304.800mm/s
}

func ExampleSpeed_Set() {
	var s unit.Speed

	if err := s.Set("10m/s"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	if err := s.Set("100kph"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	if err := s.Set("2067fps"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)

	if err := s.Set("55mph"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	// Output:
	// 10m/s
	// 27.778m/s
	// 630.022m/s
	// 24.587m/s
}

func ExampleSpeed_flag() {
	var s unit.Speed

	flag.Var(&s, "speed", "window shutter closing speed")
	flag.Parse()
}

func ExampleSpeed_float64() {
	// Speed of light in vacuum.
	v := unit.LightSpeed

	// Convert to float64 as m/s.
	fms := float64(v) / float64(unit.MetrePerSecond)

	// Convert to float64 as km/h.
	fkh := float64(v) / float64(unit.KilometrePerHour)

	// Convert to float64 as mph.
	fmh := float64(v) / float64(unit.MilePerHour)

	fmt.Println(v)
	fmt.Printf("%.0fm/s\n", fms)
	fmt.Printf("%.1fkm/h\n", fkh)
	fmt.Printf("%.1fmph\n", fmh)
	// Output:
	// 299.792Mm/s
	// 299792458m/s
	// 1079252847.9km/h
	// 670616629.4mph
}

func ExampleTemperature() {
	fmt.Println(0 * unit.Kelvin)
	fmt.Println(23010*unit.MilliCelsius + unit.ZeroCelsius)
	fmt.Println(80*unit.Fahrenheit + unit.ZeroFahrenheit)
	// Output:
	// -273.150°C
	// 23.010°C
	// 26.667°C
}

func ExampleTemperature_Celsius() {
	// Normal average human body temperature.
	v := 37*unit.Celsius + unit.ZeroCelsius

	// Convert to float64 as Celsius.
	f := v.C()

	fmt.Println(v)
	fmt.Printf("%.1f°C\n", f)
	// Output:
	// 37°C
	// 37.0°C
}

func ExampleTemperature_Fahrenheit() {
	// Normal average human body temperature.
	v := 37*unit.Celsius + unit.ZeroCelsius

	// Convert to float64 as Fahrenheit.
	f := v.F()

	fmt.Println(v)
	fmt.Printf("%.1f°F\n", f)
	// Output:
	// 37°C
	// 98.6°F
}

func ExampleTemperature_Set() {
	var t unit.Temperature

	if err := t.Set("0°C"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	if err := t.Set("1C"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	if err := t.Set("5MK"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	if err := t.Set("0°F"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	if err := t.Set("32F"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)

	// Output:
	// 0°C
	// 1°C
	// 5M°C
	// -17.778°C
	// 0°C
}

func ExampleTemperature_flag() {
	var t unit.Temperature

	flag.Var(&t, "temp", "thermostat setpoint")
	flag.Parse()
}

func ExampleTemperature_float64() {
	// Normal average human body temperature.
	v := 37*unit.Celsius + unit.ZeroCelsius

	// Convert to float64 as Kelvin.
	f := float64(v) / float64(unit.Kelvin)

	fmt.Println(v)
	fmt.Printf("%.1fK\n", f)
	// Output:
	// 37°C
	// 310.1K
}
