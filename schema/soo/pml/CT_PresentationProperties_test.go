package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PresentationPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_PresentationProperties()
	if v == nil {
		t.Errorf("pml.NewCT_PresentationProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_PresentationProperties should validate: %s", err)
	}
}

func TestCT_PresentationPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_PresentationProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_PresentationProperties()
	xml.Unmarshal(buf, v2)
}
