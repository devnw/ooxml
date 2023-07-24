package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SpacingConstructor(t *testing.T) {
	v := wml.NewCT_Spacing()
	if v == nil {
		t.Errorf("wml.NewCT_Spacing must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Spacing should validate: %s", err)
	}
}

func TestCT_SpacingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Spacing()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Spacing()
	xml.Unmarshal(buf, v2)
}
