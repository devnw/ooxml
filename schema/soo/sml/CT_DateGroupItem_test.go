package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DateGroupItemConstructor(t *testing.T) {
	v := sml.NewCT_DateGroupItem()
	if v == nil {
		t.Errorf("sml.NewCT_DateGroupItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DateGroupItem should validate: %s", err)
	}
}

func TestCT_DateGroupItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DateGroupItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DateGroupItem()
	xml.Unmarshal(buf, v2)
}
