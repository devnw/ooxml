package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExtensionListConstructor(t *testing.T) {
	v := sml.NewCT_ExtensionList()
	if v == nil {
		t.Errorf("sml.NewCT_ExtensionList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExtensionList should validate: %s", err)
	}
}

func TestCT_ExtensionListMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExtensionList()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExtensionList()
	xml.Unmarshal(buf, v2)
}
