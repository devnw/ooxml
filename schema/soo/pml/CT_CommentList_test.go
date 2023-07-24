package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommentListConstructor(t *testing.T) {
	v := pml.NewCT_CommentList()
	if v == nil {
		t.Errorf("pml.NewCT_CommentList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommentList should validate: %s", err)
	}
}

func TestCT_CommentListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommentList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommentList()
	xml.Unmarshal(buf, v2)
}
