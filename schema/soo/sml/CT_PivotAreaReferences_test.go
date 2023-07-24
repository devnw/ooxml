package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotAreaReferencesConstructor(t *testing.T) {
	v := sml.NewCT_PivotAreaReferences()
	if v == nil {
		t.Errorf("sml.NewCT_PivotAreaReferences must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotAreaReferences should validate: %s", err)
	}
}

func TestCT_PivotAreaReferencesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotAreaReferences()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotAreaReferences()
	xml.Unmarshal(buf, v2)
}
