// Copyright 2018 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Modifications 2024 Sam Anthony.

package unit

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestPicoAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1p"},
		{-1, "-1p"},
		{900, "900p"},
		{-900, "-900p"},
		{999, "999p"},
		{-999, "-999p"},
		{1000, "1n"},
		{-1000, "-1n"},
		{1100, "1.100n"},
		{-1100, "-1.100n"}, // 10
		{999999, "999.999n"},
		{-999999, "-999.999n"},
		{1000000, "1µ"},
		{-1000000, "-1µ"},
		{1000501, "1.001µ"},
		{-1000501, "-1.001µ"},
		{1100000, "1.100µ"},
		{-1100000, "-1.100µ"},
		{999999501, "1m"},
		{-999999501, "-1m"},
		{999999999, "1m"},
		{-999999999, "-1m"},
		{1000000000, "1m"},
		{-1000000000, "-1m"}, // 20
		{1100000000, "1.100m"},
		{-1100000000, "-1.100m"},
		{999999499999, "999.999m"},
		{-999999499999, "-999.999m"},
		{999999500001, "1"},
		{-999999500001, "-1"},
		{1000000000000, "1"},
		{-1000000000000, "-1"},
		{1100000000000, "1.100"},
		{-1100000000000, "-1.100"},
		{999999499999999, "999.999"},
		{-999999499999999, "-999.999"},
		{999999500000001, "1k"},
		{-999999500000001, "-1k"},
		{1000000000000000, "1k"}, //30
		{-1000000000000000, "-1k"},
		{1100000000000000, "1.100k"},
		{-1100000000000000, "-1.100k"},
		{999999499999999999, "999.999k"},
		{-999999499999999999, "-999.999k"},
		{999999500000000001, "1M"},
		{-999999500000000001, "-1M"},
		{1000000000000000000, "1M"},
		{-1000000000000000000, "-1M"},
		{1100000000000000000, "1.100M"},
		{-1100000000000000000, "-1.100M"},
		{-1999499999999999999, "-1.999M"},
		{1999499999999999999, "1.999M"},
		{-1999500000000000001, "-2M"},
		{1999500000000000001, "2M"},
		{9223372036854775807, "9.223M"},
		{-9223372036854775807, "-9.223M"},
		{-9223372036854775808, "-9.223M"},
	}
	for i, line := range data {
		if s := picoAsString(line.in); s != line.expected {
			t.Fatalf("%d: picoAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func TestNanoAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1n"},
		{-1, "-1n"},
		{900, "900n"},
		{-900, "-900n"},
		{999, "999n"},
		{-999, "-999n"},
		{1000, "1µ"},
		{-1000, "-1µ"},
		{1100, "1.100µ"},
		{-1100, "-1.100µ"}, // 10
		{999999, "999.999µ"},
		{-999999, "-999.999µ"},
		{1000000, "1m"},
		{-1000000, "-1m"},
		{1100000, "1.100m"},
		{1100100, "1.100m"},
		{1101000, "1.101m"},
		{-1100000, "-1.100m"},
		{1100499, "1.100m"},
		{1199999, "1.200m"},
		{4999501, "5m"},
		{1999501, "2m"},
		{-1100501, "-1.101m"},
		{111100501, "111.101m"},
		{999999499, "999.999m"},
		{999999501, "1"},
		{999999999, "1"},
		{1000000000, "1"},
		{-1000000000, "-1"}, // 20
		{1100000000, "1.100"},
		{-1100000000, "-1.100"},
		{1100499000, "1.100"},
		{-1100501000, "-1.101"},
		{999999499000, "999.999"},
		{999999501000, "1k"},
		{999999999999, "1k"},
		{-999999999999, "-1k"},
		{1000000000000, "1k"},
		{-1000000000000, "-1k"},
		{1100000000000, "1.100k"},
		{-1100000000000, "-1.100k"},
		{1100499000000, "1.100k"},
		{1199999000000, "1.200k"},
		{-1100501000000, "-1.101k"},
		{999999499000000, "999.999k"},
		{999999501000000, "1M"},
		{999999999999999, "1M"},
		{-999999999999999, "-1M"}, // 30
		{1000000000000000, "1M"},
		{-1000000000000000, "-1M"},
		{1100000000000000, "1.100M"},
		{-1100000000000000, "-1.100M"},
		{1100499000000000, "1.100M"},
		{-1100501000000000, "-1.101M"},
		{999999499000000000, "999.999M"},
		{999999501100000000, "1G"},
		{999999999999999999, "1G"},
		{-999999999999999999, "-1G"},
		{1000000000000000000, "1G"},
		{-1000000000000000000, "-1G"},
		{1100000000000000000, "1.100G"},
		{-1100000000000000000, "-1.100G"},
		{1999999999999999999, "2G"},
		{-1999999999999999999, "-2G"},
		{1100499000000000000, "1.100G"},
		{-1100501000000000000, "-1.101G"},
		{9223372036854775807, "9.223G"},
		{-9223372036854775807, "-9.223G"},
		{-9223372036854775808, "-9.223G"},
	}
	for i, line := range data {
		if s := nanoAsString(line.in); s != line.expected {
			t.Fatalf("%d: nanoAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func TestMicroAsString(t *testing.T) {
	data := []struct {
		in       int64
		expected string
	}{
		{0, "0"}, // 0
		{1, "1µ"},
		{-1, "-1µ"},
		{900, "900µ"},
		{-900, "-900µ"},
		{999, "999µ"},
		{-999, "-999µ"},
		{1000, "1m"},
		{-1000, "-1m"},
		{1100, "1.100m"},
		{-1100, "-1.100m"}, // 10
		{999999, "999.999m"},
		{-999999, "-999.999m"},
		{1000000, "1"},
		{-1000000, "-1"},
		{1000501, "1.001"},
		{-1000501, "-1.001"},
		{1100000, "1.100"},
		{-1100000, "-1.100"},
		{999999501, "1k"},
		{-999999501, "-1k"},
		{999999999, "1k"},
		{-999999999, "-1k"},
		{1000000000, "1k"},
		{-1000000000, "-1k"}, // 20
		{1100000000, "1.100k"},
		{-1100000000, "-1.100k"},
		{999999499999, "999.999k"},
		{-999999499999, "-999.999k"},
		{999999500001, "1M"},
		{-999999500001, "-1M"},
		{1000000000000, "1M"},
		{-1000000000000, "-1M"},
		{1100000000000, "1.100M"},
		{-1100000000000, "-1.100M"},
		{999999499999999, "999.999M"},
		{-999999499999999, "-999.999M"},
		{999999500000001, "1G"},
		{-999999500000001, "-1G"},
		{1000000000000000, "1G"}, //30
		{-1000000000000000, "-1G"},
		{1100000000000000, "1.100G"},
		{-1100000000000000, "-1.100G"},
		{999999499999999999, "999.999G"},
		{-999999499999999999, "-999.999G"},
		{999999500000000001, "1T"},
		{-999999500000000001, "-1T"},
		{1000000000000000000, "1T"},
		{-1000000000000000000, "-1T"},
		{1100000000000000000, "1.100T"},
		{-1100000000000000000, "-1.100T"},
		{-1999499999999999999, "-1.999T"},
		{1999499999999999999, "1.999T"},
		{-1999500000000000001, "-2T"},
		{1999500000000000001, "2T"},
		{9223372036854775807, "9.223T"},
		{-9223372036854775807, "-9.223T"},
		{-9223372036854775808, "-9.223T"},
	}
	for i, line := range data {
		if s := microAsString(line.in); s != line.expected {
			t.Fatalf("%d: microAsString(%d).String() = %s != %s", i, line.in, s, line.expected)
		}
	}
}

func BenchmarkCelsiusString(b *testing.B) {
	v := 10*Celsius + ZeroCelsius
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(v.String())
		buf.Reset()
	}
}

func BenchmarkCelsiusFloatf(b *testing.B) {
	v := float64(10)
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(fmt.Sprintf("%.1f°C", v))
		buf.Reset()
	}
}

func BenchmarkCelsiusFloatg(b *testing.B) {
	v := float64(10)
	buf := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(fmt.Sprintf("%g°C", v))
		buf.Reset()
	}
}

