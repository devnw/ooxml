package spreadsheet_test

import (
	"testing"

	"go.devnw.com/ooxml/spreadsheet"
)

func TestKnownHashes(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{
		{"gooxml", "DD67"},
		{"", "0000"},
	}
	for _, tc := range td {
		if got := spreadsheet.PasswordHash(tc.Inp); got != tc.Exp {
			t.Errorf("expected hash of %s = %s, got %s", tc.Inp, tc.Exp, got)
		}
	}
}
