package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionCellChangeConstructor(t *testing.T) {
	v := sml.NewCT_RevisionCellChange()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionCellChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionCellChange should validate: %s", err)
	}
}

func TestCT_RevisionCellChangeMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionCellChange()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionCellChange()
	xml.Unmarshal(buf, v2)
}
