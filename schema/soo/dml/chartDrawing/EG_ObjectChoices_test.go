package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

func TestEG_ObjectChoicesConstructor(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoices()
	if v == nil {
		t.Errorf("chartDrawing.NewEG_ObjectChoices must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.EG_ObjectChoices should validate: %s", err)
	}
}

func TestEG_ObjectChoicesMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewEG_ObjectChoices()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewEG_ObjectChoices()
	xml.Unmarshal(buf, v2)
}
