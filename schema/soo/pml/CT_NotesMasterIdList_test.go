package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_NotesMasterIdListConstructor(t *testing.T) {
	v := pml.NewCT_NotesMasterIdList()
	if v == nil {
		t.Errorf("pml.NewCT_NotesMasterIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesMasterIdList should validate: %s", err)
	}
}

func TestCT_NotesMasterIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesMasterIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesMasterIdList()
	xml.Unmarshal(buf, v2)
}
