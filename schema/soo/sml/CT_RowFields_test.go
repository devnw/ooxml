package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RowFieldsConstructor(t *testing.T) {
	v := sml.NewCT_RowFields()
	if v == nil {
		t.Errorf("sml.NewCT_RowFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RowFields should validate: %s", err)
	}
}

func TestCT_RowFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RowFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RowFields()
	xml.Unmarshal(buf, v2)
}
