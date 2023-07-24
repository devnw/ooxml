package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCommentsConstructor(t *testing.T) {
	v := sml.NewComments()
	if v == nil {
		t.Errorf("sml.NewComments must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Comments should validate: %s", err)
	}
}

func TestCommentsMarshalUnmarshal(t *testing.T) {
	v := sml.NewComments()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewComments()
	xml.Unmarshal(buf, v2)
}
