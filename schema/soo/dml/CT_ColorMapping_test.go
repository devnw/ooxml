package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ColorMappingConstructor(t *testing.T) {
	v := dml.NewCT_ColorMapping()
	if v == nil {
		t.Errorf("dml.NewCT_ColorMapping must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ColorMapping should validate: %s", err)
	}
}

func TestCT_ColorMappingMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ColorMapping()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ColorMapping()
	xml.Unmarshal(buf, v2)
}
