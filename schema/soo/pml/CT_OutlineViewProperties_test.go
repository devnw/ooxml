package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OutlineViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_OutlineViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_OutlineViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OutlineViewProperties should validate: %s", err)
	}
}

func TestCT_OutlineViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OutlineViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OutlineViewProperties()
	xml.Unmarshal(buf, v2)
}
