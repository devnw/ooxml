package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_BackgroundPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_BackgroundProperties()
	if v == nil {
		t.Errorf("pml.NewCT_BackgroundProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_BackgroundProperties should validate: %s", err)
	}
}

func TestCT_BackgroundPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_BackgroundProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_BackgroundProperties()
	xml.Unmarshal(buf, v2)
}
