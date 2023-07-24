package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WordprocessingGroupChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingGroupChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingGroupChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingGroupChoice should validate: %s", err)
	}
}

func TestWdCT_WordprocessingGroupChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingGroupChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingGroupChoice()
	xml.Unmarshal(buf, v2)
}
