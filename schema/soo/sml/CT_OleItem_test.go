package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_OleItemConstructor(t *testing.T) {
	v := sml.NewCT_OleItem()
	if v == nil {
		t.Errorf("sml.NewCT_OleItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleItem should validate: %s", err)
	}
}

func TestCT_OleItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleItem()
	xml.Unmarshal(buf, v2)
}
