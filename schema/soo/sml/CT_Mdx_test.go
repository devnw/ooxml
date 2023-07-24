package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_MdxConstructor(t *testing.T) {
	v := sml.NewCT_Mdx()
	if v == nil {
		t.Errorf("sml.NewCT_Mdx must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Mdx should validate: %s", err)
	}
}

func TestCT_MdxMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Mdx()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Mdx()
	xml.Unmarshal(buf, v2)
}
