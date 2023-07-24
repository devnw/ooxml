package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ColFieldsConstructor(t *testing.T) {
	v := sml.NewCT_ColFields()
	if v == nil {
		t.Errorf("sml.NewCT_ColFields must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ColFields should validate: %s", err)
	}
}

func TestCT_ColFieldsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ColFields()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ColFields()
	xml.Unmarshal(buf, v2)
}
