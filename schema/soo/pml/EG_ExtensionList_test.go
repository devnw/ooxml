package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestEG_ExtensionListConstructor(t *testing.T) {
	v := pml.NewEG_ExtensionList()
	if v == nil {
		t.Errorf("pml.NewEG_ExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.EG_ExtensionList should validate: %s", err)
	}
}

func TestEG_ExtensionListMarshalUnmarshal(t *testing.T) {
	v := pml.NewEG_ExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewEG_ExtensionList()
	xml.Unmarshal(buf, v2)
}
