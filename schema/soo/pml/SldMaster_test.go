package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestSldMasterConstructor(t *testing.T) {
	v := pml.NewSldMaster()
	if v == nil {
		t.Errorf("pml.NewSldMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.SldMaster should validate: %s", err)
	}
}

func TestSldMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewSldMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewSldMaster()
	xml.Unmarshal(buf, v2)
}
