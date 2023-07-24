package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualGraphicFramePropertiesConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualGraphicFrameProperties()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualGraphicFrameProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualGraphicFrameProperties should validate: %s", err)
	}
}

func TestCT_NonVisualGraphicFramePropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualGraphicFrameProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualGraphicFrameProperties()
	xml.Unmarshal(buf, v2)
}
