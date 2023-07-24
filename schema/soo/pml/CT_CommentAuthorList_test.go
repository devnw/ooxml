package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommentAuthorListConstructor(t *testing.T) {
	v := pml.NewCT_CommentAuthorList()
	if v == nil {
		t.Errorf("pml.NewCT_CommentAuthorList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommentAuthorList should validate: %s", err)
	}
}

func TestCT_CommentAuthorListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommentAuthorList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommentAuthorList()
	xml.Unmarshal(buf, v2)
}
