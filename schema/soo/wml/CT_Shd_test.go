package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ShdConstructor(t *testing.T) {
	v := wml.NewCT_Shd()
	if v == nil {
		t.Errorf("wml.NewCT_Shd must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Shd should validate: %s", err)
	}
}

func TestCT_ShdMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Shd()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Shd()
	xml.Unmarshal(buf, v2)
}
