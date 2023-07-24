package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_BookmarkRangeConstructor(t *testing.T) {
	v := wml.NewCT_BookmarkRange()
	if v == nil {
		t.Errorf("wml.NewCT_BookmarkRange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_BookmarkRange should validate: %s", err)
	}
}

func TestCT_BookmarkRangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_BookmarkRange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_BookmarkRange()
	xml.Unmarshal(buf, v2)
}
