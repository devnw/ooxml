package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_AuthorsConstructor(t *testing.T) {
	v := sml.NewCT_Authors()
	if v == nil {
		t.Errorf("sml.NewCT_Authors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Authors should validate: %s", err)
	}
}

func TestCT_AuthorsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Authors()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Authors()
	xml.Unmarshal(buf, v2)
}
