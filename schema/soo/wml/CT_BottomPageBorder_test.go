package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BottomPageBorderConstructor(t *testing.T) {
	v := wml.NewCT_BottomPageBorder()
	if v == nil {
		t.Errorf("wml.NewCT_BottomPageBorder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_BottomPageBorder should validate: %s", err)
	}
}

func TestCT_BottomPageBorderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_BottomPageBorder()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_BottomPageBorder()
	xml.Unmarshal(buf, v2)
}
