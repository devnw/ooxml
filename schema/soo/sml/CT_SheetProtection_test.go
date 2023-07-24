package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetProtectionConstructor(t *testing.T) {
	v := sml.NewCT_SheetProtection()
	if v == nil {
		t.Errorf("sml.NewCT_SheetProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetProtection should validate: %s", err)
	}
}

func TestCT_SheetProtectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetProtection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetProtection()
	xml.Unmarshal(buf, v2)
}
