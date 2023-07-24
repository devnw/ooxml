package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ExtensionListConstructor(t *testing.T) {
	v := pml.NewCT_ExtensionList()
	if v == nil {
		t.Errorf("pml.NewCT_ExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ExtensionList should validate: %s", err)
	}
}

func TestCT_ExtensionListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ExtensionList()
	xml.Unmarshal(buf, v2)
}
