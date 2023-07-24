package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ChartFormatConstructor(t *testing.T) {
	v := sml.NewCT_ChartFormat()
	if v == nil {
		t.Errorf("sml.NewCT_ChartFormat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ChartFormat should validate: %s", err)
	}
}

func TestCT_ChartFormatMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ChartFormat()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ChartFormat()
	xml.Unmarshal(buf, v2)
}
