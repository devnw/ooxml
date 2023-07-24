package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumberingConstructor(t *testing.T) {
	v := wml.NewCT_Numbering()
	if v == nil {
		t.Errorf("wml.NewCT_Numbering must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Numbering should validate: %s", err)
	}
}

func TestCT_NumberingMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Numbering()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Numbering()
	xml.Unmarshal(buf, v2)
}
