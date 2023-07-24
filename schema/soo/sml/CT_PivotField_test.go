package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotFieldConstructor(t *testing.T) {
	v := sml.NewCT_PivotField()
	if v == nil {
		t.Errorf("sml.NewCT_PivotField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotField should validate: %s", err)
	}
}

func TestCT_PivotFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotField()
	xml.Unmarshal(buf, v2)
}
