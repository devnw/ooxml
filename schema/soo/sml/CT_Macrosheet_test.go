package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MacrosheetConstructor(t *testing.T) {
	v := sml.NewCT_Macrosheet()
	if v == nil {
		t.Errorf("sml.NewCT_Macrosheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Macrosheet should validate: %s", err)
	}
}

func TestCT_MacrosheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Macrosheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Macrosheet()
	xml.Unmarshal(buf, v2)
}
