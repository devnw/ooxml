package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlTextShapeConstructor(t *testing.T) {
	v := dml.NewCT_GvmlTextShape()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlTextShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlTextShape should validate: %s", err)
	}
}

func TestCT_GvmlTextShapeMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlTextShape()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlTextShape()
	xml.Unmarshal(buf, v2)
}
