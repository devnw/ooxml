package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SectPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_SectPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_SectPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SectPrChange should validate: %s", err)
	}
}

func TestCT_SectPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SectPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SectPrChange()
	xml.Unmarshal(buf, v2)
}
