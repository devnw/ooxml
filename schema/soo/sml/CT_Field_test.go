package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FieldConstructor(t *testing.T) {
	v := sml.NewCT_Field()
	if v == nil {
		t.Errorf("sml.NewCT_Field must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Field should validate: %s", err)
	}
}

func TestCT_FieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Field()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Field()
	xml.Unmarshal(buf, v2)
}
