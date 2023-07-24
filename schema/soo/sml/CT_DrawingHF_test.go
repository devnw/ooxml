package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DrawingHFConstructor(t *testing.T) {
	v := sml.NewCT_DrawingHF()
	if v == nil {
		t.Errorf("sml.NewCT_DrawingHF must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DrawingHF should validate: %s", err)
	}
}

func TestCT_DrawingHFMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DrawingHF()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DrawingHF()
	xml.Unmarshal(buf, v2)
}
