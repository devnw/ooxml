package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GroupShapePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_GroupShapeProperties()
	if v == nil {
		t.Errorf("dml.NewCT_GroupShapeProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GroupShapeProperties should validate: %s", err)
	}
}

func TestCT_GroupShapePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GroupShapeProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GroupShapeProperties()
	xml.Unmarshal(buf, v2)
}
