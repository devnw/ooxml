package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCmAuthorLstConstructor(t *testing.T) {
	v := pml.NewCmAuthorLst()
	if v == nil {
		t.Errorf("pml.NewCmAuthorLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CmAuthorLst should validate: %s", err)
	}
}

func TestCmAuthorLstMarshalUnmarshal(t *testing.T) {
	v := pml.NewCmAuthorLst()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCmAuthorLst()
	xml.Unmarshal(buf, v2)
}
