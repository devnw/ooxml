package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryTableFieldsConstructor(t *testing.T) {
	v := sml.NewCT_QueryTableFields()
	if v == nil {
		t.Errorf("sml.NewCT_QueryTableFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_QueryTableFields should validate: %s", err)
	}
}

func TestCT_QueryTableFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_QueryTableFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_QueryTableFields()
	xml.Unmarshal(buf, v2)
}
