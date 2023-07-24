package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestEG_AnchorConstructor(t *testing.T) {
	v := chartDrawing.NewEG_Anchor()
	if v == nil {
		t.Errorf("chartDrawing.NewEG_Anchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.EG_Anchor should validate: %s", err)
	}
}

func TestEG_AnchorMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewEG_Anchor()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewEG_Anchor()
	xml.Unmarshal(buf, v2)
}
