package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_WorkbookProtectionConstructor(t *testing.T) {
	v := sml.NewCT_WorkbookProtection()
	if v == nil {
		t.Errorf("sml.NewCT_WorkbookProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_WorkbookProtection should validate: %s", err)
	}
}

func TestCT_WorkbookProtectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_WorkbookProtection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_WorkbookProtection()
	xml.Unmarshal(buf, v2)
}
