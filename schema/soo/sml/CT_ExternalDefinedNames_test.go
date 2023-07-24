package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalDefinedNamesConstructor(t *testing.T) {
	v := sml.NewCT_ExternalDefinedNames()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalDefinedNames must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalDefinedNames should validate: %s", err)
	}
}

func TestCT_ExternalDefinedNamesMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalDefinedNames()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalDefinedNames()
	xml.Unmarshal(buf, v2)
}
