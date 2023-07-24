package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_HslColorConstructor(t *testing.T) {
	v := dml.NewCT_HslColor()
	if v == nil {
		t.Errorf("dml.NewCT_HslColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_HslColor should validate: %s", err)
	}
}

func TestCT_HslColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_HslColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_HslColor()
	xml.Unmarshal(buf, v2)
}
