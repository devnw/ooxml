package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualDrawingPropsConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualDrawingProps()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualDrawingProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualDrawingProps should validate: %s", err)
	}
}

func TestCT_NonVisualDrawingPropsMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualDrawingProps()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualDrawingProps()
	xml.Unmarshal(buf, v2)
}
