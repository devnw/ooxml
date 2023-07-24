package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomSheetViewsConstructor(t *testing.T) {
	v := sml.NewCT_CustomSheetViews()
	if v == nil {
		t.Errorf("sml.NewCT_CustomSheetViews must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomSheetViews should validate: %s", err)
	}
}

func TestCT_CustomSheetViewsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomSheetViews()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomSheetViews()
	xml.Unmarshal(buf, v2)
}
