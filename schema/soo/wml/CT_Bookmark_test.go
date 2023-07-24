package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BookmarkConstructor(t *testing.T) {
	v := wml.NewCT_Bookmark()
	if v == nil {
		t.Errorf("wml.NewCT_Bookmark must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Bookmark should validate: %s", err)
	}
}

func TestCT_BookmarkMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Bookmark()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Bookmark()
	xml.Unmarshal(buf, v2)
}
