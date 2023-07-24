package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TopPageBorderConstructor(t *testing.T) {
	v := wml.NewCT_TopPageBorder()
	if v == nil {
		t.Errorf("wml.NewCT_TopPageBorder must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TopPageBorder should validate: %s", err)
	}
}

func TestCT_TopPageBorderMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TopPageBorder()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TopPageBorder()
	xml.Unmarshal(buf, v2)
}
