package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_HandoutMasterIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_HandoutMasterIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_HandoutMasterIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_HandoutMasterIdListEntry should validate: %s", err)
	}
}

func TestCT_HandoutMasterIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_HandoutMasterIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_HandoutMasterIdListEntry()
	xml.Unmarshal(buf, v2)
}
