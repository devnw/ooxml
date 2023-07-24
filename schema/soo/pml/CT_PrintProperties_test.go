package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_PrintPropertiesConstructor(t *testing.T) {
	v := pml.NewCT_PrintProperties()
	if v == nil {
		t.Errorf("pml.NewCT_PrintProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_PrintProperties should validate: %s", err)
	}
}

func TestCT_PrintPropertiesMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_PrintProperties()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_PrintProperties()
	xml.Unmarshal(buf, v2)
}
