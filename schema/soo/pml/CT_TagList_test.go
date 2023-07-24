package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_TagListConstructor(t *testing.T) {
	v := pml.NewCT_TagList()
	if v == nil {
		t.Errorf("pml.NewCT_TagList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TagList should validate: %s", err)
	}
}

func TestCT_TagListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TagList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TagList()
	xml.Unmarshal(buf, v2)
}
