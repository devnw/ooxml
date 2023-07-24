package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NormalViewPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_NormalViewProperties()
	if v == nil {
		t.Errorf("pml.NewCT_NormalViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NormalViewProperties should validate: %s", err)
	}
}

func TestCT_NormalViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NormalViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NormalViewProperties()
	xml.Unmarshal(buf, v2)
}
