package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommonViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_CommonViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_CommonViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommonViewProperties should validate: %s", err)
	}
}

func TestCT_CommonViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommonViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommonViewProperties()
	xml.Unmarshal(buf, v2)
}
