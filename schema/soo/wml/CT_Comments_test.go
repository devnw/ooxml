package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CommentsConstructor(t *testing.T) {
	v := wml.NewCT_Comments()
	if v == nil {
		t.Errorf("wml.NewCT_Comments must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Comments should validate: %s", err)
	}
}

func TestCT_CommentsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Comments()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Comments()
	xml.Unmarshal(buf, v2)
}
