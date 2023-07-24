package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_HandoutMasterIdListConstructor(t *testing.T) {
	v := pml.NewCT_HandoutMasterIdList()
	if v == nil {
		t.Errorf("pml.NewCT_HandoutMasterIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_HandoutMasterIdList should validate: %s", err)
	}
}

func TestCT_HandoutMasterIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_HandoutMasterIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_HandoutMasterIdList()
	xml.Unmarshal(buf, v2)
}
