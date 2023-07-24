package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotAreaReferenceConstructor(t *testing.T) {
	v := sml.NewCT_PivotAreaReference()
	if v == nil {
		t.Errorf("sml.NewCT_PivotAreaReference must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotAreaReference should validate: %s", err)
	}
}

func TestCT_PivotAreaReferenceMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotAreaReference()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotAreaReference()
	xml.Unmarshal(buf, v2)
}
