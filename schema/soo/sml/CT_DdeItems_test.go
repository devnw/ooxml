package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DdeItemsConstructor(t *testing.T) {
	v := sml.NewCT_DdeItems()
	if v == nil {
		t.Errorf("sml.NewCT_DdeItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DdeItems should validate: %s", err)
	}
}

func TestCT_DdeItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DdeItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DdeItems()
	xml.Unmarshal(buf, v2)
}
