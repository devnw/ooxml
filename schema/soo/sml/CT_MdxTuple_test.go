package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MdxTupleConstructor(t *testing.T) {
	v := sml.NewCT_MdxTuple()
	if v == nil {
		t.Errorf("sml.NewCT_MdxTuple must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_MdxTuple should validate: %s", err)
	}
}

func TestCT_MdxTupleMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_MdxTuple()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_MdxTuple()
	xml.Unmarshal(buf, v2)
}
