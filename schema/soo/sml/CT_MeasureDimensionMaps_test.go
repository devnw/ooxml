package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MeasureDimensionMapsConstructor(t *testing.T) {
	v := sml.NewCT_MeasureDimensionMaps()
	if v == nil {
		t.Errorf("sml.NewCT_MeasureDimensionMaps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MeasureDimensionMaps should validate: %s", err)
	}
}

func TestCT_MeasureDimensionMapsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MeasureDimensionMaps()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MeasureDimensionMaps()
	xml.Unmarshal(buf, v2)
}
