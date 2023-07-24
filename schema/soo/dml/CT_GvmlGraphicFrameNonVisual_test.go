package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlGraphicFrameNonVisualConstructor(t *testing.T) {
	v := dml.NewCT_GvmlGraphicFrameNonVisual()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlGraphicFrameNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlGraphicFrameNonVisual should validate: %s", err)
	}
}

func TestCT_GvmlGraphicFrameNonVisualMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlGraphicFrameNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlGraphicFrameNonVisual()
	xml.Unmarshal(buf, v2)
}
