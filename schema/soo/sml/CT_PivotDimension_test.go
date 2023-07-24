package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotDimensionConstructor(t *testing.T) {
	v := sml.NewCT_PivotDimension()
	if v == nil {
		t.Errorf("sml.NewCT_PivotDimension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotDimension should validate: %s", err)
	}
}

func TestCT_PivotDimensionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotDimension()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotDimension()
	xml.Unmarshal(buf, v2)
}
