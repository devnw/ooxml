package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DefinedNameConstructor(t *testing.T) {
	v := sml.NewCT_DefinedName()
	if v == nil {
		t.Errorf("sml.NewCT_DefinedName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DefinedName should validate: %s", err)
	}
}

func TestCT_DefinedNameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DefinedName()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DefinedName()
	xml.Unmarshal(buf, v2)
}
