package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CellProtectionConstructor(t *testing.T) {
	v := sml.NewCT_CellProtection()
	if v == nil {
		t.Errorf("sml.NewCT_CellProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CellProtection should validate: %s", err)
	}
}

func TestCT_CellProtectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CellProtection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CellProtection()
	xml.Unmarshal(buf, v2)
}
