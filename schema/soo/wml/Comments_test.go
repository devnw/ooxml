package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCommentsConstructor(t *testing.T) {
	v := wml.NewComments()
	if v == nil {
		t.Errorf("wml.NewComments must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Comments should validate: %s", err)
	}
}

func TestCommentsMarshalUnmarshal(t *testing.T) {
	v := wml.NewComments()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewComments()
	xml.Unmarshal(buf, v2)
}
