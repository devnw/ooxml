package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_SharedItemsConstructor(t *testing.T) {
	v := sml.NewCT_SharedItems()
	if v == nil {
		t.Errorf("sml.NewCT_SharedItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_SharedItems should validate: %s", err)
	}
}

func TestCT_SharedItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_SharedItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_SharedItems()
	xml.Unmarshal(buf, v2)
}
