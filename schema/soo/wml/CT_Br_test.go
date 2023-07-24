package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BrConstructor(t *testing.T) {
	v := wml.NewCT_Br()
	if v == nil {
		t.Errorf("wml.NewCT_Br must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Br should validate: %s", err)
	}
}

func TestCT_BrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Br()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Br()
	xml.Unmarshal(buf, v2)
}
