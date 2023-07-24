package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DivConstructor(t *testing.T) {
	v := wml.NewCT_Div()
	if v == nil {
		t.Errorf("wml.NewCT_Div must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Div should validate: %s", err)
	}
}

func TestCT_DivMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Div()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Div()
	xml.Unmarshal(buf, v2)
}
