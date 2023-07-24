package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ObjectAnchorConstructor(t *testing.T) {
	v := sml.NewCT_ObjectAnchor()
	if v == nil {
		t.Errorf("sml.NewCT_ObjectAnchor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ObjectAnchor should validate: %s", err)
	}
}

func TestCT_ObjectAnchorMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ObjectAnchor()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ObjectAnchor()
	xml.Unmarshal(buf, v2)
}
