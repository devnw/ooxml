package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestHandoutMasterConstructor(t *testing.T) {
	v := pml.NewHandoutMaster()
	if v == nil {
		t.Errorf("pml.NewHandoutMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.HandoutMaster should validate: %s", err)
	}
}

func TestHandoutMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewHandoutMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewHandoutMaster()
	xml.Unmarshal(buf, v2)
}
