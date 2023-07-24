package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_GroupItemsConstructor(t *testing.T) {
	v := sml.NewCT_GroupItems()
	if v == nil {
		t.Errorf("sml.NewCT_GroupItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_GroupItems should validate: %s", err)
	}
}

func TestCT_GroupItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_GroupItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_GroupItems()
	xml.Unmarshal(buf, v2)
}
