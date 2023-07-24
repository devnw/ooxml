package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingShapeChoice1Constructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShapeChoice1()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingShapeChoice1 must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingShapeChoice1 should validate: %s", err)
	}
}

func TestWdCT_WordprocessingShapeChoice1MarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShapeChoice1()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingShapeChoice1()
	xml.Unmarshal(buf, v2)
}
