package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RstConstructor(t *testing.T) {
	v := sml.NewCT_Rst()
	if v == nil {
		t.Errorf("sml.NewCT_Rst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Rst should validate: %s", err)
	}
}

func TestCT_RstMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Rst()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Rst()
	xml.Unmarshal(buf, v2)
}
