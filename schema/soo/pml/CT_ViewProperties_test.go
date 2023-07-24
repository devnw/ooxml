package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_ViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_ViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ViewProperties should validate: %s", err)
	}
}

func TestCT_ViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ViewProperties()
	xml.Unmarshal(buf, v2)
}
