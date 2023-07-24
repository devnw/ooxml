package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionRowColumnConstructor(t *testing.T) {
	v := sml.NewCT_RevisionRowColumn()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionRowColumn must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionRowColumn should validate: %s", err)
	}
}

func TestCT_RevisionRowColumnMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionRowColumn()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionRowColumn()
	xml.Unmarshal(buf, v2)
}
