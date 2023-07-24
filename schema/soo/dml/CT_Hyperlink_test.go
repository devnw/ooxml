package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_HyperlinkConstructor(t *testing.T) {
	v := dml.NewCT_Hyperlink()
	if v == nil {
		t.Errorf("dml.NewCT_Hyperlink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Hyperlink should validate: %s", err)
	}
}

func TestCT_HyperlinkMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Hyperlink()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Hyperlink()
	xml.Unmarshal(buf, v2)
}
