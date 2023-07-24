package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_CommentConstructor(t *testing.T) {
	v := sml.NewCT_Comment()
	if v == nil {
		t.Errorf("sml.NewCT_Comment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_Comment should validate: %s", err)
	}
}

func TestCT_CommentMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_Comment()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_Comment()
	xml.Unmarshal(buf, v2)
}
