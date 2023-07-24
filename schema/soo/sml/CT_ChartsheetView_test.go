package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ChartsheetViewConstructor(t *testing.T) {
	v := sml.NewCT_ChartsheetView()
	if v == nil {
		t.Errorf("sml.NewCT_ChartsheetView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ChartsheetView should validate: %s", err)
	}
}

func TestCT_ChartsheetViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ChartsheetView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ChartsheetView()
	xml.Unmarshal(buf, v2)
}
