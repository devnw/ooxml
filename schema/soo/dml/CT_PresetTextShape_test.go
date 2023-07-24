package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_PresetTextShapeConstructor(t *testing.T) {
	v := dml.NewCT_PresetTextShape()
	if v == nil {
		t.Errorf("dml.NewCT_PresetTextShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_PresetTextShape should validate: %s", err)
	}
}

func TestCT_PresetTextShapeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_PresetTextShape()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_PresetTextShape()
	xml.Unmarshal(buf, v2)
}
