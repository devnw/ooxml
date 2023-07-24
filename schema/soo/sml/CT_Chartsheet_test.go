package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ChartsheetConstructor(t *testing.T) {
	v := sml.NewCT_Chartsheet()
	if v == nil {
		t.Errorf("sml.NewCT_Chartsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Chartsheet should validate: %s", err)
	}
}

func TestCT_ChartsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Chartsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Chartsheet()
	xml.Unmarshal(buf, v2)
}
