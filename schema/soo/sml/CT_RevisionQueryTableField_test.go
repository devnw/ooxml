package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionQueryTableFieldConstructor(t *testing.T) {
	v := sml.NewCT_RevisionQueryTableField()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionQueryTableField must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionQueryTableField should validate: %s", err)
	}
}

func TestCT_RevisionQueryTableFieldMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionQueryTableField()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionQueryTableField()
	xml.Unmarshal(buf, v2)
}
