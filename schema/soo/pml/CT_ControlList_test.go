package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ControlListConstructor(t *testing.T) {
	v := pml.NewCT_ControlList()
	if v == nil {
		t.Errorf("pml.NewCT_ControlList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ControlList should validate: %s", err)
	}
}

func TestCT_ControlListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ControlList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ControlList()
	xml.Unmarshal(buf, v2)
}
