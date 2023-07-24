package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CommentAuthorConstructor(t *testing.T) {
	v := pml.NewCT_CommentAuthor()
	if v == nil {
		t.Errorf("pml.NewCT_CommentAuthor must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CommentAuthor should validate: %s", err)
	}
}

func TestCT_CommentAuthorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CommentAuthor()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CommentAuthor()
	xml.Unmarshal(buf, v2)
}