func TestAtod(t *testing.T) {
	const (
		negative = true
		positive = false
	)
	succeeds := []struct {
		in       string
		expected decimal
		n        int
	}{
		{"123456789", decimal{123456789, 0, positive}, 9},
		{"1nM", decimal{1, 0, positive}, 1},
		{"2.2", decimal{22, -1, positive}, 3},
		{"12.5mA", decimal{125, -1, positive}, 4},
		{"-12.5mA", decimal{125, -1, negative}, 5},
		{"1ma1", decimal{1, 0, positive}, 1},
		{"+1ma1", decimal{1, 0, positive}, 2},
		{"-1ma1", decimal{1, 0, negative}, 2},
		{"-0.00001%rH", decimal{1, -5, negative}, 8},
		{"0.00001%rH", decimal{1, -5, positive}, 7},
		{"1.0", decimal{1, 0, positive}, 3},
		{"0.10001", decimal{10001, -5, positive}, 7},
		{"+0.10001", decimal{10001, -5, positive}, 8},
		{"-0.10001", decimal{10001, -5, negative}, 8},
		{"1n", decimal{1, 0, positive}, 1},
		{"1.n", decimal{1, 0, positive}, 2},
		{"-1.n", decimal{1, 0, negative}, 3},
		{"200n", decimal{2, 2, positive}, 3},
		{".01", decimal{1, -2, positive}, 3},
		{"+.01", decimal{1, -2, positive}, 4},
		{"-.01", decimal{1, -2, negative}, 4},
		{"1-2", decimal{1, 0, positive}, 1},
		{"1+2", decimal{1, 0, positive}, 1},
		{"-1-2", decimal{1, 0, negative}, 2},
		{"-1+2", decimal{1, 0, negative}, 2},
		{"+1-2", decimal{1, 0, positive}, 2},
		{"+1+2", decimal{1, 0, positive}, 2},
		{"010", decimal{1, 1, positive}, 3},
		{"001", decimal{1, 0, positive}, 3},
	}

	fails := []struct {
		in       string
		expected decimal
		n        int
	}{
		{"1.1.1", decimal{}, 0},
		{"1a2b3a", decimal{}, 0},
		{"aba", decimal{}, 0},
		{"%-0.10001", decimal{}, 0},
		{"--100ma", decimal{}, 0},
		{"++100ma", decimal{}, 0},
		{"+-100ma", decimal{}, 0},
		{"-+100ma", decimal{}, 0},
	}

	for i, tt := range succeeds {
		got, n, err := atod(tt.in)
		if got != tt.expected {
			t.Errorf("#%d: case atod(\"%s\") got %v expected %v", i, tt.in, got, tt.expected)
		}
		if err != nil {
			t.Errorf("#%d: case atod(\"%s\") unexpected expected error %v", i, tt.in, err)
		}
		if n != tt.n {
			t.Errorf("#%d: case atod(\"%s\") expected to consume %d char but used %d", i, tt.in, tt.n, n)
		}
	}

	for i, tt := range fails {
		got, n, err := atod(tt.in)
		if got != tt.expected {
			t.Errorf("#%d: case atod(\"%s\") got %v expected %v", i, tt.in, got, tt.expected)
		}
		if err == nil {
			t.Errorf("#%d: case atod(\"%s\") expected error %v", i, tt.in, err)
		}
		if n != tt.n {
			t.Errorf("#%d: case atod(\"%s\") expected to consume %d char but used %d", i, tt.in, tt.n, n)
		}
	}
}

