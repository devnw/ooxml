package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_HandoutMasterConstructor(t *testing.T) {
	v := pml.NewCT_HandoutMaster()
	if v == nil {
		t.Errorf("pml.NewCT_HandoutMaster must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_HandoutMaster should validate: %s", err)
	}
}

func TestCT_HandoutMasterMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_HandoutMaster()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_HandoutMaster()
	xml.Unmarshal(buf, v2)
}
