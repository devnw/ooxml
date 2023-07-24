package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingGroupConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingGroup()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingGroup should validate: %s", err)
	}
}

func TestWdCT_WordprocessingGroupMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingGroup()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingGroup()
	xml.Unmarshal(buf, v2)
}
