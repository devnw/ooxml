package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OleObjectLinkConstructor(t *testing.T) {
	v := pml.NewCT_OleObjectLink()
	if v == nil {
		t.Errorf("pml.NewCT_OleObjectLink must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OleObjectLink should validate: %s", err)
	}
}

func TestCT_OleObjectLinkMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OleObjectLink()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OleObjectLink()
	xml.Unmarshal(buf, v2)
}
