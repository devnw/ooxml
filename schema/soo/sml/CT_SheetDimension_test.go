package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SheetDimensionConstructor(t *testing.T) {
	v := sml.NewCT_SheetDimension()
	if v == nil {
		t.Errorf("sml.NewCT_SheetDimension must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SheetDimension should validate: %s", err)
	}
}

func TestCT_SheetDimensionMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SheetDimension()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SheetDimension()
	xml.Unmarshal(buf, v2)
}
