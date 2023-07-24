package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CustomPropertyConstructor(t *testing.T) {
	v := sml.NewCT_CustomProperty()
	if v == nil {
		t.Errorf("sml.NewCT_CustomProperty must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CustomProperty should validate: %s", err)
	}
}

func TestCT_CustomPropertyMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CustomProperty()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CustomProperty()
	xml.Unmarshal(buf, v2)
}
