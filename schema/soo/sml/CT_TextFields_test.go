package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_TextFieldsConstructor(t *testing.T) {
	v := sml.NewCT_TextFields()
	if v == nil {
		t.Errorf("sml.NewCT_TextFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_TextFields should validate: %s", err)
	}
}

func TestCT_TextFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_TextFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_TextFields()
	xml.Unmarshal(buf, v2)
}
