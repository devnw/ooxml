package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_WriteProtectionConstructor(t *testing.T) {
	v := wml.NewCT_WriteProtection()
	if v == nil {
		t.Errorf("wml.NewCT_WriteProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_WriteProtection should validate: %s", err)
	}
}

func TestCT_WriteProtectionMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_WriteProtection()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_WriteProtection()
	xml.Unmarshal(buf, v2)
}
