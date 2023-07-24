package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MarkupConstructor(t *testing.T) {
	v := wml.NewCT_Markup()
	if v == nil {
		t.Errorf("wml.NewCT_Markup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Markup should validate: %s", err)
	}
}

func TestCT_MarkupMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Markup()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Markup()
	xml.Unmarshal(buf, v2)
}
