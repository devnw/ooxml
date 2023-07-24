package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MeasureDimensionMapConstructor(t *testing.T) {
	v := sml.NewCT_MeasureDimensionMap()
	if v == nil {
		t.Errorf("sml.NewCT_MeasureDimensionMap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MeasureDimensionMap should validate: %s", err)
	}
}

func TestCT_MeasureDimensionMapMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MeasureDimensionMap()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MeasureDimensionMap()
	xml.Unmarshal(buf, v2)
}
