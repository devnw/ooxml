package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomChartsheetViewConstructor(t *testing.T) {
	v := sml.NewCT_CustomChartsheetView()
	if v == nil {
		t.Errorf("sml.NewCT_CustomChartsheetView must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomChartsheetView should validate: %s", err)
	}
}

func TestCT_CustomChartsheetViewMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomChartsheetView()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomChartsheetView()
	xml.Unmarshal(buf, v2)
}