func TestDtoi(t *testing.T) {
	const (
		negative = true
		positive = false
	)
	succeeds := []struct {
		name     string
		in       decimal
		expected int64
	}{
		{"123", decimal{123, 0, positive}, 123},
		{"-123", decimal{123, 0, negative}, -123},
		{"1230", decimal{123, 1, positive}, 1230},
		{"-1230", decimal{123, 1, negative}, -1230},
		{"12.3", decimal{123, -1, positive}, 12},
		{"-12.3", decimal{123, -1, negative}, -12},
		{"123n", decimal{123, 0, positive}, 123},
		{"max", decimal{9223372036854775807, 0, positive}, 9223372036854775807},
		{"rounding(5.6)", decimal{56, -1, positive}, 6},
		{"rounding(5.5)", decimal{55, -1, positive}, 6},
		{"rounding(5.4)", decimal{54, -1, positive}, 5},
		{"rounding(-5.6)", decimal{56, -1, negative}, -6},
		{"rounding(-5.5)", decimal{55, -1, negative}, -6},
		{"rounding(-5.4)", decimal{54, -1, negative}, -5},
		{"rounding(0.6)", decimal{6, -1, positive}, 1},
		{"rounding(0.5)", decimal{5, -1, positive}, 1},
		{"rounding(0.4)", decimal{4, -1, positive}, 0},
		{"rounding(-0.6)", decimal{6, -1, negative}, -1},
		{"rounding(-0.5)", decimal{5, -1, negative}, -1},
		{"rounding(-0.4)", decimal{4, -1, negative}, -0},
	}

	fails := []struct {
		name string
		in   decimal
	}{
		{"max+1", decimal{9223372036854775808, 0, positive}},
		{"-max-1", decimal{9223372036854775808, 0, negative}},
		{"exponent too large for int64", decimal{123, 20, positive}},
		{"exponent too large negative for int64", decimal{123, -20, positive}},
		{"max*10^1", decimal{9223372036854775807, 1, positive}},
		{"-max*10^1", decimal{9223372036854775807, 1, negative}},
		{"overflow", decimal{7588728005190, 9, positive}},
	}

	for i, tt := range succeeds {
		got, overflow := dtoi(tt.in, 0)
		if got != tt.expected {
			t.Errorf("#%d: case dtoi() %s got %v expected %v", i, tt.name, got, tt.expected)
		}
		if overflow {
			t.Errorf("#%d: case dtoi() %s got an unexpected overflow", i, tt.name)
		}
	}

	for i, tt := range fails {
		got, overflow := dtoi(tt.in, 0)
		if got != 0 {
			t.Errorf("#%d: case dtoi() %s got %v expected %v", i, tt.name, got, 0)
		}
		if !overflow {
			t.Errorf("#%d: case dtoi() %s expected overflow", i, tt.name)
		}
	}
}

