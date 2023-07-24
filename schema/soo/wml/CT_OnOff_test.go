package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_OnOffConstructor(t *testing.T) {
	v := wml.NewCT_OnOff()
	if v == nil {
		t.Errorf("wml.NewCT_OnOff must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_OnOff should validate: %s", err)
	}
}

func TestCT_OnOffMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_OnOff()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_OnOff()
	xml.Unmarshal(buf, v2)
}
