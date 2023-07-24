package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_BuildListConstructor(t *testing.T) {
	v := pml.NewCT_BuildList()
	if v == nil {
		t.Errorf("pml.NewCT_BuildList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_BuildList should validate: %s", err)
	}
}

func TestCT_BuildListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_BuildList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_BuildList()
	xml.Unmarshal(buf, v2)
}
