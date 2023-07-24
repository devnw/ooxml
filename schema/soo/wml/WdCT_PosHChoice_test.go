package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_PosHChoiceConstructor(t *testing.T) {
	v := wml.NewWdCT_PosHChoice()
	if v == nil {
		t.Errorf("wml.NewWdCT_PosHChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_PosHChoice should validate: %s", err)
	}
}

func TestWdCT_PosHChoiceMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_PosHChoice()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_PosHChoice()
	xml.Unmarshal(buf, v2)
}
