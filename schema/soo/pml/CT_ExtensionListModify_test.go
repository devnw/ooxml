package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_ExtensionListModifyConstructor(t *testing.T) {
	v := pml.NewCT_ExtensionListModify()
	if v == nil {
		t.Errorf("pml.NewCT_ExtensionListModify must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_ExtensionListModify should validate: %s", err)
	}
}

func TestCT_ExtensionListModifyMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_ExtensionListModify()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_ExtensionListModify()
	xml.Unmarshal(buf, v2)
}
