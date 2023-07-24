package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PageFieldsConstructor(t *testing.T) {
	v := sml.NewCT_PageFields()
	if v == nil {
		t.Errorf("sml.NewCT_PageFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PageFields should validate: %s", err)
	}
}

func TestCT_PageFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PageFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PageFields()
	xml.Unmarshal(buf, v2)
}
