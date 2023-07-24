package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CommentsConstructor(t *testing.T) {
	v := sml.NewCT_Comments()
	if v == nil {
		t.Errorf("sml.NewCT_Comments must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Comments should validate: %s", err)
	}
}

func TestCT_CommentsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Comments()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Comments()
	xml.Unmarshal(buf, v2)
}
