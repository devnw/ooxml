package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TcConstructor(t *testing.T) {
	v := wml.NewCT_Tc()
	if v == nil {
		t.Errorf("wml.NewCT_Tc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Tc should validate: %s", err)
	}
}

func TestCT_TcMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Tc()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Tc()
	xml.Unmarshal(buf, v2)
}
