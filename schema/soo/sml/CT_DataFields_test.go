package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DataFieldsConstructor(t *testing.T) {
	v := sml.NewCT_DataFields()
	if v == nil {
		t.Errorf("sml.NewCT_DataFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DataFields should validate: %s", err)
	}
}

func TestCT_DataFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DataFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DataFields()
	xml.Unmarshal(buf, v2)
}