func Test_decimalMulScale(t *testing.T) {
	const (
		negative = true
		positive = false
	)
	succeeds := []struct {
		loss   uint
		a, b   decimal
		expect decimal
	}{
		{
			0,
			decimal{123, 0, positive},
			decimal{123, 0, positive},
			decimal{15129, 0, positive},
		},
		{
			0,
			decimal{123, 0, negative},
			decimal{123, 0, positive},
			decimal{15129, 0, negative},
		},
		{
			0,
			decimal{123, 0, positive},
			decimal{123, 0, negative},
			decimal{15129, 0, negative},
		},
		{
			0,
			decimal{123, 0, negative},
			decimal{123, 0, negative},
			decimal{15129, 0, positive},
		},
		{
			0,
			decimal{1000000001, 0, positive},
			decimal{1000000001, 0, positive},
			decimal{1000000002000000001, 0, positive},
		},
		{
			1,
			decimal{10000000001, 0, positive},
			decimal{10000000001, 0, positive},
			decimal{10000000001, 10, positive},
		},
		{
			2,
			decimal{10000000011, 0, positive},
			decimal{10000000001, 0, positive},
			decimal{1000000001, 11, positive},
		},
		{
			2,
			decimal{10000000011, 0, positive},
			decimal{10000000011, 0, positive},
			decimal{1000000002000000001, 2, positive},
		},
		{
			4,
			decimal{100000000111, 0, positive},
			decimal{100000000111, 0, positive},
			decimal{1000000002000000001, 4, positive},
		},
		{
			6,
			decimal{1000000001111, 0, positive},
			decimal{1000000001111, 0, positive},
			decimal{1000000002000000001, 6, positive},
		},
		{
			8,
			decimal{10000000011111, 0, positive},
			decimal{10000000011111, 0, positive},
			decimal{1000000002000000001, 8, positive},
		},
		{
			10,
			decimal{100000000111111, 0, positive},
			decimal{100000000111111, 0, positive},
			decimal{1000000002000000001, 10, positive},
		},
		{
			12,
			decimal{1000000001111111, 0, positive},
			decimal{1000000001111111, 0, positive},
			decimal{1000000002000000001, 12, positive},
		},
		{
			14,
			decimal{10000000011111111, 0, positive},
			decimal{10000000011111111, 0, positive},
			decimal{1000000002000000001, 14, positive},
		},
		{
			16,
			decimal{100000000111111111, 0, positive},
			decimal{100000000111111111, 0, positive},
			decimal{1000000002000000001, 16, positive},
		},
		{
			18,
			decimal{1000000001111111111, 0, positive},
			decimal{1000000001111111111, 0, positive},
			decimal{1000000002000000001, 18, positive},
		},
		{
			20,
			decimal{10000000011111111111, 0, positive},
			decimal{10000000011111111111, 0, positive},
			decimal{1000000002000000001, 20, positive},
		},
		{
			19,
			decimal{maxInt64, 0, positive},
			decimal{maxInt64, 0, positive},
			decimal{8507059176058364548, 19, positive},
		},
		{
			18,
			decimal{(1 << 64) - 6, 0, positive},
			decimal{(1 << 64) - 6, 0, positive},
			decimal{3402823667840801649, 20, positive},
		},
		{
			0,
			decimal{(1 << 64) - 6, 100, positive},
			decimal{0, 0, positive},
			decimal{0, 0, positive},
		},
	}

	fails := []struct {
		loss   uint
		a, b   decimal
		expect decimal
	}{
		{
			21,
			decimal{(1 << 64) - 5, 0, positive},
			decimal{(1 << 64) - 5, 0, positive},
			decimal{},
		},
	}

	for i, tt := range succeeds {
		got, loss := decimalMul(tt.a, tt.b)
		if loss != tt.loss {
			t.Errorf("#%d: decimalMulScale(%v,%v) expected %d loss but got %d", i, tt.a, tt.b, tt.loss, loss)
		}
		if got != tt.expect {
			t.Errorf("#%d: decimalMulScale(%v,%v) got: %v expected: %v", i, tt.a, tt.b, got, tt.expect)
		}
	}

	for i, tt := range fails {
		got, loss := decimalMul(tt.a, tt.b)
		if loss != tt.loss {
			t.Errorf("#%d: decimalMulScale(%v,%v) expected %d loss but got %d", i, tt.a, tt.b, tt.loss, loss)
		}
		if got != tt.expect {
			t.Errorf("#%d: decimalMulScale(%v,%v) got: %v expected: %v", i, tt.a, tt.b, got, tt.expect)
		}
	}
}

