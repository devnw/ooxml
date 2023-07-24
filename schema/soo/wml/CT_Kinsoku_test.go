package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_KinsokuConstructor(t *testing.T) {
	v := wml.NewCT_Kinsoku()
	if v == nil {
		t.Errorf("wml.NewCT_Kinsoku must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Kinsoku should validate: %s", err)
	}
}

func TestCT_KinsokuMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Kinsoku()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Kinsoku()
	xml.Unmarshal(buf, v2)
}
