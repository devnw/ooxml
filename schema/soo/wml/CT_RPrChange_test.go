package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_RPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_RPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RPrChange should validate: %s", err)
	}
}

func TestCT_RPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RPrChange()
	xml.Unmarshal(buf, v2)
}
