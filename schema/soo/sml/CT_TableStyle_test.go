package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TableStyleConstructor(t *testing.T) {
	v := sml.NewCT_TableStyle()
	if v == nil {
		t.Errorf("sml.NewCT_TableStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TableStyle should validate: %s", err)
	}
}

func TestCT_TableStyleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TableStyle()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TableStyle()
	xml.Unmarshal(buf, v2)
}
