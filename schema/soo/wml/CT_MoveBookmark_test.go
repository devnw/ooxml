package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MoveBookmarkConstructor(t *testing.T) {
	v := wml.NewCT_MoveBookmark()
	if v == nil {
		t.Errorf("wml.NewCT_MoveBookmark must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MoveBookmark should validate: %s", err)
	}
}

func TestCT_MoveBookmarkMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MoveBookmark()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MoveBookmark()
	xml.Unmarshal(buf, v2)
}
