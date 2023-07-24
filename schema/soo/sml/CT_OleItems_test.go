package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_OleItemsConstructor(t *testing.T) {
	v := sml.NewCT_OleItems()
	if v == nil {
		t.Errorf("sml.NewCT_OleItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_OleItems should validate: %s", err)
	}
}

func TestCT_OleItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_OleItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_OleItems()
	xml.Unmarshal(buf, v2)
}
