package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CommentListConstructor(t *testing.T) {
	v := sml.NewCT_CommentList()
	if v == nil {
		t.Errorf("sml.NewCT_CommentList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CommentList should validate: %s", err)
	}
}

func TestCT_CommentListMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CommentList()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CommentList()
	xml.Unmarshal(buf, v2)
}
