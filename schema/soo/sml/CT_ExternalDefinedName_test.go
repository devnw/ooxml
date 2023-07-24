package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_ExternalDefinedNameConstructor(t *testing.T) {
	v := sml.NewCT_ExternalDefinedName()
	if v == nil {
		t.Errorf("sml.NewCT_ExternalDefinedName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_ExternalDefinedName should validate: %s", err)
	}
}

func TestCT_ExternalDefinedNameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_ExternalDefinedName()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_ExternalDefinedName()
	xml.Unmarshal(buf, v2)
}
