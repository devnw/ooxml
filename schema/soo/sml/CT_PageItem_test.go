package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PageItemConstructor(t *testing.T) {
	v := sml.NewCT_PageItem()
	if v == nil {
		t.Errorf("sml.NewCT_PageItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageItem should validate: %s", err)
	}
}

func TestCT_PageItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageItem()
	xml.Unmarshal(buf, v2)
}
