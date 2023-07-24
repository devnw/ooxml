package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_rowItemsConstructor(t *testing.T) {
	v := sml.NewCT_rowItems()
	if v == nil {
		t.Errorf("sml.NewCT_rowItems must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_rowItems should validate: %s", err)
	}
}

func TestCT_rowItemsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_rowItems()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_rowItems()
	xml.Unmarshal(buf, v2)
}
