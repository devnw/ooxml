package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ChartsheetProtectionConstructor(t *testing.T) {
	v := sml.NewCT_ChartsheetProtection()
	if v == nil {
		t.Errorf("sml.NewCT_ChartsheetProtection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ChartsheetProtection should validate: %s", err)
	}
}

func TestCT_ChartsheetProtectionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ChartsheetProtection()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ChartsheetProtection()
	xml.Unmarshal(buf, v2)
}
