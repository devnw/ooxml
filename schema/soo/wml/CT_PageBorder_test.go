package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_PageBorderConstructor(t *testing.T) {
	v := wml.NewCT_PageBorder()
	if v == nil {
		t.Errorf("wml.NewCT_PageBorder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageBorder should validate: %s", err)
	}
}

func TestCT_PageBorderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageBorder()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageBorder()
	xml.Unmarshal(buf, v2)
}
