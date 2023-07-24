package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlShapeConstructor(t *testing.T) {
	v := dml.NewCT_GvmlShape()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlShape should validate: %s", err)
	}
}

func TestCT_GvmlShapeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlShape()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlShape()
	xml.Unmarshal(buf, v2)
}
