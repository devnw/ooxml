package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingShapeChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShapeChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingShapeChoice should validate: %s", err)
	}
}

func TestWdCT_WordprocessingShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingShapeChoice()
	xml.Unmarshal(buf, v2)
}
