package zippkg_test

import "testing"
import "go.devnw.com/ooxml/zippkg"

func TestRelsPathFor(t *testing.T) {
	td := []struct {
		Inp string
		Exp string
	}{{"/", "/_rels/.rels"},
		{"/xl/workbook.xml", "/xl/_rels/workbook.xml.rels"}}
	for _, tc := range td {
		if got := zippkg.RelationsPathFor(tc.Inp); got != tc.Exp {
			t.Errorf("expected RelsPathFor(%s) = %s, got %s", tc.Inp, tc.Exp, got)
		}
	}
}
