package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_SlideViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_SlideViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideViewProperties should validate: %s", err)
	}
}

func TestCT_SlideViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideViewProperties()
	xml.Unmarshal(buf, v2)
}
