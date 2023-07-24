package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_BordersConstructor(t *testing.T) {
	v := sml.NewCT_Borders()
	if v == nil {
		t.Errorf("sml.NewCT_Borders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Borders should validate: %s", err)
	}
}

func TestCT_BordersMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Borders()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Borders()
	xml.Unmarshal(buf, v2)
}
