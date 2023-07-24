package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_NumLvlConstructor(t *testing.T) {
	v := wml.NewCT_NumLvl()
	if v == nil {
		t.Errorf("wml.NewCT_NumLvl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_NumLvl should validate: %s", err)
	}
}

func TestCT_NumLvlMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_NumLvl()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_NumLvl()
	xml.Unmarshal(buf, v2)
}
