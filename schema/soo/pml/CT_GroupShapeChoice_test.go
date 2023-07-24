package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GroupShapeChoiceConstructor(t *testing.T) {
	v := pml.NewCT_GroupShapeChoice()
	if v == nil {
		t.Errorf("pml.NewCT_GroupShapeChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GroupShapeChoice should validate: %s", err)
	}
}

func TestCT_GroupShapeChoiceMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GroupShapeChoice()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GroupShapeChoice()
	xml.Unmarshal(buf, v2)
}
