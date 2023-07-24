package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideRelationshipListEntryConstructor(t *testing.T) {
	v := pml.NewCT_SlideRelationshipListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_SlideRelationshipListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideRelationshipListEntry should validate: %s", err)
	}
}

func TestCT_SlideRelationshipListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideRelationshipListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideRelationshipListEntry()
	xml.Unmarshal(buf, v2)
}
