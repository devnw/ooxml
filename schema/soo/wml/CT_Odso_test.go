package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_OdsoConstructor(t *testing.T) {
	v := wml.NewCT_Odso()
	if v == nil {
		t.Errorf("wml.NewCT_Odso must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Odso should validate: %s", err)
	}
}

func TestCT_OdsoMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Odso()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Odso()
	xml.Unmarshal(buf, v2)
}
