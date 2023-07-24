package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ColorsConstructor(t *testing.T) {
	v := sml.NewCT_Colors()
	if v == nil {
		t.Errorf("sml.NewCT_Colors must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Colors should validate: %s", err)
	}
}

func TestCT_ColorsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Colors()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Colors()
	xml.Unmarshal(buf, v2)
}
