package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlGroupShapeConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGroupShape()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGroupShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGroupShape should validate: %s", err)
	}
}

func TestCT_GvmlGroupShapeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGroupShape()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGroupShape()
	xml.Unmarshal(buf, v2)
}
