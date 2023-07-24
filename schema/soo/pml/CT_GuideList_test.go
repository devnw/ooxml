package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_GuideListConstructor(t *testing.T) {
	v := pml.NewCT_GuideList()
	if v == nil {
		t.Errorf("pml.NewCT_GuideList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_GuideList should validate: %s", err)
	}
}

func TestCT_GuideListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_GuideList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_GuideList()
	xml.Unmarshal(buf, v2)
}
