package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_DefinedNamesConstructor(t *testing.T) {
	v := sml.NewCT_DefinedNames()
	if v == nil {
		t.Errorf("sml.NewCT_DefinedNames must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_DefinedNames should validate: %s", err)
	}
}

func TestCT_DefinedNamesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_DefinedNames()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_DefinedNames()
	xml.Unmarshal(buf, v2)
}
