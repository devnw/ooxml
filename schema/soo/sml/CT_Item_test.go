package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ItemConstructor(t *testing.T) {
	v := sml.NewCT_Item()
	if v == nil {
		t.Errorf("sml.NewCT_Item must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Item should validate: %s", err)
	}
}

func TestCT_ItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Item()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Item()
	xml.Unmarshal(buf, v2)
}
