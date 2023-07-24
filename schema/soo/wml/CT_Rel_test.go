package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RelConstructor(t *testing.T) {
	v := wml.NewCT_Rel()
	if v == nil {
		t.Errorf("wml.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
