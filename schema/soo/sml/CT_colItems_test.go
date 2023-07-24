package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_colItemsConstructor(t *testing.T) {
	v := sml.NewCT_colItems()
	if v == nil {
		t.Errorf("sml.NewCT_colItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_colItems should validate: %s", err)
	}
}

func TestCT_colItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_colItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_colItems()
	xml.Unmarshal(buf, v2)
}
