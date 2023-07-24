package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_JcConstructor(t *testing.T) {
	v := wml.NewCT_Jc()
	if v == nil {
		t.Errorf("wml.NewCT_Jc must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Jc should validate: %s", err)
	}
}

func TestCT_JcMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Jc()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Jc()
	xml.Unmarshal(buf, v2)
}
