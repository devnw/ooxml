package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CalculatedItemConstructor(t *testing.T) {
	v := sml.NewCT_CalculatedItem()
	if v == nil {
		t.Errorf("sml.NewCT_CalculatedItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalculatedItem should validate: %s", err)
	}
}

func TestCT_CalculatedItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalculatedItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalculatedItem()
	xml.Unmarshal(buf, v2)
}
