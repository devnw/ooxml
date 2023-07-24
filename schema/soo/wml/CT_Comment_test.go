package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CommentConstructor(t *testing.T) {
	v := wml.NewCT_Comment()
	if v == nil {
		t.Errorf("wml.NewCT_Comment must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Comment should validate: %s", err)
	}
}

func TestCT_CommentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Comment()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Comment()
	xml.Unmarshal(buf, v2)
}
