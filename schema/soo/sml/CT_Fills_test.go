package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FillsConstructor(t *testing.T) {
	v := sml.NewCT_Fills()
	if v == nil {
		t.Errorf("sml.NewCT_Fills must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Fills should validate: %s", err)
	}
}

func TestCT_FillsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Fills()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Fills()
	xml.Unmarshal(buf, v2)
}
