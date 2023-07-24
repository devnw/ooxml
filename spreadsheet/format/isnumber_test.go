package format_test

import (
	"testing"

	"go.devnw.com/ooxml/spreadsheet/format"
)

func TestIsNumber(t *testing.T) {
	td := []struct {
		Inp string
		Exp bool
	}{
		{"123", true},
		{"1.23", true},
		{"1.23.", false},
		{"1.23E+10", true},
		{"1.23E-10", true},
		{"1.23E10", false},
		{"1213131312312312390", true},
		{"0", true},
		{"", false},
		{"abc", false},
	}
	for _, tc := range td {
		got := format.IsNumber(tc.Inp)
		if got != tc.Exp {
			t.Errorf("expected IsNumber(%s) = %v, got %v", tc.Inp, tc.Exp, got)
		}
	}
}
