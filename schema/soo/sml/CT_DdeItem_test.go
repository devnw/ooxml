package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DdeItemConstructor(t *testing.T) {
	v := sml.NewCT_DdeItem()
	if v == nil {
		t.Errorf("sml.NewCT_DdeItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DdeItem should validate: %s", err)
	}
}

func TestCT_DdeItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DdeItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DdeItem()
	xml.Unmarshal(buf, v2)
}
