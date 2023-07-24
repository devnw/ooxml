package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestCT_AbsSizeAnchorConstructor(t *testing.T) {
	v := chartDrawing.NewCT_AbsSizeAnchor()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_AbsSizeAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_AbsSizeAnchor should validate: %s", err)
	}
}

func TestCT_AbsSizeAnchorMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_AbsSizeAnchor()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_AbsSizeAnchor()
	xml.Unmarshal(buf, v2)
}
