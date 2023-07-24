package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ItemsConstructor(t *testing.T) {
	v := sml.NewCT_Items()
	if v == nil {
		t.Errorf("sml.NewCT_Items must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Items should validate: %s", err)
	}
}

func TestCT_ItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Items()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Items()
	xml.Unmarshal(buf, v2)
}
