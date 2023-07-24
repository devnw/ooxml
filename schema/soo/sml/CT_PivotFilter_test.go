package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotFilterConstructor(t *testing.T) {
	v := sml.NewCT_PivotFilter()
	if v == nil {
		t.Errorf("sml.NewCT_PivotFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotFilter should validate: %s", err)
	}
}

func TestCT_PivotFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotFilter()
	xml.Unmarshal(buf, v2)
}
