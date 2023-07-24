package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_GroupShapeConstructor(t *testing.T) {
	v := chartDrawing.NewCT_GroupShape()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_GroupShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_GroupShape should validate: %s", err)
	}
}

func TestCT_GroupShapeMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_GroupShape()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_GroupShape()
	xml.Unmarshal(buf, v2)
}
