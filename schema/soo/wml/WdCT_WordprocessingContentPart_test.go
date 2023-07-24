package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingContentPartConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingContentPart()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingContentPart must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingContentPart should validate: %s", err)
	}
}

func TestWdCT_WordprocessingContentPartMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingContentPart()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingContentPart()
	xml.Unmarshal(buf, v2)
}
