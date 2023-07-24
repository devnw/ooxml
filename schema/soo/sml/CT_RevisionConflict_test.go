package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionConflictConstructor(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionConflict must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionConflict should validate: %s", err)
	}
}

func TestCT_RevisionConflictMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionConflict()
	xml.Unmarshal(buf, v2)
}
