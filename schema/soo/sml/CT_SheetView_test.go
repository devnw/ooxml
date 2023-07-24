package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetViewConstructor(t *testing.T) {
	v := sml.NewCT_SheetView()
	if v == nil {
		t.Errorf("sml.NewCT_SheetView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetView should validate: %s", err)
	}
}

func TestCT_SheetViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetView()
	xml.Unmarshal(buf, v2)
}
