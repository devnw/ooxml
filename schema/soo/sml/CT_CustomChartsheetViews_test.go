package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomChartsheetViewsConstructor(t *testing.T) {
	v := sml.NewCT_CustomChartsheetViews()
	if v == nil {
		t.Errorf("sml.NewCT_CustomChartsheetViews must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomChartsheetViews should validate: %s", err)
	}
}

func TestCT_CustomChartsheetViewsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomChartsheetViews()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomChartsheetViews()
	xml.Unmarshal(buf, v2)
}
