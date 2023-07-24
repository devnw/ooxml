package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_CustomColorConstructor(t *testing.T) {
	v := dml.NewCT_CustomColor()
	if v == nil {
		t.Errorf("dml.NewCT_CustomColor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_CustomColor should validate: %s", err)
	}
}

func TestCT_CustomColorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_CustomColor()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_CustomColor()
	xml.Unmarshal(buf, v2)
}
