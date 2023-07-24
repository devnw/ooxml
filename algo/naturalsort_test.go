package algo_test

import (
	"testing"

	"go.devnw.com/ooxml/algo"
)

func TestSort(t *testing.T) {

	tests := []struct {
		a, b string
	}{
		{"rId1", "rId2"},
		{"rId1", "rId10"},
		{"rId2", "rId10"},
		{"rId5", "rId10"},
		{"rId5", "rId15"},
		{"rId5", "rId51"},
		{"rId1a", "rId1b"},
	}

	for _, tc := range tests {
		if !algo.NaturalLess(tc.a, tc.b) {
			t.Errorf("bad sort, expected %s < %s", tc.a, tc.b)
		} else {
			// no need to check if it failed the first time
			if algo.NaturalLess(tc.b, tc.a) {
				t.Errorf("bad sort, expected %s > %s", tc.b, tc.a)
			}
		}
	}

}
