package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_RPrDefaultConstructor(t *testing.T) {
	v := wml.NewCT_RPrDefault()
	if v == nil {
		t.Errorf("wml.NewCT_RPrDefault must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_RPrDefault should validate: %s", err)
	}
}

func TestCT_RPrDefaultMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_RPrDefault()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_RPrDefault()
	xml.Unmarshal(buf, v2)
}
