package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_RevisionCommentConstructor(t *testing.T) {
	v := sml.NewCT_RevisionComment()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionComment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionComment should validate: %s", err)
	}
}

func TestCT_RevisionCommentMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionComment()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionComment()
	xml.Unmarshal(buf, v2)
}
