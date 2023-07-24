package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCmLstConstructor(t *testing.T) {
	v := pml.NewCmLst()
	if v == nil {
		t.Errorf("pml.NewCmLst must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CmLst should validate: %s", err)
	}
}

func TestCmLstMarshalUnmarshal(t *testing.T) {
	v := pml.NewCmLst()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCmLst()
	xml.Unmarshal(buf, v2)
}
