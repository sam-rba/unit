package unit

import "testing"

func TestVolume_String(t *testing.T) {
	if s := Volume(3785 * MilliLitre).String(); s != "3.785L" {
		t.Fatalf("%#v", s)
	}
}

func TestVolume_Set(t *testing.T) {
	succeeds := []struct {
		in       string
		expected Volume
	}{
		{"1nL", NanoLitre},
		{"1uL", MicroLitre},
		{"1µL", MicroLitre},
		{"1mL", MilliLitre},
		{"1L", Litre},
		{"1kL", KiloLitre},
		{"1ML", MegaLitre},
		{"1GL", GigaLitre},
		// Maximum and minimum values that are allowed.
		{"9.223372036854775807GL", 9223372036854775807},
		{"-9.223372036854775807GL", -9223372036854775807},
	}

	fails := []struct {
		in  string
		err string
	}{
		{
			"10EL",
			"unknown unit prefix; valid prefixes for \"L\" are n,u,µ,m,k,M, or G",
		},
		{
			"10",
			"no unit provided; need L",
		},
		{
			"9.224GL",
			"maximum value is 9.223GL",
		},
		{
			"-9.224GL",
			"minimum value is -9.223GL",
		},
		{
			"9223372036854775808nL",
			"maximum value is 9.223GL",
		},
		{
			"-9223372036854775808nL",
			"minimum value is -9.223GL",
		},
		{
			"1random",
			"unknown unit provided; need L",
		},
		{
			"L",
			"not a number",
		},
		{
			"RPM",
			"does not contain number or unit L",
		},
		{
			"++1L",
			"contains multiple plus symbols",
		},
		{
			"--1L",
			"contains multiple minus symbols",
		},
		{
			"+-1L",
			"contains both plus and minus symbols",
		},
		{
			"1.1.1.1L",
			"contains multiple decimal points",
		},
		{
			string([]byte{0x33, 0x01}),
			"unexpected end of string",
		},
	}

	for i, tt := range succeeds {
		var got Volume
		if err := got.Set(tt.in); err != nil {
			t.Errorf("#%d: Volume.Set(%s) unexpected error: %v", i, tt.in, err)
		}
		if got != tt.expected {
			t.Errorf("#%d: Volume.Set(%s) wanted: %v(%d) but got: %v(%d)", i, tt.in, tt.expected, tt.expected, got, got)
		}
	}

	for i, tt := range fails {
		var got Volume
		if err := got.Set(tt.in); err == nil || err.Error() != tt.err {
			t.Errorf("#%d: Volume.Set(%s) \nexpected: %s\ngot:      %s", i, tt.in, tt.err, err)
		}
	}
}
