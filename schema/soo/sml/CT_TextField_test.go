package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TextFieldConstructor(t *testing.T) {
	v := sml.NewCT_TextField()
	if v == nil {
		t.Errorf("sml.NewCT_TextField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TextField should validate: %s", err)
	}
}

func TestCT_TextFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TextField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TextField()
	xml.Unmarshal(buf, v2)
}
