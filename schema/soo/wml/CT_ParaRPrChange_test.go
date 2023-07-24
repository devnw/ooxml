package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ParaRPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_ParaRPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_ParaRPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_ParaRPrChange should validate: %s", err)
	}
}

func TestCT_ParaRPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_ParaRPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_ParaRPrChange()
	xml.Unmarshal(buf, v2)
}
