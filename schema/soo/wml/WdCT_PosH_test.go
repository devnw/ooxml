package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_PosHConstructor(t *testing.T) {
	v := wml.NewWdCT_PosH()
	if v == nil {
		t.Errorf("wml.NewWdCT_PosH must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_PosH should validate: %s", err)
	}
}

func TestWdCT_PosHMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_PosH()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_PosH()
	xml.Unmarshal(buf, v2)
}