func TestPrefix(t *testing.T) {
	tests := []struct {
		name   string
		prefix rune
		want   prefix
		n      int
	}{
		{"pico", 'p', pico, 1},
		{"nano", 'n', nano, 1},
		{"micro", 'u', micro, 1},
		{"mu", 'µ', micro, 2},
		{"milli", 'm', milli, 1},
		{"unit", 0, unit, 0},
		{"kilo", 'k', kilo, 1},
		{"mega", 'M', mega, 1},
		{"giga", 'G', giga, 1},
		{"tera", 'T', tera, 1},
	}
	for i, tt := range tests {
		got, n := parseSIPrefix(tt.prefix)
		if got != tt.want || n != tt.n {
			t.Errorf("#%d: wanted prefix %d, and len %d, but got prefix %d, and len %d", i, tt.want, tt.n, got, n)
		}
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"empty", &parseError{errors.New("test")}, "test"},
		{"noUnits", noUnitErr("someunit"), "no unit provided; need someunit"},
	}
	for i, tt := range tests {
		if got := tt.err.Error(); got != tt.want {
			t.Errorf("#%d: not the expected error.\nwanted: %s\ngot:    %s", i, tt.want, got)
		}
	}
}

func TestMaxInt64(t *testing.T) {
	if strconv.FormatUint(maxInt64, 10) != maxInt64Str {
		t.Fatal("unexpected text representation of max")
	}
}

