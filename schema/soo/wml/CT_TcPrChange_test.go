package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcPrChangeConstructor(t *testing.T) {
	v := wml.NewCT_TcPrChange()
	if v == nil {
		t.Errorf("wml.NewCT_TcPrChange must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TcPrChange should validate: %s", err)
	}
}

func TestCT_TcPrChangeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TcPrChange()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TcPrChange()
	xml.Unmarshal(buf, v2)
}
