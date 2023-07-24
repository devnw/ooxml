package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionsConstructor(t *testing.T) {
	v := sml.NewCT_Revisions()
	if v == nil {
		t.Errorf("sml.NewCT_Revisions must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Revisions should validate: %s", err)
	}
}

func TestCT_RevisionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Revisions()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Revisions()
	xml.Unmarshal(buf, v2)
}
