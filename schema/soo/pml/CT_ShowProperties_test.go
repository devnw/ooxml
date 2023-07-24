package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ShowPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_ShowProperties()
	if v == nil {
		t.Errorf("pml.NewCT_ShowProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ShowProperties should validate: %s", err)
	}
}

func TestCT_ShowPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ShowProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ShowProperties()
	xml.Unmarshal(buf, v2)
}