func TestValueOfUnitString(t *testing.T) {
	succeeds := []struct {
		in        string
		uintbase  prefix
		expected  int64
		usedChars int
	}{
		{"1p", pico, 1, 2},
		{"1n", pico, 1000, 2},
		{"1u", pico, 1000000, 2},
		{"1µ", pico, 1000000, 3},
		{"1m", pico, 1000000000, 2},
		{"1k", pico, 1000000000000000, 2},
		{"1M", pico, 1000000000000000000, 2},
		{"9.223372036854775807M", pico, 9223372036854775807, 21},
		{"9223372036854775807p", pico, 9223372036854775807, 20},
		{"-1p", pico, -1, 3},
		{"-1n", pico, -1000, 3},
		{"-1u", pico, -1000000, 3},
		{"-1µ", pico, -1000000, 4},
		{"-1m", pico, -1000000000, 3},
		{"-1k", pico, -1000000000000000, 3},
		{"-1M", pico, -1000000000000000000, 3},
		{"-9.223372036854775807M", pico, -9223372036854775807, 22},
		{"-9223372036854775807p", pico, -9223372036854775807, 21},
		{"1p", nano, 0, 2},
		{"1n", nano, 1, 2},
		{"1u", nano, 1000, 2},
		{"1µ", nano, 1000, 3},
		{"1m", nano, 1000000, 2},
		{"1k", nano, 1000000000000, 2},
		{"1M", nano, 1000000000000000, 2},
		{"1G", nano, 1000000000000000000, 2},
		{"9.223372036854775807G", nano, 9223372036854775807, 21},
		{"9223372036854775807n", nano, 9223372036854775807, 20},
		{"-1p", nano, -0, 3},
		{"-1n", nano, -1, 3},
		{"-1u", nano, -1000, 3},
		{"-1µ", nano, -1000, 4},
		{"-1m", nano, -1000000, 3},
		{"-1k", nano, -1000000000000, 3},
		{"-1M", nano, -1000000000000000, 3},
		{"-1G", nano, -1000000000000000000, 3},
		{"-9.223372036854775807G", nano, -9223372036854775807, 22},
		{"-9223372036854775807n", nano, -9223372036854775807, 21},
		{"1p", micro, 0, 2},
		{"1n", micro, 0, 2},
		{"1u", micro, 1, 2},
		{"1µ", micro, 1, 3},
		{"1m", micro, 1000, 2},
		{"1k", micro, 1000000000, 2},
		{"1M", micro, 1000000000000, 2},
		{"1G", micro, 1000000000000000, 2},
		{"1T", micro, 1000000000000000000, 2},
		{"9.223372036854775807T", micro, 9223372036854775807, 21},
		{"9223372036854775807u", micro, 9223372036854775807, 20},
		{"-1p", micro, -0, 3},
		{"-1n", micro, -0, 3},
		{"-1u", micro, -1, 3},
		{"-1µ", micro, -1, 4},
		{"-1m", micro, -1000, 3},
		{"-1k", micro, -1000000000, 3},
		{"-1M", micro, -1000000000000, 3},
		{"-1G", micro, -1000000000000000, 3},
		{"-1T", micro, -1000000000000000000, 3},
		{"-9.223372036854775807T", micro, -9223372036854775807, 22},
		{"-9223372036854775807u", micro, -9223372036854775807, 21},
	}

	fails := []struct {
		in     string
		prefix prefix
	}{
		{"9.223372036854775808M", pico},
		{"9.223372036854775808G", nano},
		{"9.223372036854775808T", micro},
		{"9223372036854775808p", pico},
		{"9223372036854775808n", nano},
		{"9223372036854775808u", micro},
		{"-9.223372036854775808M", pico},
		{"-9.223372036854775808G", nano},
		{"-9.223372036854775808T", micro},
		{"-9223372036854775808p", pico},
		{"-9223372036854775808n", nano},
		{"-9223372036854775808u", micro},
		{"not a number", nano},
		{string([]byte{0x31, 0x01}), nano}, // 0x01 is a invalid utf8 start byte.
	}

	for i, tt := range succeeds {
		got, used, err := valueOfUnitString(tt.in, tt.uintbase)
		if got != tt.expected {
			t.Errorf("#%d: valueOfUnitString(%s,%d) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.uintbase, tt.expected, tt.expected, got, got)
		}
		if used != tt.usedChars {
			t.Errorf("#%d: valueOfUnitString(%s,%d) used %d chars but should use: %d chars", i, tt.in, tt.uintbase, used, tt.usedChars)
		}
		if err != nil {
			t.Errorf("#%d: valueOfUnitString(%s,%d) unexpected error: %v", i, tt.in, tt.uintbase, err)
		}
	}

	for i, tt := range fails {
		if _, _, err := valueOfUnitString(tt.in, tt.prefix); err == nil {
			t.Errorf("#%d: valueOfUnitString(%s,%d) expected an error", i, tt.in, tt.prefix)
		}
	}
}

// Benchmarks

func BenchmarkDecimal(b *testing.B) {
	var d decimal
	var n int
	var err error
	for i := 0; i < b.N; i++ {
		if d, n, err = atod("337.2m"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v %d", d, n)
}

func BenchmarkDecimal2Int(b *testing.B) {
	d := decimal{1234, 5, false}
	var overflow bool
	var v int64
	for i := 0; i < b.N; i++ {
		if v, overflow = dtoi(d, 0); overflow {
			b.Fatal("unexpected overflow")
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", v)
}

func BenchmarkString2Decimal2Int(b *testing.B) {
	var d decimal
	var n int
	var err error
	var overflow bool
	var v int64
	for i := 0; i < b.N; i++ {
		if d, n, err = atod("337.2m"); err != nil {
			b.Fatal(err)
		}
		if v, overflow = dtoi(d, 0); overflow {
			b.Fatal("unexpected overflow")
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d %d", v, n)
}

func BenchmarkDecimalNeg(b *testing.B) {
	var d decimal
	var n int
	var err error
	for i := 0; i < b.N; i++ {
		if d, n, err = atod("-337.2m"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%v %d", d, n)
}

func BenchmarkString2Decimal2IntNeg(b *testing.B) {
	var d decimal
	var n int
	var err error
	var overflow bool
	var v int64
	for i := 0; i < b.N; i++ {
		if d, n, err = atod("-337.2m"); err != nil {
			b.Fatal(err)
		}
		if v, overflow = dtoi(d, 0); overflow {
			b.Fatal("unexpected overflow")
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d %d", v, n)
}

func BenchmarkDistanceSet(b *testing.B) {
	var err error
	var d Distance
	for i := 0; i < b.N; i++ {
		if err = d.Set("1ft"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", d)
}

func BenchmarkElectricCurrentSet(b *testing.B) {
	var err error
	var e ElectricCurrent
	for i := 0; i < b.N; i++ {
		if err = e.Set("1A"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", e)
}

func BenchmarkForceSetMetric(b *testing.B) {
	var err error
	var f Force
	for i := 0; i < b.N; i++ {
		if err = f.Set("123N"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", f)
}

func BenchmarkForceSetImperial(b *testing.B) {
	var err error
	var f Force
	for i := 0; i < b.N; i++ {
		if err = f.Set("1.23Mlbf"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", f)
}

func BenchmarkForceSetImperialWorstCase(b *testing.B) {
	var err error
	var f Force
	for i := 0; i < b.N; i++ {
		if err = f.Set("1.0000000000101lbf"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", f)
}

func BenchmarkAngleSetRadian(b *testing.B) {
	var err error
	var a Angle
	for i := 0; i < b.N; i++ {
		if err = a.Set("1rad"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", a)
}

func BenchmarkAngleSet1Degree(b *testing.B) {
	var err error
	var a Angle
	for i := 0; i < b.N; i++ {
		if err = a.Set("1deg"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", a)
}

func BenchmarkAngleSet2Degree(b *testing.B) {
	var err error
	var a Angle
	for i := 0; i < b.N; i++ {
		if err = a.Set("2deg"); err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	_ = fmt.Sprintf("%d", a)
}
