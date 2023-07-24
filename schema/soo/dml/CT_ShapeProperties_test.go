package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ShapePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_ShapeProperties()
	if v == nil {
		t.Errorf("dml.NewCT_ShapeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ShapeProperties should validate: %s", err)
	}
}

func TestCT_ShapePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ShapeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ShapeProperties()
	xml.Unmarshal(buf, v2)
}
