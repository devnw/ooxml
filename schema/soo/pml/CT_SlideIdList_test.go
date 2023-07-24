package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideIdListConstructor(t *testing.T) {
	v := pml.NewCT_SlideIdList()
	if v == nil {
		t.Errorf("pml.NewCT_SlideIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideIdList should validate: %s", err)
	}
}

func TestCT_SlideIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideIdList()
	xml.Unmarshal(buf, v2)
}
