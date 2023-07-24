package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionInsertSheetConstructor(t *testing.T) {
	v := sml.NewCT_RevisionInsertSheet()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionInsertSheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionInsertSheet should validate: %s", err)
	}
}

func TestCT_RevisionInsertSheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionInsertSheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionInsertSheet()
	xml.Unmarshal(buf, v2)
}
