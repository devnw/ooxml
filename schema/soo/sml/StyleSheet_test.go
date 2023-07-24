package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestStyleSheetConstructor(t *testing.T) {
	v := sml.NewStyleSheet()
	if v == nil {
		t.Errorf("sml.NewStyleSheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.StyleSheet should validate: %s", err)
	}
}

func TestStyleSheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewStyleSheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewStyleSheet()
	xml.Unmarshal(buf, v2)
}
