package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingShapeConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShape()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingShape must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingShape should validate: %s", err)
	}
}

func TestWdCT_WordprocessingShapeMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingShape()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingShape()
	xml.Unmarshal(buf, v2)
}
