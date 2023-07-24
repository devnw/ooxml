package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_IconFilterConstructor(t *testing.T) {
	v := sml.NewCT_IconFilter()
	if v == nil {
		t.Errorf("sml.NewCT_IconFilter must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_IconFilter should validate: %s", err)
	}
}

func TestCT_IconFilterMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_IconFilter()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_IconFilter()
	xml.Unmarshal(buf, v2)
}
