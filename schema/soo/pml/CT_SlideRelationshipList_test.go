package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideRelationshipListConstructor(t *testing.T) {
	v := pml.NewCT_SlideRelationshipList()
	if v == nil {
		t.Errorf("pml.NewCT_SlideRelationshipList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideRelationshipList should validate: %s", err)
	}
}

func TestCT_SlideRelationshipListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideRelationshipList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideRelationshipList()
	xml.Unmarshal(buf, v2)
}
