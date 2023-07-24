package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideMasterIdListConstructor(t *testing.T) {
	v := pml.NewCT_SlideMasterIdList()
	if v == nil {
		t.Errorf("pml.NewCT_SlideMasterIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideMasterIdList should validate: %s", err)
	}
}

func TestCT_SlideMasterIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideMasterIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideMasterIdList()
	xml.Unmarshal(buf, v2)
}
