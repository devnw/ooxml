package office_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml"
)

func TestAddPreserveSpaceAttr(t *testing.T) {
	td := []struct {
		Input   string
		HasAttr bool
	}{
		{"", false},
		{"foo", false},
		{"f o o", false},
		{"foo ", true},
		{" foo ", true},
		{" foo ", true},
		{"\tfoo", true},
		{"\nfoo", true},
	}
	for _, tc := range td {
		se := &xml.StartElement{}
		office.AddPreserveSpaceAttr(se, tc.Input)
		if tc.HasAttr && len(se.Attr) == 0 {
			t.Errorf("expected a preserve space attribute for %s", tc.Input)
		} else if !tc.HasAttr && len(se.Attr) != 0 {
			t.Errorf("expected no preserve space attribute for %s", tc.Input)
		}
		if tc.HasAttr {
			if se.Attr[0].Name.Local != "xml:space" {
				t.Errorf("expected name = xml:space, got %s", se.Attr[0].Name.Local)
			}
			if se.Attr[0].Value != "preserve" {
				t.Errorf("expected name = preserve, got %s", se.Attr[0].Value)
			}
		}
	}
}
