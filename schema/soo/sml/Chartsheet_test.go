package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestChartsheetConstructor(t *testing.T) {
	v := sml.NewChartsheet()
	if v == nil {
		t.Errorf("sml.NewChartsheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Chartsheet should validate: %s", err)
	}
}

func TestChartsheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewChartsheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewChartsheet()
	xml.Unmarshal(buf, v2)
}
