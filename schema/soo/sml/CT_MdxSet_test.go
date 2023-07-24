package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MdxSetConstructor(t *testing.T) {
	v := sml.NewCT_MdxSet()
	if v == nil {
		t.Errorf("sml.NewCT_MdxSet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MdxSet should validate: %s", err)
	}
}

func TestCT_MdxSetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MdxSet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MdxSet()
	xml.Unmarshal(buf, v2)
}
