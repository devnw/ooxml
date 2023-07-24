package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PageBordersConstructor(t *testing.T) {
	v := wml.NewCT_PageBorders()
	if v == nil {
		t.Errorf("wml.NewCT_PageBorders must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageBorders should validate: %s", err)
	}
}

func TestCT_PageBordersMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageBorders()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageBorders()
	xml.Unmarshal(buf, v2)
}
