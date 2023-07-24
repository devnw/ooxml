package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlGroupShapeNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGroupShapeNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGroupShapeNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlGroupShapeNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGroupShapeNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGroupShapeNonVisual()
	xml.Unmarshal(buf, v2)
}
