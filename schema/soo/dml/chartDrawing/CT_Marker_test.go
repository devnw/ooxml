package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_MarkerConstructor(t *testing.T) {
	v := chartDrawing.NewCT_Marker()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_Marker must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_Marker should validate: %s", err)
	}
}

func TestCT_MarkerMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_Marker()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_Marker()
	xml.Unmarshal(buf, v2)
}
