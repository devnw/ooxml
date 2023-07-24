package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomSheetViewConstructor(t *testing.T) {
	v := sml.NewCT_CustomSheetView()
	if v == nil {
		t.Errorf("sml.NewCT_CustomSheetView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomSheetView should validate: %s", err)
	}
}

func TestCT_CustomSheetViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomSheetView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomSheetView()
	xml.Unmarshal(buf, v2)
}
