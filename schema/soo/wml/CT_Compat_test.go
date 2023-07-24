package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CompatConstructor(t *testing.T) {
	v := wml.NewCT_Compat()
	if v == nil {
		t.Errorf("wml.NewCT_Compat must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Compat should validate: %s", err)
	}
}

func TestCT_CompatMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Compat()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Compat()
	xml.Unmarshal(buf, v2)
}
