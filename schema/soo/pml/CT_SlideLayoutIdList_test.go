package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_SlideLayoutIdListConstructor(t *testing.T) {
	v := pml.NewCT_SlideLayoutIdList()
	if v == nil {
		t.Errorf("pml.NewCT_SlideLayoutIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_SlideLayoutIdList should validate: %s", err)
	}
}

func TestCT_SlideLayoutIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_SlideLayoutIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_SlideLayoutIdList()
	xml.Unmarshal(buf, v2)
}
