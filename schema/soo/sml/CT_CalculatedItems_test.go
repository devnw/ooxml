package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalculatedItemsConstructor(t *testing.T) {
	v := sml.NewCT_CalculatedItems()
	if v == nil {
		t.Errorf("sml.NewCT_CalculatedItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalculatedItems should validate: %s", err)
	}
}

func TestCT_CalculatedItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalculatedItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalculatedItems()
	xml.Unmarshal(buf, v2)
}
