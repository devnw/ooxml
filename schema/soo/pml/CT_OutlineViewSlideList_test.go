package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_OutlineViewSlideListConstructor(t *testing.T) {
	v := pml.NewCT_OutlineViewSlideList()
	if v == nil {
		t.Errorf("pml.NewCT_OutlineViewSlideList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_OutlineViewSlideList should validate: %s", err)
	}
}

func TestCT_OutlineViewSlideListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_OutlineViewSlideList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_OutlineViewSlideList()
	xml.Unmarshal(buf, v2)
}
