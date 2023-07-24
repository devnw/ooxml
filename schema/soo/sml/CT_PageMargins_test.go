package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PageMarginsConstructor(t *testing.T) {
	v := sml.NewCT_PageMargins()
	if v == nil {
		t.Errorf("sml.NewCT_PageMargins must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageMargins should validate: %s", err)
	}
}

func TestCT_PageMarginsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageMargins()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageMargins()
	xml.Unmarshal(buf, v2)
}
