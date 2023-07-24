package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionSheetRenameConstructor(t *testing.T) {
	v := sml.NewCT_RevisionSheetRename()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionSheetRename must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionSheetRename should validate: %s", err)
	}
}

func TestCT_RevisionSheetRenameMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionSheetRename()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionSheetRename()
	xml.Unmarshal(buf, v2)
}
