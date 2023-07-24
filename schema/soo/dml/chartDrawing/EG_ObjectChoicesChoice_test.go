package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestEG_ObjectChoicesChoiceConstructor(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoicesChoice()
	if v == nil {
		t.Errorf("chartDrawing.NewEG_ObjectChoicesChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.EG_ObjectChoicesChoice should validate: %s", err)
	}
}

func TestEG_ObjectChoicesChoiceMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoicesChoice()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewEG_ObjectChoicesChoice()
	xml.Unmarshal(buf, v2)
}
