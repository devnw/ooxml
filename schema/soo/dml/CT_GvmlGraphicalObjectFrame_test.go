package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlGraphicalObjectFrameConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGraphicalObjectFrame()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGraphicalObjectFrame must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGraphicalObjectFrame should validate: %s", err)
	}
}

func TestCT_GvmlGraphicalObjectFrameMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGraphicalObjectFrame()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGraphicalObjectFrame()
	xml.Unmarshal(buf, v2)
}
