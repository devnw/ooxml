package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideMasterIdListEntryConstructor(t *testing.T) {
	v := pml.NewCT_SlideMasterIdListEntry()
	if v == nil {
		t.Errorf("pml.NewCT_SlideMasterIdListEntry must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideMasterIdListEntry should validate: %s", err)
	}
}

func TestCT_SlideMasterIdListEntryMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideMasterIdListEntry()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideMasterIdListEntry()
	xml.Unmarshal(buf, v2)
}
