package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_WebPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_WebProperties()
	if v == nil {
		t.Errorf("pml.NewCT_WebProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_WebProperties should validate: %s", err)
	}
}

func TestCT_WebPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_WebProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_WebProperties()
	xml.Unmarshal(buf, v2)
}
