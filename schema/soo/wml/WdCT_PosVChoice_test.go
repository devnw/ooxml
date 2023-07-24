package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_PosVChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_PosVChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_PosVChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_PosVChoice should validate: %s", err)
	}
}

func TestWdCT_PosVChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_PosVChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_PosVChoice()
	xml.Unmarshal(buf, v2)
}
