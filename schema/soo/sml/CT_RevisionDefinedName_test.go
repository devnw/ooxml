package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionDefinedNameConstructor(t *testing.T) {
	v := sml.NewCT_RevisionDefinedName()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionDefinedName must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionDefinedName should validate: %s", err)
	}
}

func TestCT_RevisionDefinedNameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionDefinedName()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionDefinedName()
	xml.Unmarshal(buf, v2)
}
